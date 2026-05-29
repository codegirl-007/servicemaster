// Package types contains transport types for external integrations.
package types

import "time"

// RefundReceiptPaymentType represents the documented refund receipt payment types.
type RefundReceiptPaymentType string

const (
	RefundReceiptPaymentTypeCheck      RefundReceiptPaymentType = "Check"
	RefundReceiptPaymentTypeCreditCard RefundReceiptPaymentType = "CreditCard"
)

// RefundReceiptResponse represents the QuickBooks refund receipt response envelope.
type RefundReceiptResponse struct {
	RefundReceipt RefundReceipt `json:"RefundReceipt"`
	Time          time.Time     `json:"time"`
}

// RefundReceipt represents a QuickBooks refund receipt object.
type RefundReceipt struct {
	ID                      string                     `json:"Id"`
	SyncToken               string                     `json:"SyncToken,omitempty"`
	DepositToAccountRef     Reference                  `json:"DepositToAccountRef"`
	Line                    []SalesReceiptLine         `json:"Line,omitempty"`
	CustomerRef             *Reference                 `json:"CustomerRef,omitempty"`
	TxnDate                 *Date                      `json:"TxnDate,omitempty"`
	DocNumber               string                     `json:"DocNumber,omitempty"`
	CurrencyRef             *Reference                 `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation       `json:"GlobalTaxCalculation,omitempty"`
	ProjectRef              *Reference                 `json:"ProjectRef,omitempty"`
	BillEmail               *EmailAddress              `json:"BillEmail,omitempty"`
	HomeBalance             float64                    `json:"HomeBalance,omitempty"`
	RecurDataRef            *Reference                 `json:"RecurDataRef,omitempty"`
	TotalAmt                float64                    `json:"TotalAmt,omitempty"`
	TaxExemptionRef         *Reference                 `json:"TaxExemptionRef,omitempty"`
	Balance                 float64                    `json:"Balance,omitempty"`
	HomeTotalAmt            float64                    `json:"HomeTotalAmt,omitempty"`
	CustomField             []CustomField              `json:"CustomField,omitempty"`
	ClassRef                *Reference                 `json:"ClassRef,omitempty"`
	PrintStatus             PrintStatus                `json:"PrintStatus,omitempty"`
	PaymentType             RefundReceiptPaymentType   `json:"PaymentType,omitempty"`
	CheckPayment            *RefundReceiptCheckPayment `json:"CheckPayment,omitempty"`
	CreditCardPayment       *PaymentCreditCardPayment  `json:"CreditCardPayment,omitempty"`
	TxnSource               string                     `json:"TxnSource,omitempty"`
	TransactionLocationType TransactionLocationType    `json:"TransactionLocationType,omitempty"`
	PrivateNote             string                     `json:"PrivateNote,omitempty"`
	CustomerMemo            *Memo                      `json:"CustomerMemo,omitempty"`
	TxnTaxDetail            *TxnTaxDetail              `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef        *Reference                 `json:"PaymentMethodRef,omitempty"`
	ExchangeRate            float64                    `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Reference                 `json:"DepartmentRef,omitempty"`
	PaymentRefNum           string                     `json:"PaymentRefNum,omitempty"`
	ApplyTaxAfterDiscount   *bool                      `json:"ApplyTaxAfterDiscount,omitempty"`
	BillAddr                *PhysicalAddress           `json:"BillAddr,omitempty"`
	ShipAddr                *PhysicalAddress           `json:"ShipAddr,omitempty"`
	EmailStatus             EmailStatus                `json:"EmailStatus,omitempty"`
	MetaData                *MetaData                  `json:"MetaData,omitempty"`
	Domain                  string                     `json:"domain,omitempty"`
	Sparse                  *bool                      `json:"sparse,omitempty"`
}

// CreateRefundReceiptRequest represents the documented create refund receipt payload.
type CreateRefundReceiptRequest struct {
	DepositToAccountRef   Reference                  `json:"DepositToAccountRef"`
	Line                  []SalesReceiptLine         `json:"Line"`
	CustomerRef           Reference                  `json:"CustomerRef"`
	TxnDate               *Date                      `json:"TxnDate,omitempty"`
	DocNumber             string                     `json:"DocNumber,omitempty"`
	CurrencyRef           *Reference                 `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation       `json:"GlobalTaxCalculation,omitempty"`
	ProjectRef            *Reference                 `json:"ProjectRef,omitempty"`
	BillEmail             *EmailAddress              `json:"BillEmail,omitempty"`
	CustomField           []CustomField              `json:"CustomField,omitempty"`
	ClassRef              *Reference                 `json:"ClassRef,omitempty"`
	PrintStatus           PrintStatus                `json:"PrintStatus,omitempty"`
	PaymentType           RefundReceiptPaymentType   `json:"PaymentType,omitempty"`
	CheckPayment          *RefundReceiptCheckPayment `json:"CheckPayment,omitempty"`
	CreditCardPayment     *PaymentCreditCardPayment  `json:"CreditCardPayment,omitempty"`
	PrivateNote           string                     `json:"PrivateNote,omitempty"`
	TxnTaxDetail          *TxnTaxDetail              `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef      *Reference                 `json:"PaymentMethodRef,omitempty"`
	ExchangeRate          float64                    `json:"ExchangeRate,omitempty"`
	DepartmentRef         *Reference                 `json:"DepartmentRef,omitempty"`
	PaymentRefNum         string                     `json:"PaymentRefNum,omitempty"`
	CustomerMemo          *Memo                      `json:"CustomerMemo,omitempty"`
	ApplyTaxAfterDiscount *bool                      `json:"ApplyTaxAfterDiscount,omitempty"`
	BillAddr              *PhysicalAddress           `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress           `json:"ShipAddr,omitempty"`
	EmailStatus           EmailStatus                `json:"EmailStatus,omitempty"`
}

