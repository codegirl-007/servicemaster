package oauth

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"servicemaster/internal/qbo/tokens"
	"servicemaster/internal/store"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var testTenantID = uuid.MustParse("00000000-0000-0000-0000-000000000001")
var testEncryptor = &noopEncryptor{}

type noopEncryptor struct{}

func (noopEncryptor) Encrypt(plaintext []byte) ([]byte, error) {
	return plaintext, nil
}

func (noopEncryptor) Decrypt(encrypted []byte) ([]byte, error) {
	return encrypted, nil
}

// stubStateStore implements StateStore with a simple map.
// It does not reproduce SQL filtering semantics — tests configure
// exactly what GetActiveOAuthStateByChecksum returns.
type stubStateStore struct {
	states map[string]store.QboOauthState
}

func newStubStateStore() *stubStateStore {
	return &stubStateStore{states: make(map[string]store.QboOauthState)}
}

func (s *stubStateStore) CreateOAuthState(_ context.Context, arg store.CreateOAuthStateParams) (store.QboOauthState, error) {
	st := store.QboOauthState{
		ID:             arg.ID,
		TenantID:       arg.TenantID,
		StateChecksum:  arg.StateChecksum,
		EncryptedState: arg.EncryptedState,
		ExpiresAt:      arg.ExpiresAt,
		CreatedAt:      time.Now(),
	}
	s.states[arg.StateChecksum] = st
	return st, nil
}

func (s *stubStateStore) GetActiveOAuthStateByChecksum(_ context.Context, checksum string) (store.QboOauthState, error) {
	st, ok := s.states[checksum]
	if !ok {
		return store.QboOauthState{}, sql.ErrNoRows
	}
	return st, nil
}

// stubTokenLoader implements TokenLoader with a simple map.
type stubTokenLoader struct {
	tokens map[uuid.UUID]tokens.Tokens
}

func newStubTokenLoader() *stubTokenLoader {
	return &stubTokenLoader{tokens: make(map[uuid.UUID]tokens.Tokens)}
}

func (l *stubTokenLoader) Load(_ context.Context, connectionID uuid.UUID) (tokens.Tokens, error) {
	t, ok := l.tokens[connectionID]
	if !ok {
		return tokens.Tokens{}, sql.ErrNoRows
	}
	return t, nil
}

// stubTxStore implements TxDataStore with configurable function fields.
// Each test sets function fields for the methods it exercises and leaves
// unused methods at their defaults (returns sql.ErrNoRows or no-ops).
type stubTxStore struct {
	getConnByTenantFn func(context.Context, uuid.UUID) (store.QboConnection, error)
	createConnFn      func(context.Context, store.CreateQBOConnectionParams) (store.QboConnection, error)
	updateStateFn     func(context.Context, store.UpdateQBOConnectionStateParams) (store.QboConnection, error)
	updateNameFn      func(context.Context, store.UpdateQBOConnectionCompanyNameParams) (store.QboConnection, error)
	consumeStateFn    func(context.Context, uuid.UUID) (uuid.UUID, error)
	storeTokensFn     func(context.Context, uuid.UUID, uuid.UUID, string, string, time.Time, time.Time) error
	createEventFn     func(context.Context, store.CreateQBOConnectionEventParams) (store.QboConnectionEvent, error)

	// Recording fields populated by the default success functions.
	CreatedConn   *store.CreateQBOConnectionParams
	ConsumedID    *uuid.UUID
	StoredConnID  *uuid.UUID
	StoredAccess  string
	StoredRefresh string
	CreatedEvent  *store.CreateQBOConnectionEventParams
	UpdatedConnID *uuid.UUID
	UpdatedNameID *uuid.UUID
}

