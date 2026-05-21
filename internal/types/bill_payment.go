// Package types contains transport types for external integrations.
package types

import "time"

// BillPaymentType represents the documented bill payment types.
type BillPaymentType string

const (
	BillPaymentTypeCheck      BillPaymentType = "Check"
	BillPaymentTypeCreditCard BillPaymentType = "CreditCard"
)

// BillPaymentResponse represents the QuickBooks bill payment response envelope.
type BillPaymentResponse struct {
	BillPayment BillPayment `json:"BillPayment"`
	Time        time.Time   `json:"time"`
}

// BillPayment represents a QuickBooks bill payment object.
type BillPayment struct {
	ID                      string                  `json:"Id"`
	VendorRef               Reference               `json:"VendorRef"`
	Line                    []BillPaymentLine       `json:"Line"`
	TotalAmt                float64                 `json:"TotalAmt"`
	PayType                 BillPaymentType         `json:"PayType"`
	SyncToken               string                  `json:"SyncToken,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	CheckPayment            *BillPaymentCheck       `json:"CheckPayment,omitempty"`
	CreditCardPayment       *BillPaymentCreditCard  `json:"CreditCardPayment,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	APAccountRef            *Reference              `json:"APAccountRef,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	ProcessBillPayment      *bool                   `json:"ProcessBillPayment,omitempty"`
	MetaData                *MetaData               `json:"MetaData,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
	Sparse                  *bool                   `json:"sparse,omitempty"`
}

// CreateBillPaymentRequest represents the documented create bill payment payload.
type CreateBillPaymentRequest struct {
	VendorRef               Reference               `json:"VendorRef"`
	Line                    []BillPaymentLine       `json:"Line"`
	TotalAmt                float64                 `json:"TotalAmt"`
	PayType                 BillPaymentType         `json:"PayType"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	CheckPayment            *BillPaymentCheck       `json:"CheckPayment,omitempty"`
	CreditCardPayment       *BillPaymentCreditCard  `json:"CreditCardPayment,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	APAccountRef            *Reference              `json:"APAccountRef,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	ProcessBillPayment      *bool                   `json:"ProcessBillPayment,omitempty"`
}

// BillPaymentLine represents an applied bill payment line.
type BillPaymentLine struct {
	Amount    float64     `json:"Amount"`
	LinkedTxn []LinkedTxn `json:"LinkedTxn"`
}

// BillPaymentCheck represents check payment details for a bill payment.
type BillPaymentCheck struct {
	BankAccountRef Reference   `json:"BankAccountRef"`
	PrintStatus    PrintStatus `json:"PrintStatus,omitempty"`
}

// BillPaymentCreditCard represents credit card payment details for a bill payment.
type BillPaymentCreditCard struct {
	CCAccountRef Reference `json:"CCAccountRef"`
}

// VoidBillPaymentRequest represents the documented void bill payment payload.
type VoidBillPaymentRequest struct {
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
	Sparse    bool   `json:"sparse"`
}

// VoidBillPaymentResponse represents the QuickBooks voided bill payment response envelope.
type VoidBillPaymentResponse struct {
	BillPayment BillPayment `json:"BillPayment"`
	Time        time.Time   `json:"time"`
}

// DeleteBillPaymentRequest represents the documented delete bill payment payload.
type DeleteBillPaymentRequest struct {
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
}

// BillPaymentDeleteResponse represents the QuickBooks deleted bill payment response envelope.
type BillPaymentDeleteResponse struct {
	BillPayment DeletedEntity `json:"BillPayment"`
	Time        time.Time     `json:"time"`
}
