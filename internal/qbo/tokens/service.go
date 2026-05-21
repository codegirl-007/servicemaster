package tokens

import (
	"context"
	"fmt"
	"servicemaster/internal/store"
	"time"

	"github.com/google/uuid"
	"database/sql"
)

type Encryptor interface {
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}

type Service struct {
	queries   *store.Queries
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

func NewService(queries *store.Queries, encryptor Encryptor) *Service {
	return &Service{
		queries:   queries,
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

	_, err = s.queries.UpsertQBOConnectionTokens(ctx, store.UpsertQBOConnectionTokensParams{
		QboConnectionID: connectionID,
		TenantID: tenantID,
		EncryptedAccessToken: encryptedAccessToken,
		EncryptedRefreshToken: encryptedRefreshToken,
		AccessTokenExpiresAt: accessExpiresAt,
		RefreshTokenExpiresAt: sql.NullTime{
			Time: refreshExpiresAt,
			Valid: true,
		},
	})

	if err != nil {
		return fmt.Errorf("upsert qbo connection tokens: %w", err)
	}

	return nil
}

func (s *Service) Load(ctx context.Context, connectionID uuid.UUID) (Tokens, error) {
	row, err := s.queries.GetQBOConnectionTokens(ctx, connectionID);
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
		ConnectionID:    row.QboConnectionID,
		AccessToken:     string(accessToken),
		RefreshToken:    string(refreshToken),
		AccessExpiresAt: row.AccessTokenExpiresAt,
		RefreshExpiresAt: row.RefreshTokenExpiresAt.Time,
		Version:         row.Version,
	}

	return tokens, nil
}