func newStubTxStore() *stubTxStore {
	s := &stubTxStore{}
	s.getConnByTenantFn = func(_ context.Context, _ uuid.UUID) (store.QboConnection, error) {
		return store.QboConnection{}, sql.ErrNoRows
	}
	s.createConnFn = func(_ context.Context, p store.CreateQBOConnectionParams) (store.QboConnection, error) {
		s.CreatedConn = &p
		return store.QboConnection{
			ID: p.ID, TenantID: p.TenantID, RealmID: p.RealmID,
			CompanyName: p.CompanyName, State: p.State,
			CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}, nil
	}
	s.updateStateFn = func(_ context.Context, p store.UpdateQBOConnectionStateParams) (store.QboConnection, error) {
		s.UpdatedConnID = &p.ID
		return store.QboConnection{}, nil
	}
	s.updateNameFn = func(_ context.Context, p store.UpdateQBOConnectionCompanyNameParams) (store.QboConnection, error) {
		s.UpdatedNameID = &p.ID
		return store.QboConnection{}, nil
	}
	s.consumeStateFn = func(_ context.Context, id uuid.UUID) (uuid.UUID, error) {
		s.ConsumedID = &id
		return id, nil
	}
	s.storeTokensFn = func(_ context.Context, connID uuid.UUID, _ uuid.UUID, access, refresh string, _, _ time.Time) error {
		s.StoredConnID = &connID
		s.StoredAccess = access
		s.StoredRefresh = refresh
		return nil
	}
	s.createEventFn = func(_ context.Context, p store.CreateQBOConnectionEventParams) (store.QboConnectionEvent, error) {
		s.CreatedEvent = &p
		return store.QboConnectionEvent{}, nil
	}
	return s
}

func (s *stubTxStore) GetQBOConnectionByTenant(ctx context.Context, tenantID uuid.UUID) (store.QboConnection, error) {
	return s.getConnByTenantFn(ctx, tenantID)
}

func (s *stubTxStore) CreateQBOConnection(ctx context.Context, arg store.CreateQBOConnectionParams) (store.QboConnection, error) {
	return s.createConnFn(ctx, arg)
}

func (s *stubTxStore) UpdateQBOConnectionState(ctx context.Context, arg store.UpdateQBOConnectionStateParams) (store.QboConnection, error) {
	return s.updateStateFn(ctx, arg)
}

func (s *stubTxStore) UpdateQBOConnectionCompanyName(ctx context.Context, arg store.UpdateQBOConnectionCompanyNameParams) (store.QboConnection, error) {
	return s.updateNameFn(ctx, arg)
}

func (s *stubTxStore) ConsumeOAuthState(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	return s.consumeStateFn(ctx, id)
}

func (s *stubTxStore) StoreTokens(ctx context.Context, connectionID uuid.UUID, tenantID uuid.UUID, accessToken string, refreshToken string, accessExpiresAt time.Time, refreshExpiresAt time.Time) error {
	return s.storeTokensFn(ctx, connectionID, tenantID, accessToken, refreshToken, accessExpiresAt, refreshExpiresAt)
}

func (s *stubTxStore) CreateQBOConnectionEvent(ctx context.Context, arg store.CreateQBOConnectionEventParams) (store.QboConnectionEvent, error) {
	return s.createEventFn(ctx, arg)
}

// simpleTxRunner wraps a TxDataStore and passes it through to the
// closure without any transaction overhead.
type simpleTxRunner struct {
	txStore TxDataStore
}

func (r *simpleTxRunner) RunInTx(ctx context.Context, fn func(context.Context, TxDataStore) error) error {
	return fn(ctx, r.txStore)
}

// errorTxRunner is a TxRunner that always returns the configured error.
type errorTxRunner struct {
	err error
}

func (r *errorTxRunner) RunInTx(_ context.Context, _ func(context.Context, TxDataStore) error) error {
	return r.err
}

