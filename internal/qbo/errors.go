package qbo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Sentinel errors classify failures for import orchestration without parsing bodies.
var (
	ErrRateLimited   = errors.New("qbo: rate limited")
	ErrUnauthorized  = errors.New("qbo: unauthorized")
	ErrNotFound      = errors.New("qbo: not found")
	ErrBadRequest    = errors.New("qbo: bad request")
	ErrServerFault   = errors.New("qbo: server fault")
	ErrEmptyResponse = errors.New("qbo: empty response body")
)

// Fault mirrors the JSON Fault object returned by QuickBooks on validation errors.
type Fault struct {
	Type  string       `json:"type"`
	Error []FaultError `json:"Error"`
}

// FaultError is a single error entry inside a Fault.
type FaultError struct {
	Message string `json:"Message"`
	Detail  string `json:"Detail"`
	Code    string `json:"code"`
	Element string `json:"element"`
}

// APIError is the primary failure surface from Client.Do and query helpers.
type APIError struct {
	StatusCode  int
	IntuitTID   string
	Fault       *Fault
	Body        []byte
	Retryable   bool
	Err         error
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
		Body:       append([]byte(nil), body...),
	}

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
		return nil
	}

	if len(envelope.Fault.Error) == 0 && envelope.Fault.Type == "" {
		return nil
	}

	return &envelope.Fault
}

// IsRetryable reports whether an error is safe to retry at the HTTP layer.
func IsRetryable(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Retryable
	}

	return false
}
