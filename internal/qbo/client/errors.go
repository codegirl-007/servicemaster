package client

import "fmt"

// ErrUnauthorized indicates the request was not authorized (e.g., token expired).
var ErrUnauthorized = fmt.Errorf("unauthorized")

// ErrRateLimited indicates the request hit a QBO rate limit.
var ErrRateLimited = fmt.Errorf("rate limited")

// ErrNotFound indicates the requested resource was not found.
var ErrNotFound = fmt.Errorf("not found")

// APIError wraps a QBO HTTP response with useful metadata.
type APIError struct {
	StatusCode int    // HTTP status code returned by QBO.
	IntuitTID  string // Intuit request ID header for tracing.
	Fault      []byte // Raw fault payload if provided.
	Body       []byte // Raw response body for debugging.
	Err        error  // Underlying error, if any.
}

func (e *APIError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return fmt.Sprintf("qbo api error status=%d", e.StatusCode)
}

func (e *APIError) Unwrap() error {
	return e.Err
}
