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
