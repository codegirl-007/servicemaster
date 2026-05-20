// Package types contains transport types for external integrations.
package types

import "time"

// InvoiceLineDetailType represents the documented invoice line detail types.
type InvoiceLineDetailType string

const (
	InvoiceLineDetailTypeSalesItem       InvoiceLineDetailType = "SalesItemLineDetail"
	InvoiceLineDetailTypeGroup           InvoiceLineDetailType = "GroupLineDetail"
	InvoiceLineDetailTypeDescriptionOnly InvoiceLineDetailType = "DescriptionOnly"
	InvoiceLineDetailTypeSubTotal        InvoiceLineDetailType = "SubTotalLineDetail"
	InvoiceLineDetailTypeTaxLine         InvoiceLineDetailType = "TaxLineDetail"
	InvoiceLineDetailTypeDiscountLine    InvoiceLineDetailType = "DiscountLineDetail"
)

// InvoiceResponse represents the QuickBooks invoice response envelope.
type InvoiceResponse struct {
	Invoice Invoice   `json:"Invoice"`
	Time    time.Time `json:"time"`
}

// Invoice represents a QuickBooks invoice object.
type Invoice struct {
	ID                    string               `json:"Id"`
	SyncToken             string               `json:"SyncToken"`
	CustomerRef           *Reference           `json:"CustomerRef,omitempty"`
	Line                  []InvoiceLine        `json:"Line,omitempty"`
	TxnDate               *Date                `json:"TxnDate,omitempty"`
	DueDate               *Date                `json:"DueDate,omitempty"`
	DocNumber             string               `json:"DocNumber,omitempty"`
	SalesTermRef          *Reference           `json:"SalesTermRef,omitempty"`
	CurrencyRef           *Reference           `json:"CurrencyRef,omitempty"`
	ExchangeRate          float64              `json:"ExchangeRate,omitempty"`
	BillAddr              *PhysicalAddress     `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress     `json:"ShipAddr,omitempty"`
	ShipFromAddr          *PhysicalAddress     `json:"ShipFromAddr,omitempty"`
	BillEmail             *EmailAddress        `json:"BillEmail,omitempty"`
	EmailStatus           EmailStatus          `json:"EmailStatus,omitempty"`
	PrintStatus           PrintStatus          `json:"PrintStatus,omitempty"`
	TxnTaxDetail          *InvoiceTxnTaxDetail `json:"TxnTaxDetail,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation `json:"GlobalTaxCalculation,omitempty"`
	ApplyTaxAfterDiscount *bool                `json:"ApplyTaxAfterDiscount,omitempty"`
	TotalAmt              float64              `json:"TotalAmt,omitempty"`
	Balance               float64              `json:"Balance,omitempty"`
	Deposit               float64              `json:"Deposit,omitempty"`
	PrivateNote           string               `json:"PrivateNote,omitempty"`
	CustomerMemo          *InvoiceMemo         `json:"CustomerMemo,omitempty"`
	ClassRef              *Reference           `json:"ClassRef,omitempty"`
	DepartmentRef         *Reference           `json:"DepartmentRef,omitempty"`
	ProjectRef            *Reference           `json:"ProjectRef,omitempty"`
	LinkedTxn             []LinkedTxn          `json:"LinkedTxn,omitempty"`
	HomeTotalAmt          float64              `json:"HomeTotalAmt,omitempty"`
	HomeBalance           float64              `json:"HomeBalance,omitempty"`
	MetaData              *MetaData            `json:"MetaData,omitempty"`
	Domain                string               `json:"domain,omitempty"`
	Sparse                *bool                `json:"sparse,omitempty"`
}

