package client

import "time"

// RetryPolicy defines retry behavior for transient failures.
type RetryPolicy struct {
	MaxAttempts   int
	BaseDelay     time.Duration
	MaxDelay      time.Duration
	JitterPercent float64
}

// DefaultRetryPolicy provides conservative defaults for QBO.
var DefaultRetryPolicy = RetryPolicy{
	MaxAttempts:   3,
	BaseDelay:     500 * time.Millisecond,
	MaxDelay:      30 * time.Second,
	JitterPercent: 0.2,
}
