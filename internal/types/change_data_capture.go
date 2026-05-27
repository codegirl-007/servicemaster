// Package types contains transport types for external integrations.
package types

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
// Additional entity arrays use the same pattern as Customer and Estimate.
type CDCQueryResponse struct {
	StartPosition int        `json:"startPosition,omitempty"`
	MaxResults    int        `json:"maxResults,omitempty"`
	TotalCount    int        `json:"totalCount,omitempty"`
	Customer      []Customer `json:"Customer,omitempty"`
	Estimate      []Estimate `json:"Estimate,omitempty"`
}
