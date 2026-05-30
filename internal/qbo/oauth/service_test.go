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

type fakeStore struct {
	states     map[string]store.QboOauthState
	connections map[uuid.UUID]store.QboConnection
	events     []store.QboConnectionEvent
}

func newFakeStore() *fakeStore {
	return &fakeStore{
		states:      make(map[string]store.QboOauthState),
		connections: make(map[uuid.UUID]store.QboConnection),
	}
}

func (f *fakeStore) CreateOAuthState(_ context.Context, arg store.CreateOAuthStateParams) (store.QboOauthState, error) {
	s := store.QboOauthState{
		ID:             arg.ID,
		TenantID:       arg.TenantID,
		StateChecksum:  arg.StateChecksum,
		EncryptedState: arg.EncryptedState,
		ExpiresAt:      arg.ExpiresAt,
		CreatedAt:      time.Now(),
	}
	f.states[arg.StateChecksum] = s
	return s, nil
}

func (f *fakeStore) GetActiveOAuthStateByChecksum(_ context.Context, stateChecksum string) (store.QboOauthState, error) {
	s, ok := f.states[stateChecksum]
	if !ok {
		return store.QboOauthState{}, sql.ErrNoRows
	}
	if s.ConsumedAt.Valid {
		return store.QboOauthState{}, sql.ErrNoRows
	}
	if time.Now().After(s.ExpiresAt) {
		return store.QboOauthState{}, sql.ErrNoRows
	}
	return s, nil
}

func (f *fakeStore) ConsumeOAuthState(_ context.Context, id uuid.UUID) (uuid.UUID, error) {
	for checksum, s := range f.states {
		if s.ID == id {
			if s.ConsumedAt.Valid {
				return uuid.UUID{}, sql.ErrNoRows
			}
			s.ConsumedAt = sql.NullTime{Time: time.Now(), Valid: true}
			f.states[checksum] = s
			return id, nil
		}
	}
	return uuid.UUID{}, sql.ErrNoRows
}

func (f *fakeStore) GetQBOConnectionByTenant(_ context.Context, tenantID uuid.UUID) (store.QboConnection, error) {
	for _, c := range f.connections {
		if c.TenantID == tenantID {
			return c, nil
		}
	}
	return store.QboConnection{}, sql.ErrNoRows
}

