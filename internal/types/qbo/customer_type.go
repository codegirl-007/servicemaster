// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// CustomerTypeResponse represents the QuickBooks customer type response envelope.
type CustomerTypeResponse struct {
	CustomerType CustomerType `json:"CustomerType"`
	Time         time.Time    `json:"time"`
}

// CustomerType represents a QuickBooks customer type object.
type CustomerType struct {
	ID        string    `json:"Id"`
	SyncToken string    `json:"SyncToken,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Active    *bool     `json:"Active,omitempty"`
	MetaData  *MetaData `json:"MetaData,omitempty"`
	Domain    string    `json:"domain,omitempty"`
	Sparse    *bool     `json:"sparse,omitempty"`
}
