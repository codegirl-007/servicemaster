// Package types contains transport types for external integrations.
package types

import "time"

// VendorContactInfoType represents the documented vendor contact info types.
type VendorContactInfoType string

const (
	VendorContactInfoTypeTelephoneNumber VendorContactInfoType = "TelephoneNumber"
)

// VendorSource represents the documented QuickBooks Commerce source value.
type VendorSource string

const (
	VendorSourceQBCommerce VendorSource = "QBCommerce"
)

// VendorTaxReportingBasis represents the documented FR tax reporting basis values.
type VendorTaxReportingBasis string

const (
	VendorTaxReportingBasisCash    VendorTaxReportingBasis = "Cash"
	VendorTaxReportingBasisAccrual VendorTaxReportingBasis = "Accrual"
)

// VendorGSTRegistrationType represents the documented IN GST registration types.
type VendorGSTRegistrationType string

const (
	VendorGSTRegistrationTypeRegisteredRegular     VendorGSTRegistrationType = "GST_REG_REG"
	VendorGSTRegistrationTypeRegisteredComposition VendorGSTRegistrationType = "GST_REG_COMP"
	VendorGSTRegistrationTypeUnregistered          VendorGSTRegistrationType = "GST_UNREG"
	VendorGSTRegistrationTypeConsumer              VendorGSTRegistrationType = "CONSUMER"
	VendorGSTRegistrationTypeOverseas              VendorGSTRegistrationType = "OVERSEAS"
	VendorGSTRegistrationTypeSEZ                   VendorGSTRegistrationType = "SEZ"
	VendorGSTRegistrationTypeDeemed                VendorGSTRegistrationType = "DEEMED"
)

// VendorResponse represents the QuickBooks vendor response envelope.
type VendorResponse struct {
	Vendor Vendor    `json:"Vendor"`
	Time   time.Time `json:"time"`
}

// Vendor represents a QuickBooks vendor object.
type Vendor struct {
	ID                      string                    `json:"Id"`
	SyncToken               string                    `json:"SyncToken"`
	Title                   string                    `json:"Title,omitempty"`
	GivenName               string                    `json:"GivenName,omitempty"`
	MiddleName              string                    `json:"MiddleName,omitempty"`
	Suffix                  string                    `json:"Suffix,omitempty"`
	FamilyName              string                    `json:"FamilyName,omitempty"`
	Balance                 float64                   `json:"Balance,omitempty"`
	PrimaryEmailAddr        *EmailAddress             `json:"PrimaryEmailAddr,omitempty"`
	DisplayName             string                    `json:"DisplayName,omitempty"`
	OtherContactInfo        []VendorContactInfo       `json:"OtherContactInfo,omitempty"`
	APAccountRef            *Reference                `json:"APAccountRef,omitempty"`
	TermRef                 *Reference                `json:"TermRef,omitempty"`
	Source                  VendorSource              `json:"Source,omitempty"`
	GSTIN                   string                    `json:"GSTIN,omitempty"`
	T4AEligible             *bool                     `json:"T4AEligible,omitempty"`
	Fax                     *TelephoneNumber          `json:"Fax,omitempty"`
	BusinessNumber          string                    `json:"BusinessNumber,omitempty"`
	CurrencyRef             *Reference                `json:"CurrencyRef,omitempty"`
	HasTPAR                 *bool                     `json:"HasTPAR,omitempty"`
	TaxReportingBasis       VendorTaxReportingBasis   `json:"TaxReportingBasis,omitempty"`
	Mobile                  *TelephoneNumber          `json:"Mobile,omitempty"`
	PrimaryPhone            *TelephoneNumber          `json:"PrimaryPhone,omitempty"`
	Active                  *bool                     `json:"Active,omitempty"`
	AlternatePhone          *TelephoneNumber          `json:"AlternatePhone,omitempty"`
	MetaData                *MetaData                 `json:"MetaData,omitempty"`
	Vendor1099              *bool                     `json:"Vendor1099,omitempty"`
	CostRate                float64                   `json:"CostRate,omitempty"`
	BillRate                float64                   `json:"BillRate,omitempty"`
	WebAddr                 *WebSiteAddress           `json:"WebAddr,omitempty"`
	T5018Eligible           *bool                     `json:"T5018Eligible,omitempty"`
	CompanyName             string                    `json:"CompanyName,omitempty"`
	VendorPaymentBankDetail *VendorPaymentBankDetail  `json:"VendorPaymentBankDetail,omitempty"`
	TaxIdentifier           string                    `json:"TaxIdentifier,omitempty"`
	AcctNum                 string                    `json:"AcctNum,omitempty"`
	GSTRegistrationType     VendorGSTRegistrationType `json:"GSTRegistrationType,omitempty"`
	PrintOnCheckName        string                    `json:"PrintOnCheckName,omitempty"`
	BillAddr                *PhysicalAddress          `json:"BillAddr,omitempty"`
	Domain                  string                    `json:"domain,omitempty"`
	Sparse                  *bool                     `json:"sparse,omitempty"`
}

