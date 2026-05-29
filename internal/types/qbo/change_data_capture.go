// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// ChangeDataCaptureResponse represents the QuickBooks CDC response envelope.
type ChangeDataCaptureResponse struct {
	CDCResponse []CDCResponseGroup `json:"CDCResponse"`
	Time        time.Time          `json:"time"`
}

// CDCResponseGroup wraps query responses for one CDC poll.
type CDCResponseGroup struct {
	QueryResponse []CDCQueryResponse `json:"QueryResponse"`
}

// CDCQueryResponse represents changed entities for one entity type in a CDC response.
// CDC is not supported for JournalCode, TaxAgency, TimeActivity, TaxCode, or TaxRate.
type CDCQueryResponse struct {
	StartPosition        int                    `json:"startPosition,omitempty"`
	MaxResults           int                    `json:"maxResults,omitempty"`
	TotalCount           int                    `json:"totalCount,omitempty"`
	Account              []Account              `json:"Account,omitempty"`
	Attachable           []Attachable           `json:"Attachable,omitempty"`
	Bill                 []Bill                 `json:"Bill,omitempty"`
	BillPayment          []BillPayment          `json:"BillPayment,omitempty"`
	Budget               []Budget               `json:"Budget,omitempty"`
	Class                []Class                `json:"Class,omitempty"`
	CompanyCurrency      []CompanyCurrency      `json:"CompanyCurrency,omitempty"`
	CompanyInfo          []CompanyInfo          `json:"CompanyInfo,omitempty"`
	CreditCardPayment    []CreditCardPayment    `json:"CreditCardPayment,omitempty"`
	CreditMemo           []CreditMemo           `json:"CreditMemo,omitempty"`
	Customer             []Customer             `json:"Customer,omitempty"`
	CustomerType         []CustomerType         `json:"CustomerType,omitempty"`
	Department           []Department           `json:"Department,omitempty"`
	Deposit              []Deposit              `json:"Deposit,omitempty"`
	Employee             []Employee             `json:"Employee,omitempty"`
	Estimate             []Estimate             `json:"Estimate,omitempty"`
	ExchangeRate         []ExchangeRate         `json:"ExchangeRate,omitempty"`
	InventoryAdjustment  []InventoryAdjustment  `json:"InventoryAdjustment,omitempty"`
	Invoice              []Invoice              `json:"Invoice,omitempty"`
	Item                 []Item                 `json:"Item,omitempty"`
	JournalEntry         []JournalEntry         `json:"JournalEntry,omitempty"`
	Payment              []Payment              `json:"Payment,omitempty"`
	PaymentMethod        []PaymentMethod        `json:"PaymentMethod,omitempty"`
	Preferences          []Preferences          `json:"Preferences,omitempty"`
	Purchase             []Purchase             `json:"Purchase,omitempty"`
	PurchaseOrder        []PurchaseOrder        `json:"PurchaseOrder,omitempty"`
	RecurringTransaction []RecurringTransaction `json:"RecurringTransaction,omitempty"`
	RefundReceipt        []RefundReceipt        `json:"RefundReceipt,omitempty"`
	ReimburseCharge      []ReimburseCharge      `json:"ReimburseCharge,omitempty"`
	SalesReceipt         []SalesReceipt         `json:"SalesReceipt,omitempty"`
	TaxClassification    []TaxClassification    `json:"TaxClassification,omitempty"`
	TaxPayment           []TaxPayment           `json:"TaxPayment,omitempty"`
	Term                 []Term                 `json:"Term,omitempty"`
	Transfer             []Transfer             `json:"Transfer,omitempty"`
	Vendor               []Vendor               `json:"Vendor,omitempty"`
	VendorCredit         []VendorCredit         `json:"VendorCredit,omitempty"`
}
