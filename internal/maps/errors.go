// Package maps defines provider-neutral types and interfaces for geocoding,
// routing, and places search.
package maps

import "errors"

var (
	// ErrNotFound is returned when a geocode, place, or route result is not found.
	ErrNotFound = errors.New("maps: not found")
	// ErrNoResults is returned when a query produces zero results.
	ErrNoResults = errors.New("maps: no results")
	// ErrInvalidInput is returned when the request parameters are invalid.
	ErrInvalidInput = errors.New("maps: invalid input")
	// ErrProvider is returned when the underlying provider returns an error.
	ErrProvider = errors.New("maps: provider error")
	// ErrOverQueryLimit is returned when the provider's query quota is exceeded.
	ErrOverQueryLimit = errors.New("maps: provider query limit exceeded")
)
