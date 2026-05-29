// Package types contains transport types for external integrations.
package types

import (
	"encoding/json"
	"time"
)

// BatchOperation represents the documented batch item operation values.
type BatchOperation string

const (
	BatchOperationCreate BatchOperation = "create"
	BatchOperationUpdate BatchOperation = "update"
	BatchOperationDelete BatchOperation = "delete"
)

// BatchOptionsData represents documented optionsData values for batch updates.
type BatchOptionsData string

const (
	BatchOptionsDataVoid BatchOptionsData = "void"
)

// BatchRequest represents the QuickBooks batch request envelope.
type BatchRequest struct {
	BatchItemRequest []BatchItemRequest `json:"BatchItemRequest"`
}

// BatchResponse represents the QuickBooks batch response envelope.
type BatchResponse struct {
	BatchItemResponse []BatchItemResponse `json:"BatchItemResponse"`
	Time              time.Time           `json:"time"`
}

// BatchItemRequest represents one operation in a batch request.
// Set Query for query operations, or operation with exactly one entity payload field.
type BatchItemRequest struct {
	BID                  string           `json:"bId"`
	Operation            BatchOperation   `json:"operation,omitempty"`
	OptionsData          BatchOptionsData `json:"optionsData,omitempty"`
	Query                string           `json:"Query,omitempty"`
	Account              json.RawMessage  `json:"Account,omitempty"`
	Attachable           json.RawMessage  `json:"Attachable,omitempty"`
	Bill                 json.RawMessage  `json:"Bill,omitempty"`
	BillPayment          json.RawMessage  `json:"BillPayment,omitempty"`
	Budget               json.RawMessage  `json:"Budget,omitempty"`
	Class                json.RawMessage  `json:"Class,omitempty"`
	CompanyCurrency      json.RawMessage  `json:"CompanyCurrency,omitempty"`
	CompanyInfo          json.RawMessage  `json:"CompanyInfo,omitempty"`
	CreditCardPayment    json.RawMessage  `json:"CreditCardPayment,omitempty"`
	CreditMemo           json.RawMessage  `json:"CreditMemo,omitempty"`
	Customer             json.RawMessage  `json:"Customer,omitempty"`
	CustomerType         json.RawMessage  `json:"CustomerType,omitempty"`
	Department           json.RawMessage  `json:"Department,omitempty"`
	Deposit              json.RawMessage  `json:"Deposit,omitempty"`
	Employee             json.RawMessage  `json:"Employee,omitempty"`
	Estimate             json.RawMessage  `json:"Estimate,omitempty"`
	ExchangeRate         json.RawMessage  `json:"ExchangeRate,omitempty"`
	InventoryAdjustment  json.RawMessage  `json:"InventoryAdjustment,omitempty"`
	Invoice              json.RawMessage  `json:"Invoice,omitempty"`
	Item                 json.RawMessage  `json:"Item,omitempty"`
	JournalEntry         json.RawMessage  `json:"JournalEntry,omitempty"`
	Payment              json.RawMessage  `json:"Payment,omitempty"`
	PaymentMethod        json.RawMessage  `json:"PaymentMethod,omitempty"`
	Preferences          json.RawMessage  `json:"Preferences,omitempty"`
	Purchase             json.RawMessage  `json:"Purchase,omitempty"`
	PurchaseOrder        json.RawMessage  `json:"PurchaseOrder,omitempty"`
	RecurringTransaction json.RawMessage  `json:"RecurringTransaction,omitempty"`
	RefundReceipt        json.RawMessage  `json:"RefundReceipt,omitempty"`
	ReimburseCharge      json.RawMessage  `json:"ReimburseCharge,omitempty"`
	SalesReceipt         json.RawMessage  `json:"SalesReceipt,omitempty"`
	TaxClassification    json.RawMessage  `json:"TaxClassification,omitempty"`
	TaxPayment           json.RawMessage  `json:"TaxPayment,omitempty"`
	TaxService           json.RawMessage  `json:"TaxService,omitempty"`
	Term                 json.RawMessage  `json:"Term,omitempty"`
	Transfer             json.RawMessage  `json:"Transfer,omitempty"`
	Vendor               json.RawMessage  `json:"Vendor,omitempty"`
	VendorCredit         json.RawMessage  `json:"VendorCredit,omitempty"`
}

