// Package types contains transport types for external integrations.
package types

import "time"

// CreditCardPaymentResponse represents the QuickBooks credit card payment response envelope.
type CreditCardPaymentResponse struct {
	CreditCardPaymentTxn CreditCardPayment `json:"CreditCardPaymentTxn"`
	Time                 time.Time         `json:"time"`
}

// CreditCardPaymentDeleteResponse represents the QuickBooks deleted credit card payment response envelope.
type CreditCardPaymentDeleteResponse struct {
	CreditCardPaymentTxn DeletedEntity `json:"CreditCardPaymentTxn"`
	Time                 time.Time     `json:"time"`
}

// CreditCardPayment represents a QuickBooks credit card payment object.
type CreditCardPayment struct {
	ID                   string      `json:"Id"`
	CreditCardAccountRef *Reference  `json:"CreditCardAccountRef,omitempty"`
	Amount               float64     `json:"Amount,omitempty"`
	BankAccountRef       *Reference  `json:"BankAccountRef,omitempty"`
	SyncToken            string      `json:"SyncToken,omitempty"`
	PrivateNote          string      `json:"PrivateNote,omitempty"`
	VendorRef            *Reference  `json:"VendorRef,omitempty"`
	TxnDate              *Date       `json:"TxnDate,omitempty"`
	Memo                 string      `json:"Memo,omitempty"`
	PrintStatus          PrintStatus `json:"PrintStatus,omitempty"`
	CheckNum             string      `json:"CheckNum,omitempty"`
	MetaData             *MetaData   `json:"MetaData,omitempty"`
	CurrencyRef          *Reference  `json:"CurrencyRef,omitempty"`
	Domain               string      `json:"domain,omitempty"`
	Sparse               *bool       `json:"sparse,omitempty"`
}

// CreateCreditCardPaymentRequest represents the documented create credit card payment payload.
type CreateCreditCardPaymentRequest struct {
	TxnDate              *Date     `json:"TxnDate,omitempty"`
	Amount               float64   `json:"Amount"`
	BankAccountRef       Reference `json:"BankAccountRef"`
	CreditCardAccountRef Reference `json:"CreditCardAccountRef"`
	PrivateNote          string    `json:"PrivateNote,omitempty"`
}

// SparseUpdateCreditCardPaymentRequest represents the documented sparse update payload.
type SparseUpdateCreditCardPaymentRequest struct {
	ID                   string      `json:"Id"`
	SyncToken            string      `json:"SyncToken"`
	Sparse               bool        `json:"sparse"`
	TxnDate              *Date       `json:"TxnDate,omitempty"`
	Amount               float64     `json:"Amount,omitempty"`
	BankAccountRef       *Reference  `json:"BankAccountRef,omitempty"`
	CreditCardAccountRef *Reference  `json:"CreditCardAccountRef,omitempty"`
	PrivateNote          string      `json:"PrivateNote,omitempty"`
	VendorRef            *Reference  `json:"VendorRef,omitempty"`
	Memo                 string      `json:"Memo,omitempty"`
	PrintStatus          PrintStatus `json:"PrintStatus,omitempty"`
	CheckNum             string      `json:"CheckNum,omitempty"`
	Domain               string      `json:"domain,omitempty"`
}

// DeletedEntity represents a deleted QBO entity response body.
type DeletedEntity struct {
	Status string `json:"status,omitempty"`
	Domain string `json:"domain,omitempty"`
	ID     string `json:"Id,omitempty"`
}
