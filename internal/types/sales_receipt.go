// Package types contains transport types for external integrations.
package types

import "time"

// SalesReceiptLineDetailType represents the documented sales receipt line detail types.
type SalesReceiptLineDetailType string

const (
	SalesReceiptLineDetailTypeSalesItem       SalesReceiptLineDetailType = "SalesItemLineDetail"
	SalesReceiptLineDetailTypeGroup           SalesReceiptLineDetailType = "GroupLineDetail"
	SalesReceiptLineDetailTypeDescriptionOnly SalesReceiptLineDetailType = "DescriptionOnly"
	SalesReceiptLineDetailTypeSubTotal        SalesReceiptLineDetailType = "SubTotalLineDetail"
	SalesReceiptLineDetailTypeTaxLine         SalesReceiptLineDetailType = "TaxLineDetail"
	SalesReceiptLineDetailTypeDiscountLine    SalesReceiptLineDetailType = "DiscountLineDetail"
)

// SalesReceiptResponse represents the QuickBooks sales receipt response envelope.
type SalesReceiptResponse struct {
	SalesReceipt SalesReceipt `json:"SalesReceipt"`
	Time         time.Time    `json:"time"`
}

// SalesReceipt represents a QuickBooks sales receipt object.
type SalesReceipt struct {
	ID                    string                    `json:"Id"`
	Line                  []SalesReceiptLine        `json:"Line,omitempty"`
	CustomerRef           *Reference                `json:"CustomerRef,omitempty"`
	SyncToken             string                    `json:"SyncToken,omitempty"`
	ShipFromAddr          *PhysicalAddress          `json:"ShipFromAddr,omitempty"`
	CurrencyRef           *Reference                `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation      `json:"GlobalTaxCalculation,omitempty"`
	ProjectRef            *Reference                `json:"ProjectRef,omitempty"`
	BillEmail             *EmailAddress             `json:"BillEmail,omitempty"`
	HomeBalance           float64                   `json:"HomeBalance,omitempty"`
	DeliveryInfo          *SalesReceiptDeliveryInfo `json:"DeliveryInfo,omitempty"`
	RecurDataRef          *Reference                `json:"RecurDataRef,omitempty"`
	TotalAmt              float64                   `json:"TotalAmt,omitempty"`
	TxnDate               *Date                     `json:"TxnDate,omitempty"`
	CustomField           []NameValue               `json:"CustomField,omitempty"`
	ClassRef              *Reference                `json:"ClassRef,omitempty"`
	PrintStatus           PrintStatus               `json:"PrintStatus,omitempty"`
	DocNumber             string                    `json:"DocNumber,omitempty"`
	PrivateNote           string                    `json:"PrivateNote,omitempty"`
	TxnTaxDetail          *TxnTaxDetail             `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef      *Reference                `json:"PaymentMethodRef,omitempty"`
	ExchangeRate          float64                   `json:"ExchangeRate,omitempty"`
	DepartmentRef         *Reference                `json:"DepartmentRef,omitempty"`
	DepositToAccountRef   *Reference                `json:"DepositToAccountRef,omitempty"`
	PaymentRefNum         string                    `json:"PaymentRefNum,omitempty"`
	CustomerMemo          *InvoiceMemo              `json:"CustomerMemo,omitempty"`
	ApplyTaxAfterDiscount *bool                     `json:"ApplyTaxAfterDiscount,omitempty"`
	BillAddr              *PhysicalAddress          `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress          `json:"ShipAddr,omitempty"`
	EmailStatus           EmailStatus               `json:"EmailStatus,omitempty"`
	MetaData              *MetaData                 `json:"MetaData,omitempty"`
	Domain                string                    `json:"domain,omitempty"`
	Sparse                *bool                     `json:"sparse,omitempty"`
}

// CreateSalesReceiptRequest represents the documented create sales receipt payload.
type CreateSalesReceiptRequest struct {
	CustomerRef           Reference            `json:"CustomerRef"`
	Line                  []SalesReceiptLine   `json:"Line"`
	TxnDate               *Date                `json:"TxnDate,omitempty"`
	DocNumber             string               `json:"DocNumber,omitempty"`
	CurrencyRef           *Reference           `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation `json:"GlobalTaxCalculation,omitempty"`
	ProjectRef            *Reference           `json:"ProjectRef,omitempty"`
	BillEmail             *EmailAddress        `json:"BillEmail,omitempty"`
	CustomField           []NameValue          `json:"CustomField,omitempty"`
	ClassRef              *Reference           `json:"ClassRef,omitempty"`
	PrintStatus           PrintStatus          `json:"PrintStatus,omitempty"`
	PrivateNote           string               `json:"PrivateNote,omitempty"`
	TxnTaxDetail          *TxnTaxDetail        `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef      *Reference           `json:"PaymentMethodRef,omitempty"`
	ExchangeRate          float64              `json:"ExchangeRate,omitempty"`
	DepartmentRef         *Reference           `json:"DepartmentRef,omitempty"`
	DepositToAccountRef   *Reference           `json:"DepositToAccountRef,omitempty"`
	PaymentRefNum         string               `json:"PaymentRefNum,omitempty"`
	CustomerMemo          *InvoiceMemo         `json:"CustomerMemo,omitempty"`
	ApplyTaxAfterDiscount *bool                `json:"ApplyTaxAfterDiscount,omitempty"`
	BillAddr              *PhysicalAddress     `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress     `json:"ShipAddr,omitempty"`
	EmailStatus           EmailStatus          `json:"EmailStatus,omitempty"`
}