func (f *fakeStore) CreateQBOConnection(_ context.Context, arg store.CreateQBOConnectionParams) (store.QboConnection, error) {
	c := store.QboConnection{
		ID:          arg.ID,
		TenantID:    arg.TenantID,
		RealmID:     arg.RealmID,
		CompanyName: arg.CompanyName,
		State:       arg.State,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	f.connections[arg.ID] = c
	return c, nil
}

func (f *fakeStore) UpdateQBOConnectionState(_ context.Context, arg store.UpdateQBOConnectionStateParams) (store.QboConnection, error) {
	c, ok := f.connections[arg.ID]
	if !ok {
		return store.QboConnection{}, sql.ErrNoRows
	}
	c.State = arg.State
	c.UpdatedAt = time.Now()
	f.connections[arg.ID] = c
	return c, nil
}

func (f *fakeStore) UpdateQBOConnectionCompanyName(_ context.Context, arg store.UpdateQBOConnectionCompanyNameParams) (store.QboConnection, error) {
	c, ok := f.connections[arg.ID]
	if !ok {
		return store.QboConnection{}, sql.ErrNoRows
	}
	c.CompanyName = arg.CompanyName
	c.UpdatedAt = time.Now()
	f.connections[arg.ID] = c
	return c, nil
}

func (f *fakeStore) CreateQBOConnectionEvent(_ context.Context, arg store.CreateQBOConnectionEventParams) (store.QboConnectionEvent, error) {
	e := store.QboConnectionEvent{
		ID:              arg.ID,
		QboConnectionID: arg.QboConnectionID,
		TenantID:        arg.TenantID,
		EventType:       arg.EventType,
		Message:         arg.Message,
		Metadata:        arg.Metadata,
		OccurredAt:      time.Now(),
	}
	f.events = append(f.events, e)
	return e, nil
}

type fakeTokenDataStore struct {
	rows map[uuid.UUID]store.QboConnectionToken
}

func newFakeTokenDataStore() *fakeTokenDataStore {
	return &fakeTokenDataStore{rows: make(map[uuid.UUID]store.QboConnectionToken)}
}

func (f *fakeTokenDataStore) UpsertQBOConnectionTokens(_ context.Context, arg store.UpsertQBOConnectionTokensParams) (store.QboConnectionToken, error) {
	row := store.QboConnectionToken{
		QboConnectionID:       arg.QboConnectionID,
		TenantID:              arg.TenantID,
		EncryptedAccessToken:  arg.EncryptedAccessToken,
		EncryptedRefreshToken: arg.EncryptedRefreshToken,
		AccessTokenExpiresAt:  arg.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: arg.RefreshTokenExpiresAt,
		TokenType:             arg.TokenType,
		Scope:                 arg.Scope,
		Version:               1,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}
	f.rows[arg.QboConnectionID] = row
	return row, nil
}

func (f *fakeTokenDataStore) GetQBOConnectionTokens(_ context.Context, connectionID uuid.UUID) (store.QboConnectionToken, error) {
	row, ok := f.rows[connectionID]
	if !ok {
		return store.QboConnectionToken{}, sql.ErrNoRows
	}
	return row, nil
}

func (f *fakeTokenDataStore) ReplaceQBOConnectionTokensIfVersion(_ context.Context, _ store.ReplaceQBOConnectionTokensIfVersionParams) (store.QboConnectionToken, error) {
	return store.QboConnectionToken{}, sql.ErrNoRows
}

func newTestService(fs *fakeStore, tokenStore *fakeTokenDataStore, tokenURL string) *Service {
	ep := oauth2.Endpoint{
		AuthURL:  "https://example.com/auth",
		TokenURL: tokenURL,
	}
	return newService(
		ep,
		"test-client",
		"test-secret",
		"http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		fs,
		tokens.NewService(tokenStore, testEncryptor),
		testEncryptor,
		nil,
	)
}

func fakeTokenServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "method not allowed", 405)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token":  "test-access-token",
			"refresh_token": "test-refresh-token",
			"token_type":    "bearer",
			"expires_in":    3600,
		})
	}))
}

func TestAuthURL(t *testing.T) {
	fs := newFakeStore()
	svc := newTestService(fs, newFakeTokenDataStore(), "http://example.com/token")

	url, err := svc.AuthURL(context.Background(), testTenantID)
	if err != nil {
		t.Fatal(err)
	}

	// URL contains the expected OAuth params.
	if !strings.Contains(url, "client_id=test-client") {
		t.Error("url missing client_id")
	}
	if !strings.Contains(url, "redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fqbo%2Fcallback") {
		t.Error("url missing redirect_uri")
	}
	if !strings.Contains(url, "response_type=code") {
		t.Error("url missing response_type=code")
	}

	// State was persisted.
	if len(fs.states) != 1 {
		t.Fatalf("expected 1 state, got %d", len(fs.states))
	}
	for _, s := range fs.states {
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

	fs := newFakeStore()
	fts := newFakeTokenDataStore()
	svc := newTestService(fs, fts, ts.URL)

	// Pre-populate a valid state.
	stateChecksum := "test-state-checksum"
	expiresAt := time.Now().Add(10 * time.Minute)
	encrypted := encryptState(t, testTenantID, expiresAt)
	fs.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encrypted,
		ExpiresAt:      expiresAt,
		CreatedAt:      time.Now(),
	}

	err := svc.Exchange(context.Background(), "auth-code", stateChecksum, "123456789", "Test Company")
	if err != nil {
		t.Fatal(err)
	}

	// State was consumed.
	s := fs.states[stateChecksum]
	if !s.ConsumedAt.Valid {
		t.Error("state was not consumed")
	}

	// Connection was created.
	if len(fs.connections) != 1 {
		t.Fatalf("expected 1 connection, got %d", len(fs.connections))
	}
	for _, c := range fs.connections {
		if c.State != "connected" {
			t.Errorf("connection state = %q, want \"connected\"", c.State)
		}
		if c.RealmID != "123456789" {
			t.Errorf("realm_id = %q, want \"123456789\"", c.RealmID)
		}
		if c.CompanyName.String != "Test Company" {
			t.Errorf("company_name = %q, want \"Test Company\"", c.CompanyName.String)
		}
	}

	// Tokens were stored.
	if len(fts.rows) != 1 {
		t.Fatalf("expected 1 token row, got %d", len(fts.rows))
	}
	for _, row := range fts.rows {
		if string(row.EncryptedAccessToken) != "test-access-token" {
			t.Errorf("encrypted_access_token = %q, want \"test-access-token\"", string(row.EncryptedAccessToken))
		}
		if string(row.EncryptedRefreshToken) != "test-refresh-token" {
			t.Errorf("encrypted_refresh_token = %q, want \"test-refresh-token\"", string(row.EncryptedRefreshToken))
		}
	}

	// Event was logged.
	if len(fs.events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(fs.events))
	}
	if fs.events[0].EventType != "connected" {
		t.Errorf("event type = %q, want \"connected\"", fs.events[0].EventType)
	}
}

