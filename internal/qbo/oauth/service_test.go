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

// nullExchangeStore implements ExchangeStore with safe defaults:
// query methods return sql.ErrNoRows, command methods are no-ops.
// Tests embed this and override only the methods they exercise.
type nullExchangeStore struct{}

func (nullExchangeStore) GetQBOConnectionByTenant(_ context.Context, _ uuid.UUID) (store.QboConnection, error) {
	return store.QboConnection{}, sql.ErrNoRows
}

func (nullExchangeStore) CreateQBOConnection(_ context.Context, arg store.CreateQBOConnectionParams) (store.QboConnection, error) {
	return store.QboConnection{
		ID: arg.ID, TenantID: arg.TenantID, RealmID: arg.RealmID,
		CompanyName: arg.CompanyName, State: arg.State,
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}, nil
}

func (nullExchangeStore) UpdateQBOConnectionState(_ context.Context, _ store.UpdateQBOConnectionStateParams) (store.QboConnection, error) {
	return store.QboConnection{}, nil
}

func (nullExchangeStore) UpdateQBOConnectionCompanyName(_ context.Context, _ store.UpdateQBOConnectionCompanyNameParams) (store.QboConnection, error) {
	return store.QboConnection{}, nil
}

func (nullExchangeStore) ConsumeOAuthState(_ context.Context, id uuid.UUID) (uuid.UUID, error) {
	return id, nil
}

func (nullExchangeStore) StoreTokens(_ context.Context, _ uuid.UUID, _ uuid.UUID, _, _ string, _, _ time.Time) error {
	return nil
}

func (nullExchangeStore) CreateQBOConnectionEvent(_ context.Context, _ store.CreateQBOConnectionEventParams) (store.QboConnectionEvent, error) {
	return store.QboConnectionEvent{}, nil
}

// happyPathStore records all ExchangeStore calls for verification
// in the happy path test.
type happyPathStore struct {
	nullExchangeStore

	ConsumedID    *uuid.UUID
	CreatedConn   *store.QboConnection
	StoredConnID  *uuid.UUID
	StoredAccess  string
	StoredRefresh string
	CreatedEvent  *store.CreateQBOConnectionEventParams
}

func (s *happyPathStore) CreateQBOConnection(ctx context.Context, arg store.CreateQBOConnectionParams) (store.QboConnection, error) {
	conn := store.QboConnection{
		ID: arg.ID, TenantID: arg.TenantID, RealmID: arg.RealmID,
		CompanyName: arg.CompanyName, State: arg.State,
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}
	s.CreatedConn = &conn
	return conn, nil
}

func (s *happyPathStore) ConsumeOAuthState(_ context.Context, id uuid.UUID) (uuid.UUID, error) {
	s.ConsumedID = &id
	return id, nil
}

func (s *happyPathStore) StoreTokens(_ context.Context, connID uuid.UUID, _ uuid.UUID, access, refresh string, _, _ time.Time) error {
	s.StoredConnID = &connID
	s.StoredAccess = access
	s.StoredRefresh = refresh
	return nil
}

func (s *happyPathStore) CreateQBOConnectionEvent(_ context.Context, arg store.CreateQBOConnectionEventParams) (store.QboConnectionEvent, error) {
	s.CreatedEvent = &arg
	return store.QboConnectionEvent{}, nil
}

// reconnectStore records calls made during the reconnect scenario.
type reconnectStore struct {
	nullExchangeStore

	ExistingConn   store.QboConnection
	UpdatedConnID  *uuid.UUID
	UpdatedNameID  *uuid.UUID
	StoredConnID   *uuid.UUID
	StoredAccess   string
	StoredRefresh  string
}

func (s *reconnectStore) GetQBOConnectionByTenant(_ context.Context, _ uuid.UUID) (store.QboConnection, error) {
	return s.ExistingConn, nil
}

func (s *reconnectStore) UpdateQBOConnectionState(_ context.Context, arg store.UpdateQBOConnectionStateParams) (store.QboConnection, error) {
	s.UpdatedConnID = &arg.ID
	return store.QboConnection{}, nil
}

