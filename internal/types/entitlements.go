// Package types contains transport types for external integrations.
package types

import "time"

// EntitlementsResponse represents the QuickBooks entitlements response.
type EntitlementsResponse struct {
	PlanName             string               `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	Entitlements         []Entitlement        `json:"Entitlement,omitempty" xml:"Entitlement,omitempty"`
	SupportedLanguages   string               `json:"SupportedLanguages,omitempty" xml:"SupportedLanguages,omitempty"`
	CompanyStartDate     *time.Time           `json:"CompanyStartDate,omitempty" xml:"CompanyStartDate,omitempty"`
	EmployerID           string               `json:"EmployerId,omitempty" xml:"EmployerId,omitempty"`
	QBOCompany           *bool                `json:"QboCompany,omitempty" xml:"QboCompany,omitempty"`
	Email                *EmailAddress        `json:"Email,omitempty" xml:"Email,omitempty"`
	WebAddr              *WebSiteAddress      `json:"WebAddr,omitempty" xml:"WebAddr,omitempty"`
	FiscalYearStartMonth FiscalYearStartMonth `json:"FiscalYearStartMonth,omitempty" xml:"FiscalYearStartMonth,omitempty"`
	Thresholds           []Threshold          `json:"Thresholds,omitempty" xml:"Thresholds,omitempty"`
	DaysRemainingTrial   int                  `json:"DaysRemainingTrial,omitempty" xml:"DaysRemainingTrial,omitempty"`
	MaxUsers             int                  `json:"MaxUsers,omitempty" xml:"MaxUsers,omitempty"`
	CurrentUsers         int                  `json:"CurrentUsers,omitempty" xml:"CurrentUsers,omitempty"`
}

// Entitlement represents an entitlement flag for a company.
type Entitlement struct {
	ID   int    `json:"id,omitempty" xml:"id,attr,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	Term string `json:"term,omitempty" xml:"term,omitempty"`
}

// Threshold represents a company threshold value.
type Threshold struct {
	CurrentCount   string `json:"currentCount,omitempty" xml:"currentCount,omitempty"`
	AboveThreshold string `json:"aboveThreshold,omitempty" xml:"aboveThreshold,omitempty"`
	Enforced       string `json:"enforced,omitempty" xml:"enforced,omitempty"`
	Limit          string `json:"limit,omitempty" xml:"limit,omitempty"`
	Name           string `json:"name,omitempty" xml:"name,omitempty"`
}
