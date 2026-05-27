// Package types contains transport types for external integrations.
package types

import "time"

// TransferResponse represents the QuickBooks transfer response envelope.
type TransferResponse struct {
	Transfer Transfer  `json:"Transfer"`
	Time     time.Time `json:"time"`
}

// Transfer represents a QuickBooks transfer object.
type Transfer struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken,omitempty"`
	FromAccountRef          Reference               `json:"FromAccountRef"`
	ToAccountRef            Reference               `json:"ToAccountRef"`
	Amount                  float64                 `json:"Amount"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	RecurDataRef            *Reference              `json:"RecurDataRef,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	MetaData                *MetaData               `json:"MetaData,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
	Sparse                  *bool                   `json:"sparse,omitempty"`
}

// CreateTransferRequest represents the documented create transfer payload.
type CreateTransferRequest struct {
	FromAccountRef Reference `json:"FromAccountRef"`
	ToAccountRef   Reference `json:"ToAccountRef"`
	Amount         float64   `json:"Amount"`
	TxnDate        *Date     `json:"TxnDate,omitempty"`
	PrivateNote    string    `json:"PrivateNote,omitempty"`
}

// UpdateTransferRequest represents the documented full update transfer payload.
type UpdateTransferRequest struct {
	ID             string    `json:"Id"`
	SyncToken      string    `json:"SyncToken"`
	FromAccountRef Reference `json:"FromAccountRef"`
	ToAccountRef   Reference `json:"ToAccountRef"`
	Amount         float64   `json:"Amount"`
	TxnDate        *Date     `json:"TxnDate,omitempty"`
	PrivateNote    string    `json:"PrivateNote,omitempty"`
	Domain         string    `json:"domain,omitempty"`
}

// SparseUpdateTransferRequest represents the documented sparse update transfer payload.
type SparseUpdateTransferRequest struct {
	ID             string     `json:"Id"`
	SyncToken      string     `json:"SyncToken"`
	Sparse         bool       `json:"sparse"`
	FromAccountRef *Reference `json:"FromAccountRef,omitempty"`
	ToAccountRef   *Reference `json:"ToAccountRef,omitempty"`
	Amount         float64    `json:"Amount,omitempty"`
	TxnDate        *Date      `json:"TxnDate,omitempty"`
	PrivateNote    string     `json:"PrivateNote,omitempty"`
	Domain         string     `json:"domain,omitempty"`
}

// DeleteTransferRequest is the full transfer object returned from read.
// QBO requires the complete transfer payload for delete, not just Id and SyncToken.
type DeleteTransferRequest = Transfer

// TransferDeleteResponse represents the QuickBooks deleted transfer response envelope.
type TransferDeleteResponse struct {
	Transfer DeletedEntity `json:"Transfer"`
	Time     time.Time     `json:"time"`
}