func newTestService(
	ep oauth2.Endpoint,
	clientID, clientSecret, redirectURL string,
	scopes []string,
	stateStore StateStore,
	tokenLoader TokenLoader,
	encryptor tokens.Encryptor,
	txRunner TxRunner,
) *Service {
	return newServiceForTest(ep, clientID, clientSecret, redirectURL, scopes, stateStore, tokenLoader, encryptor, txRunner)
}

func fakeTokenServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]interface{}{
			"access_token":  "test-access-token",
			"refresh_token": "test-refresh-token",
			"token_type":    "bearer",
			"expires_in":    3600,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))
}

func encryptState(t *testing.T, tenantID uuid.UUID, expiresAt time.Time) []byte {
	t.Helper()
	b, err := testEncryptor.Encrypt([]byte(tenantID.String() + "|" + fmt.Sprintf("%d", expiresAt.Unix())))
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestAuthURL(t *testing.T) {
	stateStore := newStubStateStore()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: "http://example.com/token"},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&simpleTxRunner{txStore: newStubTxStore()},
	)

	url, err := svc.AuthURL(context.Background(), testTenantID)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(url, "client_id=test-client") {
		t.Error("url missing client_id")
	}
	if !strings.Contains(url, "redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fqbo%2Fcallback") {
		t.Error("url missing redirect_uri")
	}
	if !strings.Contains(url, "response_type=code") {
		t.Error("url missing response_type=code")
	}

	if len(stateStore.states) != 1 {
		t.Fatalf("expected 1 state, got %d", len(stateStore.states))
	}
	for _, s := range stateStore.states {
		if s.TenantID != testTenantID {
			t.Errorf("state tenant = %v, want %v", s.TenantID, testTenantID)
		}
		if time.Now().After(s.ExpiresAt) {
			t.Error("state already expired")
		}
	}
}

func TestExchange_HappyPath(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	stateStore := newStubStateStore()
	txStore := newStubTxStore()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: ts.URL},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&simpleTxRunner{txStore: txStore},
	)

	stateChecksum := "test-state-checksum"
	expiresAt := time.Now().Add(10 * time.Minute)
	stateStore.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encryptState(t, testTenantID, expiresAt),
		ExpiresAt:      expiresAt,
		CreatedAt:      time.Now(),
	}

	err := svc.Exchange(context.Background(), "auth-code", stateChecksum, "123456789", "Test Company")
	if err != nil {
		t.Fatal(err)
	}

	if txStore.ConsumedID == nil {
		t.Error("state was not consumed")
	}
	if txStore.CreatedConn == nil {
		t.Fatal("connection was not created")
	}
	if txStore.CreatedConn.State != "connected" {
		t.Errorf("connection state = %q, want \"connected\"", txStore.CreatedConn.State)
	}
	if txStore.CreatedConn.RealmID != "123456789" {
		t.Errorf("realm_id = %q, want \"123456789\"", txStore.CreatedConn.RealmID)
	}
	if txStore.CreatedConn.CompanyName.String != "Test Company" {
		t.Errorf("company_name = %q, want \"Test Company\"", txStore.CreatedConn.CompanyName.String)
	}
	if txStore.StoredConnID == nil {
		t.Error("tokens were not stored")
	}
	if txStore.StoredAccess != "test-access-token" {
		t.Errorf("access token = %q, want \"test-access-token\"", txStore.StoredAccess)
	}
	if txStore.StoredRefresh != "test-refresh-token" {
		t.Errorf("refresh token = %q, want \"test-refresh-token\"", txStore.StoredRefresh)
	}
	if txStore.CreatedEvent == nil {
		t.Fatal("event was not created")
	}
	if txStore.CreatedEvent.EventType != "connected" {
		t.Errorf("event type = %q, want \"connected\"", txStore.CreatedEvent.EventType)
	}
}

func TestExchange_StateNotFound(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	stateStore := newStubStateStore()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: ts.URL},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&simpleTxRunner{txStore: newStubTxStore()},
	)

	err := svc.Exchange(context.Background(), "code", "nonexistent-state", "r1", "C")
	if !errors.Is(err, ErrStateNotFound) {
		t.Errorf("expected ErrStateNotFound, got %v", err)
	}
}

