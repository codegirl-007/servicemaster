// Package oauth implements the QuickBooks Online OAuth 2.0 connection flow.
package oauth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"servicemaster/internal/config"
	"servicemaster/internal/qbo/tokens"
	"servicemaster/internal/store"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var (
	// ErrStateNotFound is returned when the OAuth state is not found or expired.
	ErrStateNotFound = errors.New("oauth state not found or expired")
	// ErrStateConsumed is returned when the OAuth state was already used.
	ErrStateConsumed = errors.New("oauth state already consumed")
	// ErrTenantMismatch is returned when the decrypted tenant does not match the stored tenant.
	ErrTenantMismatch = errors.New("oauth state tenant mismatch")
)

// intuitEndpoint is the Intuit OAuth 2.0 endpoint for QuickBooks Online.
var intuitEndpoint = oauth2.Endpoint{
	AuthURL:  "https://appcenter.intuit.com/connect/oauth2",
	TokenURL: "https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer",
}

// stateTTL is how long an OAuth state is valid before it expires.
const stateTTL = 10 * time.Minute

// refreshTokenLifetime is how long QBO refresh tokens are valid (~6 months).
const refreshTokenLifetime = 180 * 24 * time.Hour

// StateStore handles non-transactional OAuth state read and write operations.
type StateStore interface {
	CreateOAuthState(context.Context, store.CreateOAuthStateParams) (store.QboOauthState, error)
	GetActiveOAuthStateByChecksum(context.Context, string) (store.QboOauthState, error)
}

// TokenLoader loads stored OAuth tokens for a connection.
type TokenLoader interface {
	Load(ctx context.Context, connectionID uuid.UUID) (tokens.Tokens, error)
}

// TxRunner executes a function within a database transaction.
type TxRunner interface {
	RunInTx(ctx context.Context, fn func(context.Context, TxDataStore) error) error
}

// TxDataStore provides data access methods available within the Exchange
// transaction. The production implementation wraps both *store.Queries and
// *tokens.Service so that all Exchange writes go through a single sql.Tx.
type TxDataStore interface {
	GetQBOConnectionByTenant(context.Context, uuid.UUID) (store.QboConnection, error)
	CreateQBOConnection(context.Context, store.CreateQBOConnectionParams) (store.QboConnection, error)
	UpdateQBOConnectionState(context.Context, store.UpdateQBOConnectionStateParams) (store.QboConnection, error)
	UpdateQBOConnectionCompanyName(context.Context, store.UpdateQBOConnectionCompanyNameParams) (store.QboConnection, error)
	ConsumeOAuthState(context.Context, uuid.UUID) (uuid.UUID, error)
	StoreTokens(ctx context.Context, connectionID uuid.UUID, tenantID uuid.UUID, accessToken string, refreshToken string, accessExpiresAt time.Time, refreshExpiresAt time.Time) error
	CreateQBOConnectionEvent(context.Context, store.CreateQBOConnectionEventParams) (store.QboConnectionEvent, error)
}

// dbTxRunner is the production TxRunner backed by *sql.DB.
type dbTxRunner struct {
	db        *sql.DB
	encryptor tokens.Encryptor
}

// RunInTx begins a transaction, constructs transactional data stores, calls fn,
// and commits on success or rolls back on failure.
func (r *dbTxRunner) RunInTx(ctx context.Context, fn func(context.Context, TxDataStore) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
				slog.Warn("exchange transaction rollback", "error", err)
			}
		}
	}()
	qtx := store.New(tx)
	ts := tokens.NewService(qtx, r.encryptor)
	txStore := &txDataStore{qtx: qtx, tokenSvc: ts}
	if err := fn(ctx, txStore); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	committed = true
	return nil
}

// txDataStore is the production TxDataStore backed by transactional
// *store.Queries and *tokens.Service instances.
type txDataStore struct {
	qtx      *store.Queries
	tokenSvc *tokens.Service
}

func (s *txDataStore) GetQBOConnectionByTenant(ctx context.Context, tenantID uuid.UUID) (store.QboConnection, error) {
	return s.qtx.GetQBOConnectionByTenant(ctx, tenantID)
}

func (s *txDataStore) CreateQBOConnection(ctx context.Context, arg store.CreateQBOConnectionParams) (store.QboConnection, error) {
	return s.qtx.CreateQBOConnection(ctx, arg)
}

func (s *txDataStore) UpdateQBOConnectionState(ctx context.Context, arg store.UpdateQBOConnectionStateParams) (store.QboConnection, error) {
	return s.qtx.UpdateQBOConnectionState(ctx, arg)
}

func (s *txDataStore) UpdateQBOConnectionCompanyName(ctx context.Context, arg store.UpdateQBOConnectionCompanyNameParams) (store.QboConnection, error) {
	return s.qtx.UpdateQBOConnectionCompanyName(ctx, arg)
}

