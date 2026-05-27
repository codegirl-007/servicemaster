// Package types contains transport types for external integrations.
package types

import "time"

// FiscalYearStartMonth represents the documented fiscal year start month values.
type FiscalYearStartMonth string

const (
	FiscalYearStartMonthJanuary   FiscalYearStartMonth = "January"
	FiscalYearStartMonthFebruary  FiscalYearStartMonth = "February"
	FiscalYearStartMonthMarch     FiscalYearStartMonth = "March"
	FiscalYearStartMonthApril     FiscalYearStartMonth = "April"
	FiscalYearStartMonthMay       FiscalYearStartMonth = "May"
	FiscalYearStartMonthJune      FiscalYearStartMonth = "June"
	FiscalYearStartMonthJuly      FiscalYearStartMonth = "July"
	FiscalYearStartMonthAugust    FiscalYearStartMonth = "August"
	FiscalYearStartMonthSeptember FiscalYearStartMonth = "September"
	FiscalYearStartMonthOctober   FiscalYearStartMonth = "October"
	FiscalYearStartMonthNovember  FiscalYearStartMonth = "November"
	FiscalYearStartMonthDecember  FiscalYearStartMonth = "December"
)

// CompanyInfoResponse represents the QuickBooks company info response envelope.
type CompanyInfoResponse struct {
	CompanyInfo CompanyInfo `json:"CompanyInfo"`
	Time        time.Time   `json:"time"`
}

// CompanyInfo represents a QuickBooks company info object.
type CompanyInfo struct {
	ID                        string               `json:"Id"`
	SyncToken                 string               `json:"SyncToken"`
	CompanyName               string               `json:"CompanyName"`
	CompanyAddr               PhysicalAddress      `json:"CompanyAddr"`
	CompanyStartDate          *Date                `json:"CompanyStartDate,omitempty"`
	LegalAddr                 *PhysicalAddress     `json:"LegalAddr,omitempty"`
	SupportedLanguages        string               `json:"SupportedLanguages,omitempty"`
	Country                   string               `json:"Country,omitempty"`
	Email                     *EmailAddress        `json:"Email,omitempty"`
	WebAddr                   *WebSiteAddress      `json:"WebAddr,omitempty"`
	NameValue                 []NameValue          `json:"NameValue,omitempty"`
	FiscalYearStartMonth      FiscalYearStartMonth `json:"FiscalYearStartMonth,omitempty"`
	CustomerCommunicationAddr *PhysicalAddress     `json:"CustomerCommunicationAddr,omitempty"`
	PrimaryPhone              *TelephoneNumber     `json:"PrimaryPhone,omitempty"`
	LegalName                 string               `json:"LegalName,omitempty"`
	EmployerID                string               `json:"EmployerId,omitempty"`
	MetaData                  *MetaData            `json:"MetaData,omitempty"`
	Domain                    string               `json:"domain,omitempty"`
	Sparse                    *bool                `json:"sparse,omitempty"`
}

// SparseUpdateCompanyInfoRequest represents the documented sparse update payload.
type SparseUpdateCompanyInfoRequest struct {
	// ID is required for update.
	ID string `json:"Id"`
	// SyncToken is required for update.
	SyncToken string `json:"SyncToken"`
	// Sparse must be true for sparse update.
	Sparse                    bool                 `json:"sparse"`
	CompanyName               string               `json:"CompanyName,omitempty"`
	CompanyAddr               *PhysicalAddress     `json:"CompanyAddr,omitempty"`
	CompanyStartDate          *Date                `json:"CompanyStartDate,omitempty"`
	LegalAddr                 *PhysicalAddress     `json:"LegalAddr,omitempty"`
	SupportedLanguages        string               `json:"SupportedLanguages,omitempty"`
	Country                   string               `json:"Country,omitempty"`
	Email                     *EmailAddress        `json:"Email,omitempty"`
	WebAddr                   *WebSiteAddress      `json:"WebAddr,omitempty"`
	NameValue                 []NameValue          `json:"NameValue,omitempty"`
	FiscalYearStartMonth      FiscalYearStartMonth `json:"FiscalYearStartMonth,omitempty"`
	CustomerCommunicationAddr *PhysicalAddress     `json:"CustomerCommunicationAddr,omitempty"`
	PrimaryPhone              *TelephoneNumber     `json:"PrimaryPhone,omitempty"`
	LegalName                 string               `json:"LegalName,omitempty"`
	EmployerID                string               `json:"EmployerId,omitempty"`
	Domain                    string               `json:"domain,omitempty"`
}

// PhysicalAddress represents a QuickBooks physical address.
type PhysicalAddress struct {
	ID                     string `json:"Id,omitempty"`
	PostalCode             string `json:"PostalCode,omitempty"`
	City                   string `json:"City,omitempty"`
	Country                string `json:"Country,omitempty"`
	Line5                  string `json:"Line5,omitempty"`
	Line4                  string `json:"Line4,omitempty"`
	Line3                  string `json:"Line3,omitempty"`
	Line2                  string `json:"Line2,omitempty"`
	Line1                  string `json:"Line1,omitempty"`
	Lat                    string `json:"Lat,omitempty"`
	Long                   string `json:"Long,omitempty"`
	CountrySubDivisionCode string `json:"CountrySubDivisionCode,omitempty"`
}
