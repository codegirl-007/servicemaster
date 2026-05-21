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
	BID         string           `json:"bId"`
	Operation   BatchOperation   `json:"operation,omitempty"`
	OptionsData BatchOptionsData `json:"optionsData,omitempty"`
	Query       string           `json:"Query,omitempty"`
	Vendor      json.RawMessage  `json:"Vendor,omitempty"`
	Customer    json.RawMessage  `json:"Customer,omitempty"`
	Invoice     json.RawMessage  `json:"Invoice,omitempty"`
	Payment     json.RawMessage  `json:"Payment,omitempty"`
	SalesReceipt  json.RawMessage `json:"SalesReceipt,omitempty"`
	RefundReceipt json.RawMessage `json:"RefundReceipt,omitempty"`
	Deposit     json.RawMessage  `json:"Deposit,omitempty"`
	JournalEntry  json.RawMessage `json:"JournalEntry,omitempty"`
	Transfer    json.RawMessage  `json:"Transfer,omitempty"`
	Bill        json.RawMessage  `json:"Bill,omitempty"`
}

// BatchItemResponse represents one result in a batch response.
type BatchItemResponse struct {
	BID           string              `json:"bId"`
	Fault         *BatchFault         `json:"Fault,omitempty"`
	QueryResponse *BatchQueryResponse `json:"QueryResponse,omitempty"`
	Vendor        json.RawMessage     `json:"Vendor,omitempty"`
	Customer      json.RawMessage     `json:"Customer,omitempty"`
	Invoice       json.RawMessage     `json:"Invoice,omitempty"`
	Payment       json.RawMessage     `json:"Payment,omitempty"`
	SalesReceipt  json.RawMessage     `json:"SalesReceipt,omitempty"`
	RefundReceipt json.RawMessage     `json:"RefundReceipt,omitempty"`
	Deposit       json.RawMessage     `json:"Deposit,omitempty"`
	JournalEntry  json.RawMessage     `json:"JournalEntry,omitempty"`
	Transfer      json.RawMessage     `json:"Transfer,omitempty"`
	Bill          json.RawMessage     `json:"Bill,omitempty"`
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
	StartPosition int             `json:"startPosition,omitempty"`
	MaxResults    int             `json:"maxResults,omitempty"`
	TotalCount    int             `json:"totalCount,omitempty"`
	SalesReceipt  json.RawMessage `json:"SalesReceipt,omitempty"`
	Customer      json.RawMessage `json:"Customer,omitempty"`
	Invoice       json.RawMessage `json:"Invoice,omitempty"`
}
