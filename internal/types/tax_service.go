// Package types contains transport types for external integrations.
package types

// TaxServiceResponse represents the QuickBooks tax service response body.
type TaxServiceResponse struct {
	TaxCode        string                 `json:"TaxCode,omitempty"`
	TaxCodeID      string                 `json:"TaxCodeId,omitempty"`
	TaxRateDetails []TaxServiceRateDetail `json:"TaxRateDetails,omitempty"`
}

// CreateTaxServiceRequest represents the documented create tax service payload.
type CreateTaxServiceRequest struct {
	TaxCode        string                 `json:"TaxCode"`
	TaxCodeID      string                 `json:"TaxCodeId,omitempty"`
	TaxRateDetails []TaxServiceRateDetail `json:"TaxRateDetails"`
}

// TaxServiceRateDetail represents a tax rate detail in a tax service request or response.
type TaxServiceRateDetail struct {
	RateValue       string          `json:"RateValue,omitempty"`
	TaxRateID       string          `json:"TaxRateId,omitempty"`
	TaxApplicableOn TaxApplicableOn `json:"TaxApplicableOn,omitempty"`
	TaxAgencyID     string          `json:"TaxAgencyId,omitempty"`
	TaxRateName     string          `json:"TaxRateName,omitempty"`
}