func TestExchange_StateExpired(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	stateStore := newStubStateStore()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: ts.URL},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&simpleTxRunner{txStore: newStubTxStore()},
	)

	err := svc.Exchange(context.Background(), "code", "expired-state", "r1", "C")
	if !errors.Is(err, ErrStateNotFound) {
		t.Errorf("expected ErrStateNotFound, got %v", err)
	}
}

func TestExchange_TenantMismatch(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	stateStore := newStubStateStore()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: ts.URL},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&simpleTxRunner{txStore: newStubTxStore()},
	)

	otherTenant := uuid.MustParse("00000000-0000-0000-0000-000000000099")
	stateChecksum := "mismatch-state"
	expiresAt := time.Now().Add(10 * time.Minute)
	stateStore.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encryptState(t, otherTenant, expiresAt),
		ExpiresAt:      expiresAt,
		CreatedAt:      time.Now(),
	}

	err := svc.Exchange(context.Background(), "code", stateChecksum, "r1", "C")
	if !errors.Is(err, ErrTenantMismatch) {
		t.Errorf("expected ErrTenantMismatch, got %v", err)
	}
}

func TestExchange_StateAlreadyConsumed(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	stateStore := newStubStateStore()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: ts.URL},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&simpleTxRunner{txStore: newStubTxStore()},
	)

	err := svc.Exchange(context.Background(), "code", "already-consumed", "r1", "C")
	if !errors.Is(err, ErrStateNotFound) {
		t.Errorf("expected ErrStateNotFound, got %v", err)
	}
}

func TestExchange_TokenEndpointFails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	stateStore := newStubStateStore()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: ts.URL},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&simpleTxRunner{txStore: newStubTxStore()},
	)

	stateChecksum := "token-fail-state"
	expiresAt := time.Now().Add(10 * time.Minute)
	stateStore.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encryptState(t, testTenantID, expiresAt),
		ExpiresAt:      expiresAt,
		CreatedAt:      time.Now(),
	}

	err := svc.Exchange(context.Background(), "code", stateChecksum, "r1", "C")
	if err == nil {
		t.Fatal("expected error from failed token exchange")
	}
}

func TestExchange_MalformedState(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	stateStore := newStubStateStore()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: ts.URL},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&simpleTxRunner{txStore: newStubTxStore()},
	)

	stateChecksum := "malformed-state"
	stateStore.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: []byte("garbage-no-pipe"),
		ExpiresAt:      time.Now().Add(10 * time.Minute),
		CreatedAt:      time.Now(),
	}

	err := svc.Exchange(context.Background(), "code", stateChecksum, "r1", "C")
	if err == nil {
		t.Fatal("expected error from malformed state")
	}
}

func TestExchange_Reconnect(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	stateStore := newStubStateStore()
	txStore := newStubTxStore()
	existingID := uuid.New()

	txStore.getConnByTenantFn = func(_ context.Context, tenantID uuid.UUID) (store.QboConnection, error) {
		return store.QboConnection{
			ID:       existingID,
			TenantID: tenantID,
			RealmID:  "old-realm",
			CompanyName: sql.NullString{
				String: "Old Company",
				Valid:  true,
			},
			State: "disconnected",
		}, nil
	}

	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: ts.URL},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&simpleTxRunner{txStore: txStore},
	)

	stateChecksum := "reconnect-state"
	expiresAt := time.Now().Add(10 * time.Minute)
	stateStore.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encryptState(t, testTenantID, expiresAt),
		ExpiresAt:      expiresAt,
		CreatedAt:      time.Now(),
	}

	err := svc.Exchange(context.Background(), "auth-code", stateChecksum, "new-realm", "New Company")
	if err != nil {
		t.Fatal(err)
	}

	if txStore.UpdatedConnID == nil {
		t.Error("connection state was not updated")
	}
	if txStore.UpdatedConnID != nil && *txStore.UpdatedConnID != existingID {
		t.Errorf("updated connection = %v, want %v", *txStore.UpdatedConnID, existingID)
	}
	if txStore.UpdatedNameID == nil {
		t.Error("company name was not updated")
	}
	if txStore.UpdatedNameID != nil && *txStore.UpdatedNameID != existingID {
		t.Errorf("updated name connection = %v, want %v", *txStore.UpdatedNameID, existingID)
	}
	if txStore.StoredConnID == nil {
		t.Error("tokens were not stored")
	}
	if txStore.StoredConnID != nil && *txStore.StoredConnID != existingID {
		t.Errorf("tokens stored under connection = %v, want %v", *txStore.StoredConnID, existingID)
	}
	if txStore.StoredAccess != "test-access-token" {
		t.Errorf("access token = %q, want \"test-access-token\"", txStore.StoredAccess)
	}
}