// BatchItemResponse represents one result in a batch response.
type BatchItemResponse struct {
	BID                  string              `json:"bId"`
	Fault                *BatchFault         `json:"Fault,omitempty"`
	QueryResponse        *BatchQueryResponse `json:"QueryResponse,omitempty"`
	Account              json.RawMessage     `json:"Account,omitempty"`
	Attachable           json.RawMessage     `json:"Attachable,omitempty"`
	Bill                 json.RawMessage     `json:"Bill,omitempty"`
	BillPayment          json.RawMessage     `json:"BillPayment,omitempty"`
	Budget               json.RawMessage     `json:"Budget,omitempty"`
	Class                json.RawMessage     `json:"Class,omitempty"`
	CompanyCurrency      json.RawMessage     `json:"CompanyCurrency,omitempty"`
	CompanyInfo          json.RawMessage     `json:"CompanyInfo,omitempty"`
	CreditCardPayment    json.RawMessage     `json:"CreditCardPayment,omitempty"`
	CreditMemo           json.RawMessage     `json:"CreditMemo,omitempty"`
	Customer             json.RawMessage     `json:"Customer,omitempty"`
	CustomerType         json.RawMessage     `json:"CustomerType,omitempty"`
	Department           json.RawMessage     `json:"Department,omitempty"`
	Deposit              json.RawMessage     `json:"Deposit,omitempty"`
	Employee             json.RawMessage     `json:"Employee,omitempty"`
	Estimate             json.RawMessage     `json:"Estimate,omitempty"`
	ExchangeRate         json.RawMessage     `json:"ExchangeRate,omitempty"`
	InventoryAdjustment  json.RawMessage     `json:"InventoryAdjustment,omitempty"`
	Invoice              json.RawMessage     `json:"Invoice,omitempty"`
	Item                 json.RawMessage     `json:"Item,omitempty"`
	JournalEntry         json.RawMessage     `json:"JournalEntry,omitempty"`
	Payment              json.RawMessage     `json:"Payment,omitempty"`
	PaymentMethod        json.RawMessage     `json:"PaymentMethod,omitempty"`
	Preferences          json.RawMessage     `json:"Preferences,omitempty"`
	Purchase             json.RawMessage     `json:"Purchase,omitempty"`
	PurchaseOrder        json.RawMessage     `json:"PurchaseOrder,omitempty"`
	RecurringTransaction json.RawMessage     `json:"RecurringTransaction,omitempty"`
	RefundReceipt        json.RawMessage     `json:"RefundReceipt,omitempty"`
	ReimburseCharge      json.RawMessage     `json:"ReimburseCharge,omitempty"`
	SalesReceipt         json.RawMessage     `json:"SalesReceipt,omitempty"`
	TaxClassification    json.RawMessage     `json:"TaxClassification,omitempty"`
	TaxPayment           json.RawMessage     `json:"TaxPayment,omitempty"`
	TaxService           json.RawMessage     `json:"TaxService,omitempty"`
	Term                 json.RawMessage     `json:"Term,omitempty"`
	Transfer             json.RawMessage     `json:"Transfer,omitempty"`
	Vendor               json.RawMessage     `json:"Vendor,omitempty"`
	VendorCredit         json.RawMessage     `json:"VendorCredit,omitempty"`
}

// BatchFault represents a batch item error payload.
type BatchFault struct {
	Type  string       `json:"type,omitempty"`
	Error []BatchError `json:"Error,omitempty"`
}

// BatchError represents one error in a batch fault.
type BatchError struct {
	Message string `json:"Message,omitempty"`
	Code    string `json:"code,omitempty"`
	Detail  string `json:"Detail,omitempty"`
	Element string `json:"element,omitempty"`
}

// BatchQueryResponse represents query results returned from a batch query item.
type BatchQueryResponse struct {
	StartPosition        int             `json:"startPosition,omitempty"`
	MaxResults           int             `json:"maxResults,omitempty"`
	TotalCount           int             `json:"totalCount,omitempty"`
	Account              json.RawMessage `json:"Account,omitempty"`
	Attachable           json.RawMessage `json:"Attachable,omitempty"`
	Bill                 json.RawMessage `json:"Bill,omitempty"`
	BillPayment          json.RawMessage `json:"BillPayment,omitempty"`
	Budget               json.RawMessage `json:"Budget,omitempty"`
	Class                json.RawMessage `json:"Class,omitempty"`
	CompanyCurrency      json.RawMessage `json:"CompanyCurrency,omitempty"`
	CompanyInfo          json.RawMessage `json:"CompanyInfo,omitempty"`
	CreditCardPayment    json.RawMessage `json:"CreditCardPayment,omitempty"`
	CreditMemo           json.RawMessage `json:"CreditMemo,omitempty"`
	Customer             json.RawMessage `json:"Customer,omitempty"`
	CustomerType         json.RawMessage `json:"CustomerType,omitempty"`
	Department           json.RawMessage `json:"Department,omitempty"`
	Deposit              json.RawMessage `json:"Deposit,omitempty"`
	Employee             json.RawMessage `json:"Employee,omitempty"`
	Estimate             json.RawMessage `json:"Estimate,omitempty"`
	ExchangeRate         json.RawMessage `json:"ExchangeRate,omitempty"`
	InventoryAdjustment  json.RawMessage `json:"InventoryAdjustment,omitempty"`
	Invoice              json.RawMessage `json:"Invoice,omitempty"`
	Item                 json.RawMessage `json:"Item,omitempty"`
	JournalEntry         json.RawMessage `json:"JournalEntry,omitempty"`
	Payment              json.RawMessage `json:"Payment,omitempty"`
	PaymentMethod        json.RawMessage `json:"PaymentMethod,omitempty"`
	Preferences          json.RawMessage `json:"Preferences,omitempty"`
	Purchase             json.RawMessage `json:"Purchase,omitempty"`
	PurchaseOrder        json.RawMessage `json:"PurchaseOrder,omitempty"`
	RecurringTransaction json.RawMessage `json:"RecurringTransaction,omitempty"`
	RefundReceipt        json.RawMessage `json:"RefundReceipt,omitempty"`
	ReimburseCharge      json.RawMessage `json:"ReimburseCharge,omitempty"`
	SalesReceipt         json.RawMessage `json:"SalesReceipt,omitempty"`
	TaxClassification    json.RawMessage `json:"TaxClassification,omitempty"`
	TaxPayment           json.RawMessage `json:"TaxPayment,omitempty"`
	TaxService           json.RawMessage `json:"TaxService,omitempty"`
	Term                 json.RawMessage `json:"Term,omitempty"`
	Transfer             json.RawMessage `json:"Transfer,omitempty"`
	Vendor               json.RawMessage `json:"Vendor,omitempty"`
	VendorCredit         json.RawMessage `json:"VendorCredit,omitempty"`
}
