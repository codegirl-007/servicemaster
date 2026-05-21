package qbo

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// RetryPolicy defines HTTP-layer retry behavior only.
//
// This is separate from import-batch retries on purpose:
//   - HTTP retries are fast and idempotent for GET /query.
//   - Batch retries resume from a saved startPosition and may span minutes.
//
// Mixing both into one counter causes double-retries (transport + job) and
// makes failures hard to reason about in audit logs.
type RetryPolicy struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
}

// DefaultRetryPolicy is tuned for bulk import: a few quick retries, then let
// the job scheduler back off cooperatively on sustained 429s.
var DefaultRetryPolicy = RetryPolicy{
	MaxAttempts: 3,
	BaseDelay:   500 * time.Millisecond,
	MaxDelay:    30 * time.Second,
}

func (p RetryPolicy) normalized() RetryPolicy {
	if p.MaxAttempts <= 0 {
		p.MaxAttempts = DefaultRetryPolicy.MaxAttempts
	}
	if p.BaseDelay <= 0 {
		p.BaseDelay = DefaultRetryPolicy.BaseDelay
	}
	if p.MaxDelay <= 0 {
		p.MaxDelay = DefaultRetryPolicy.MaxDelay
	}

	return p
}

// shouldRetry encodes the retry boundary table from DESIGN.md.
//
// We intentionally exclude 401: the access token may be expired but refresh is
// a state transition (DB write + connection event), not "send the same Bearer
// token again and hope."
func (p RetryPolicy) shouldRetry(statusCode int, attempt int) bool {
	if attempt >= p.normalized().MaxAttempts {
		return false
	}

	switch statusCode {
	case http.StatusTooManyRequests:
		return true
	case http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return true
	default:
		return false
	}
}

func (p RetryPolicy) delay(attempt int, retryAfter string) time.Duration {
	p = p.normalized()

	// Prefer Intuit's Retry-After when it is a positive integer of seconds.
	// Ignore zero/invalid values and fall back to exponential backoff so we
	// do not busy-loop when the header is "0".
	if retryAfter != "" {
		if seconds, err := strconv.Atoi(retryAfter); err == nil && seconds > 0 {
			delay := time.Duration(seconds) * time.Second
			if delay <= p.MaxDelay {
				return delay
			}

			return p.MaxDelay
		}
	}

	// attempt is 1-based in Do's loop; shift gives 500ms, 1s, 2s before cap.
	delay := p.BaseDelay << (attempt - 1)
	if delay > p.MaxDelay {
		delay = p.MaxDelay
	}

	// Jitter spreads retries when multiple tenants import concurrently and hit
	// 429 together (thundering herd on our side, not just Intuit's).
	jitterWindow := delay / 5
	if jitterWindow <= 0 {
		return delay
	}

	jitter := time.Duration(rand.Int63n(int64(jitterWindow)))
	return delay + jitter
}

func (p RetryPolicy) wait(ctx context.Context, attempt int, retryAfter string) error {
	timer := time.NewTimer(p.delay(attempt, retryAfter))
	defer timer.Stop()

	select {
	case <-ctx.Done():
		// Respect import job cancellation during backoff.
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
