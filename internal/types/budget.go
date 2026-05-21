// Package types contains transport types for external integrations.
package types

import "time"

// BudgetEntryType represents the documented budget entry period values.
type BudgetEntryType string

const (
	BudgetEntryTypeMonthly   BudgetEntryType = "Monthly"
	BudgetEntryTypeQuarterly BudgetEntryType = "Quarterly"
	BudgetEntryTypeAnnually  BudgetEntryType = "Annually"
)

// BudgetType represents the documented budget type values.
type BudgetType string

const (
	BudgetTypeProfitAndLoss BudgetType = "ProfitAndLoss"
)

// BudgetResponse represents the QuickBooks budget response envelope.
type BudgetResponse struct {
	Budget Budget    `json:"Budget"`
	Time   time.Time `json:"time"`
}

// Budget represents a QuickBooks budget object.
type Budget struct {
	ID              string         `json:"Id"`
	SyncToken       string         `json:"SyncToken,omitempty"`
	StartDate       *Date          `json:"StartDate,omitempty"`
	EndDate         *Date          `json:"EndDate,omitempty"`
	Name            string         `json:"Name,omitempty"`
	BudgetEntryType BudgetEntryType `json:"BudgetEntryType,omitempty"`
	BudgetType      BudgetType     `json:"BudgetType,omitempty"`
	Active          *bool          `json:"Active,omitempty"`
	BudgetDetail    []BudgetDetail `json:"BudgetDetail,omitempty"`
	MetaData        *MetaData      `json:"MetaData,omitempty"`
	Domain          string         `json:"domain,omitempty"`
	Sparse          *bool          `json:"sparse,omitempty"`
}

// CreateBudgetRequest represents the documented create budget payload.
type CreateBudgetRequest struct {
	StartDate       *Date          `json:"StartDate"`
	EndDate         *Date          `json:"EndDate"`
	Name            string         `json:"Name,omitempty"`
	BudgetEntryType BudgetEntryType `json:"BudgetEntryType,omitempty"`
	BudgetType      BudgetType     `json:"BudgetType,omitempty"`
	BudgetDetail    []BudgetDetail `json:"BudgetDetail,omitempty"`
}

// UpdateBudgetRequest represents the documented full update budget payload.
type UpdateBudgetRequest struct {
	ID              string         `json:"Id"`
	SyncToken       string         `json:"SyncToken"`
	StartDate       *Date          `json:"StartDate,omitempty"`
	EndDate         *Date          `json:"EndDate,omitempty"`
	Name            string         `json:"Name,omitempty"`
	BudgetEntryType BudgetEntryType `json:"BudgetEntryType,omitempty"`
	BudgetType      BudgetType     `json:"BudgetType,omitempty"`
	Active          *bool          `json:"Active,omitempty"`
	BudgetDetail    []BudgetDetail `json:"BudgetDetail,omitempty"`
	Domain          string         `json:"domain,omitempty"`
}

// DeleteBudgetRequest represents the documented delete budget payload.
type DeleteBudgetRequest struct {
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
}

// BudgetDeleteResponse represents the QuickBooks budget delete response envelope.
type BudgetDeleteResponse struct {
	Time time.Time `json:"time"`
}

// BudgetDetail represents one budget line item.
type BudgetDetail struct {
	Amount        float64     `json:"Amount,omitempty"`
	BudgetDate    *Date       `json:"BudgetDate,omitempty"`
	AccountRef    *Reference  `json:"AccountRef,omitempty"`
	CustomerRef   *Reference  `json:"CustomerRef,omitempty"`
	ClassRef      *Reference  `json:"ClassRef,omitempty"`
	DepartmentRef *Reference  `json:"DepartmentRef,omitempty"`
}
