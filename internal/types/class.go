// Package types contains transport types for external integrations.
package types

import "time"

// ClassResponse represents the QuickBooks class response envelope.
type ClassResponse struct {
	Class Class     `json:"Class"`
	Time  time.Time `json:"time"`
}

// Class represents a QuickBooks class object.
type Class struct {
	ID                 string     `json:"Id"`
	Name               string     `json:"Name,omitempty"`
	SyncToken          string     `json:"SyncToken,omitempty"`
	ParentRef          *Reference `json:"ParentRef,omitempty"`
	FullyQualifiedName string     `json:"FullyQualifiedName,omitempty"`
	SubClass           *bool      `json:"SubClass,omitempty"`
	Active             *bool      `json:"Active,omitempty"`
	MetaData           *MetaData  `json:"MetaData,omitempty"`
	Domain             string     `json:"domain,omitempty"`
	Sparse             *bool      `json:"sparse,omitempty"`
}

// CreateClassRequest represents the documented create class payload.
type CreateClassRequest struct {
	// Name is required.
	Name      string     `json:"Name"`
	ParentRef *Reference `json:"ParentRef,omitempty"`
}
