// Package types contains transport types for external integrations.
package types

import "time"

// TaxCodeConfigType represents the documented tax code config values.
type TaxCodeConfigType string

const (
	TaxCodeConfigTypeUserDefined     TaxCodeConfigType = "USER_DEFINED"
	TaxCodeConfigTypeSystemGenerated TaxCodeConfigType = "SYSTEM_GENERATED"
)

// TaxCodeResponse represents the QuickBooks tax code response envelope.
type TaxCodeResponse struct {
	TaxCode TaxCode   `json:"TaxCode"`
	Time    time.Time `json:"time"`
}

// TaxCode represents a QuickBooks tax code object.
type TaxCode struct {
	ID                  string            `json:"Id"`
	Name                string            `json:"Name,omitempty"`
	SyncToken           string            `json:"SyncToken,omitempty"`
	PurchaseTaxRateList *TaxRateList      `json:"PurchaseTaxRateList,omitempty"`
	SalesTaxRateList    *TaxRateList      `json:"SalesTaxRateList,omitempty"`
	TaxCodeConfigType   TaxCodeConfigType `json:"TaxCodeConfigType,omitempty"`
	TaxGroup            *bool             `json:"TaxGroup,omitempty"`
	Taxable             *bool             `json:"Taxable,omitempty"`
	Active              *bool             `json:"Active,omitempty"`
	Description         string            `json:"Description,omitempty"`
	Hidden              *bool             `json:"Hidden,omitempty"`
	MetaData            *MetaData         `json:"MetaData,omitempty"`
	Domain              string            `json:"domain,omitempty"`
	Sparse              *bool             `json:"sparse,omitempty"`
}

// TaxRateList represents a QuickBooks tax rate list.
type TaxRateList struct {
	TaxRateDetail []TaxRateDetail `json:"TaxRateDetail,omitempty"`
}

// TaxRateDetail represents a tax rate reference and ordering in a tax code.
type TaxRateDetail struct {
	TaxRateRef        *Reference `json:"TaxRateRef,omitempty"`
	TaxTypeApplicable string     `json:"TaxTypeApplicable,omitempty"`
	TaxOrder          int        `json:"TaxOrder,omitempty"`
}
