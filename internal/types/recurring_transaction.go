// Package types contains transport types for external integrations.
package types

import "time"

// RecurType represents the documented recurring transaction recur types.
type RecurType string

const (
	RecurTypeAutomated RecurType = "Automated"
	RecurTypeReminded  RecurType = "Reminded"
)

// RecurringIntervalType represents the documented recurring schedule interval types.
type RecurringIntervalType string

const (
	RecurringIntervalTypeDaily   RecurringIntervalType = "Daily"
	RecurringIntervalTypeWeekly  RecurringIntervalType = "Weekly"
	RecurringIntervalTypeMonthly RecurringIntervalType = "Monthly"
	RecurringIntervalTypeYearly  RecurringIntervalType = "Yearly"
)

// RecurringTransactionResponse represents the QuickBooks recurring transaction response envelope.
type RecurringTransactionResponse struct {
	RecurringTransaction RecurringTransaction `json:"RecurringTransaction"`
	Time                 time.Time            `json:"time"`
}

// RecurringTransaction represents a QuickBooks recurring transaction wrapper object.
type RecurringTransaction struct {
	Status  string            `json:"status,omitempty"`
	Domain  string            `json:"domain,omitempty"`
	Invoice *RecurringInvoice `json:"Invoice,omitempty"`
	Bill    *RecurringBill    `json:"Bill,omitempty"`
}

// RecurringInvoice embeds an invoice template with recurring schedule metadata.
type RecurringInvoice struct {
	Invoice
	RecurringInfo RecurringInfo `json:"RecurringInfo"`
}

// RecurringBill embeds a bill template with recurring schedule metadata.
type RecurringBill struct {
	Bill
	RecurringInfo RecurringInfo `json:"RecurringInfo"`
}

// CreateRecurringTransactionRequest represents the documented create recurring transaction payload.
type CreateRecurringTransactionRequest struct {
	Invoice *RecurringInvoice `json:"Invoice,omitempty"`
	Bill    *RecurringBill    `json:"Bill,omitempty"`
}

// DeleteRecurringTransactionRequest represents the documented delete recurring transaction payload.
type DeleteRecurringTransactionRequest struct {
	Invoice *DeleteRecurringInvoiceTemplate `json:"Invoice,omitempty"`
	Bill    *DeleteRecurringBillTemplate    `json:"Bill,omitempty"`
}

// DeleteRecurringInvoiceTemplate identifies a recurring invoice template to delete.
type DeleteRecurringInvoiceTemplate struct {
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
}

// DeleteRecurringBillTemplate identifies a recurring bill template to delete.
type DeleteRecurringBillTemplate struct {
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
}

// RecurringTransactionDeleteResponse represents the QuickBooks deleted recurring transaction response envelope.
type RecurringTransactionDeleteResponse struct {
	RecurringTransaction RecurringTransaction `json:"RecurringTransaction"`
	Time                 time.Time            `json:"time"`
}

// RecurringInfo represents recurring schedule metadata on a transaction template.
type RecurringInfo struct {
	Active       *bool                    `json:"Active,omitempty"`
	RecurType    RecurType                `json:"RecurType,omitempty"`
	Name         string                   `json:"Name,omitempty"`
	ScheduleInfo *RecurringScheduleInfo   `json:"ScheduleInfo,omitempty"`
}

// RecurringScheduleInfo represents the schedule for a recurring transaction.
type RecurringScheduleInfo struct {
	StartDate      *Date                 `json:"StartDate,omitempty"`
	EndDate        *Date                 `json:"EndDate,omitempty"`
	NextDate       *Date                 `json:"NextDate,omitempty"`
	PreviousDate   *Date                 `json:"PreviousDate,omitempty"`
	IntervalType   RecurringIntervalType `json:"IntervalType,omitempty"`
	NumInterval    int                   `json:"NumInterval,omitempty"`
	DayOfMonth     int                   `json:"DayOfMonth,omitempty"`
	DaysBefore     int                   `json:"DaysBefore,omitempty"`
	RemindDays     int                   `json:"RemindDays,omitempty"`
	MaxOccurrences int                   `json:"MaxOccurrences,omitempty"`
	MonthOfYear    int                   `json:"MonthOfYear,omitempty"`
}
