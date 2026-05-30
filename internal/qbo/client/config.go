// Package client provides a minimal QuickBooks Online HTTP client scaffold.
package client

import "net/http"

// TokenSource supplies bearer tokens for authenticated QBO requests.
type TokenSource interface {
	AccessToken() (string, error)
}

// Config holds construction parameters for the QBO client.
type Config struct {
	// BaseURL is the QBO company base URL (sandbox or production).
	BaseURL string
	// RealmID identifies the QuickBooks company.
	RealmID string
	// MinorVersion sets the optional Intuit minorversion query parameter.
	MinorVersion int
	// HTTPClient is the transport; if nil, http.DefaultClient is used.
	HTTPClient *http.Client
	// TokenSource provides bearer tokens for requests.
	TokenSource TokenSource
	// RetryPolicy controls retry behavior for transient errors.
	RetryPolicy RetryPolicy
	// UserAgent is sent on outbound requests.
	UserAgent string
}