func TestExchange_StateNotFound(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	svc := newTestService(newFakeStore(), newFakeTokenDataStore(), ts.URL)
	err := svc.Exchange(context.Background(), "code", "nonexistent-state", "r1", "C")
	if !errors.Is(err, ErrStateNotFound) {
		t.Errorf("expected ErrStateNotFound, got %v", err)
	}
}

func TestExchange_StateExpired(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	fs := newFakeStore()
	svc := newTestService(fs, newFakeTokenDataStore(), ts.URL)

	stateChecksum := "expired-state"
	encrypted := encryptState(t, testTenantID, time.Now().Add(-1*time.Hour))
	fs.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encrypted,
		ExpiresAt:      time.Now().Add(-1 * time.Hour),
		CreatedAt:      time.Now().Add(-2 * time.Hour),
	}

	err := svc.Exchange(context.Background(), "code", stateChecksum, "r1", "C")
	if !errors.Is(err, ErrStateNotFound) {
		t.Errorf("expected ErrStateNotFound, got %v", err)
	}
}

func TestExchange_TenantMismatch(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	fs := newFakeStore()
	svc := newTestService(fs, newFakeTokenDataStore(), ts.URL)

	otherTenant := uuid.MustParse("00000000-0000-0000-0000-000000000099")
	stateChecksum := "mismatch-state"
	expiresAt := time.Now().Add(10 * time.Minute)
	// Encrypt with otherTenant but store the state row with testTenantID.
	encrypted := encryptState(t, otherTenant, expiresAt)
	fs.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encrypted,
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

	fs := newFakeStore()
	fts := newFakeTokenDataStore()
	svc := newTestService(fs, fts, ts.URL)

	// Pre-populate a state that is already consumed.
	stateChecksum := "already-consumed"
	expiresAt := time.Now().Add(10 * time.Minute)
	encrypted := encryptState(t, testTenantID, expiresAt)
	fs.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encrypted,
		ConsumedAt:     sql.NullTime{Time: time.Now(), Valid: true},
		ExpiresAt:      expiresAt,
		CreatedAt:      time.Now(),
	}

	err := svc.Exchange(context.Background(), "code", stateChecksum, "r1", "C")
	// GetActiveOAuthStateByChecksum filters consumed states, so this returns not found.
	if !errors.Is(err, ErrStateNotFound) {
		t.Errorf("expected ErrStateNotFound, got %v", err)
	}
}

