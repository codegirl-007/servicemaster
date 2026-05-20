// Package types contains transport types for external integrations.
package types

import "time"

// Date represents a date-only value.
type Date struct {
	time.Time
}

// EmailStatus represents the documented QuickBooks email delivery status values.
type EmailStatus string

const (
	EmailStatusNotSet     EmailStatus = "NotSet"
	EmailStatusNeedToSend EmailStatus = "NeedToSend"
	EmailStatusEmailSent  EmailStatus = "EmailSent"
)

// PrintStatus represents the documented QuickBooks print status values.
type PrintStatus string

const (
	PrintStatusNotSet        PrintStatus = "NotSet"
	PrintStatusNeedToPrint   PrintStatus = "NeedToPrint"
	PrintStatusPrintComplete PrintStatus = "PrintComplete"
)

// LinkedTxn represents a transaction linked to another QuickBooks entity.
type LinkedTxn struct {
	TxnID     string `json:"TxnId,omitempty"`
	TxnType   string `json:"TxnType,omitempty"`
	TxnLineID string `json:"TxnLineId,omitempty"`
}

// TxnTaxDetail represents transaction-level tax details.
type TxnTaxDetail struct {
	TxnTaxCodeRef *Reference `json:"TxnTaxCodeRef,omitempty"`
	TotalTax      float64    `json:"TotalTax,omitempty"`
	TaxLine       []TaxLine  `json:"TaxLine,omitempty"`
}

// TaxLine represents a tax line in transaction tax detail
type TaxLine struct {
	Amount        float64        `json:"Amount,omitempty"`
	DetailType    string         `json:"DetailType,omitempty"`
	TaxLineDetail *TaxLineDetail `json:"TaxLineDetail,omitempty"`
}

// TaxLineDetail represents tax details where TaxRateRef is optional.
type TaxLineDetail struct {
	TaxRateRef          *Reference `json:"TaxRateRef,omitempty"`
	NetAmountTaxable    float64    `json:"NetAmountTaxable,omitempty"`
	PercentBased        *bool      `json:"PercentBased,omitempty"`
	TaxInclusiveAmount  float64    `json:"TaxInclusiveAmount,omitempty"`
	OverrideDeltaAmount float64    `json:"OverrideDeltaAmount,omitempty"`
	TaxPercent          float64    `json:"TaxPercent,omitempty"`
}
