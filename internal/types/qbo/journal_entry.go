// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// JournalEntryLineDetailType represents the documented journal entry line detail types.
type JournalEntryLineDetailType string

const (
	JournalEntryLineDetailTypeJournalEntry JournalEntryLineDetailType = "JournalEntryLineDetail"
	JournalEntryLineDetailTypeDescription  JournalEntryLineDetailType = "DescriptionOnly"
)

// PostingType represents debit or credit on a journal entry line.
type PostingType string

const (
	PostingTypeDebit  PostingType = "Debit"
	PostingTypeCredit PostingType = "Credit"
)

// JournalEntryEntityType represents the documented journal entry entity types.
type JournalEntryEntityType string

const (
	JournalEntryEntityTypeVendor   JournalEntryEntityType = "Vendor"
	JournalEntryEntityTypeEmployee JournalEntryEntityType = "Employee"
	JournalEntryEntityTypeCustomer JournalEntryEntityType = "Customer"
)

// JournalEntryResponse represents the QuickBooks journal entry response envelope.
type JournalEntryResponse struct {
	JournalEntry JournalEntry `json:"JournalEntry"`
	Time         time.Time    `json:"time"`
}

// JournalEntry represents a QuickBooks journal entry object.
type JournalEntry struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken,omitempty"`
	Line                    []JournalEntryLine      `json:"Line"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail            *TxnTaxDetail           `json:"TxnTaxDetail,omitempty"`
	Adjustment              *bool                   `json:"Adjustment,omitempty"`
	TaxRateRef              *Reference              `json:"TaxRateRef,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	RecurDataRef            *Reference              `json:"RecurDataRef,omitempty"`
	TotalAmt                float64                 `json:"TotalAmt,omitempty"`
	HomeTotalAmt            float64                 `json:"HomeTotalAmt,omitempty"`
	MetaData                *MetaData               `json:"MetaData,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
	Sparse                  *bool                   `json:"sparse,omitempty"`
}

// CreateJournalEntryRequest represents the documented create journal entry payload.
type CreateJournalEntryRequest struct {
	Line                 []JournalEntryLine   `json:"Line"`
	TxnDate              *Date                `json:"TxnDate,omitempty"`
	DocNumber            string               `json:"DocNumber,omitempty"`
	PrivateNote          string               `json:"PrivateNote,omitempty"`
	CurrencyRef          *Reference           `json:"CurrencyRef,omitempty"`
	ExchangeRate         float64              `json:"ExchangeRate,omitempty"`
	GlobalTaxCalculation GlobalTaxCalculation `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail         *TxnTaxDetail        `json:"TxnTaxDetail,omitempty"`
	Adjustment           *bool                `json:"Adjustment,omitempty"`
	TaxRateRef           *Reference           `json:"TaxRateRef,omitempty"`
}

// UpdateJournalEntryRequest represents the documented full update journal entry payload.
type UpdateJournalEntryRequest struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken"`
	Line                    []JournalEntryLine      `json:"Line"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail            *TxnTaxDetail           `json:"TxnTaxDetail,omitempty"`
	Adjustment              *bool                   `json:"Adjustment,omitempty"`
	TaxRateRef              *Reference              `json:"TaxRateRef,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
}

// SparseUpdateJournalEntryRequest represents the documented sparse update journal entry payload.
type SparseUpdateJournalEntryRequest struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken"`
	Sparse                  bool                    `json:"sparse"`
	Line                    []JournalEntryLine      `json:"Line,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail            *TxnTaxDetail           `json:"TxnTaxDetail,omitempty"`
	Adjustment              *bool                   `json:"Adjustment,omitempty"`
	TaxRateRef              *Reference              `json:"TaxRateRef,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
}

// DeleteJournalEntryRequest represents the documented delete journal entry payload.
type DeleteJournalEntryRequest struct {
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
}

// JournalEntryDeleteResponse represents the QuickBooks deleted journal entry response envelope.
type JournalEntryDeleteResponse struct {
	JournalEntry DeletedEntity `json:"JournalEntry"`
	Time         time.Time     `json:"time"`
}

// JournalEntryLine represents a journal entry line.
type JournalEntryLine struct {
	ID                     string                     `json:"Id,omitempty"`
	LineNum                float64                    `json:"LineNum,omitempty"`
	Amount                 float64                    `json:"Amount,omitempty"`
	Description            string                     `json:"Description,omitempty"`
	DetailType             JournalEntryLineDetailType `json:"DetailType,omitempty"`
	JournalEntryLineDetail *JournalEntryLineDetail    `json:"JournalEntryLineDetail,omitempty"`
	DescriptionLineDetail  *DescriptionLineDetail     `json:"DescriptionLineDetail,omitempty"`
	ProjectRef             *Reference                 `json:"ProjectRef,omitempty"`
}

// JournalEntryLineDetail represents journal entry line details.
type JournalEntryLineDetail struct {
	PostingType     PostingType         `json:"PostingType,omitempty"`
	AccountRef      Reference           `json:"AccountRef"`
	Entity          *JournalEntryEntity `json:"Entity,omitempty"`
	TaxApplicableOn TaxApplicableOn     `json:"TaxApplicableOn,omitempty"`
	TaxAmount       float64             `json:"TaxAmount,omitempty"`
	TaxInclusiveAmt float64             `json:"TaxInclusiveAmt,omitempty"`
	ClassRef        *Reference          `json:"ClassRef,omitempty"`
	DepartmentRef   *Reference          `json:"DepartmentRef,omitempty"`
	TaxCodeRef      *Reference          `json:"TaxCodeRef,omitempty"`
	BillableStatus  BillableStatus      `json:"BillableStatus,omitempty"`
	JournalCodeRef  *Reference          `json:"JournalCodeRef,omitempty"`
}

// JournalEntryEntity represents the party on a journal entry line.
type JournalEntryEntity struct {
	Type      JournalEntryEntityType `json:"Type,omitempty"`
	EntityRef *Reference             `json:"EntityRef,omitempty"`
}
