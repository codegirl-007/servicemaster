package qbo

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// RetryPolicy defines which HTTP statuses are retried and how long to wait.
type RetryPolicy struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
}

// DefaultRetryPolicy is the recommended starting point for import workloads.
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

	if retryAfter != "" {
		if seconds, err := strconv.Atoi(retryAfter); err == nil && seconds > 0 {
			delay := time.Duration(seconds) * time.Second
			if delay <= p.MaxDelay {
				return delay
			}

			return p.MaxDelay
		}
	}

	delay := p.BaseDelay << (attempt - 1)
	if delay > p.MaxDelay {
		delay = p.MaxDelay
	}

	// Add up to 20% jitter to reduce synchronized retries across workers.
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
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
