// Package types contains transport types for external integrations.
package types

import "time"

// PurchaseOrderStatus represents the documented purchase order statuses.
type PurchaseOrderStatus string

const (
	PurchaseOrderStatusOpen   PurchaseOrderStatus = "Open"
	PurchaseOrderStatusClosed PurchaseOrderStatus = "Closed"
)

// PurchaseOrderResponse represents the QuickBooks purchase order response envelope.
type PurchaseOrderResponse struct {
	PurchaseOrder PurchaseOrder `json:"PurchaseOrder"`
	Time          time.Time     `json:"time"`
}

// PurchaseOrder represents a QuickBooks purchase order object.
type PurchaseOrder struct {
	ID                      string                  `json:"Id"`
	APAccountRef            Reference               `json:"APAccountRef"`
	VendorRef               Reference               `json:"VendorRef"`
	Line                    []BillLine              `json:"Line"`
	SyncToken               string                  `json:"SyncToken,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TotalAmt                float64                 `json:"TotalAmt,omitempty"`
	RecurDataRef            *Reference              `json:"RecurDataRef,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	CustomField             []NameValue             `json:"CustomField,omitempty"`
	POEmail                 *EmailAddress           `json:"POEmail,omitempty"`
	ClassRef                *Reference              `json:"ClassRef,omitempty"`
	SalesTermRef            *Reference              `json:"SalesTermRef,omitempty"`
	LinkedTxn               []LinkedTxn             `json:"LinkedTxn,omitempty"`
	Memo                    string                  `json:"Memo,omitempty"`
	POStatus                PurchaseOrderStatus     `json:"POStatus,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	DueDate                 *Date                   `json:"DueDate,omitempty"`
	MetaData                *MetaData               `json:"MetaData,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	ShipMethodRef           *Reference              `json:"ShipMethodRef,omitempty"`
	TxnTaxDetail            *BillTxnTaxDetail       `json:"TxnTaxDetail,omitempty"`
	ShipTo                  *Reference              `json:"ShipTo,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	ShipAddr                *PhysicalAddress        `json:"ShipAddr,omitempty"`
	VendorAddr              *PhysicalAddress        `json:"VendorAddr,omitempty"`
	EmailStatus             EmailStatus             `json:"EmailStatus,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
	Sparse                  *bool                   `json:"sparse,omitempty"`
}

// CreatePurchaseOrderRequest represents the documented create purchase order payload.
type CreatePurchaseOrderRequest struct {
	APAccountRef            Reference               `json:"APAccountRef"`
	VendorRef               Reference               `json:"VendorRef"`
	Line                    []BillLine              `json:"Line"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	CustomField             []NameValue             `json:"CustomField,omitempty"`
	POEmail                 *EmailAddress           `json:"POEmail,omitempty"`
	ClassRef                *Reference              `json:"ClassRef,omitempty"`
	SalesTermRef            *Reference              `json:"SalesTermRef,omitempty"`
	LinkedTxn               []LinkedTxn             `json:"LinkedTxn,omitempty"`
	Memo                    string                  `json:"Memo,omitempty"`
	POStatus                PurchaseOrderStatus     `json:"POStatus,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	DueDate                 *Date                   `json:"DueDate,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	ShipMethodRef           *Reference              `json:"ShipMethodRef,omitempty"`
	TxnTaxDetail            *BillTxnTaxDetail       `json:"TxnTaxDetail,omitempty"`
	ShipTo                  *Reference              `json:"ShipTo,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	ShipAddr                *PhysicalAddress        `json:"ShipAddr,omitempty"`
	VendorAddr              *PhysicalAddress        `json:"VendorAddr,omitempty"`
	EmailStatus             EmailStatus             `json:"EmailStatus,omitempty"`
}

// SparseUpdatePurchaseOrderRequest represents the documented sparse update payload.
type SparseUpdatePurchaseOrderRequest struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken"`
	Sparse                  bool                    `json:"sparse"`
	APAccountRef            *Reference              `json:"APAccountRef,omitempty"`
	VendorRef               *Reference              `json:"VendorRef,omitempty"`
	Line                    []BillLine              `json:"Line,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	POEmail                 *EmailAddress           `json:"POEmail,omitempty"`
	ClassRef                *Reference              `json:"ClassRef,omitempty"`
	SalesTermRef            *Reference              `json:"SalesTermRef,omitempty"`
	LinkedTxn               []LinkedTxn             `json:"LinkedTxn,omitempty"`
	Memo                    string                  `json:"Memo,omitempty"`
	POStatus                PurchaseOrderStatus     `json:"POStatus,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	DueDate                 *Date                   `json:"DueDate,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	ShipMethodRef           *Reference              `json:"ShipMethodRef,omitempty"`
	TxnTaxDetail            *BillTxnTaxDetail       `json:"TxnTaxDetail,omitempty"`
	ShipTo                  *Reference              `json:"ShipTo,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	ShipAddr                *PhysicalAddress        `json:"ShipAddr,omitempty"`
	VendorAddr              *PhysicalAddress        `json:"VendorAddr,omitempty"`
	EmailStatus             EmailStatus             `json:"EmailStatus,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
}