func (s *reconnectStore) UpdateQBOConnectionCompanyName(_ context.Context, arg store.UpdateQBOConnectionCompanyNameParams) (store.QboConnection, error) {
	s.UpdatedNameID = &arg.ID
	return store.QboConnection{}, nil
}

func (s *reconnectStore) StoreTokens(_ context.Context, connID uuid.UUID, _ uuid.UUID, access, refresh string, _, _ time.Time) error {
	s.StoredConnID = &connID
	s.StoredAccess = access
	s.StoredRefresh = refresh
	return nil
}

// failConsumeStore fails ConsumeOAuthState with sql.ErrNoRows to
// simulate a concurrent state consumption race.
type failConsumeStore struct {
	nullExchangeStore
}

func (failConsumeStore) ConsumeOAuthState(_ context.Context, _ uuid.UUID) (uuid.UUID, error) {
	return uuid.UUID{}, sql.ErrNoRows
}

// stubStateStore implements StateStore with a simple map.
// It does not reproduce SQL filtering semantics: GetActiveOAuthStateByChecksum
// returns exactly what the test puts in the map.
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

// simpleTxRunner wraps an ExchangeStore and passes it through to the
// closure without any transaction overhead.
type simpleTxRunner struct {
	store ExchangeStore
}

func (r *simpleTxRunner) WithinTx(ctx context.Context, fn func(context.Context, ExchangeStore) error) error {
	return fn(ctx, r.store)
}

// errorTxRunner is a TxRunner that always returns the configured error.
type errorTxRunner struct {
	err error
}

func (r *errorTxRunner) WithinTx(_ context.Context, _ func(context.Context, ExchangeStore) error) error {
	return r.err
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

func testEndpoint(tokenURL string) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  "https://example.com/auth",
		TokenURL: tokenURL,
	}
}

func baseDeps() Dependencies {
	return Dependencies{
		StateStore:  newStubStateStore(),
		TokenLoader: newStubTokenLoader(),
		Encryptor:   testEncryptor,
		TxRunner:    &simpleTxRunner{store: &nullExchangeStore{}},
		OAuthConfig: oauth2.Config{
			ClientID:     "test-client",
			ClientSecret: "test-secret",
			RedirectURL:  "http://localhost:8080/qbo/callback",
			Scopes:       []string{"com.intuit.quickbooks.accounting"},
		},
	}
}

func TestAuthURL(t *testing.T) {
	stateStore := newStubStateStore()
	deps := baseDeps()
	deps.StateStore = stateStore
	deps.TxRunner = &simpleTxRunner{store: &nullExchangeStore{}}
	svc := NewService(deps)

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
	exchangeStore := &happyPathStore{}
	deps := baseDeps()
	deps.StateStore = stateStore
	deps.TxRunner = &simpleTxRunner{store: exchangeStore}
	deps.OAuthConfig.Endpoint = testEndpoint(ts.URL)
	svc := NewService(deps)

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

	if exchangeStore.ConsumedID == nil {
		t.Error("state was not consumed")
	}
	if exchangeStore.CreatedConn == nil {
		t.Fatal("connection was not created")
	}
	if exchangeStore.CreatedConn.State != "connected" {
		t.Errorf("connection state = %q, want \"connected\"", exchangeStore.CreatedConn.State)
	}
	if exchangeStore.CreatedConn.RealmID != "123456789" {
		t.Errorf("realm_id = %q, want \"123456789\"", exchangeStore.CreatedConn.RealmID)
	}
	if exchangeStore.CreatedConn.CompanyName.String != "Test Company" {
		t.Errorf("company_name = %q, want \"Test Company\"", exchangeStore.CreatedConn.CompanyName.String)
	}
	if exchangeStore.StoredConnID == nil {
		t.Error("tokens were not stored")
	}
	if exchangeStore.StoredAccess != "test-access-token" {
		t.Errorf("access token = %q, want \"test-access-token\"", exchangeStore.StoredAccess)
	}
	if exchangeStore.StoredRefresh != "test-refresh-token" {
		t.Errorf("refresh token = %q, want \"test-refresh-token\"", exchangeStore.StoredRefresh)
	}
	if exchangeStore.CreatedEvent == nil {
		t.Fatal("event was not created")
	}
	if exchangeStore.CreatedEvent.EventType != "connected" {
		t.Errorf("event type = %q, want \"connected\"", exchangeStore.CreatedEvent.EventType)
	}
}

