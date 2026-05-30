# QuickBooks API client foundation (demonstration)

This package is a **spike** for COD-29. It is not wired into import jobs, OAuth flows, or token refresh yet. The goal is to show how transport concerns can stay separate from import orchestration while still meeting the acceptance criteria.

## Problem

`internal/types` holds QuickBooks entity DTOs, and the database stores encrypted tokens on `qbo_connection_tokens`, but nothing yet performs authenticated HTTP calls with consistent pagination, retries, and error surfaces.

## Recommendation

Introduce a thin `internal/qbo` HTTP client that:

1. Accepts credentials through a `TokenSource` interface (implemented by an adapter over `internal/qbo/tokens.Service` — see [PR #97](https://github.com/codegirl-007/servicemaster/pull/97)).
2. Owns request construction, Intuit headers, retry boundaries, and response decoding.
3. Exposes pagination helpers that append `STARTPOSITION` / `MAXRESULTS` to SQL queries.
4. Maps HTTP and Fault payloads into typed errors import jobs can branch on.

Import orchestration (batch creation, staging writes, review tasks) should depend on `*qbo.Client` methods like `QueryPages` and decode into `internal/types` at the job layer—not inside the client.

## Package layout

| File | Responsibility |
|------|----------------|
| `auth.go` | `TokenSource`, static token for tests |
| `errors.go` | `APIError`, fault parsing, sentinel errors |
| `retry.go` | Retry policy and backoff boundaries |
| `pagination.go` | Query page decoding and `PageIterator` |
| `client.go` | HTTP client, `Query`, `QueryPages` |
| `client_test.go` | `httptest` coverage for auth, retry, pagination, errors |
| `tokens/` (PR #97, WIP) | Encrypted `Store` / `Load` for `qbo_connection_tokens` — not HTTP |

## Relationship to PR #97 (token service, WIP)

[PR #97](https://github.com/codegirl-007/servicemaster/pull/97) adds `internal/qbo/tokens.Service` with:

- `Store` — encrypt and upsert access/refresh tokens after OAuth callback
- `Load` — decrypt and return `tokens.Tokens` (includes `Version` for optimistic locking)

**We do not modify PR #97.** This spike (#99) assumes it merges first (or in parallel) and fills the gap *above* persistence:

```
  OAuth callback ──► tokens.Service.Store
                            │
  import job ──► TokenSource adapter ──► tokens.Service.Load (+ refresh when added)
                            │
                     qbo.Client.Do / QueryPages
```

### Division of responsibility

| Concern | Owner | PR |
|---------|--------|-----|
| Encrypt tokens, read/write `qbo_connection_tokens` | `tokens.Service` | #97 |
| Implement `qbo.TokenSource` for `Client` | New adapter (e.g. `tokens.Source`) | After #97 |
| Intuit OAuth refresh HTTP, bump `version` | Same adapter + connection events | After #97 |
| HTTP query transport, pagination, retries | `qbo.Client` | #99 (this spike) |

PR #97 today is **persistence only** — no `AccessToken`, no expiry check, no refresh. That is correct: refresh must not live in `Store`/`Load` alone, or import code would bypass the `TokenSource` boundary and call `Load` directly with stale tokens.

### Planned `TokenSource` adapter (not in PR #97)

After #97 merges, a small type in `internal/qbo/tokens` (or `internal/qbo`) should implement `qbo.TokenSource`:

```go
// Pseudocode — follow-up PR after #97, not part of either open spike.
type Source struct {
    svc          *tokens.Service
    connectionID uuid.UUID
    refresher    OAuthRefresher // Intuit token endpoint; injected for tests
}

func (s *Source) AccessToken(ctx context.Context) (string, error) {
    tok, err := s.svc.Load(ctx, s.connectionID)
    if err != nil {
        return "", err
    }
    if time.Now().Before(tok.AccessExpiresAt.Add(-refreshSkew)) {
        return tok.AccessToken, nil
    }
    // Use tok.RefreshToken via refresher, then svc.Store with version check.
    // On failure: return ErrUnauthorized; connection worker sets reconnect_required.
}
```

`qbo.Client` stays unchanged: it only calls `TokenSource.AccessToken` once per `Do` (see `client.go` comments).

### Merge order suggestion

1. **#97** — token `Store`/`Load` (WIP: fix compile nits, add tests, no scope creep).
2. **#99** — client foundation (this spike; do not merge as-is).
3. **Follow-up** — `tokens.Source` implementing `TokenSource` + OAuth refresh.
4. **Follow-up** — first import job wiring `Client` + `Source` + staging.

### Gaps to close after both land

- `GetQBOConnectionTokensForUpdate` + `version` check on refresh write (PR #97 uses `Load` only today).
- Proactive refresh before expiry (skew window), not only reactive 401.
- Wire `tokens.Encryptor` to `internal/platform/crypto.Encryptor` (same shape, duplicate interface in #97 for now).

## Auth-aware client shape

```go
type Config struct {
    BaseURL      string        // sandbox or production company base
    RealmID      string
    MinorVersion int
    HTTPClient   *http.Client
    TokenSource  TokenSource   // supplies Bearer access tokens
    RetryPolicy  RetryPolicy
}
```

**Why `TokenSource` instead of passing a raw token string?**

- Access tokens expire; refresh belongs in a dedicated service that updates `qbo_connection_tokens` and bumps `version` for optimistic locking.
- The client stays stateless and safe to share across goroutines.
- Tests inject `StaticTokenSource`; production injects a `TokenSource` adapter over `tokens.Service` (PR #97).

**401 handling boundary:** The client does **not** auto-refresh on 401. It returns `ErrUnauthorized` (or an `APIError` wrapping it). The token service / connection state machine decides whether to refresh, mark `reconnect_required`, or fail the import batch. Mixing refresh into the transport layer creates inconsistent retry behavior and hides audit-relevant connection transitions.

## Pagination handling

QuickBooks query pagination uses SQL clauses:

```sql
SELECT * FROM Customer STARTPOSITION 1 MAXRESULTS 100
```

Response metadata lives under `QueryResponse`:

- `startPosition`, `maxResults`, `totalCount` (count in **this page**, not global total)
- One or more entity array keys (`Customer`, `Invoice`, …)

`PageIterator` appends pagination clauses to a base query and stops when `totalCount < pageSize`.

**Import jobs** should call `QueryPages` and decode `Page.Entities` into `[]types.Customer` (etc.) in the job package. Keeping entity typing out of `qbo` avoids import cycles and prevents the client from knowing about every entity we might query.

To get a global entity count, run `SELECT COUNT(*) FROM Customer` as a separate query (documented QBO behavior).

## Retry boundaries

| Condition | Retry? | Rationale |
|-----------|--------|-----------|
| Network / context errors | No (caller retries job) | Idempotency lives at import batch level |
| HTTP 429 | Yes | Rate limit; honor `Retry-After` when present |
| HTTP 500, 502, 503, 504 | Yes | Transient upstream failures |
| HTTP 400, 403, 404 | No | Client/query bug; fix before retry |
| HTTP 401 | No | Token refresh is explicit, not blind retry |
| HTTP 409 / validation Fault | No | Data conflict; surface to review |

Default: 3 attempts, exponential backoff from 500ms with 20% jitter, capped at 30s.

**Rate limits:** Intuit documents per-realm limits (commonly cited around 500 requests/minute). Import jobs should prefer paginated bulk queries over per-entity GETs and backoff cooperatively on 429 rather than hammering.

## Error classes and failure surfaces

| Type | When | Import layer action |
|------|------|---------------------|
| `ErrRateLimited` | 429 after retries exhausted | Pause job, reschedule with delay |
| `ErrUnauthorized` | 401 | Trigger token refresh or `reconnect_required` |
| `ErrNotFound` | 404 | Skip or mark entity missing |
| `APIError` with `Fault` | 400 + JSON Fault | Log validation detail, create review task |
| Wrapped network errors | Dial/timeouts | Fail batch with retryable flag |

`APIError` carries `StatusCode`, `IntuitTID` (from `intuit_tid` header), optional parsed `Fault`, and raw body for support/debug.

## What stays outside this package

- OAuth authorization URL and callback handling
- Token encryption/decryption and DB persistence (**PR #97** `tokens.Service`; not `Client`)
- OAuth refresh and `TokenSource` adapter (**follow-up** after #97)
- Import batch lifecycle, staging tables, review queues
- Mapping `types.Customer` → staging models
- CDC / change-data-capture cursor management

## Failure modes

1. **Partial page failure mid-import:** Job should persist last successful `startPosition` on the batch metadata so resume does not re-fetch from 1.
2. **Inconsistent retry if 401 is retried blindly:** Could amplify lock contention on token rows; avoided by policy above.
3. **Assuming `totalCount` is global:** Leads to premature stop or infinite loops; documented in pagination code.
4. **XML Fault on some endpoints:** This spike parses JSON Fault only; expand `decodeFault` when we hit XML error bodies in the wild.

## Lean version (this PR)

- `Client` + `TokenSource` + `Query` / `QueryPages`
- Retry policy with tests
- Typed errors from status + JSON Fault
- `httptest` demonstrations

## Later version

- `tokens.Source` implementing `TokenSource` (Load via PR #97 + refresh + `version` checks)
- MinorVersion from tenant preferences
- Metrics (request count, 429 rate, latency)
- Optional request logging with redacted tokens
- XML fault fallback
- Per-entity typed query helpers only if repetition justifies them

## Example usage (future import job)

```go
tokenSvc := tokens.NewService(queries, encryptor)
tokenSrc := tokens.NewSource(tokenSvc, connectionID, refresher) // follow-up after PR #97
client, err := qbo.NewClient(qbo.Config{
    BaseURL:     qbo.SandboxBaseURL(realmID),
    RealmID:     realmID,
    TokenSource: tokenSrc,
})
iter, err := client.QueryPages(ctx, "SELECT * FROM Customer", 100)
for iter.Next(ctx) {
    page, err := iter.Page()
    var customers []types.Customer
    json.Unmarshal(page.Entities["Customer"], &customers)
    // stage customers...
}
```