// CreateInvoiceRequest represents the documented create invoice payload.
type CreateInvoiceRequest struct {
	CustomerRef           Reference            `json:"CustomerRef"`
	Line                  []InvoiceLine        `json:"Line"`
	TxnDate               *Date                `json:"TxnDate,omitempty"`
	DueDate               *Date                `json:"DueDate,omitempty"`
	DocNumber             string               `json:"DocNumber,omitempty"`
	SalesTermRef          *Reference           `json:"SalesTermRef,omitempty"`
	CurrencyRef           *Reference           `json:"CurrencyRef,omitempty"`
	ExchangeRate          float64              `json:"ExchangeRate,omitempty"`
	BillAddr              *PhysicalAddress     `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress     `json:"ShipAddr,omitempty"`
	ShipFromAddr          *PhysicalAddress     `json:"ShipFromAddr,omitempty"`
	BillEmail             *EmailAddress        `json:"BillEmail,omitempty"`
	EmailStatus           EmailStatus          `json:"EmailStatus,omitempty"`
	PrintStatus           PrintStatus          `json:"PrintStatus,omitempty"`
	TxnTaxDetail          *InvoiceTxnTaxDetail `json:"TxnTaxDetail,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation `json:"GlobalTaxCalculation,omitempty"`
	ApplyTaxAfterDiscount *bool                `json:"ApplyTaxAfterDiscount,omitempty"`
	Deposit               float64              `json:"Deposit,omitempty"`
	PrivateNote           string               `json:"PrivateNote,omitempty"`
	CustomerMemo          *InvoiceMemo         `json:"CustomerMemo,omitempty"`
	ClassRef              *Reference           `json:"ClassRef,omitempty"`
	DepartmentRef         *Reference           `json:"DepartmentRef,omitempty"`
	ProjectRef            *Reference           `json:"ProjectRef,omitempty"`
	LinkedTxn             []LinkedTxn          `json:"LinkedTxn,omitempty"`
}

// SparseUpdateInvoiceRequest represents the documented sparse update payload.
type SparseUpdateInvoiceRequest struct {
	ID                    string               `json:"Id"`
	SyncToken             string               `json:"SyncToken"`
	Sparse                bool                 `json:"sparse"`
	CustomerRef           *Reference           `json:"CustomerRef,omitempty"`
	Line                  []InvoiceLine        `json:"Line,omitempty"`
	TxnDate               *Date                `json:"TxnDate,omitempty"`
	DueDate               *Date                `json:"DueDate,omitempty"`
	DocNumber             string               `json:"DocNumber,omitempty"`
	SalesTermRef          *Reference           `json:"SalesTermRef,omitempty"`
	CurrencyRef           *Reference           `json:"CurrencyRef,omitempty"`
	ExchangeRate          float64              `json:"ExchangeRate,omitempty"`
	BillAddr              *PhysicalAddress     `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress     `json:"ShipAddr,omitempty"`
	ShipFromAddr          *PhysicalAddress     `json:"ShipFromAddr,omitempty"`
	BillEmail             *EmailAddress        `json:"BillEmail,omitempty"`
	EmailStatus           EmailStatus          `json:"EmailStatus,omitempty"`
	PrintStatus           PrintStatus          `json:"PrintStatus,omitempty"`
	TxnTaxDetail          *InvoiceTxnTaxDetail `json:"TxnTaxDetail,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation `json:"GlobalTaxCalculation,omitempty"`
	ApplyTaxAfterDiscount *bool                `json:"ApplyTaxAfterDiscount,omitempty"`
	Deposit               float64              `json:"Deposit,omitempty"`
	PrivateNote           string               `json:"PrivateNote,omitempty"`
	CustomerMemo          *InvoiceMemo         `json:"CustomerMemo,omitempty"`
	ClassRef              *Reference           `json:"ClassRef,omitempty"`
	DepartmentRef         *Reference           `json:"DepartmentRef,omitempty"`
	ProjectRef            *Reference           `json:"ProjectRef,omitempty"`
	LinkedTxn             []LinkedTxn          `json:"LinkedTxn,omitempty"`
	Domain                string               `json:"domain,omitempty"`
}

// InvoiceMemo represents a QuickBooks invoice memo object.
type InvoiceMemo struct {
	Value string `json:"value,omitempty"`
}

