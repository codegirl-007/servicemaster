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
// Production persistence lives in internal/qbo/tokens (PR #97, WIP):
//   - tokens.Service.Store after OAuth callback
//   - tokens.Service.Load to decrypt tokens.Tokens (includes Version)
//
// A separate type should implement TokenSource on top of that service (follow-up
// PR, not PR #97): check AccessExpiresAt, refresh via Intuit when needed, Store
// with version guard, return ErrUnauthorized if refresh fails so the connection
// worker can set reconnect_required.
//
// qbo.Client must not call tokens.Service directly — only TokenSource — so HTTP
// retries never skip the refresh boundary.
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