func (s *txDataStore) ConsumeOAuthState(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	return s.qtx.ConsumeOAuthState(ctx, id)
}

func (s *txDataStore) StoreTokens(ctx context.Context, connectionID uuid.UUID, tenantID uuid.UUID, accessToken string, refreshToken string, accessExpiresAt time.Time, refreshExpiresAt time.Time) error {
	return s.tokenSvc.Store(ctx, connectionID, tenantID, accessToken, refreshToken, accessExpiresAt, refreshExpiresAt)
}

func (s *txDataStore) CreateQBOConnectionEvent(ctx context.Context, arg store.CreateQBOConnectionEventParams) (store.QboConnectionEvent, error) {
	return s.qtx.CreateQBOConnectionEvent(ctx, arg)
}

// Service manages the QBO OAuth 2.0 connection flow.
type Service struct {
	stateStore   StateStore
	tokenLoader  TokenLoader
	encryptor    tokens.Encryptor
	oauth2Config oauth2.Config
	txRunner     TxRunner
}

// NewService creates a new QBO OAuth service wired to the Intuit endpoint.
func NewService(
	queries *store.Queries,
	tokenService *tokens.Service,
	encryptor tokens.Encryptor,
	cfg config.Config,
	db *sql.DB,
) *Service {
	scopes := strings.Fields(cfg.QBOScopes)
	return &Service{
		stateStore:  queries,
		tokenLoader: tokenService,
		encryptor:   encryptor,
		oauth2Config: oauth2.Config{
			ClientID:     cfg.QBOClientID,
			ClientSecret: cfg.QBOClientSecret,
			RedirectURL:  cfg.QBORedirectURI,
			Endpoint:     intuitEndpoint,
			Scopes:       scopes,
		},
		txRunner: &dbTxRunner{db: db, encryptor: encryptor},
	}
}

// newServiceForTest is the test-only constructor that accepts direct
// interface implementations instead of wiring production dependencies.
func newServiceForTest(
	ep oauth2.Endpoint,
	clientID string,
	clientSecret string,
	redirectURL string,
	scopes []string,
	stateStore StateStore,
	tokenLoader TokenLoader,
	encryptor tokens.Encryptor,
	txRunner TxRunner,
) *Service {
	return &Service{
		stateStore:  stateStore,
		tokenLoader: tokenLoader,
		encryptor:   encryptor,
		oauth2Config: oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Endpoint:     ep,
			Scopes:       scopes,
		},
		txRunner: txRunner,
	}
}

// AuthURL generates an OAuth state, persists it to the database, and returns
// the Intuit authorization URL the user should be redirected to.
func (s *Service) AuthURL(ctx context.Context, tenantID uuid.UUID) (string, error) {
	stateBytes := make([]byte, 32)
	if _, err := rand.Read(stateBytes); err != nil {
		return "", fmt.Errorf("generate state: %w", err)
	}
	state := hex.EncodeToString(stateBytes)

	expiresAt := time.Now().Add(stateTTL)

	// Encrypt the tenant binding into the state so the callback can verify it.
	plaintext := []byte(tenantID.String() + "|" + fmt.Sprintf("%d", expiresAt.Unix()))
	encrypted, err := s.encryptor.Encrypt(plaintext)
	if err != nil {
		return "", fmt.Errorf("encrypt state: %w", err)
	}

	if _, err := s.stateStore.CreateOAuthState(ctx, store.CreateOAuthStateParams{
		ID:             uuid.New(),
		TenantID:       tenantID,
		StateChecksum:  state,
		EncryptedState: encrypted,
		ExpiresAt:      expiresAt,
	}); err != nil {
		return "", fmt.Errorf("persist state: %w", err)
	}

	return s.oauth2Config.AuthCodeURL(state), nil
}

// Exchange validates the OAuth state, exchanges the authorization code for
// tokens, creates the QBO connection, and persists the encrypted tokens.
// All DB writes are performed atomically in a single transaction.
func (s *Service) Exchange(
	ctx context.Context,
	code string,
	stateChecksum string,
	realmID string,
	companyName string,
) error {
	oauthState, token, err := s.prequelExchange(ctx, code, stateChecksum)
	if err != nil {
		return err
	}

	return s.txRunner.RunInTx(ctx, func(ctx context.Context, tx TxDataStore) error {
		return s.writeExchangeResults(ctx, tx, oauthState, token, realmID, companyName)
	})
}

