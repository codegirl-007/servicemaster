// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// EmployeeGender represents the documented employee gender values.
type EmployeeGender string

const (
	EmployeeGenderMale   EmployeeGender = "Male"
	EmployeeGenderFemale EmployeeGender = "Female"
)

// EmployeeResponse represents the QuickBooks employee response envelope.
type EmployeeResponse struct {
	Employee Employee  `json:"Employee"`
	Time     time.Time `json:"time"`
}

// Employee represents a QuickBooks employee object.
type Employee struct {
	ID               string           `json:"Id"`
	SyncToken        string           `json:"SyncToken"`
	PrimaryAddr      *PhysicalAddress `json:"PrimaryAddr,omitempty"`
	V4IDPseudonym    string           `json:"V4IDPseudonym,omitempty"`
	PrimaryEmailAddr *EmailAddress    `json:"PrimaryEmailAddr,omitempty"`
	DisplayName      string           `json:"DisplayName,omitempty"`
	Title            string           `json:"Title,omitempty"`
	BillableTime     *bool            `json:"BillableTime,omitempty"`
	GivenName        string           `json:"GivenName,omitempty"`
	BirthDate        *Date            `json:"BirthDate,omitempty"`
	MiddleName       string           `json:"MiddleName,omitempty"`
	SSN              string           `json:"SSN,omitempty"`
	PrimaryPhone     *TelephoneNumber `json:"PrimaryPhone,omitempty"`
	Active           *bool            `json:"Active,omitempty"`
	ReleasedDate     *Date            `json:"ReleasedDate,omitempty"`
	MetaData         *MetaData        `json:"MetaData,omitempty"`
	CostRate         float64          `json:"CostRate,omitempty"`
	Mobile           *TelephoneNumber `json:"Mobile,omitempty"`
	Gender           EmployeeGender   `json:"Gender,omitempty"`
	HiredDate        *Date            `json:"HiredDate,omitempty"`
	BillRate         float64          `json:"BillRate,omitempty"`
	Organization     *bool            `json:"Organization,omitempty"`
	Suffix           string           `json:"Suffix,omitempty"`
	FamilyName       string           `json:"FamilyName,omitempty"`
	PrintOnCheckName string           `json:"PrintOnCheckName,omitempty"`
	EmployeeNumber   string           `json:"EmployeeNumber,omitempty"`
	Domain           string           `json:"domain,omitempty"`
	Sparse           *bool            `json:"sparse,omitempty"`
}

// CreateEmployeeRequest represents the documented create employee payload.
type CreateEmployeeRequest struct {
	PrimaryAddr      *PhysicalAddress `json:"PrimaryAddr,omitempty"`
	GivenName        string           `json:"GivenName,omitempty"`
	FamilyName       string           `json:"FamilyName,omitempty"`
	PrimaryEmailAddr *EmailAddress    `json:"PrimaryEmailAddr,omitempty"`
	DisplayName      string           `json:"DisplayName,omitempty"`
	Title            string           `json:"Title,omitempty"`
	BillableTime     *bool            `json:"BillableTime,omitempty"`
	BirthDate        *Date            `json:"BirthDate,omitempty"`
	MiddleName       string           `json:"MiddleName,omitempty"`
	SSN              string           `json:"SSN,omitempty"`
	PrimaryPhone     *TelephoneNumber `json:"PrimaryPhone,omitempty"`
	Active           *bool            `json:"Active,omitempty"`
	ReleasedDate     *Date            `json:"ReleasedDate,omitempty"`
	CostRate         float64          `json:"CostRate,omitempty"`
	Mobile           *TelephoneNumber `json:"Mobile,omitempty"`
	Gender           EmployeeGender   `json:"Gender,omitempty"`
	HiredDate        *Date            `json:"HiredDate,omitempty"`
	BillRate         float64          `json:"BillRate,omitempty"`
	Organization     *bool            `json:"Organization,omitempty"`
	Suffix           string           `json:"Suffix,omitempty"`
	PrintOnCheckName string           `json:"PrintOnCheckName,omitempty"`
	EmployeeNumber   string           `json:"EmployeeNumber,omitempty"`
}

// SparseUpdateEmployeeRequest represents the documented sparse update payload.
type SparseUpdateEmployeeRequest struct {
	// ID is required for update.
	ID string `json:"Id"`
	// SyncToken is required for update.
	SyncToken string `json:"SyncToken"`
	// Sparse must be true for sparse update.
	Sparse           bool             `json:"sparse"`
	PrimaryAddr      *PhysicalAddress `json:"PrimaryAddr,omitempty"`
	PrimaryEmailAddr *EmailAddress    `json:"PrimaryEmailAddr,omitempty"`
	DisplayName      string           `json:"DisplayName,omitempty"`
	Title            string           `json:"Title,omitempty"`
	BillableTime     *bool            `json:"BillableTime,omitempty"`
	GivenName        string           `json:"GivenName,omitempty"`
	BirthDate        *Date            `json:"BirthDate,omitempty"`
	MiddleName       string           `json:"MiddleName,omitempty"`
	SSN              string           `json:"SSN,omitempty"`
	PrimaryPhone     *TelephoneNumber `json:"PrimaryPhone,omitempty"`
	Active           *bool            `json:"Active,omitempty"`
	ReleasedDate     *Date            `json:"ReleasedDate,omitempty"`
	CostRate         float64          `json:"CostRate,omitempty"`
	Mobile           *TelephoneNumber `json:"Mobile,omitempty"`
	Gender           EmployeeGender   `json:"Gender,omitempty"`
	HiredDate        *Date            `json:"HiredDate,omitempty"`
	BillRate         float64          `json:"BillRate,omitempty"`
	Organization     *bool            `json:"Organization,omitempty"`
	Suffix           string           `json:"Suffix,omitempty"`
	FamilyName       string           `json:"FamilyName,omitempty"`
	PrintOnCheckName string           `json:"PrintOnCheckName,omitempty"`
	EmployeeNumber   string           `json:"EmployeeNumber,omitempty"`
	Domain           string           `json:"domain,omitempty"`
}
