// Package types contains transport types for external integrations.
package types

import "time"

// CompanyCurrencyResponse represents the QuickBooks company currency response envelope.
type CompanyCurrencyResponse struct {
	CompanyCurrency CompanyCurrency `json:"CompanyCurrency"`
	Time            time.Time       `json:"time"`
}

// CompanyCurrency represents a QuickBooks company currency object.
type CompanyCurrency struct {
	ID          string                 `json:"Id"`
	Code        string                 `json:"Code,omitempty"`
	SyncToken   string                 `json:"SyncToken,omitempty"`
	Name        string                 `json:"Name,omitempty"`
	CustomField []CompanyCurrencyField `json:"CustomField,omitempty"`
	Active      *bool                  `json:"Active,omitempty"`
	MetaData    *MetaData              `json:"MetaData,omitempty"`
	Domain      string                 `json:"domain,omitempty"`
	Sparse      *bool                  `json:"sparse,omitempty"`
}

// CreateCompanyCurrencyRequest represents the documented create company currency payload.
type CreateCompanyCurrencyRequest struct {
	// Code is required.
	Code string `json:"Code"`
}

// CompanyCurrencyField represents a custom field on a company currency object.
type CompanyCurrencyField struct {
	DefinitionID string `json:"DefinitionId,omitempty"`
	Type         string `json:"Type,omitempty"`
	StringValue  string `json:"StringValue,omitempty"`
	Name         string `json:"Name,omitempty"`
}