// CreateVendorRequest represents the documented create vendor payload.
type CreateVendorRequest struct {
	// DisplayName is conditionally required when no name components are provided.
	DisplayName             string                    `json:"DisplayName,omitempty"`
	Title                   string                    `json:"Title,omitempty"`
	GivenName               string                    `json:"GivenName,omitempty"`
	MiddleName              string                    `json:"MiddleName,omitempty"`
	FamilyName              string                    `json:"FamilyName,omitempty"`
	Suffix                  string                    `json:"Suffix,omitempty"`
	PrimaryEmailAddr        *EmailAddress             `json:"PrimaryEmailAddr,omitempty"`
	WebAddr                 *WebSiteAddress           `json:"WebAddr,omitempty"`
	PrimaryPhone            *TelephoneNumber          `json:"PrimaryPhone,omitempty"`
	Mobile                  *TelephoneNumber          `json:"Mobile,omitempty"`
	Fax                     *TelephoneNumber          `json:"Fax,omitempty"`
	CompanyName             string                    `json:"CompanyName,omitempty"`
	PrintOnCheckName        string                    `json:"PrintOnCheckName,omitempty"`
	BillAddr                *PhysicalAddress          `json:"BillAddr,omitempty"`
	TaxIdentifier           string                    `json:"TaxIdentifier,omitempty"`
	AcctNum                 string                    `json:"AcctNum,omitempty"`
	Vendor1099              *bool                     `json:"Vendor1099,omitempty"`
	Active                  *bool                     `json:"Active,omitempty"`
	TermRef                 *Reference                `json:"TermRef,omitempty"`
	OtherContactInfo        []VendorContactInfo       `json:"OtherContactInfo,omitempty"`
	APAccountRef            *Reference                `json:"APAccountRef,omitempty"`
	Source                  VendorSource              `json:"Source,omitempty"`
	GSTIN                   string                    `json:"GSTIN,omitempty"`
	T4AEligible             *bool                     `json:"T4AEligible,omitempty"`
	BusinessNumber          string                    `json:"BusinessNumber,omitempty"`
	HasTPAR                 *bool                     `json:"HasTPAR,omitempty"`
	TaxReportingBasis       VendorTaxReportingBasis   `json:"TaxReportingBasis,omitempty"`
	AlternatePhone          *TelephoneNumber          `json:"AlternatePhone,omitempty"`
	CostRate                float64                   `json:"CostRate,omitempty"`
	BillRate                float64                   `json:"BillRate,omitempty"`
	T5018Eligible           *bool                     `json:"T5018Eligible,omitempty"`
	VendorPaymentBankDetail *VendorPaymentBankDetail  `json:"VendorPaymentBankDetail,omitempty"`
	GSTRegistrationType     VendorGSTRegistrationType `json:"GSTRegistrationType,omitempty"`
}

// VendorContactInfo represents vendor contact info of any supported type.
type VendorContactInfo struct {
	Type      VendorContactInfoType `json:"Type,omitempty"`
	Telephone *TelephoneNumber      `json:"Telephone,omitempty"`
}

// VendorPaymentBankDetail represents AU vendor payment bank details.
type VendorPaymentBankDetail struct {
	BankAccountName      string `json:"BankAccountName,omitempty"`
	BankBranchIdentifier string `json:"BankBranchIdentifier,omitempty"`
	BankAccountNumber    string `json:"BankAccountNumber,omitempty"`
	StatementText        string `json:"StatementText,omitempty"`
}
