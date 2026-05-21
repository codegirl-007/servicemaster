package qbo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Sentinel errors let import orchestration branch with errors.Is without
// inspecting HTTP bodies or Fault JSON. They are the stable contract between
// transport and jobs.
//
// Example: on ErrUnauthorized, the job pauses the batch and emits a
// reconnect_required connection event instead of retrying the same stale token.
var (
	ErrRateLimited   = errors.New("qbo: rate limited")
	ErrUnauthorized  = errors.New("qbo: unauthorized")
	ErrNotFound      = errors.New("qbo: not found")
	ErrBadRequest    = errors.New("qbo: bad request")
	ErrServerFault   = errors.New("qbo: server fault")
	ErrEmptyResponse = errors.New("qbo: empty response body")
)

// Fault mirrors the JSON Fault object QuickBooks returns on validation failures.
// Some endpoints still return XML faults; this spike only parses JSON. When we
// hit XML in production, extend decodeFault rather than pushing XML into jobs.
type Fault struct {
	Type  string       `json:"type"`
	Error []FaultError `json:"Error"`
}

// FaultError is one element inside Fault.Error (Intuit uses a JSON array even
// for a single validation problem).
type FaultError struct {
	Message string `json:"Message"`
	Detail  string `json:"Detail"`
	Code    string `json:"code"`
	Element string `json:"element"`
}

// APIError is the rich failure surface from Client.Do and query helpers.
//
// Jobs should prefer errors.Is(err, ErrUnauthorized) for control flow and use
// APIError fields for logging/support (IntuitTID, Fault, raw Body).
//
// Retryable is set here—not in RetryPolicy—so callers that catch an error after
// retries are exhausted still know whether the batch itself should be retried
// later (429/5xx) vs failed permanently (400/401).
type APIError struct {
	StatusCode int
	IntuitTID  string
	Fault      *Fault
	Body       []byte
	Retryable  bool
	Err        error
}

func (e *APIError) Error() string {
	if e == nil {
		return "qbo: unknown api error"
	}

	parts := []string{fmt.Sprintf("qbo api error: status=%d", e.StatusCode)}
	if e.IntuitTID != "" {
		parts = append(parts, "intuit_tid="+e.IntuitTID)
	}
	if e.Fault != nil && len(e.Fault.Error) > 0 {
		parts = append(parts, e.Fault.Error[0].Message)
	}
	if e.Err != nil {
		parts = append(parts, e.Err.Error())
	}

	return strings.Join(parts, "; ")
}

func (e *APIError) Unwrap() error {
	return e.Err
}

func newAPIError(status int, intuitTID string, body []byte) *APIError {
	fault := decodeFault(body)

	apiErr := &APIError{
		StatusCode: status,
		IntuitTID:  intuitTID,
		Fault:      fault,
		// Copy body so callers can log it after the HTTP response is gone.
		Body: append([]byte(nil), body...),
	}

	// Map status codes to sentinel errors and retry hints.
	//
	// Notable omissions:
	//   - 401: never Retryable; refresh is explicit (see TokenSource doc).
	//   - 403/404: permanent for this request; job decides skip vs fail.
	//   - 400: usually bad query or validation; retry without fixing is wasteful.
	switch status {
	case http.StatusTooManyRequests:
		apiErr.Retryable = true
		apiErr.Err = ErrRateLimited
	case http.StatusUnauthorized:
		apiErr.Err = ErrUnauthorized
	case http.StatusNotFound:
		apiErr.Err = ErrNotFound
	case http.StatusBadRequest:
		apiErr.Err = ErrBadRequest
	default:
		if status >= http.StatusInternalServerError {
			apiErr.Retryable = true
			apiErr.Err = ErrServerFault
		}
	}

	return apiErr
}

func decodeFault(body []byte) *Fault {
	if len(body) == 0 {
		return nil
	}

	var envelope struct {
		Fault Fault `json:"Fault"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		// Intuit sometimes returns non-JSON bodies; treat as no structured fault.
		return nil
	}

	if len(envelope.Fault.Error) == 0 && envelope.Fault.Type == "" {
		return nil
	}

	return &envelope.Fault
}

// IsRetryable reports whether the error is safe to retry at the HTTP layer.
//
// Import batches have their own idempotency boundary; this helper is for code
// still inside the client stack or deciding whether to reschedule a whole job.
func IsRetryable(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Retryable
	}

	return false
}
