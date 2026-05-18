// Package types contains transport types for external integrations.
package types

import "time"

// Date represents a date-only value.
type Date struct {
	time.Time
}
