// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// PreferredDeliveryMethod represents the documented preferred delivery method values.
type PreferredDeliveryMethod string

const (
	PreferredDeliveryMethodPrint PreferredDeliveryMethod = "Print"
	PreferredDeliveryMethodEmail PreferredDeliveryMethod = "Email"
	PreferredDeliveryMethodNone  PreferredDeliveryMethod = "None"
)

// CustomerResponse represents the QuickBooks customer response envelope.
type CustomerResponse struct {
	Customer Customer  `json:"Customer"`
	Time     time.Time `json:"time"`
}

// Customer represents a QuickBooks customer object.
type Customer struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken"`
	DisplayName             string                  `json:"DisplayName,omitempty"`
	Title                   string                  `json:"Title,omitempty"`
	GivenName               string                  `json:"GivenName,omitempty"`
	MiddleName              string                  `json:"MiddleName,omitempty"`
	FamilyName              string                  `json:"FamilyName,omitempty"`
	Suffix                  string                  `json:"Suffix,omitempty"`
	CompanyName             string                  `json:"CompanyName,omitempty"`
	PrintOnCheckName        string                  `json:"PrintOnCheckName,omitempty"`
	Notes                   string                  `json:"Notes,omitempty"`
	Active                  *bool                   `json:"Active,omitempty"`
	Job                     *bool                   `json:"Job,omitempty"`
	ParentRef               *Reference              `json:"ParentRef,omitempty"`
	BillWithParent          *bool                   `json:"BillWithParent,omitempty"`
	PrimaryEmailAddr        *EmailAddress           `json:"PrimaryEmailAddr,omitempty"`
	PrimaryPhone            *TelephoneNumber        `json:"PrimaryPhone,omitempty"`
	AlternatePhone          *TelephoneNumber        `json:"AlternatePhone,omitempty"`
	Mobile                  *TelephoneNumber        `json:"Mobile,omitempty"`
	Fax                     *TelephoneNumber        `json:"Fax,omitempty"`
	WebAddr                 *WebSiteAddress         `json:"WebAddr,omitempty"`
	BillAddr                *PhysicalAddress        `json:"BillAddr,omitempty"`
	ShipAddr                *PhysicalAddress        `json:"ShipAddr,omitempty"`
	Taxable                 *bool                   `json:"Taxable,omitempty"`
	DefaultTaxCodeRef       *Reference              `json:"DefaultTaxCodeRef,omitempty"`
	PaymentMethodRef        *Reference              `json:"PaymentMethodRef,omitempty"`
	SalesTermRef            *Reference              `json:"SalesTermRef,omitempty"`
	Balance                 float64                 `json:"Balance,omitempty"`
	OpenBalanceDate         *Date                   `json:"OpenBalanceDate,omitempty"`
	FullyQualifiedName      string                  `json:"FullyQualifiedName,omitempty"`
	Level                   int                     `json:"Level,omitempty"`
	BalanceWithJobs         float64                 `json:"BalanceWithJobs,omitempty"`
	PreferredDeliveryMethod PreferredDeliveryMethod `json:"PreferredDeliveryMethod,omitempty"`
	MetaData                *MetaData               `json:"MetaData,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
	Sparse                  *bool                   `json:"sparse,omitempty"`
}

// CreateCustomerRequest represents the documented create customer payload.
type CreateCustomerRequest struct {
	// DisplayName is conditionally required when no name components are provided.
	DisplayName             string                  `json:"DisplayName,omitempty"`
	Title                   string                  `json:"Title,omitempty"`
	GivenName               string                  `json:"GivenName,omitempty"`
	MiddleName              string                  `json:"MiddleName,omitempty"`
	FamilyName              string                  `json:"FamilyName,omitempty"`
	Suffix                  string                  `json:"Suffix,omitempty"`
	CompanyName             string                  `json:"CompanyName,omitempty"`
	PrintOnCheckName        string                  `json:"PrintOnCheckName,omitempty"`
	Notes                   string                  `json:"Notes,omitempty"`
	Active                  *bool                   `json:"Active,omitempty"`
	Job                     *bool                   `json:"Job,omitempty"`
	ParentRef               *Reference              `json:"ParentRef,omitempty"`
	BillWithParent          *bool                   `json:"BillWithParent,omitempty"`
	PrimaryEmailAddr        *EmailAddress           `json:"PrimaryEmailAddr,omitempty"`
	PrimaryPhone            *TelephoneNumber        `json:"PrimaryPhone,omitempty"`
	AlternatePhone          *TelephoneNumber        `json:"AlternatePhone,omitempty"`
	Mobile                  *TelephoneNumber        `json:"Mobile,omitempty"`
	Fax                     *TelephoneNumber        `json:"Fax,omitempty"`
	WebAddr                 *WebSiteAddress         `json:"WebAddr,omitempty"`
	BillAddr                *PhysicalAddress        `json:"BillAddr,omitempty"`
	ShipAddr                *PhysicalAddress        `json:"ShipAddr,omitempty"`
	Taxable                 *bool                   `json:"Taxable,omitempty"`
	DefaultTaxCodeRef       *Reference              `json:"DefaultTaxCodeRef,omitempty"`
	PaymentMethodRef        *Reference              `json:"PaymentMethodRef,omitempty"`
	SalesTermRef            *Reference              `json:"SalesTermRef,omitempty"`
	Balance                 float64                 `json:"Balance,omitempty"`
	OpenBalanceDate         *Date                   `json:"OpenBalanceDate,omitempty"`
	PreferredDeliveryMethod PreferredDeliveryMethod `json:"PreferredDeliveryMethod,omitempty"`
}

// SparseUpdateCustomerRequest represents the documented sparse update payload.
type SparseUpdateCustomerRequest struct {
	// ID is required for update.
	ID string `json:"Id"`
	// SyncToken is required for update.
	SyncToken string `json:"SyncToken"`
	// Sparse must be true for sparse update.
	Sparse                  bool                    `json:"sparse"`
	DisplayName             string                  `json:"DisplayName,omitempty"`
	Title                   string                  `json:"Title,omitempty"`
	GivenName               string                  `json:"GivenName,omitempty"`
	MiddleName              string                  `json:"MiddleName,omitempty"`
	FamilyName              string                  `json:"FamilyName,omitempty"`
	Suffix                  string                  `json:"Suffix,omitempty"`
	CompanyName             string                  `json:"CompanyName,omitempty"`
	PrintOnCheckName        string                  `json:"PrintOnCheckName,omitempty"`
	Notes                   string                  `json:"Notes,omitempty"`
	Active                  *bool                   `json:"Active,omitempty"`
	Job                     *bool                   `json:"Job,omitempty"`
	ParentRef               *Reference              `json:"ParentRef,omitempty"`
	BillWithParent          *bool                   `json:"BillWithParent,omitempty"`
	PrimaryEmailAddr        *EmailAddress           `json:"PrimaryEmailAddr,omitempty"`
	PrimaryPhone            *TelephoneNumber        `json:"PrimaryPhone,omitempty"`
	AlternatePhone          *TelephoneNumber        `json:"AlternatePhone,omitempty"`
	Mobile                  *TelephoneNumber        `json:"Mobile,omitempty"`
	Fax                     *TelephoneNumber        `json:"Fax,omitempty"`
	WebAddr                 *WebSiteAddress         `json:"WebAddr,omitempty"`
	BillAddr                *PhysicalAddress        `json:"BillAddr,omitempty"`
	ShipAddr                *PhysicalAddress        `json:"ShipAddr,omitempty"`
	Taxable                 *bool                   `json:"Taxable,omitempty"`
	DefaultTaxCodeRef       *Reference              `json:"DefaultTaxCodeRef,omitempty"`
	PaymentMethodRef        *Reference              `json:"PaymentMethodRef,omitempty"`
	SalesTermRef            *Reference              `json:"SalesTermRef,omitempty"`
	PreferredDeliveryMethod PreferredDeliveryMethod `json:"PreferredDeliveryMethod,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
}