// InvoiceLine represents a QuickBooks invoice line.
type InvoiceLine struct {
	ID                    string                        `json:"Id,omitempty"`
	LineNum               float64                       `json:"LineNum,omitempty"`
	Amount                float64                       `json:"Amount,omitempty"`
	Description           string                        `json:"Description,omitempty"`
	DetailType            InvoiceLineDetailType         `json:"DetailType,omitempty"`
	SalesItemLineDetail   *InvoiceSalesItemLineDetail   `json:"SalesItemLineDetail,omitempty"`
	GroupLineDetail       *InvoiceGroupLineDetail       `json:"GroupLineDetail,omitempty"`
	DescriptionLineDetail *InvoiceDescriptionLineDetail `json:"DescriptionLineDetail,omitempty"`
	SubTotalLineDetail    *InvoiceSubTotalLineDetail    `json:"SubTotalLineDetail,omitempty"`
	TaxLineDetail         *InvoiceTaxLineDetail         `json:"TaxLineDetail,omitempty"`
}

// InvoiceSalesItemLineDetail represents sales-item invoice line details.
type InvoiceSalesItemLineDetail struct {
	ItemRef         *Reference `json:"ItemRef,omitempty"`
	ClassRef        *Reference `json:"ClassRef,omitempty"`
	TaxCodeRef      *Reference `json:"TaxCodeRef,omitempty"`
	ServiceDate     *Date      `json:"ServiceDate,omitempty"`
	Qty             float64    `json:"Qty,omitempty"`
	UnitPrice       float64    `json:"UnitPrice,omitempty"`
	TaxInclusiveAmt float64    `json:"TaxInclusiveAmt,omitempty"`
	DiscountAmt     float64    `json:"DiscountAmt,omitempty"`
	DiscountRate    float64    `json:"DiscountRate,omitempty"`
	ItemAccountRef  *Reference `json:"ItemAccountRef,omitempty"`
}

// InvoiceGroupLineDetail represents grouped invoice line details.
type InvoiceGroupLineDetail struct {
	GroupItemRef *Reference    `json:"GroupItemRef,omitempty"`
	Quantity     float64       `json:"Quantity,omitempty"`
	Line         []InvoiceLine `json:"Line,omitempty"`
}

// InvoiceDescriptionLineDetail represents descriptive invoice line details.
type InvoiceDescriptionLineDetail struct {
	TaxCodeRef  *Reference `json:"TaxCodeRef,omitempty"`
	ServiceDate *Date      `json:"ServiceDate,omitempty"`
}

// InvoiceSubTotalLineDetail represents subtotal invoice line details.
type InvoiceSubTotalLineDetail struct{}

// InvoiceTxnTaxDetail represents invoice-level tax details.
type InvoiceTxnTaxDetail struct {
	TxnTaxCodeRef *Reference       `json:"TxnTaxCodeRef,omitempty"`
	TotalTax      float64          `json:"TotalTax,omitempty"`
	TaxLine       []InvoiceTaxLine `json:"TaxLine,omitempty"`
}

// InvoiceTaxLine represents an invoice tax line.
type InvoiceTaxLine struct {
	DetailType    InvoiceLineDetailType `json:"DetailType,omitempty"`
	TaxLineDetail InvoiceTaxLineDetail  `json:"TaxLineDetail"`
	Amount        float64               `json:"Amount,omitempty"`
}

// InvoiceTaxLineDetail represents invoice tax line details.
type InvoiceTaxLineDetail struct {
	TaxRateRef          *Reference `json:"TaxRateRef,omitempty"`
	NetAmountTaxable    float64    `json:"NetAmountTaxable,omitempty"`
	PercentBased        *bool      `json:"PercentBased,omitempty"`
	TaxInclusiveAmount  float64    `json:"TaxInclusiveAmount,omitempty"`
	OverrideDeltaAmount float64    `json:"OverrideDeltaAmount,omitempty"`
	TaxPercent          float64    `json:"TaxPercent,omitempty"`
}

// InvoiceLinkedTxn represents transactions linked to an invoice.
type InvoiceLinkedTxn struct {
	TxnID     string `json:"TxnId,omitempty"`
	TxnType   string `json:"TxnType,omitempty"`
	TxnLineID string `json:"TxnLineId,omitempty"`
}
