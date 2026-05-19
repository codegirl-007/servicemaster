// Package types contains transport types for external integrations.
package types

import "time"

// TaxAgencyConfig represents the documented tax agency configuration values.
type TaxAgencyConfig string

const (
	TaxAgencyConfigUserDefined     TaxAgencyConfig = "USER_DEFINED"
	TaxAgencyConfigSystemGenerated TaxAgencyConfig = "SYSTEM_GENERATED"
)

// TaxAgencyResponse represents the QuickBooks tax agency response envelope.
type TaxAgencyResponse struct {
	TaxAgency TaxAgency `json:"TaxAgency"`
	Time      time.Time `json:"time"`
}

// TaxAgencyQueryResponse represents the QuickBooks tax agency query response.
type TaxAgencyQueryResponse struct {
	QueryResponse TaxAgencyQueryResult `json:"QueryResponse"`
	Time          time.Time            `json:"time"`
}

// TaxAgencyQueryResult represents query results for tax agencies.
type TaxAgencyQueryResult struct {
	TaxAgency     []TaxAgency `json:"TaxAgency,omitempty"`
	StartPosition int         `json:"startPosition,omitempty"`
	MaxResults    int         `json:"maxResults,omitempty"`
}

// TaxAgency represents a QuickBooks tax agency object.
type TaxAgency struct {
	ID                    string          `json:"Id"`
	SyncToken             string          `json:"SyncToken,omitempty"`
	DisplayName           string          `json:"DisplayName,omitempty"`
	TaxTrackedOnSales     *bool           `json:"TaxTrackedOnSales,omitempty"`
	TaxTrackedOnPurchases *bool           `json:"TaxTrackedOnPurchases,omitempty"`
	TaxRegistrationNumber string          `json:"TaxRegistrationNumber,omitempty"`
	ReportingPeriod       string          `json:"ReportingPeriod,omitempty"`
	LastFileDate          *Date           `json:"LastFileDate,omitempty"`
	TaxAgencyConfig       TaxAgencyConfig `json:"TaxAgencyConfig,omitempty"`
	MetaData              *MetaData       `json:"MetaData,omitempty"`
	Domain                string          `json:"domain,omitempty"`
	Sparse                *bool           `json:"sparse,omitempty"`
}

// CreateTaxAgencyRequest represents the documented create tax agency payload.
type CreateTaxAgencyRequest struct {
	// DisplayName is required.
	DisplayName string `json:"DisplayName"`
}
