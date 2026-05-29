// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// ExchangeRateResponse represents the QuickBooks exchange rate response envelope.
type ExchangeRateResponse struct {
	ExchangeRate ExchangeRate `json:"ExchangeRate"`
	Time         time.Time    `json:"time"`
}

// ExchangeRate represents a QuickBooks exchange rate object.
type ExchangeRate struct {
	SyncToken          string        `json:"SyncToken,omitempty"`
	AsOfDate           *Date         `json:"AsOfDate,omitempty"`
	SourceCurrencyCode string        `json:"SourceCurrencyCode,omitempty"`
	Rate               float64       `json:"Rate,omitempty"`
	CustomField        []CustomField `json:"CustomField,omitempty"`
	TargetCurrencyCode string        `json:"TargetCurrencyCode,omitempty"`
	MetaData           *MetaData     `json:"MetaData,omitempty"`
	Domain             string        `json:"domain,omitempty"`
	Sparse             *bool         `json:"sparse,omitempty"`
}

// UpdateExchangeRateRequest represents the documented update exchange rate payload.
type UpdateExchangeRateRequest struct {
	// SyncToken is required for update.
	SyncToken string `json:"SyncToken"`
	// AsOfDate is required for update.
	AsOfDate *Date `json:"AsOfDate"`
	// SourceCurrencyCode is required for update.
	SourceCurrencyCode string `json:"SourceCurrencyCode"`
	// Rate is required for update.
	Rate               float64 `json:"Rate"`
	TargetCurrencyCode string  `json:"TargetCurrencyCode,omitempty"`
}
