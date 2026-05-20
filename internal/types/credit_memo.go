// Package types contains transport types for external integrations.
package types

import "time"

// CreditMemoPrintStatus represents the documented credit memo print status values.
type CreditMemoPrintStatus string

const (
	CreditMemoPrintStatusNotSet        CreditMemoPrintStatus = "NotSet"
	CreditMemoPrintStatusNeedToPrint   CreditMemoPrintStatus = "NeedToPrint"
	CreditMemoPrintStatusPrintComplete CreditMemoPrintStatus = "PrintComplete"
)

// CreditMemoLineDetailType represents the documented credit memo line detail types.
type CreditMemoLineDetailType string

const (
	CreditMemoLineDetailTypeSalesItem       CreditMemoLineDetailType = "SalesItemLineDetail"
	CreditMemoLineDetailTypeGroup           CreditMemoLineDetailType = "GroupLineDetail"
	CreditMemoLineDetailTypeDescriptionOnly CreditMemoLineDetailType = "DescriptionOnly"
	CreditMemoLineDetailTypeSubTotal        CreditMemoLineDetailType = "SubTotalLineDetail"
	CreditMemoLineDetailTypeTaxLine         CreditMemoLineDetailType = "TaxLineDetail"
	CreditMemoLineDetailTypeDiscountLine    CreditMemoLineDetailType = "DiscountLineDetail"
)

// CreditMemoResponse represents the QuickBooks credit memo response envelope.
type CreditMemoResponse struct {
	CreditMemo CreditMemo `json:"CreditMemo"`
	Time       time.Time  `json:"time"`
}

