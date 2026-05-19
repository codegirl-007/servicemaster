// Package types contains transport types for external integrations.
package types

// TaxClassificationResponse represents the QuickBooks tax classification response envelope.
type TaxClassificationResponse struct {
	TaxClassification TaxClassification `json:"TaxClassification"`
}

// TaxClassificationQueryResponse represents a query response for tax classifications.
type TaxClassificationQueryResponse struct {
	QueryResponse TaxClassificationQueryResult `json:"QueryResponse"`
}

// TaxClassificationQueryResult represents tax classification query results.
type TaxClassificationQueryResult struct {
	TaxClassification []TaxClassification `json:"TaxClassification,omitempty"`
}

// TaxClassification represents a QuickBooks tax classification object.
type TaxClassification struct {
	ParentRef    Reference `json:"ParentRef"`
	Level        string    `json:"level,omitempty"`
	ApplicableTo []string  `json:"applicableTo,omitempty"`
	Code         string    `json:"code,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	ID           string    `json:"id,omitempty"`
}