// prequelExchange performs the read-only validation and Intuit API call
// that must happen before any DB writes.
func (s *Service) prequelExchange(
	ctx context.Context,
	code string,
	stateChecksum string,
) (store.QboOauthState, *oauth2.Token, error) {
	oauthState, err := s.stateStore.GetActiveOAuthStateByChecksum(ctx, stateChecksum)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return store.QboOauthState{}, nil, ErrStateNotFound
		}
		return store.QboOauthState{}, nil, fmt.Errorf("lookup state: %w", err)
	}

	plaintext, err := s.encryptor.Decrypt(oauthState.EncryptedState)
	if err != nil {
		return store.QboOauthState{}, nil, fmt.Errorf("decrypt state: %w", err)
	}

	parts := strings.SplitN(string(plaintext), "|", 2)
	if len(parts) != 2 {
		return store.QboOauthState{}, nil, fmt.Errorf("invalid state format")
	}

	stateTenantID, err := uuid.Parse(parts[0])
	if err != nil {
		return store.QboOauthState{}, nil, fmt.Errorf("parse state tenant: %w", err)
	}

	if stateTenantID != oauthState.TenantID {
		return store.QboOauthState{}, nil, ErrTenantMismatch
	}

	token, err := s.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return store.QboOauthState{}, nil, fmt.Errorf("exchange code: %w", err)
	}

	return oauthState, token, nil
}

// writeExchangeResults performs all DB writes for the exchange using the
// provided transactional data store so they commit or roll back atomically.
func (s *Service) writeExchangeResults(
	ctx context.Context,
	tx TxDataStore,
	oauthState store.QboOauthState,
	token *oauth2.Token,
	realmID string,
	companyName string,
) error {
	// Determine the connection ID — reuse existing or create new.
	connectionID := uuid.New()
	existingConn, err := tx.GetQBOConnectionByTenant(ctx, oauthState.TenantID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("lookup existing connection: %w", err)
		}

		if _, err := tx.CreateQBOConnection(ctx, store.CreateQBOConnectionParams{
			ID:       connectionID,
			TenantID: oauthState.TenantID,
			RealmID:  realmID,
			CompanyName: sql.NullString{
				String: companyName,
				Valid:  companyName != "",
			},
			State: "connected",
		}); err != nil {
			return fmt.Errorf("create connection: %w", err)
		}
	} else {
		connectionID = existingConn.ID

		if _, err := tx.UpdateQBOConnectionState(ctx, store.UpdateQBOConnectionStateParams{
			ID:    connectionID,
			State: "connected",
		}); err != nil {
			return fmt.Errorf("update connection state: %w", err)
		}

		if _, err := tx.UpdateQBOConnectionCompanyName(ctx, store.UpdateQBOConnectionCompanyNameParams{
			ID: connectionID,
			CompanyName: sql.NullString{
				String: companyName,
				Valid:  companyName != "",
			},
		}); err != nil {
			return fmt.Errorf("update connection company name: %w", err)
		}
	}

	// Encrypt and store the tokens.
	if err := tx.StoreTokens(
		ctx,
		connectionID,
		oauthState.TenantID,
		token.AccessToken,
		token.RefreshToken,
		token.Expiry,
		token.Expiry.Add(refreshTokenLifetime),
	); err != nil {
		return fmt.Errorf("store tokens: %w", err)
	}

	// Consume the state — single-use enforcement.
	if _, err := tx.ConsumeOAuthState(ctx, oauthState.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrStateConsumed
		}
		return fmt.Errorf("consume state: %w", err)
	}

	// Log a connected event for the audit trail.
	metadata, err := json.Marshal(map[string]string{
		"realm_id": realmID,
	})
	if err != nil {
		return fmt.Errorf("marshal event metadata: %w", err)
	}

	msg := "Connected to " + companyName
	if companyName == "" {
		msg = "Connected"
	}

	if _, err := tx.CreateQBOConnectionEvent(ctx, store.CreateQBOConnectionEventParams{
		ID:              uuid.New(),
		QboConnectionID: connectionID,
		TenantID:        oauthState.TenantID,
		EventType:       "connected",
		Message:         sql.NullString{String: msg, Valid: true},
		Metadata:        metadata,
	}); err != nil {
		return fmt.Errorf("create event: %w", err)
	}

	return nil
}

// NewClient returns an HTTP client authenticated with the stored QBO tokens.
// The caller is responsible for handling token expiry (the oauth2 transport
// sets the Authorization header automatically).
func (s *Service) NewClient(ctx context.Context, connectionID uuid.UUID) (*http.Client, error) {
	tok, err := s.tokenLoader.Load(ctx, connectionID)
	if err != nil {
		return nil, fmt.Errorf("load tokens: %w", err)
	}

	token := &oauth2.Token{
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
		Expiry:       tok.AccessExpiresAt,
	}

	return oauth2.NewClient(ctx, oauth2.StaticTokenSource(token)), nil
}
