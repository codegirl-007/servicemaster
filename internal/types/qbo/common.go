// Package qbo contains transport types for the QuickBooks Online API.
package qbo

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

// ReportAPPaid represents the appaid query parameter for AP balance reports.
type ReportAPPaid string

const (
	ReportAPPaidAll    ReportAPPaid = "All"
	ReportAPPaidPaid   ReportAPPaid = "Paid"
	ReportAPPaidUnpaid ReportAPPaid = "Unpaid"
)

// ReportTransactionType represents the transaction_type query parameter for transaction list reports.
type ReportTransactionType string

const (
	ReportTransactionTypeCreditCardCharge             ReportTransactionType = "CreditCardCharge"
	ReportTransactionTypeCheck                        ReportTransactionType = "Check"
	ReportTransactionTypeInvoice                      ReportTransactionType = "Invoice"
	ReportTransactionTypeReceivePayment               ReportTransactionType = "ReceivePayment"
	ReportTransactionTypeJournalEntry                 ReportTransactionType = "JournalEntry"
	ReportTransactionTypeBill                         ReportTransactionType = "Bill"
	ReportTransactionTypeCreditCardCredit             ReportTransactionType = "CreditCardCredit"
	ReportTransactionTypeVendorCredit                 ReportTransactionType = "VendorCredit"
	ReportTransactionTypeCredit                       ReportTransactionType = "Credit"
	ReportTransactionTypeBillPaymentCheck             ReportTransactionType = "BillPaymentCheck"
	ReportTransactionTypeBillPaymentCreditCard        ReportTransactionType = "BillPaymentCreditCard"
	ReportTransactionTypeCharge                       ReportTransactionType = "Charge"
	ReportTransactionTypeTransfer                     ReportTransactionType = "Transfer"
	ReportTransactionTypeDeposit                      ReportTransactionType = "Deposit"
	ReportTransactionTypeStatement                    ReportTransactionType = "Statement"
	ReportTransactionTypeBillableCharge               ReportTransactionType = "BillableCharge"
	ReportTransactionTypeTimeActivity                 ReportTransactionType = "TimeActivity"
	ReportTransactionTypeCashPurchase                 ReportTransactionType = "CashPurchase"
	ReportTransactionTypeSalesReceipt                 ReportTransactionType = "SalesReceipt"
	ReportTransactionTypeCreditMemo                   ReportTransactionType = "CreditMemo"
	ReportTransactionTypeCreditRefund                 ReportTransactionType = "CreditRefund"
	ReportTransactionTypeEstimate                     ReportTransactionType = "Estimate"
	ReportTransactionTypeInventoryQuantityAdjustment  ReportTransactionType = "InventoryQuantityAdjustment"
	ReportTransactionTypePurchaseOrder                ReportTransactionType = "PurchaseOrder"
	ReportTransactionTypeGlobalTaxPayment             ReportTransactionType = "GlobalTaxPayment"
	ReportTransactionTypeGlobalTaxAdjustment          ReportTransactionType = "GlobalTaxAdjustment"
	ReportTransactionTypeServiceTaxRefund             ReportTransactionType = "Service Tax Refund"
	ReportTransactionTypeServiceTaxGrossAdjustment    ReportTransactionType = "Service Tax Gross Adjustment"
	ReportTransactionTypeServiceTaxReversal           ReportTransactionType = "Service Tax Reversal"
	ReportTransactionTypeServiceTaxDefer              ReportTransactionType = "Service Tax Defer"
	ReportTransactionTypeServiceTaxPartialUtilisation ReportTransactionType = "Service Tax Partial Utilisation"
)

// ReportGroupBy represents the group_by query parameter for transaction list reports.
type ReportGroupBy string

const (
	ReportGroupByName            ReportGroupBy = "Name"
	ReportGroupByAccount         ReportGroupBy = "Account"
	ReportGroupByTransactionType ReportGroupBy = "Transaction Type"
	ReportGroupByCustomer        ReportGroupBy = "Customer"
	ReportGroupByVendor          ReportGroupBy = "Vendor"
	ReportGroupByEmployee        ReportGroupBy = "Employee"
	ReportGroupByLocation        ReportGroupBy = "Location"
	ReportGroupByPaymentMethod   ReportGroupBy = "Payment Method"
	ReportGroupByDay             ReportGroupBy = "Day"
	ReportGroupByWeek            ReportGroupBy = "Week"
	ReportGroupByMonth           ReportGroupBy = "Month"
	ReportGroupByQuarter         ReportGroupBy = "Quarter"
	ReportGroupByYear            ReportGroupBy = "Year"
	ReportGroupByFiscalYear      ReportGroupBy = "Fiscal Year"
	ReportGroupByFiscalQuarter   ReportGroupBy = "Fiscal Quarter"
	ReportGroupByNone            ReportGroupBy = "None"
)

// ReportPaymentMethod represents the payment_method query parameter for transaction list reports.
type ReportPaymentMethod string

const (
	ReportPaymentMethodCash            ReportPaymentMethod = "Cash"
	ReportPaymentMethodCheck           ReportPaymentMethod = "Check"
	ReportPaymentMethodDinnersClub     ReportPaymentMethod = "Dinners Club"
	ReportPaymentMethodAmericanExpress ReportPaymentMethod = "American Express"
	ReportPaymentMethodDiscover        ReportPaymentMethod = "Discover"
	ReportPaymentMethodMasterCard      ReportPaymentMethod = "MasterCard"
	ReportPaymentMethodVisa            ReportPaymentMethod = "Visa"
	ReportPaymentMethodCreditCard      ReportPaymentMethod = "Credit Card"
)

// ReportPrinted represents the printed query parameter for transaction list reports.
type ReportPrinted string

const (
	ReportPrintedPrinted     ReportPrinted = "Printed"
	ReportPrintedToBePrinted ReportPrinted = "To_be_printed"
)

// ReportCleared represents the cleared query parameter for transaction list reports.
type ReportCleared string

const (
	ReportClearedCleared    ReportCleared = "Cleared"
	ReportClearedUncleared  ReportCleared = "Uncleared"
	ReportClearedReconciled ReportCleared = "Reconciled"
	ReportClearedDeposited  ReportCleared = "Deposited"
)

// ReportAgingMethod represents the aging_method query parameter for aging reports.
type ReportAgingMethod string

const (
	ReportAgingMethodReportDate ReportAgingMethod = "Report_Date"
	ReportAgingMethodCurrent    ReportAgingMethod = "Current"
)
