// Package types contains transport types for external integrations.
package types

import "time"

// PurchasePaymentType represents the documented purchase payment types.
type PurchasePaymentType string

const (
	PurchasePaymentTypeCash       PurchasePaymentType = "Cash"
	PurchasePaymentTypeCheck      PurchasePaymentType = "Check"
	PurchasePaymentTypeCreditCard PurchasePaymentType = "CreditCard"
)

// PurchaseResponse represents the QuickBooks purchase response envelope.
type PurchaseResponse struct {
	Purchase Purchase  `json:"Purchase"`
	Time     time.Time `json:"time"`
}

// Purchase represents a QuickBooks purchase object.
type Purchase struct {
	ID                      string                  `json:"Id"`
	Line                    []BillLine              `json:"Line,omitempty"`
	PaymentType             PurchasePaymentType     `json:"PaymentType,omitempty"`
	AccountRef              *Reference              `json:"AccountRef,omitempty"`
	SyncToken               string                  `json:"SyncToken,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	TotalAmt                float64                 `json:"TotalAmt,omitempty"`
	RecurDataRef            *Reference              `json:"RecurDataRef,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	PrintStatus             PrintStatus             `json:"PrintStatus,omitempty"`
	RemitToAddr             *PhysicalAddress        `json:"RemitToAddr,omitempty"`
	TxnSource               string                  `json:"TxnSource,omitempty"`
	LinkedTxn               []LinkedTxn             `json:"LinkedTxn,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	MetaData                *MetaData               `json:"MetaData,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	TxnTaxDetail            *BillTxnTaxDetail       `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef        *Reference              `json:"PaymentMethodRef,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	EntityRef               *PurchaseEntityRef      `json:"EntityRef,omitempty"`
	IncludeInAnnualTPAR     *bool                   `json:"IncludeInAnnualTPAR,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
	Sparse                  *bool                   `json:"sparse,omitempty"`
}

// CreatePurchaseRequest represents the documented create purchase payload.
type CreatePurchaseRequest struct {
	PaymentType             PurchasePaymentType     `json:"PaymentType"`
	AccountRef              Reference               `json:"AccountRef"`
	Line                    []BillLine              `json:"Line"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	PrintStatus             PrintStatus             `json:"PrintStatus,omitempty"`
	TxnSource               string                  `json:"TxnSource,omitempty"`
	LinkedTxn               []LinkedTxn             `json:"LinkedTxn,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	TxnTaxDetail            *BillTxnTaxDetail       `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef        *Reference              `json:"PaymentMethodRef,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	EntityRef               *PurchaseEntityRef      `json:"EntityRef,omitempty"`
	IncludeInAnnualTPAR     *bool                   `json:"IncludeInAnnualTPAR,omitempty"`
}

// SparseUpdatePurchaseRequest represents the documented sparse update payload.
type SparseUpdatePurchaseRequest struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken"`
	Sparse                  bool                    `json:"sparse"`
	Line                    []BillLine              `json:"Line,omitempty"`
	PaymentType             PurchasePaymentType     `json:"PaymentType,omitempty"`
	AccountRef              *Reference              `json:"AccountRef,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	PrintStatus             PrintStatus             `json:"PrintStatus,omitempty"`
	TxnSource               string                  `json:"TxnSource,omitempty"`
	LinkedTxn               []LinkedTxn             `json:"LinkedTxn,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	TxnTaxDetail            *BillTxnTaxDetail       `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef        *Reference              `json:"PaymentMethodRef,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	EntityRef               *PurchaseEntityRef      `json:"EntityRef,omitempty"`
	IncludeInAnnualTPAR     *bool                   `json:"IncludeInAnnualTPAR,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
}

// PurchaseEntityRef represents the party associated with a purchase.
type PurchaseEntityRef struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
}
