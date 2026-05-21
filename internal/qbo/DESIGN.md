# QuickBooks API client foundation (demonstration)

This package is a **spike** for COD-29. It is not wired into import jobs, OAuth flows, or token refresh yet. The goal is to show how transport concerns can stay separate from import orchestration while still meeting the acceptance criteria.

## Problem

`internal/types` holds QuickBooks entity DTOs, and the database stores encrypted tokens on `qbo_connection_tokens`, but nothing yet performs authenticated HTTP calls with consistent pagination, retries, and error surfaces.

## Recommendation

Introduce a thin `internal/qbo` HTTP client that:

1. Accepts credentials through a `TokenSource` interface (implemented later by a token service using `store` + `crypto`).
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
- Tests inject `StaticTokenSource`; production injects a store-backed implementation.

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
- Token encryption/decryption and DB persistence
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

- Store-backed `TokenSource` with refresh and `version` checks
- MinorVersion from tenant preferences
- Metrics (request count, 429 rate, latency)
- Optional request logging with redacted tokens
- XML fault fallback
- Per-entity typed query helpers only if repetition justifies them

## Example usage (future import job)

```go
tokens := token.NewStoreSource(store, encryptor, connectionID)
client, err := qbo.NewClient(qbo.Config{
    BaseURL:     qbo.SandboxBaseURL(realmID),
    RealmID:     realmID,
    TokenSource: tokens,
})
iter, err := client.QueryPages(ctx, "SELECT * FROM Customer", 100)
for iter.Next(ctx) {
    page, err := iter.Page()
    var customers []types.Customer
    json.Unmarshal(page.Entities["Customer"], &customers)
    // stage customers...
}
```
