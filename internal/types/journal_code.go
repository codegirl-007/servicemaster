// Package types contains transport types for external integrations.
package types

import "time"

// JournalCodeType represents the documented journal code types.
type JournalCodeType string

const (
	JournalCodeTypeExpenses JournalCodeType = "Expenses"
	JournalCodeTypeSales    JournalCodeType = "Sales"
	JournalCodeTypeBank     JournalCodeType = "Bank"
	JournalCodeTypeNouveaux JournalCodeType = "Nouveaux"
	JournalCodeTypeWages    JournalCodeType = "Wages"
	JournalCodeTypeCash     JournalCodeType = "Cash"
	JournalCodeTypeOthers   JournalCodeType = "Others"
)

// JournalCodeResponse represents the QuickBooks journal code response envelope.
type JournalCodeResponse struct {
	JournalCode JournalCode `json:"JournalCode"`
	Time        time.Time   `json:"time"`
}

// JournalCode represents a QuickBooks journal code object.
type JournalCode struct {
	ID          string             `json:"Id"`
	Name        string             `json:"Name,omitempty"`
	SyncToken   string             `json:"SyncToken,omitempty"`
	Description string             `json:"Description,omitempty"`
	CustomField []JournalCodeField `json:"CustomField,omitempty"`
	Type        JournalCodeType    `json:"Type,omitempty"`
	MetaData    *MetaData          `json:"MetaData,omitempty"`
	Active      *bool              `json:"Active,omitempty"`
	Domain      string             `json:"domain,omitempty"`
	Sparse      *bool              `json:"sparse,omitempty"`
}

// CreateJournalCodeRequest represents the documented create journal code payload.
type CreateJournalCodeRequest struct {
	// Name is required.
	Name string          `json:"Name"`
	Type JournalCodeType `json:"Type,omitempty"`
}

// JournalCodeField represents a custom field on a journal code object.
type JournalCodeField struct {
	DefinitionID string `json:"DefinitionId,omitempty"`
	Type         string `json:"Type,omitempty"`
	StringValue  string `json:"StringValue,omitempty"`
	Name         string `json:"Name,omitempty"`
}
