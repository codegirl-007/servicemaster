// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// DepositLineDetailType represents the documented QuickBooks deposit line detail types.
type DepositLineDetailType string

const (
	DepositLineDetailTypeDepositLine DepositLineDetailType = "DepositLineDetail"
)

// DepositLinkedTxnType represents the documented linked transaction types for deposits.
type DepositLinkedTxnType string

const (
	DepositLinkedTxnTypeTransfer      DepositLinkedTxnType = "Transfer"
	DepositLinkedTxnTypePayment       DepositLinkedTxnType = "Payment"
	DepositLinkedTxnTypeSalesReceipt  DepositLinkedTxnType = "SalesReceipt"
	DepositLinkedTxnTypeRefundReceipt DepositLinkedTxnType = "RefundReceipt"
	DepositLinkedTxnTypeJournalEntry  DepositLinkedTxnType = "JournalEntry"
)

// TaxApplicableOn represents whether tax on a deposit line applies to sales or purchase.
type TaxApplicableOn string

const (
	TaxApplicableOnSales    TaxApplicableOn = "Sales"
	TaxApplicableOnPurchase TaxApplicableOn = "Purchase"
)

// DepositResponse represents the QuickBooks deposit response envelope.
type DepositResponse struct {
	Deposit Deposit   `json:"Deposit"`
	Time    time.Time `json:"time"`
}

// Deposit represents a QuickBooks deposit object.
type Deposit struct {
	ID                      string                  `json:"Id"`
	DepositToAccountRef     Reference               `json:"DepositToAccountRef"`
	Line                    []DepositLine           `json:"Line"`
	SyncToken               string                  `json:"SyncToken,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	RecurDataRef            *Reference              `json:"RecurDataRef,omitempty"`
	TotalAmt                float64                 `json:"TotalAmt,omitempty"`
	HomeTotalAmt            float64                 `json:"HomeTotalAmt,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	TxnSource               string                  `json:"TxnSource,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	CashBack                *DepositCashBackInfo    `json:"CashBack,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	TxnTaxDetail            *TxnTaxDetail           `json:"TxnTaxDetail,omitempty"`
	MetaData                *MetaData               `json:"MetaData,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
	Sparse                  *bool                   `json:"sparse,omitempty"`
}

// CreateDepositRequest represents the documented create deposit payload.
type CreateDepositRequest struct {
	DepositToAccountRef     Reference               `json:"DepositToAccountRef"`
	Line                    []DepositLine           `json:"Line"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail            *TxnTaxDetail           `json:"TxnTaxDetail,omitempty"`
	CashBack                *DepositCashBackInfo    `json:"CashBack,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	TxnSource               string                  `json:"TxnSource,omitempty"`
}

// UpdateDepositRequest represents the documented full update deposit payload.
type UpdateDepositRequest struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken"`
	DepositToAccountRef     Reference               `json:"DepositToAccountRef"`
	Line                    []DepositLine           `json:"Line"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail            *TxnTaxDetail           `json:"TxnTaxDetail,omitempty"`
	CashBack                *DepositCashBackInfo    `json:"CashBack,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	TxnSource               string                  `json:"TxnSource,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
}

// SparseUpdateDepositRequest represents the documented sparse update deposit payload.
type SparseUpdateDepositRequest struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken"`
	Sparse                  bool                    `json:"sparse"`
	DepositToAccountRef     *Reference              `json:"DepositToAccountRef,omitempty"`
	Line                    []DepositLine           `json:"Line,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail            *TxnTaxDetail           `json:"TxnTaxDetail,omitempty"`
	CashBack                *DepositCashBackInfo    `json:"CashBack,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	TxnSource               string                  `json:"TxnSource,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
}

// DeleteDepositRequest represents the documented delete deposit payload.
type DeleteDepositRequest struct {
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
}

// DepositDeleteResponse represents the QuickBooks deleted deposit response envelope.
type DepositDeleteResponse struct {
	Deposit DeletedEntity `json:"Deposit"`
	Time    time.Time     `json:"time"`
}

// DepositLine represents a deposit line.
type DepositLine struct {
	ID                string                `json:"Id,omitempty"`
	Amount            float64               `json:"Amount,omitempty"`
	DetailType        DepositLineDetailType `json:"DetailType,omitempty"`
	LinkedTxn         []DepositLinkedTxn    `json:"LinkedTxn,omitempty"`
	DepositLineDetail *DepositLineDetail    `json:"DepositLineDetail,omitempty"`
	Description       string                `json:"Description,omitempty"`
	LineNum           float64               `json:"LineNum,omitempty"`
	CustomField       []CustomField         `json:"CustomField,omitempty"`
	ProjectRef        *Reference            `json:"ProjectRef,omitempty"`
}

// DepositLinkedTxn represents a transaction linked to a deposit line.
type DepositLinkedTxn struct {
	TxnID     string               `json:"TxnId,omitempty"`
	TxnType   DepositLinkedTxnType `json:"TxnType,omitempty"`
	TxnLineID string               `json:"TxnLineId,omitempty"`
}

// DepositLineDetail represents direct deposit line details.
type DepositLineDetail struct {
	AccountRef       Reference       `json:"AccountRef"`
	PaymentMethodRef *Reference      `json:"PaymentMethodRef,omitempty"`
	ClassRef         *Reference      `json:"ClassRef,omitempty"`
	CheckNum         string          `json:"CheckNum,omitempty"`
	TaxCodeRef       *Reference      `json:"TaxCodeRef,omitempty"`
	TaxApplicableOn  TaxApplicableOn `json:"TaxApplicableOn,omitempty"`
	TxnType          string          `json:"TxnType,omitempty"`
	Entity           *Reference      `json:"Entity,omitempty"`
}

// DepositCashBackInfo represents cash back recorded on a deposit.
type DepositCashBackInfo struct {
	AccountRef Reference `json:"AccountRef"`
	Amount     float64   `json:"Amount"`
	Memo       string    `json:"Memo,omitempty"`
}
