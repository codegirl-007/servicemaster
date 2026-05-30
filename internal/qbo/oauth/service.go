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

// Service manages the QBO OAuth 2.0 connection flow.
type Service struct {
	queries      *store.Queries
	tokenService *tokens.Service
	encryptor    tokens.Encryptor
	oauth2Config oauth2.Config
}

// NewService creates a new QBO OAuth service.
func NewService(
	queries *store.Queries,
	tokenService *tokens.Service,
	encryptor tokens.Encryptor,
	cfg config.Config,
) *Service {
	return &Service{
		queries:      queries,
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

	if _, err := s.queries.CreateOAuthState(ctx, store.CreateOAuthStateParams{
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
func (s *Service) Exchange(
	ctx context.Context,
	code string,
	stateChecksum string,
	realmID string,
	companyName string,
) error {
	oauthState, err := s.queries.GetActiveOAuthStateByChecksum(ctx, stateChecksum)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrStateNotFound
		}
		return fmt.Errorf("lookup state: %w", err)
	}

	// Decrypt the state to verify the tenant binding.
	plaintext, err := s.encryptor.Decrypt(oauthState.EncryptedState)
	if err != nil {
		return fmt.Errorf("decrypt state: %w", err)
	}

	parts := strings.SplitN(string(plaintext), "|", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid state format")
	}

	stateTenantID, err := uuid.Parse(parts[0])
	if err != nil {
		return fmt.Errorf("parse state tenant: %w", err)
	}

	if stateTenantID != oauthState.TenantID {
		return ErrTenantMismatch
	}

	// Exchange the authorization code for an OAuth token.
	token, err := s.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return fmt.Errorf("exchange code: %w", err)
	}

	// Determine the connection ID — reuse existing or create new.
	connectionID := uuid.New()
	existingConn, err := s.queries.GetQBOConnectionByTenant(ctx, oauthState.TenantID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("lookup existing connection: %w", err)
		}

		// No existing connection — create one.
		if _, err := s.queries.CreateQBOConnection(ctx, store.CreateQBOConnectionParams{
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
		// Existing connection found — update it.
		connectionID = existingConn.ID

		if _, err := s.queries.UpdateQBOConnectionState(ctx, store.UpdateQBOConnectionStateParams{
			ID:    connectionID,
			State: "connected",
		}); err != nil {
			return fmt.Errorf("update connection state: %w", err)
		}

		if _, err := s.queries.UpdateQBOConnectionCompanyName(ctx, store.UpdateQBOConnectionCompanyNameParams{
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
	if err := s.tokenService.Store(
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

	// Consume the state — single-use enforcement (last irreversible step).
	if _, err := s.queries.ConsumeOAuthState(ctx, oauthState.ID); err != nil {
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

	if _, err := s.queries.CreateQBOConnectionEvent(ctx, store.CreateQBOConnectionEventParams{
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
