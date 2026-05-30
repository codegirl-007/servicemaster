package tokens

import (
	"context"
	"errors"
	"fmt"
	"servicemaster/internal/store"
	"time"

	"database/sql"

	"github.com/google/uuid"
)

type Encryptor interface {
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}

// DataStore narrows *store.Queries to the methods the token service needs.
type DataStore interface {
	UpsertQBOConnectionTokens(context.Context, store.UpsertQBOConnectionTokensParams) (store.QboConnectionToken, error)
	GetQBOConnectionTokens(context.Context, uuid.UUID) (store.QboConnectionToken, error)
	ReplaceQBOConnectionTokensIfVersion(context.Context, store.ReplaceQBOConnectionTokensIfVersionParams) (store.QboConnectionToken, error)
}

type Service struct {
	store     DataStore
	encryptor Encryptor
}

type Tokens struct {
	ConnectionID     uuid.UUID
	AccessToken      string
	RefreshToken     string
	AccessExpiresAt  time.Time
	RefreshExpiresAt time.Time
	Version          int64
}

var ErrConcurrentRefresh = errors.New("qbo token refresh lost concurrent update")

func NewService(store DataStore, encryptor Encryptor) *Service {
	return &Service{
		store:     store,
		encryptor: encryptor,
	}
}

func (s *Service) Store(
	ctx context.Context,
	connectionID uuid.UUID,
	tenantID uuid.UUID,
	accessToken string,
	refreshToken string,
	accessExpiresAt time.Time,
	refreshExpiresAt time.Time) error {

	encryptedAccessToken, err := s.encryptor.Encrypt([]byte(accessToken))
	if err != nil {
		return fmt.Errorf("encrypt access token: %w", err)
	}

	encryptedRefreshToken, err := s.encryptor.Encrypt([]byte(refreshToken))
	if err != nil {
		return fmt.Errorf("encrypt refresh token: %w", err)
	}

	_, err = s.store.UpsertQBOConnectionTokens(ctx, store.UpsertQBOConnectionTokensParams{
		QboConnectionID:       connectionID,
		TenantID:              tenantID,
		EncryptedAccessToken:  encryptedAccessToken,
		EncryptedRefreshToken: encryptedRefreshToken,
		AccessTokenExpiresAt:  accessExpiresAt,
		RefreshTokenExpiresAt: sql.NullTime{
			Time:  refreshExpiresAt,
			Valid: true,
		},
		TokenType: "bearer",
		Scope:     sql.NullString{},
	})

	if err != nil {
		return fmt.Errorf("upsert qbo connection tokens: %w", err)
	}

	return nil
}

func (s *Service) Load(ctx context.Context, connectionID uuid.UUID) (Tokens, error) {
	row, err := s.store.GetQBOConnectionTokens(ctx, connectionID)
	if err != nil {
		return Tokens{}, fmt.Errorf("GetQBOConnectionTokens: %w", err)
	}

	accessToken, err := s.encryptor.Decrypt(row.EncryptedAccessToken)
	if err != nil {
		return Tokens{}, fmt.Errorf("decrypt access token: %w", err)
	}

	refreshToken, err := s.encryptor.Decrypt(row.EncryptedRefreshToken)
	if err != nil {
		return Tokens{}, fmt.Errorf("decrypt refresh token: %w", err)
	}

	tokens := Tokens{
		ConnectionID:     row.QboConnectionID,
		AccessToken:      string(accessToken),
		RefreshToken:     string(refreshToken),
		AccessExpiresAt:  row.AccessTokenExpiresAt,
		RefreshExpiresAt: row.RefreshTokenExpiresAt.Time,
		Version:          row.Version,
	}

	return tokens, nil
}

func (s *Service) ReplaceIfVersion(
	ctx context.Context,
	connectionID uuid.UUID,
	version int64,
	accessToken string,
	refreshToken string,
	accessExpiresAt time.Time,
	refreshExpiresAt time.Time,
) error {
	encryptedAccessToken, err := s.encryptor.Encrypt([]byte(accessToken))
	if err != nil {
		return fmt.Errorf("encrypt access token: %w", err)
	}

	encryptedRefreshToken, err := s.encryptor.Encrypt([]byte(refreshToken))
	if err != nil {
		return fmt.Errorf("encrypt refresh token: %w", err)
	}

	_, err = s.store.ReplaceQBOConnectionTokensIfVersion(ctx, store.ReplaceQBOConnectionTokensIfVersionParams{
		QboConnectionID:       connectionID,
		Version:               version,
		EncryptedAccessToken:  encryptedAccessToken,
		EncryptedRefreshToken: encryptedRefreshToken,
		AccessTokenExpiresAt:  accessExpiresAt,
		RefreshTokenExpiresAt: sql.NullTime{
			Time:  refreshExpiresAt,
			Valid: true,
		},
		TokenType: "bearer",
		Scope:     sql.NullString{},
	})

	if errors.Is(err, sql.ErrNoRows) {
		return ErrConcurrentRefresh
	}

	if err != nil {
		return fmt.Errorf("replace qbo connection tokens: %w", err)
	}

	return nil
}