func TestNewClient_NoTokens(t *testing.T) {
	tokenLoader := newStubTokenLoader()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: "http://example.com/token"},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		newStubStateStore(), tokenLoader, testEncryptor,
		&simpleTxRunner{txStore: newStubTxStore()},
	)

	_, err := svc.NewClient(context.Background(), uuid.New())
	if err == nil {
		t.Fatal("expected error for connection with no stored tokens")
	}
}

func TestNewClient(t *testing.T) {
	tokenLoader := newStubTokenLoader()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: "http://example.com/token"},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		newStubStateStore(), tokenLoader, testEncryptor,
		&simpleTxRunner{txStore: newStubTxStore()},
	)

	connID := uuid.New()
	expiry := time.Now().Add(1 * time.Hour)
	tokenLoader.tokens[connID] = tokens.Tokens{
		ConnectionID:    connID,
		AccessToken:     "stored-access",
		RefreshToken:    "stored-refresh",
		AccessExpiresAt: expiry,
		Version:         1,
	}

	client, err := svc.NewClient(context.Background(), connID)
	if err != nil {
		t.Fatal(err)
	}
	if client == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestExchange_BeginTxFails(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	stateStore := newStubStateStore()
	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: ts.URL},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&errorTxRunner{err: errors.New("db connection lost")},
	)

	stateChecksum := "begin-tx-fail"
	expiresAt := time.Now().Add(10 * time.Minute)
	stateStore.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encryptState(t, testTenantID, expiresAt),
		ExpiresAt:      expiresAt,
		CreatedAt:      time.Now(),
	}

	err := svc.Exchange(context.Background(), "auth-code", stateChecksum, "r1", "C")
	if err == nil {
		t.Fatal("expected error from failed beginTx")
	}
}

func TestExchange_WriteTimeStateConsumed(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	stateStore := newStubStateStore()
	txStore := newStubTxStore()
	txStore.consumeStateFn = func(_ context.Context, _ uuid.UUID) (uuid.UUID, error) {
		return uuid.UUID{}, sql.ErrNoRows
	}

	svc := newTestService(
		oauth2.Endpoint{AuthURL: "https://example.com/auth", TokenURL: ts.URL},
		"test-client", "test-secret", "http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		stateStore, newStubTokenLoader(), testEncryptor,
		&simpleTxRunner{txStore: txStore},
	)

	stateChecksum := "write-time-consumed"
	expiresAt := time.Now().Add(10 * time.Minute)
	stateStore.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encryptState(t, testTenantID, expiresAt),
		ExpiresAt:      expiresAt,
		CreatedAt:      time.Now(),
	}

	err := svc.Exchange(context.Background(), "auth-code", stateChecksum, "r1", "C")
	if !errors.Is(err, ErrStateConsumed) {
		t.Errorf("expected ErrStateConsumed, got %v", err)
	}
}