func TestExchange_StateNotFound(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	deps := baseDeps()
	deps.OAuthConfig.Endpoint = testEndpoint(ts.URL)
	svc := NewService(deps)

	err := svc.Exchange(context.Background(), "code", "nonexistent-state", "r1", "C")
	if !errors.Is(err, ErrStateNotFound) {
		t.Errorf("expected ErrStateNotFound, got %v", err)
	}
}

func TestExchange_TenantMismatch(t *testing.T) {
	ts := fakeTokenServer()
	defer ts.Close()

	stateStore := newStubStateStore()
	deps := baseDeps()
	deps.StateStore = stateStore
	deps.OAuthConfig.Endpoint = testEndpoint(ts.URL)
	svc := NewService(deps)

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

func TestExchange_TokenEndpointFails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	stateStore := newStubStateStore()
	deps := baseDeps()
	deps.StateStore = stateStore
	deps.OAuthConfig.Endpoint = testEndpoint(ts.URL)
	svc := NewService(deps)

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
	stateStore := newStubStateStore()
	deps := baseDeps()
	deps.StateStore = stateStore
	svc := NewService(deps)

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
	exchangeStore := &reconnectStore{
		ExistingConn: store.QboConnection{
			ID:       uuid.New(),
			TenantID: testTenantID,
			RealmID:  "old-realm",
			CompanyName: sql.NullString{
				String: "Old Company",
				Valid:  true,
			},
			State: "disconnected",
		},
	}
	deps := baseDeps()
	deps.StateStore = stateStore
	deps.TxRunner = &simpleTxRunner{store: exchangeStore}
	deps.OAuthConfig.Endpoint = testEndpoint(ts.URL)
	svc := NewService(deps)

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

	existingID := exchangeStore.ExistingConn.ID
	if exchangeStore.UpdatedConnID == nil {
		t.Error("connection state was not updated")
	}
	if exchangeStore.UpdatedConnID != nil && *exchangeStore.UpdatedConnID != existingID {
		t.Errorf("updated connection = %v, want %v", *exchangeStore.UpdatedConnID, existingID)
	}
	if exchangeStore.UpdatedNameID == nil {
		t.Error("company name was not updated")
	}
	if exchangeStore.UpdatedNameID != nil && *exchangeStore.UpdatedNameID != existingID {
		t.Errorf("updated name connection = %v, want %v", *exchangeStore.UpdatedNameID, existingID)
	}
	if exchangeStore.StoredConnID == nil {
		t.Error("tokens were not stored")
	}
	if exchangeStore.StoredConnID != nil && *exchangeStore.StoredConnID != existingID {
		t.Errorf("tokens stored under connection = %v, want %v", *exchangeStore.StoredConnID, existingID)
	}
	if exchangeStore.StoredAccess != "test-access-token" {
		t.Errorf("access token = %q, want \"test-access-token\"", exchangeStore.StoredAccess)
	}
}

func TestNewClient_NoTokens(t *testing.T) {
	deps := baseDeps()
	svc := NewService(deps)

	_, err := svc.NewClient(context.Background(), uuid.New())
	if err == nil {
		t.Fatal("expected error for connection with no stored tokens")
	}
}

func TestNewClient(t *testing.T) {
	tokenLoader := newStubTokenLoader()
	deps := baseDeps()
	deps.TokenLoader = tokenLoader
	svc := NewService(deps)

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
	stateStore := newStubStateStore()
	deps := baseDeps()
	deps.StateStore = stateStore
	deps.TxRunner = &errorTxRunner{err: errors.New("db connection lost")}
	svc := NewService(deps)

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
	deps := baseDeps()
	deps.StateStore = stateStore
	deps.TxRunner = &simpleTxRunner{store: &failConsumeStore{}}
	deps.OAuthConfig.Endpoint = testEndpoint(ts.URL)
	svc := NewService(deps)

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
