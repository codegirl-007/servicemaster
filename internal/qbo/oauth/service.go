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

// dataStore narrows *store.Queries to the methods the OAuth service needs.
type dataStore interface {
	CreateOAuthState(context.Context, store.CreateOAuthStateParams) (store.QboOauthState, error)
	GetActiveOAuthStateByChecksum(context.Context, string) (store.QboOauthState, error)
	ConsumeOAuthState(context.Context, uuid.UUID) (uuid.UUID, error)
	GetQBOConnectionByTenant(context.Context, uuid.UUID) (store.QboConnection, error)
	CreateQBOConnection(context.Context, store.CreateQBOConnectionParams) (store.QboConnection, error)
	UpdateQBOConnectionState(context.Context, store.UpdateQBOConnectionStateParams) (store.QboConnection, error)
	UpdateQBOConnectionCompanyName(context.Context, store.UpdateQBOConnectionCompanyNameParams) (store.QboConnection, error)
	CreateQBOConnectionEvent(context.Context, store.CreateQBOConnectionEventParams) (store.QboConnectionEvent, error)
}

// tokenService narrows *tokens.Service to the methods the OAuth service needs.
type tokenService interface {
	Store(context.Context, uuid.UUID, uuid.UUID, string, string, time.Time, time.Time) error
	Load(context.Context, uuid.UUID) (tokens.Tokens, error)
}

// Service manages the QBO OAuth 2.0 connection flow.
type Service struct {
	store        dataStore
	tokenService tokenService
	encryptor    tokens.Encryptor
	oauth2Config oauth2.Config
	beginTx      func(context.Context, *sql.TxOptions) (*sql.Tx, error)
}

// NewService creates a new QBO OAuth service. Pass nil for db in tests.
func NewService(
	queries *store.Queries,
	tokenService *tokens.Service,
	encryptor tokens.Encryptor,
	cfg config.Config,
	db *sql.DB,
) *Service {
	s := &Service{
		store:        queries,
		tokenService: tokenService,
		encryptor:    encryptor,
		oauth2Config: oauth2.Config{
			ClientID:     cfg.QBOClientID,
			ClientSecret: cfg.QBOClientSecret,
			RedirectURL:  cfg.QBORedirectURI,
			Endpoint:     intuitEndpoint,
			Scopes:       strings.Fields(cfg.QBOScopes),
		},
	}
	if db != nil {
		s.beginTx = db.BeginTx
	}
	return s
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

	if _, err := s.store.CreateOAuthState(ctx, store.CreateOAuthStateParams{
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
	// Steps 1-3: validate state and exchange code with Intuit (no DB writes).
	oauthState, token, err := s.prequelExchange(ctx, code, stateChecksum)
	if err != nil {
		return err
	}

	// Steps 4-7: all DB writes inside a single transaction.
	if s.beginTx != nil {
		tx, err := s.beginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("begin transaction: %w", err)
		}
		defer tx.Rollback()

		qtx := store.New(tx)
		txTokenSvc := tokens.NewService(qtx, s.encryptor)

		if err := s.writeExchangeResults(ctx, qtx, txTokenSvc, oauthState, token, realmID, companyName); err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit transaction: %w", err)
		}
		return nil
	}

	// Test path — no transaction (in-memory fake store).
	return s.writeExchangeResults(ctx, s.store, s.tokenService, oauthState, token, realmID, companyName)
}

// prequelExchange performs the read-only validation and Intuit API call
// that must happen before any DB writes.
func (s *Service) prequelExchange(
	ctx context.Context,
	code string,
	stateChecksum string,
) (store.QboOauthState, *oauth2.Token, error) {
	oauthState, err := s.store.GetActiveOAuthStateByChecksum(ctx, stateChecksum)
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
// provided store and token service (these may be transactional or not).
func (s *Service) writeExchangeResults(
	ctx context.Context,
	ds dataStore,
	tokenSvc tokenService,
	oauthState store.QboOauthState,
	token *oauth2.Token,
	realmID string,
	companyName string,
) error {
	// Determine the connection ID — reuse existing or create new.
	connectionID := uuid.New()
	existingConn, err := ds.GetQBOConnectionByTenant(ctx, oauthState.TenantID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("lookup existing connection: %w", err)
		}

		if _, err := ds.CreateQBOConnection(ctx, store.CreateQBOConnectionParams{
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

		if _, err := ds.UpdateQBOConnectionState(ctx, store.UpdateQBOConnectionStateParams{
			ID:    connectionID,
			State: "connected",
		}); err != nil {
			return fmt.Errorf("update connection state: %w", err)
		}

		if _, err := ds.UpdateQBOConnectionCompanyName(ctx, store.UpdateQBOConnectionCompanyNameParams{
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
	if err := tokenSvc.Store(
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
	if _, err := ds.ConsumeOAuthState(ctx, oauthState.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrStateConsumed
		}
		return fmt.Errorf("consume state: %w", err)
	}

	// Log a connected event for the audit trail.
	metadata, _ := json.Marshal(map[string]string{
		"realm_id": realmID,
	})

	msg := "Connected to " + companyName
	if companyName == "" {
		msg = "Connected"
	}

	if _, err := ds.CreateQBOConnectionEvent(ctx, store.CreateQBOConnectionEventParams{
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
	tok, err := s.tokenService.Load(ctx, connectionID)
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
