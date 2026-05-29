// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// VendorCreditResponse represents the QuickBooks vendor credit response envelope.
type VendorCreditResponse struct {
	VendorCredit VendorCredit `json:"VendorCredit"`
	Time         time.Time    `json:"time"`
}

// VendorCredit represents a QuickBooks vendor credit object.
type VendorCredit struct {
	ID                      string                  `json:"Id"`
	VendorRef               Reference               `json:"VendorRef"`
	Line                    []BillLine              `json:"Line"`
	SyncToken               string                  `json:"SyncToken,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	RecurDataRef            *Reference              `json:"RecurDataRef,omitempty"`
	TotalAmt                float64                 `json:"TotalAmt,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	LinkedTxn               []LinkedTxn             `json:"LinkedTxn,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	APAccountRef            *Reference              `json:"APAccountRef,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	IncludeInAnnualTPAR     *bool                   `json:"IncludeInAnnualTPAR,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	Balance                 float64                 `json:"Balance,omitempty"`
	MetaData                *MetaData               `json:"MetaData,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
	Sparse                  *bool                   `json:"sparse,omitempty"`
}

// CreateVendorCreditRequest represents the documented create vendor credit payload.
type CreateVendorCreditRequest struct {
	VendorRef               Reference               `json:"VendorRef"`
	Line                    []BillLine              `json:"Line"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	LinkedTxn               []LinkedTxn             `json:"LinkedTxn,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	APAccountRef            *Reference              `json:"APAccountRef,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	IncludeInAnnualTPAR     *bool                   `json:"IncludeInAnnualTPAR,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
}

// SparseUpdateVendorCreditRequest represents the documented sparse update payload.
type SparseUpdateVendorCreditRequest struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken"`
	Sparse                  bool                    `json:"sparse"`
	VendorRef               *Reference              `json:"VendorRef,omitempty"`
	Line                    []BillLine              `json:"Line,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	LinkedTxn               []LinkedTxn             `json:"LinkedTxn,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	APAccountRef            *Reference              `json:"APAccountRef,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	IncludeInAnnualTPAR     *bool                   `json:"IncludeInAnnualTPAR,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
}