// CreditMemo represents a QuickBooks credit memo object.
type CreditMemo struct {
	ID                    string                `json:"Id"`
	Line                  []CreditMemoLine      `json:"Line,omitempty"`
	CustomerRef           *Reference            `json:"CustomerRef,omitempty"`
	SyncToken             string                `json:"SyncToken,omitempty"`
	CurrencyRef           *Reference            `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation  `json:"GlobalTaxCalculation,omitempty"`
	ProjectRef            *Reference            `json:"ProjectRef,omitempty"`
	BillEmail             *EmailAddress         `json:"BillEmail,omitempty"`
	HomeBalance           float64               `json:"HomeBalance,omitempty"`
	RemainingCredit       float64               `json:"RemainingCredit,omitempty"`
	RecurDataRef          *Reference            `json:"RecurDataRef,omitempty"`
	TaxExemptionRef       *Reference            `json:"TaxExemptionRef,omitempty"`
	Balance               float64               `json:"Balance,omitempty"`
	HomeTotalAmt          float64               `json:"HomeTotalAmt,omitempty"`
	TxnDate               *Date                 `json:"TxnDate,omitempty"`
	ClassRef              *Reference            `json:"ClassRef,omitempty"`
	PrintStatus           CreditMemoPrintStatus `json:"PrintStatus,omitempty"`
	SalesTermRef          *Reference            `json:"SalesTermRef,omitempty"`
	DocNumber             string                `json:"DocNumber,omitempty"`
	PrivateNote           string                `json:"PrivateNote,omitempty"`
	CustomerMemo          *InvoiceMemo          `json:"CustomerMemo,omitempty"`
	TxnTaxDetail          *InvoiceTxnTaxDetail  `json:"TxnTaxDetail,omitempty"`
	ApplyTaxAfterDiscount *bool                 `json:"ApplyTaxAfterDiscount,omitempty"`
	ExchangeRate          float64               `json:"ExchangeRate,omitempty"`
	DepartmentRef         *Reference            `json:"DepartmentRef,omitempty"`
	MetaData              *MetaData             `json:"MetaData,omitempty"`
	Domain                string                `json:"domain,omitempty"`
	Sparse                *bool                 `json:"sparse,omitempty"`
}

// CreateCreditMemoRequest represents the documented create credit memo payload.
type CreateCreditMemoRequest struct {
	CustomerRef           Reference             `json:"CustomerRef"`
	Line                  []CreditMemoLine      `json:"Line"`
	TxnDate               *Date                 `json:"TxnDate,omitempty"`
	DocNumber             string                `json:"DocNumber,omitempty"`
	CustomerMemo          *InvoiceMemo          `json:"CustomerMemo,omitempty"`
	PrivateNote           string                `json:"PrivateNote,omitempty"`
	BillEmail             *EmailAddress         `json:"BillEmail,omitempty"`
	EmailStatus           EmailStatus           `json:"EmailStatus,omitempty"`
	PrintStatus           CreditMemoPrintStatus `json:"PrintStatus,omitempty"`
	TxnTaxDetail          *InvoiceTxnTaxDetail  `json:"TxnTaxDetail,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation  `json:"GlobalTaxCalculation,omitempty"`
	ApplyTaxAfterDiscount *bool                 `json:"ApplyTaxAfterDiscount,omitempty"`
	CurrencyRef           *Reference            `json:"CurrencyRef,omitempty"`
	ExchangeRate          float64               `json:"ExchangeRate,omitempty"`
	ClassRef              *Reference            `json:"ClassRef,omitempty"`
	DepartmentRef         *Reference            `json:"DepartmentRef,omitempty"`
	ProjectRef            *Reference            `json:"ProjectRef,omitempty"`
	BillAddr              *PhysicalAddress      `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress      `json:"ShipAddr,omitempty"`
}

// SparseUpdateCreditMemoRequest represents the documented sparse update payload.
type SparseUpdateCreditMemoRequest struct {
	ID                    string                `json:"Id"`
	SyncToken             string                `json:"SyncToken"`
	Sparse                bool                  `json:"sparse"`
	CustomerRef           *Reference            `json:"CustomerRef,omitempty"`
	Line                  []CreditMemoLine      `json:"Line,omitempty"`
	TxnDate               *Date                 `json:"TxnDate,omitempty"`
	DocNumber             string                `json:"DocNumber,omitempty"`
	CustomerMemo          *InvoiceMemo          `json:"CustomerMemo,omitempty"`
	PrivateNote           string                `json:"PrivateNote,omitempty"`
	BillEmail             *EmailAddress         `json:"BillEmail,omitempty"`
	EmailStatus           EmailStatus           `json:"EmailStatus,omitempty"`
	PrintStatus           CreditMemoPrintStatus `json:"PrintStatus,omitempty"`
	TxnTaxDetail          *InvoiceTxnTaxDetail  `json:"TxnTaxDetail,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation  `json:"GlobalTaxCalculation,omitempty"`
	ApplyTaxAfterDiscount *bool                 `json:"ApplyTaxAfterDiscount,omitempty"`
	CurrencyRef           *Reference            `json:"CurrencyRef,omitempty"`
	ExchangeRate          float64               `json:"ExchangeRate,omitempty"`
	ClassRef              *Reference            `json:"ClassRef,omitempty"`
	DepartmentRef         *Reference            `json:"DepartmentRef,omitempty"`
	ProjectRef            *Reference            `json:"ProjectRef,omitempty"`
	Domain                string                `json:"domain,omitempty"`
}

// CreditMemoLine represents a QuickBooks credit memo line.
type CreditMemoLine struct {
	ID                    string                        `json:"Id,omitempty"`
	LineNum               float64                       `json:"LineNum,omitempty"`
	Amount                float64                       `json:"Amount,omitempty"`
	Description           string                        `json:"Description,omitempty"`
	DetailType            CreditMemoLineDetailType      `json:"DetailType,omitempty"`
	SalesItemLineDetail   *InvoiceSalesItemLineDetail   `json:"SalesItemLineDetail,omitempty"`
	GroupLineDetail       *InvoiceGroupLineDetail       `json:"GroupLineDetail,omitempty"`
	DescriptionLineDetail *InvoiceDescriptionLineDetail `json:"DescriptionLineDetail,omitempty"`
	SubTotalLineDetail    *InvoiceSubTotalLineDetail    `json:"SubTotalLineDetail,omitempty"`
	TaxLineDetail         *InvoiceTaxLineDetail         `json:"TaxLineDetail,omitempty"`
}
