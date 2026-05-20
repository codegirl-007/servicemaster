// Package types contains transport types for external integrations.
package types

import "time"

// EffectiveTaxRate represents an effective rate period for a tax rate.
type EffectiveTaxRate struct {
	RateValue     float64 `json:"RateValue,omitempty"`
	EndDate       string  `json:"EndDate,omitempty"`
	EffectiveDate string  `json:"EffectiveDate,omitempty"`
}

// TaxRateResponse represents the QuickBooks tax rate response envelope.
type TaxRateResponse struct {
	TaxRate TaxRate   `json:"TaxRate"`
	Time    time.Time `json:"time"`
}

// TaxRate represents a QuickBooks tax rate object.
type TaxRate struct {
	ID               string             `json:"Id"`
	SyncToken        string             `json:"SyncToken,omitempty"`
	RateValue        string             `json:"RateValue,omitempty"`
	Name             string             `json:"Name,omitempty"`
	AgencyRef        *Reference         `json:"AgencyRef,omitempty"`
	SpecialTaxType   string             `json:"SpecialTaxType,omitempty"`
	EffectiveTaxRate []EffectiveTaxRate `json:"EffectiveTaxRate,omitempty"`
	DisplayType      string             `json:"DisplayType,omitempty"`
	TaxReturnLineRef *Reference         `json:"TaxReturnLineRef,omitempty"`
	Active           *bool              `json:"Active,omitempty"`
	MetaData         *MetaData          `json:"MetaData,omitempty"`
	OriginalTaxRate  string             `json:"OriginalTaxRate,omitempty"`
	Description      string             `json:"Description,omitempty"`
	Domain           string             `json:"domain,omitempty"`
	Sparse           *bool              `json:"sparse,omitempty"`
}
