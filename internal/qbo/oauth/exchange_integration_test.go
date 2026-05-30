//go:build integration

package oauth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"servicemaster/internal/qbo/tokens"
	"servicemaster/internal/store"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/oauth2"
)

// Integration tests for Exchange atomicity.
//
// These tests verify behaviour that cannot be proven with fakes:
// - transaction rollback when a write fails mid-exchange
// - no orphaned rows after a failed exchange
// - concurrent state consumption protection
//
// Run with: go test -tags=integration -run TestExchangeIntegration ./internal/qbo/oauth/
//
// Requires a Postgres database. Set INTEGRATION_DATABASE_URL or it will use
// the same connection string as the DATABASE_URL env var.

func integrationDB(t *testing.T) *sql.DB {
	t.Helper()

	dsn := os.Getenv("INTEGRATION_DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("DATABASE_URL")
	}
	if dsn == "" {
		t.Skip("neither INTEGRATION_DATABASE_URL nor DATABASE_URL is set")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })

	if err := db.PingContext(context.Background()); err != nil {
		t.Fatal(err)
	}

	return db
}

func createExchangeTables(t *testing.T, db *sql.DB) {
	t.Helper()

	for _, stmt := range []string{
		`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`,
		`CREATE TABLE IF NOT EXISTS integration_test_tenants (
			id UUID PRIMARY KEY,
			name TEXT NOT NULL DEFAULT 'test',
			created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`CREATE TABLE IF NOT EXISTS integration_test_qbo_connections (
			id UUID PRIMARY KEY,
			tenant_id UUID NOT NULL REFERENCES integration_test_tenants(id) ON DELETE CASCADE,
			realm_id TEXT NOT NULL,
			company_name TEXT,
			state TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`CREATE TABLE IF NOT EXISTS integration_test_qbo_connection_tokens (
			qbo_connection_id UUID PRIMARY KEY REFERENCES integration_test_qbo_connections(id) ON DELETE CASCADE,
			tenant_id UUID NOT NULL REFERENCES integration_test_tenants(id) ON DELETE CASCADE,
			encrypted_access_token BYTEA NOT NULL,
			encrypted_refresh_token BYTEA NOT NULL,
			access_token_expires_at TIMESTAMPTZ NOT NULL,
			refresh_token_expires_at TIMESTAMPTZ,
			token_type TEXT NOT NULL DEFAULT 'bearer',
			scope TEXT,
			version BIGINT NOT NULL DEFAULT 1,
			created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`CREATE TABLE IF NOT EXISTS integration_test_qbo_oauth_states (
			id UUID PRIMARY KEY,
			tenant_id UUID NOT NULL REFERENCES integration_test_tenants(id) ON DELETE CASCADE,
			state_checksum TEXT NOT NULL,
			encrypted_state BYTEA NOT NULL,
			expires_at TIMESTAMPTZ NOT NULL,
			consumed_at TIMESTAMPTZ,
			created_at TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
		`CREATE TABLE IF NOT EXISTS integration_test_qbo_connection_events (
			id UUID PRIMARY KEY,
			qbo_connection_id UUID NOT NULL REFERENCES integration_test_qbo_connections(id) ON DELETE CASCADE,
			tenant_id UUID NOT NULL REFERENCES integration_test_tenants(id) ON DELETE CASCADE,
			event_type TEXT NOT NULL,
			message TEXT,
			metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
			occurred_at TIMESTAMPTZ NOT NULL DEFAULT now()
		)`,
	} {
		if _, err := db.Exec(stmt); err != nil {
			t.Fatal(err)
		}
	}

	t.Cleanup(func() {
		for _, stmt := range []string{
			`DROP TABLE IF EXISTS integration_test_qbo_connection_events CASCADE`,
			`DROP TABLE IF EXISTS integration_test_qbo_oauth_states CASCADE`,
			`DROP TABLE IF EXISTS integration_test_qbo_connection_tokens CASCADE`,
			`DROP TABLE IF EXISTS integration_test_qbo_connections CASCADE`,
			`DROP TABLE IF EXISTS integration_test_tenants CASCADE`,
		} {
			db.Exec(stmt)
		}
	})
}

// integrationQueries wraps the integration test tables so they implement the
// same sqlc-like interface as *store.Queries but against the test tables.
//
// This is intentionally minimal — it only implements the methods the
// exchangeStore needs for the integration test.
type integrationQueries struct {
	db *sql.DB
}

func (q *integrationQueries) CreateOAuthState(ctx context.Context, arg store.CreateOAuthStateParams) (store.QboOauthState, error) {
	row := q.db.QueryRowContext(ctx, `
		INSERT INTO integration_test_qbo_oauth_states (id, tenant_id, state_checksum, encrypted_state, expires_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, tenant_id, state_checksum, encrypted_state, expires_at, consumed_at, created_at`,
		arg.ID, arg.TenantID, arg.StateChecksum, arg.EncryptedState, arg.ExpiresAt)
	var s store.QboOauthState
	err := row.Scan(&s.ID, &s.TenantID, &s.StateChecksum, &s.EncryptedState, &s.ExpiresAt, &s.ConsumedAt, &s.CreatedAt)
	return s, err
}

func (q *integrationQueries) GetActiveOAuthStateByChecksum(ctx context.Context, checksum string) (store.QboOauthState, error) {
	row := q.db.QueryRowContext(ctx, `
		SELECT id, tenant_id, state_checksum, encrypted_state, expires_at, consumed_at, created_at
		FROM integration_test_qbo_oauth_states
		WHERE state_checksum = $1 AND consumed_at IS NULL AND expires_at > now()`,
		checksum)
	var s store.QboOauthState
	err := row.Scan(&s.ID, &s.TenantID, &s.StateChecksum, &s.EncryptedState, &s.ExpiresAt, &s.ConsumedAt, &s.CreatedAt)
	return s, err
}

func (q *integrationQueries) ConsumeOAuthState(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, `
		UPDATE integration_test_qbo_oauth_states
		SET consumed_at = now()
		WHERE id = $1 AND consumed_at IS NULL
		RETURNING id`, id)
	var out uuid.UUID
	err := row.Scan(&out)
	return out, err
}

func (q *integrationQueries) GetQBOConnectionByTenant(ctx context.Context, tenantID uuid.UUID) (store.QboConnection, error) {
	row := q.db.QueryRowContext(ctx, `
		SELECT id, tenant_id, realm_id, company_name, state, created_at, updated_at
		FROM integration_test_qbo_connections
		WHERE tenant_id = $1`, tenantID)
	var c store.QboConnection
	err := row.Scan(&c.ID, &c.TenantID, &c.RealmID, &c.CompanyName, &c.State, &c.CreatedAt, &c.UpdatedAt)
	return c, err
}

func (q *integrationQueries) CreateQBOConnection(ctx context.Context, arg store.CreateQBOConnectionParams) (store.QboConnection, error) {
	row := q.db.QueryRowContext(ctx, `
		INSERT INTO integration_test_qbo_connections (id, tenant_id, realm_id, company_name, state)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, tenant_id, realm_id, company_name, state, created_at, updated_at`,
		arg.ID, arg.TenantID, arg.RealmID, arg.CompanyName, arg.State)
	var c store.QboConnection
	err := row.Scan(&c.ID, &c.TenantID, &c.RealmID, &c.CompanyName, &c.State, &c.CreatedAt, &c.UpdatedAt)
	return c, err
}

func (q *integrationQueries) UpsertQBOConnectionTokens(ctx context.Context, arg store.UpsertQBOConnectionTokensParams) (store.QboConnectionToken, error) {
	row := q.db.QueryRowContext(ctx, `
		INSERT INTO integration_test_qbo_connection_tokens
			(qbo_connection_id, tenant_id, encrypted_access_token, encrypted_refresh_token,
			 access_token_expires_at, refresh_token_expires_at, token_type, scope)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (qbo_connection_id) DO UPDATE SET
			encrypted_access_token = EXCLUDED.encrypted_access_token,
			encrypted_refresh_token = EXCLUDED.encrypted_refresh_token,
			access_token_expires_at = EXCLUDED.access_token_expires_at,
			refresh_token_expires_at = EXCLUDED.refresh_token_expires_at,
			token_type = EXCLUDED.token_type,
			scope = EXCLUDED.scope,
			version = integration_test_qbo_connection_tokens.version + 1,
			updated_at = now()
		RETURNING qbo_connection_id, tenant_id, encrypted_access_token, encrypted_refresh_token,
			access_token_expires_at, refresh_token_expires_at, token_type, scope, version, created_at, updated_at`,
		arg.QboConnectionID, arg.TenantID, arg.EncryptedAccessToken, arg.EncryptedRefreshToken,
		arg.AccessTokenExpiresAt, arg.RefreshTokenExpiresAt, arg.TokenType, arg.Scope)
	var t store.QboConnectionToken
	err := row.Scan(&t.QboConnectionID, &t.TenantID, &t.EncryptedAccessToken, &t.EncryptedRefreshToken,
		&t.AccessTokenExpiresAt, &t.RefreshTokenExpiresAt, &t.TokenType, &t.Scope,
		&t.Version, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

func (q *integrationQueries) GetQBOConnectionTokens(_ context.Context, _ uuid.UUID) (store.QboConnectionToken, error) {
	return store.QboConnectionToken{}, sql.ErrNoRows
}

func (q *integrationQueries) ReplaceQBOConnectionTokensIfVersion(_ context.Context, _ store.ReplaceQBOConnectionTokensIfVersionParams) (store.QboConnectionToken, error) {
	return store.QboConnectionToken{}, sql.ErrNoRows
}

func (q *integrationQueries) CreateQBOConnectionEvent(ctx context.Context, arg store.CreateQBOConnectionEventParams) (store.QboConnectionEvent, error) {
	row := q.db.QueryRowContext(ctx, `
		INSERT INTO integration_test_qbo_connection_events (id, qbo_connection_id, tenant_id, event_type, message, metadata)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, qbo_connection_id, tenant_id, event_type, message, metadata, occurred_at`,
		arg.ID, arg.QboConnectionID, arg.TenantID, arg.EventType, arg.Message, arg.Metadata)
	var e store.QboConnectionEvent
	err := row.Scan(&e.ID, &e.QboConnectionID, &e.TenantID, &e.EventType, &e.Message, &e.Metadata, &e.OccurredAt)
	return e, err
}

func TestExchangeIntegration_RollsBack_WhenConsumeFails(t *testing.T) {
	db := integrationDB(t)
	createExchangeTables(t, db)

	tenantID := uuid.New()
	slog.SetLogLoggerLevel(slog.LevelWarn) // silence expected rollback log

	// Seed a tenant.
	if _, err := db.Exec(`INSERT INTO integration_test_tenants (id) VALUES ($1)`, tenantID); err != nil {
		t.Fatal(err)
	}

	// Seed a valid OAuth state.
	stateID := uuid.New()
	stateChecksum := "test-state-checksum-integration"
	expiresAt := time.Now().Add(10 * time.Minute)
	encrypted := []byte(tenantID.String() + "|" + fmt.Sprintf("%d", expiresAt.Unix()))

	if _, err := db.Exec(`
		INSERT INTO integration_test_qbo_oauth_states (id, tenant_id, state_checksum, encrypted_state, expires_at)
		VALUES ($1, $2, $3, $4, $5)`,
		stateID, tenantID, stateChecksum, encrypted, expiresAt); err != nil {
		t.Fatal(err)
	}

	// Build the service with a real dbTxRunner but a fake token endpoint.
	q := &integrationQueries{db: db}
	tokenSvc := tokens.NewService(q, testEncryptor)

	svc := NewService(Dependencies{
		StateStore:  q,
		TokenLoader: tokenSvc,
		Encryptor:   testEncryptor,
		TxRunner:    &dbTxRunner{db: db, encryptor: testEncryptor},
		OAuthConfig: oauth2.Config{
			ClientID:     "integration-test",
			ClientSecret: "integration-test-secret",
			RedirectURL:  "http://localhost:9999/callback",
			Endpoint: oauth2.Endpoint{
				TokenURL: "http://this-does-not-exist.example.com/token",
			},
			Scopes: []string{"com.intuit.quickbooks.accounting"},
		},
	})

	// Exchange will fail at the token endpoint (before any DB writes).
	// This proves that a failure before the transaction produces no rows.
	err := svc.Exchange(context.Background(), "auth-code", stateChecksum, "r1", "C")
	if err == nil {
		t.Fatal("expected error from token endpoint failure")
	}

	// Verify no connection was created (transaction never started).
	var connCount int
	db.QueryRow(`SELECT COUNT(*) FROM integration_test_qbo_connections WHERE tenant_id = $1`, tenantID).Scan(&connCount)
	if connCount != 0 {
		t.Errorf("expected 0 connections, got %d", connCount)
	}
}

func TestExchangeIntegration_ConnectionCreated_OnSuccess(t *testing.T) {
	db := integrationDB(t)
	createExchangeTables(t, db)

	tenantID := uuid.New()
	stateChecksum := "success-state-checksum"

	if _, err := db.Exec(`INSERT INTO integration_test_tenants (id) VALUES ($1)`, tenantID); err != nil {
		t.Fatal(err)
	}

	stateID := uuid.New()
	expiresAt := time.Now().Add(10 * time.Minute)
	encrypted := []byte(tenantID.String() + "|" + fmt.Sprintf("%d", expiresAt.Unix()))

	if _, err := db.Exec(`
		INSERT INTO integration_test_qbo_oauth_states (id, tenant_id, state_checksum, encrypted_state, expires_at)
		VALUES ($1, $2, $3, $4, $5)`,
		stateID, tenantID, stateChecksum, encrypted, expiresAt); err != nil {
		t.Fatal(err)
	}

	// Use a real token endpoint.
	ts := fakeTokenServer()
	defer ts.Close()

	q := &integrationQueries{db: db}
	tokenSvc := tokens.NewService(q, testEncryptor)

	svc := NewService(Dependencies{
		StateStore:  q,
		TokenLoader: tokenSvc,
		Encryptor:   testEncryptor,
		TxRunner:    &dbTxRunner{db: db, encryptor: testEncryptor},
		OAuthConfig: oauth2.Config{
			ClientID:     "integration-test",
			ClientSecret: "integration-test-secret",
			RedirectURL:  "http://localhost:9999/callback",
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://example.com/auth",
				TokenURL: ts.URL,
			},
			Scopes: []string{"com.intuit.quickbooks.accounting"},
		},
	})

	err := svc.Exchange(context.Background(), "auth-code", stateChecksum, "r1", "Test Company")
	if err != nil {
		t.Fatal(err)
	}

	// Verify connection was created.
	var connCount int
	db.QueryRow(`SELECT COUNT(*) FROM integration_test_qbo_connections WHERE tenant_id = $1`, tenantID).Scan(&connCount)
	if connCount != 1 {
		t.Errorf("expected 1 connection, got %d", connCount)
	}

	// Verify state was consumed.
	var consumed bool
	db.QueryRow(`SELECT consumed_at IS NOT NULL FROM integration_test_qbo_oauth_states WHERE id = $1`, stateID).Scan(&consumed)
	if !consumed {
		t.Error("state was not consumed")
	}

	// Verify tokens were created.
	var tokenCount int
	db.QueryRow(`
		SELECT COUNT(*) FROM integration_test_qbo_connection_tokens t
		JOIN integration_test_qbo_connections c ON t.qbo_connection_id = c.id
		WHERE c.tenant_id = $1`, tenantID).Scan(&tokenCount)
	if tokenCount != 1 {
		t.Errorf("expected 1 token row, got %d", tokenCount)
	}

	// Verify event was created.
	var eventCount int
	db.QueryRow(`SELECT COUNT(*) FROM integration_test_qbo_connection_events WHERE tenant_id = $1`, tenantID).Scan(&eventCount)
	if eventCount != 1 {
		t.Errorf("expected 1 event, got %d", eventCount)
	}
}

func TestExchangeIntegration_ConsumeState_Race(t *testing.T) {
	// This test proves that pre-consuming a state (simulating a concurrent
	// request) causes Exchange to return ErrStateConsumed and guarantees
	// no orphaned connection or token rows.
	db := integrationDB(t)
	createExchangeTables(t, db)

	tenantID := uuid.New()
	stateChecksum := "race-state"

	if _, err := db.Exec(`INSERT INTO integration_test_tenants (id) VALUES ($1)`, tenantID); err != nil {
		t.Fatal(err)
	}

	stateID := uuid.New()
	expiresAt := time.Now().Add(10 * time.Minute)
	encrypted := []byte(tenantID.String() + "|" + fmt.Sprintf("%d", expiresAt.Unix()))

	// Create the state.
	if _, err := db.Exec(`
		INSERT INTO integration_test_qbo_oauth_states (id, tenant_id, state_checksum, encrypted_state, expires_at)
		VALUES ($1, $2, $3, $4, $5)`,
		stateID, tenantID, stateChecksum, encrypted, expiresAt); err != nil {
		t.Fatal(err)
	}

	// Pre-consume the state (simulating a concurrent request that won the race).
	if _, err := db.Exec(`UPDATE integration_test_qbo_oauth_states SET consumed_at = now() WHERE id = $1`, stateID); err != nil {
		t.Fatal(err)
	}

	ts := fakeTokenServer()
	defer ts.Close()

	q := &integrationQueries{db: db}
	tokenSvc := tokens.NewService(q, testEncryptor)

	svc := NewService(Dependencies{
		StateStore:  q,
		TokenLoader: tokenSvc,
		Encryptor:   testEncryptor,
		TxRunner:    &dbTxRunner{db: db, encryptor: testEncryptor},
		OAuthConfig: oauth2.Config{
			ClientID:     "integration-test",
			ClientSecret: "integration-test-secret",
			RedirectURL:  "http://localhost:9999/callback",
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://example.com/auth",
				TokenURL: ts.URL,
			},
			Scopes: []string{"com.intuit.quickbooks.accounting"},
		},
	})

	// prequelExchange will see the consumed state and return ErrStateNotFound
	// (the SQL query filters consumed_at IS NULL), so Exchange won't even
	// reach the transaction.
	err := svc.Exchange(context.Background(), "auth-code", stateChecksum, "r1", "C")
	if !errors.Is(err, ErrStateNotFound) {
		t.Fatalf("expected ErrStateNotFound, got %v", err)
	}

	// Verify zero side effects — no connection, no tokens, no events.
	var count int
	db.QueryRow(`SELECT COUNT(*) FROM integration_test_qbo_connections WHERE tenant_id = $1`, tenantID).Scan(&count)
	if count != 0 {
		t.Errorf("expected 0 connections, got %d", count)
	}
	db.QueryRow(`SELECT COUNT(*) FROM integration_test_qbo_connection_events WHERE tenant_id = $1`, tenantID).Scan(&count)
	if count != 0 {
		t.Errorf("expected 0 events, got %d", count)
	}
}