// SparseUpdateSalesReceiptRequest represents the documented sparse update payload.
type SparseUpdateSalesReceiptRequest struct {
	ID                    string               `json:"Id"`
	SyncToken             string               `json:"SyncToken"`
	Sparse                bool                 `json:"sparse"`
	CustomerRef           *Reference           `json:"CustomerRef,omitempty"`
	Line                  []SalesReceiptLine   `json:"Line,omitempty"`
	TxnDate               *Date                `json:"TxnDate,omitempty"`
	DocNumber             string               `json:"DocNumber,omitempty"`
	CurrencyRef           *Reference           `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation `json:"GlobalTaxCalculation,omitempty"`
	ProjectRef            *Reference           `json:"ProjectRef,omitempty"`
	BillEmail             *EmailAddress        `json:"BillEmail,omitempty"`
	ClassRef              *Reference           `json:"ClassRef,omitempty"`
	PrintStatus           PrintStatus          `json:"PrintStatus,omitempty"`
	PrivateNote           string               `json:"PrivateNote,omitempty"`
	TxnTaxDetail          *TxnTaxDetail        `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef      *Reference           `json:"PaymentMethodRef,omitempty"`
	ExchangeRate          float64              `json:"ExchangeRate,omitempty"`
	DepartmentRef         *Reference           `json:"DepartmentRef,omitempty"`
	DepositToAccountRef   *Reference           `json:"DepositToAccountRef,omitempty"`
	PaymentRefNum         string               `json:"PaymentRefNum,omitempty"`
	CustomerMemo          *InvoiceMemo         `json:"CustomerMemo,omitempty"`
	ApplyTaxAfterDiscount *bool                `json:"ApplyTaxAfterDiscount,omitempty"`
	BillAddr              *PhysicalAddress     `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress     `json:"ShipAddr,omitempty"`
	EmailStatus           EmailStatus          `json:"EmailStatus,omitempty"`
	Domain                string               `json:"domain,omitempty"`
}

// SalesReceiptLine represents a QuickBooks sales receipt line.
type SalesReceiptLine struct {
	ID                    string                        `json:"Id,omitempty"`
	LineNum               float64                       `json:"LineNum,omitempty"`
	Amount                float64                       `json:"Amount,omitempty"`
	Description           string                        `json:"Description,omitempty"`
	DetailType            SalesReceiptLineDetailType    `json:"DetailType,omitempty"`
	SalesItemLineDetail   *InvoiceSalesItemLineDetail   `json:"SalesItemLineDetail,omitempty"`
	GroupLineDetail       *InvoiceGroupLineDetail       `json:"GroupLineDetail,omitempty"`
	DescriptionLineDetail *InvoiceDescriptionLineDetail `json:"DescriptionLineDetail,omitempty"`
	SubTotalLineDetail    *InvoiceSubTotalLineDetail    `json:"SubTotalLineDetail,omitempty"`
	TaxLineDetail         *TaxLineDetail                `json:"TaxLineDetail,omitempty"`
}

// SalesReceiptDeliveryInfo represents delivery info for a sales receipt.
type SalesReceiptDeliveryInfo struct {
	DeliveryType string    `json:"DeliveryType,omitempty"`
	DeliveryTime time.Time `json:"DeliveryTime,omitempty"`
}
