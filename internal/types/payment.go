// Package types contains transport types for external integrations.
package types

import "time"

// PaymentLinkedTxnType represents the documented linked transaction types for payments.
type PaymentLinkedTxnType string

const (
	PaymentLinkedTxnTypeInvoice          PaymentLinkedTxnType = "Invoice"
	PaymentLinkedTxnTypeCreditMemo       PaymentLinkedTxnType = "CreditMemo"
	PaymentLinkedTxnTypeJournalEntry     PaymentLinkedTxnType = "JournalEntry"
	PaymentLinkedTxnTypeExpense          PaymentLinkedTxnType = "Expense"
	PaymentLinkedTxnTypeCheck            PaymentLinkedTxnType = "Check"
	PaymentLinkedTxnTypeCreditCardCredit PaymentLinkedTxnType = "CreditCardCredit"
)

// CCPaymentStatus represents the documented credit card payment status values.
type CCPaymentStatus string

const (
	CCPaymentStatusCompleted CCPaymentStatus = "Completed"
	CCPaymentStatusUnknown   CCPaymentStatus = "Unknown"
)

// PaymentResponse represents the QuickBooks payment response envelope.
type PaymentResponse struct {
	Payment Payment   `json:"Payment"`
	Time    time.Time `json:"time"`
}

// Payment represents a QuickBooks payment object.
type Payment struct {
	ID                  string                    `json:"Id"`
	SyncToken           string                    `json:"SyncToken"`
	TotalAmt            float64                   `json:"TotalAmt,omitempty"`
	TxnDate             *Date                     `json:"TxnDate,omitempty"`
	CustomerRef         *Reference                `json:"CustomerRef,omitempty"`
	Line                []PaymentLine             `json:"Line,omitempty"`
	PaymentMethodRef    *Reference                `json:"PaymentMethodRef,omitempty"`
	DepositToAccountRef *Reference                `json:"DepositToAccountRef,omitempty"`
	PaymentRefNum       string                    `json:"PaymentRefNum,omitempty"`
	PrivateNote         string                    `json:"PrivateNote,omitempty"`
	CurrencyRef         *Reference                `json:"CurrencyRef,omitempty"`
	ExchangeRate        float64                   `json:"ExchangeRate,omitempty"`
	UnappliedAmt        float64                   `json:"UnappliedAmt,omitempty"`
	ProjectRef          *Reference                `json:"ProjectRef,omitempty"`
	CreditCardPayment   *PaymentCreditCardPayment `json:"CreditCardPayment,omitempty"`
	TxnSource           string                    `json:"TxnSource,omitempty"`
	MetaData            *MetaData                 `json:"MetaData,omitempty"`
	Domain              string                    `json:"domain,omitempty"`
	Sparse              *bool                     `json:"sparse,omitempty"`
}

// CreatePaymentRequest represents the documented create payment payload.
type CreatePaymentRequest struct {
	// TotalAmt is required.
	TotalAmt float64 `json:"TotalAmt"`
	// CustomerRef is required.
	CustomerRef         Reference                 `json:"CustomerRef"`
	TxnDate             *Date                     `json:"TxnDate,omitempty"`
	Line                []PaymentLine             `json:"Line,omitempty"`
	PaymentMethodRef    *Reference                `json:"PaymentMethodRef,omitempty"`
	DepositToAccountRef *Reference                `json:"DepositToAccountRef,omitempty"`
	PaymentRefNum       string                    `json:"PaymentRefNum,omitempty"`
	PrivateNote         string                    `json:"PrivateNote,omitempty"`
	CurrencyRef         *Reference                `json:"CurrencyRef,omitempty"`
	ExchangeRate        float64                   `json:"ExchangeRate,omitempty"`
	ProjectRef          *Reference                `json:"ProjectRef,omitempty"`
	CreditCardPayment   *PaymentCreditCardPayment `json:"CreditCardPayment,omitempty"`
	TxnSource           string                    `json:"TxnSource,omitempty"`
}

// SparseUpdatePaymentRequest represents the documented sparse update payload.
type SparseUpdatePaymentRequest struct {
	ID                  string                    `json:"Id"`
	SyncToken           string                    `json:"SyncToken"`
	Sparse              bool                      `json:"sparse"`
	TotalAmt            float64                   `json:"TotalAmt,omitempty"`
	TxnDate             *Date                     `json:"TxnDate,omitempty"`
	CustomerRef         *Reference                `json:"CustomerRef,omitempty"`
	Line                []PaymentLine             `json:"Line,omitempty"`
	PaymentMethodRef    *Reference                `json:"PaymentMethodRef,omitempty"`
	DepositToAccountRef *Reference                `json:"DepositToAccountRef,omitempty"`
	PaymentRefNum       string                    `json:"PaymentRefNum,omitempty"`
	PrivateNote         string                    `json:"PrivateNote,omitempty"`
	CurrencyRef         *Reference                `json:"CurrencyRef,omitempty"`
	ExchangeRate        float64                   `json:"ExchangeRate,omitempty"`
	ProjectRef          *Reference                `json:"ProjectRef,omitempty"`
	CreditCardPayment   *PaymentCreditCardPayment `json:"CreditCardPayment,omitempty"`
	TxnSource           string                    `json:"TxnSource,omitempty"`
	Domain              string                    `json:"domain,omitempty"`
}

// PaymentLine represents an applied payment line.
type PaymentLine struct {
	Amount    float64        `json:"Amount,omitempty"`
	LinkedTxn []LinkedTxn    `json:"LinkedTxn,omitempty"`
	LineEx    *PaymentLineEx `json:"LineEx,omitempty"`
}

// PaymentLinkedTxn represents a transaction linked to a payment line.
type PaymentLinkedTxn struct {
	TxnID     string               `json:"TxnId,omitempty"`
	TxnType   PaymentLinkedTxnType `json:"TxnType,omitempty"`
	TxnLineID string               `json:"TxnLineId,omitempty"`
}

// PaymentLineEx represents additional payment line metadata.
type PaymentLineEx struct {
	Any []NameValue `json:"any,omitempty"`
}

// PaymentCreditCardPayment represents Intuit Payments credit card details.
type PaymentCreditCardPayment struct {
	Status CCPaymentStatus `json:"Status,omitempty"`
}
