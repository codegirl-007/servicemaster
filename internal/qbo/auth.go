package qbo

import "context"

// TokenSource supplies OAuth access tokens for QuickBooks API calls.
//
// Production implementations should load encrypted tokens from the store,
// refresh when near expiry, and persist updates with optimistic locking on
// qbo_connection_tokens.version. This package intentionally does not perform
// refresh so transport retries never fight token-service transactions.
type TokenSource interface {
	AccessToken(ctx context.Context) (string, error)
}

// StaticTokenSource returns a fixed token. Useful for tests and local spikes.
type StaticTokenSource struct {
	Token string
}

// AccessToken implements TokenSource.
func (s StaticTokenSource) AccessToken(context.Context) (string, error) {
	return s.Token, nil
}
