// Package types contains transport types for external integrations.
package types

import "time"

// TermType represents the documented sales term types.
type TermType string

const (
	TermTypeStandard   TermType = "STANDARD"
	TermTypeDateDriven TermType = "DATE_DRIVEN"
)

// TermResponse represents the QuickBooks term response envelope.
type TermResponse struct {
	Term Term      `json:"Term"`
	Time time.Time `json:"time"`
}

// Term represents a QuickBooks term object.
type Term struct {
	ID                 string    `json:"Id"`
	Name               string    `json:"Name,omitempty"`
	SyncToken          string    `json:"SyncToken,omitempty"`
	DayOfMonthDue      int       `json:"DayOfMonthDue,omitempty"`
	DiscountDayOfMonth int       `json:"DiscountDayOfMonth,omitempty"`
	DueNextMonthDays   int       `json:"DueNextMonthDays,omitempty"`
	DueDays            int       `json:"DueDays,omitempty"`
	DiscountPercent    float64   `json:"DiscountPercent,omitempty"`
	DiscountDays       int       `json:"DiscountDays,omitempty"`
	Active             *bool     `json:"Active,omitempty"`
	Type               TermType  `json:"Type,omitempty"`
	MetaData           *MetaData `json:"MetaData,omitempty"`
	Domain             string    `json:"domain,omitempty"`
	Sparse             *bool     `json:"sparse,omitempty"`
}

// CreateTermRequest represents the documented create term payload.
type CreateTermRequest struct {
	// Name is required.
	Name               string  `json:"Name"`
	DayOfMonthDue      int     `json:"DayOfMonthDue,omitempty"`
	DiscountDayOfMonth int     `json:"DiscountDayOfMonth,omitempty"`
	DueNextMonthDays   int     `json:"DueNextMonthDays,omitempty"`
	DueDays            int     `json:"DueDays,omitempty"`
	DiscountPercent    float64 `json:"DiscountPercent,omitempty"`
	DiscountDays       int     `json:"DiscountDays,omitempty"`
}