func TestExchange_TokenEndpointFails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	fs := newFakeStore()
	svc := newTestService(fs, newFakeTokenDataStore(), ts.URL)

	stateChecksum := "token-fail-state"
	expiresAt := time.Now().Add(10 * time.Minute)
	fs.states[stateChecksum] = store.QboOauthState{
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

	fs := newFakeStore()
	svc := newTestService(fs, newFakeTokenDataStore(), ts.URL)

	stateChecksum := "malformed-state"
	fs.states[stateChecksum] = store.QboOauthState{
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

	fs := newFakeStore()
	fts := newFakeTokenDataStore()
	svc := newTestService(fs, fts, ts.URL)

	// Pre-populate an existing disconnected connection.
	existingID := uuid.New()
	fs.connections[existingID] = store.QboConnection{
		ID:       existingID,
		TenantID: testTenantID,
		RealmID:  "old-realm",
		CompanyName: sql.NullString{
			String: "Old Company",
			Valid:  true,
		},
		State: "disconnected",
	}

	// Pre-populate a valid state.
	stateChecksum := "reconnect-state"
	expiresAt := time.Now().Add(10 * time.Minute)
	encrypted := encryptState(t, testTenantID, expiresAt)
	fs.states[stateChecksum] = store.QboOauthState{
		ID:             uuid.New(),
		TenantID:       testTenantID,
		StateChecksum:  stateChecksum,
		EncryptedState: encrypted,
		ExpiresAt:      expiresAt,
		CreatedAt:      time.Now(),
	}

	err := svc.Exchange(context.Background(), "auth-code", stateChecksum, "new-realm", "New Company")
	if err != nil {
		t.Fatal(err)
	}

	// Same connection ID was reused.
	if len(fs.connections) != 1 {
		t.Fatalf("expected 1 connection, got %d", len(fs.connections))
	}
	c := fs.connections[existingID]
	if c.State != "connected" {
		t.Errorf("connection state = %q, want \"connected\"", c.State)
	}
	if c.CompanyName.String != "New Company" {
		t.Errorf("company_name = %q, want \"New Company\"", c.CompanyName.String)
	}
	// RealmID stays from the original connection (we don't update it).
	if c.RealmID != "old-realm" {
		t.Errorf("realm_id = %q, want \"old-realm\" (should not change on reconnect)", c.RealmID)
	}

	// Tokens were stored under the existing connection ID.
	row, ok := fts.rows[existingID]
	if !ok {
		t.Fatal("tokens not stored under existing connection ID")
	}
	if string(row.EncryptedAccessToken) != "test-access-token" {
		t.Errorf("encrypted_access_token = %q, want \"test-access-token\"", string(row.EncryptedAccessToken))
	}
}

func TestNewClient_NoTokens(t *testing.T) {
	fts := newFakeTokenDataStore()
	svc := newTestService(newFakeStore(), fts, "http://example.com/token")

	_, err := svc.NewClient(context.Background(), uuid.New())
	if err == nil {
		t.Fatal("expected error for connection with no stored tokens")
	}
}

func TestNewClient(t *testing.T) {
	fts := newFakeTokenDataStore()
	svc := newTestService(newFakeStore(), fts, "http://example.com/token")

	connID := uuid.New()
	expiry := time.Now().Add(1 * time.Hour)
	fts.rows[connID] = store.QboConnectionToken{
		QboConnectionID:      connID,
		EncryptedAccessToken:  []byte("stored-access"),
		EncryptedRefreshToken: []byte("stored-refresh"),
		AccessTokenExpiresAt:  expiry,
		Version:               1,
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

	fs := newFakeStore()
	fts := newFakeTokenDataStore()
	svc := newTestService(fs, fts, ts.URL)

	// Override beginTx to simulate a DB connection failure.
	svc.beginTx = func(_ context.Context) (*exchangeTx, error) {
		return nil, errors.New("db connection lost")
	}

	stateChecksum := "begin-tx-fail"
	expiresAt := time.Now().Add(10 * time.Minute)
	fs.states[stateChecksum] = store.QboOauthState{
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

// errConsumeStore wraps fakeStore but always fails ConsumeOAuthState
// to simulate a concurrent-consumption race.
type errConsumeStore struct {
	*fakeStore
}

func (e *errConsumeStore) ConsumeOAuthState(_ context.Context, _ uuid.UUID) (uuid.UUID, error) {
	return uuid.UUID{}, sql.ErrNoRows
}

func TestExchange_WriteTimeStateConsumed(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	fs := &errConsumeStore{newFakeStore()}
	fts := newFakeTokenDataStore()

	ep := oauth2.Endpoint{
		AuthURL:  "https://example.com/auth",
		TokenURL: ts.URL,
	}
	svc := newService(
		ep,
		"test-client",
		"test-secret",
		"http://localhost:8080/qbo/callback",
		[]string{"com.intuit.quickbooks.accounting"},
		fs,
		tokens.NewService(fts, testEncryptor),
		testEncryptor,
		nil,
	)

	stateChecksum := "write-time-consumed"
	expiresAt := time.Now().Add(10 * time.Minute)
	fs.states[stateChecksum] = store.QboOauthState{
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

func encryptState(t *testing.T, tenantID uuid.UUID, expiresAt time.Time) []byte {
	t.Helper()
	b, err := testEncryptor.Encrypt([]byte(tenantID.String() + "|" + fmt.Sprintf("%d", expiresAt.Unix())))
	if err != nil {
		t.Fatal(err)
	}
	return b
}


