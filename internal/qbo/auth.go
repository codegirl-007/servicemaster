// Package qbo is the HTTP transport layer for QuickBooks Online API v3.
//
// It deliberately does not import internal/types or internal/store so import
// jobs can depend on it without cycles. Entity DTOs stay in types; token
// persistence stays in store; this package only knows about HTTP, pagination,
// retries, and Intuit error shapes.
//
// See DESIGN.md for the full spike rationale and merge path.
package qbo

import "context"

// TokenSource supplies OAuth access tokens for QuickBooks API calls.
//
// We use an interface instead of passing a string because tokens expire mid-import.
// A long-running customer sync cannot hold one access token for its whole lifetime.
//
// Production flow (not implemented here):
//  1. Read encrypted row from qbo_connection_tokens.
//  2. If access token is still valid, decrypt and return.
//  3. If expired, refresh against Intuit, encrypt new tokens, bump version.
//  4. On refresh failure, mark connection reconnect_required and return ErrUnauthorized.
//
// Keeping refresh out of Client.Do is intentional: transport retries must not
// silently re-auth, or we get duplicate refresh calls, lost audit events, and
// races on the token row when multiple import workers run concurrently.
type TokenSource interface {
	AccessToken(ctx context.Context) (string, error)
}

// StaticTokenSource returns a fixed token.
//
// Used in httptest and local spikes. Never use in production; it hides expiry
// and refresh behavior that import correctness depends on.
type StaticTokenSource struct {
	Token string
}

// AccessToken implements TokenSource.
func (s StaticTokenSource) AccessToken(context.Context) (string, error) {
	return s.Token, nil
}
