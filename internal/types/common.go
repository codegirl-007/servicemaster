// Package types contains transport types for external integrations.
package types

import "time"

// Date represents a date-only value.
type Date struct {
	time.Time
}

// EmailStatus represents the documented QuickBooks email delivery status values.
type EmailStatus string

const (
	EmailStatusNotSet     EmailStatus = "NotSet"
	EmailStatusNeedToSend EmailStatus = "NeedToSend"
	EmailStatusEmailSent  EmailStatus = "EmailSent"
)

// PrintStatus represents the documented QuickBooks print status values.
type PrintStatus string

const (
	PrintStatusNotSet        PrintStatus = "NotSet"
	PrintStatusNeedToPrint   PrintStatus = "NeedToPrint"
	PrintStatusPrintComplete PrintStatus = "PrintComplete"
)

// LinkedTxn represents a transaction linked to another QuickBooks entity.
type LinkedTxn struct {
	TxnID     string `json:"TxnId,omitempty"`
	TxnType   string `json:"TxnType,omitempty"`
	TxnLineID string `json:"TxnLineId,omitempty"`
}

// TxnTaxDetail represents transaction-level tax details.
type TxnTaxDetail struct {
	TxnTaxCodeRef *Reference `json:"TxnTaxCodeRef,omitempty"`
	TotalTax      float64    `json:"TotalTax,omitempty"`
	TaxLine       []TaxLine  `json:"TaxLine,omitempty"`
}

// TaxLine represents a tax line in transaction tax detail
type TaxLine struct {
	Amount        float64        `json:"Amount,omitempty"`
	DetailType    string         `json:"DetailType,omitempty"`
	TaxLineDetail *TaxLineDetail `json:"TaxLineDetail,omitempty"`
}

// TaxLineDetail represents tax details where TaxRateRef is optional.
type TaxLineDetail struct {
	TaxRateRef          *Reference `json:"TaxRateRef,omitempty"`
	NetAmountTaxable    float64    `json:"NetAmountTaxable,omitempty"`
	PercentBased        *bool      `json:"PercentBased,omitempty"`
	TaxInclusiveAmount  float64    `json:"TaxInclusiveAmount,omitempty"`
	OverrideDeltaAmount float64    `json:"OverrideDeltaAmount,omitempty"`
	TaxPercent          float64    `json:"TaxPercent,omitempty"`
}

// Memo represents a QuickBooks memo value object.
type Memo struct {
	Value string `json:"value,omitempty"`
}

// EmailAddress represents a QuickBooks email address object.
type EmailAddress struct {
	Address string `json:"Address,omitempty"`
}

// WebSiteAddress represents a QuickBooks website address object.
type WebSiteAddress struct {
	URI string `json:"URI,omitempty"`
}

// TelephoneNumber represents a QuickBooks telephone number object.
type TelephoneNumber struct {
	FreeFormNumber string `json:"FreeFormNumber,omitempty"`
}

// NameValue represents a QuickBooks name/value extension pair.
type NameValue struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// DeletedEntity represents a deleted QBO entity response body.
type DeletedEntity struct {
	Status string `json:"status,omitempty"`
	Domain string `json:"domain,omitempty"`
	ID     string `json:"Id,omitempty"`
}

// CustomField represents a custom field value on a QuickBooks entity.
type CustomField struct {
	DefinitionID string `json:"DefinitionId,omitempty"`
	Type         string `json:"Type,omitempty"`
	StringValue  string `json:"StringValue,omitempty"`
	Name         string `json:"Name,omitempty"`
}

// SummarizeColumnBy represents the documented summarize_column_by query parameter values.
type SummarizeColumnBy string

const (
	SummarizeColumnByTotal               SummarizeColumnBy = "Total"
	SummarizeColumnByMonth               SummarizeColumnBy = "Month"
	SummarizeColumnByWeek                SummarizeColumnBy = "Week"
	SummarizeColumnByDays                SummarizeColumnBy = "Days"
	SummarizeColumnByQuarter             SummarizeColumnBy = "Quarter"
	SummarizeColumnByYear                SummarizeColumnBy = "Year"
	SummarizeColumnByCustomers           SummarizeColumnBy = "Customers"
	SummarizeColumnByVendors             SummarizeColumnBy = "Vendors"
	SummarizeColumnByClasses             SummarizeColumnBy = "Classes"
	SummarizeColumnByDepartments         SummarizeColumnBy = "Departments"
	SummarizeColumnByEmployees           SummarizeColumnBy = "Employees"
	SummarizeColumnByProductsAndServices SummarizeColumnBy = "ProductsAndServices"
)

// ReportARPaid represents the arpaid query parameter for AR balance reports.
type ReportARPaid string

const (
	ReportARPaidAll    ReportARPaid = "All"
	ReportARPaidPaid   ReportARPaid = "Paid"
	ReportARPaidUnpaid ReportARPaid = "Unpaid"
)