// UpdateRefundReceiptRequest represents the documented full update refund receipt payload.
type UpdateRefundReceiptRequest struct {
	ID                    string                     `json:"Id"`
	SyncToken             string                     `json:"SyncToken"`
	DepositToAccountRef   Reference                  `json:"DepositToAccountRef"`
	Line                  []SalesReceiptLine         `json:"Line"`
	CustomerRef           *Reference                 `json:"CustomerRef,omitempty"`
	TxnDate               *Date                      `json:"TxnDate,omitempty"`
	DocNumber             string                     `json:"DocNumber,omitempty"`
	CurrencyRef           *Reference                 `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation       `json:"GlobalTaxCalculation,omitempty"`
	ProjectRef            *Reference                 `json:"ProjectRef,omitempty"`
	BillEmail             *EmailAddress              `json:"BillEmail,omitempty"`
	CustomField           []CustomField              `json:"CustomField,omitempty"`
	ClassRef              *Reference                 `json:"ClassRef,omitempty"`
	PrintStatus           PrintStatus                `json:"PrintStatus,omitempty"`
	PaymentType           RefundReceiptPaymentType   `json:"PaymentType,omitempty"`
	CheckPayment          *RefundReceiptCheckPayment `json:"CheckPayment,omitempty"`
	CreditCardPayment     *PaymentCreditCardPayment  `json:"CreditCardPayment,omitempty"`
	PrivateNote           string                     `json:"PrivateNote,omitempty"`
	TxnTaxDetail          *TxnTaxDetail              `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef      *Reference                 `json:"PaymentMethodRef,omitempty"`
	ExchangeRate          float64                    `json:"ExchangeRate,omitempty"`
	DepartmentRef         *Reference                 `json:"DepartmentRef,omitempty"`
	PaymentRefNum         string                     `json:"PaymentRefNum,omitempty"`
	CustomerMemo          *Memo                      `json:"CustomerMemo,omitempty"`
	ApplyTaxAfterDiscount *bool                      `json:"ApplyTaxAfterDiscount,omitempty"`
	BillAddr              *PhysicalAddress           `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress           `json:"ShipAddr,omitempty"`
	EmailStatus           EmailStatus                `json:"EmailStatus,omitempty"`
	Domain                string                     `json:"domain,omitempty"`
}

// SparseUpdateRefundReceiptRequest represents the documented sparse update refund receipt payload.
type SparseUpdateRefundReceiptRequest struct {
	ID                    string                     `json:"Id"`
	SyncToken             string                     `json:"SyncToken"`
	Sparse                bool                       `json:"sparse"`
	DepositToAccountRef   *Reference                 `json:"DepositToAccountRef,omitempty"`
	Line                  []SalesReceiptLine         `json:"Line,omitempty"`
	CustomerRef           *Reference                 `json:"CustomerRef,omitempty"`
	TxnDate               *Date                      `json:"TxnDate,omitempty"`
	DocNumber             string                     `json:"DocNumber,omitempty"`
	CurrencyRef           *Reference                 `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation       `json:"GlobalTaxCalculation,omitempty"`
	ProjectRef            *Reference                 `json:"ProjectRef,omitempty"`
	BillEmail             *EmailAddress              `json:"BillEmail,omitempty"`
	ClassRef              *Reference                 `json:"ClassRef,omitempty"`
	PrintStatus           PrintStatus                `json:"PrintStatus,omitempty"`
	PaymentType           RefundReceiptPaymentType   `json:"PaymentType,omitempty"`
	CheckPayment          *RefundReceiptCheckPayment `json:"CheckPayment,omitempty"`
	CreditCardPayment     *PaymentCreditCardPayment  `json:"CreditCardPayment,omitempty"`
	PrivateNote           string                     `json:"PrivateNote,omitempty"`
	TxnTaxDetail          *TxnTaxDetail              `json:"TxnTaxDetail,omitempty"`
	PaymentMethodRef      *Reference                 `json:"PaymentMethodRef,omitempty"`
	ExchangeRate          float64                    `json:"ExchangeRate,omitempty"`
	DepartmentRef         *Reference                 `json:"DepartmentRef,omitempty"`
	PaymentRefNum         string                     `json:"PaymentRefNum,omitempty"`
	CustomerMemo          *Memo                      `json:"CustomerMemo,omitempty"`
	ApplyTaxAfterDiscount *bool                      `json:"ApplyTaxAfterDiscount,omitempty"`
	BillAddr              *PhysicalAddress           `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress           `json:"ShipAddr,omitempty"`
	EmailStatus           EmailStatus                `json:"EmailStatus,omitempty"`
	Domain                string                     `json:"domain,omitempty"`
}

// DeleteRefundReceiptRequest represents the documented delete refund receipt payload.
type DeleteRefundReceiptRequest struct {
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
}

// RefundReceiptDeleteResponse represents the QuickBooks deleted refund receipt response envelope.
type RefundReceiptDeleteResponse struct {
	RefundReceipt DeletedEntity `json:"RefundReceipt"`
	Time          time.Time     `json:"time"`
}

// RefundReceiptCheckPayment represents check payment details on a refund receipt.
type RefundReceiptCheckPayment struct {
	CheckNum string `json:"CheckNum,omitempty"`
	Status   string `json:"Status,omitempty"`
}
