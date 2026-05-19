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
	ID          string     `json:"Id"`
	SyncToken   string     `json:"SyncToken"`
	Name        string     `json:"Name"`
	FullyQualifiedName string `json:"FullyQualifiedName,omitempty"`
	SubClass    bool       `json:"SubClass,omitempty"`
	ParentRef   *Reference `json:"ParentRef,omitempty"`
	Active      *bool      `json:"Active,omitempty"`
	Domain      string     `json:"domain,omitempty"`
	Sparse      *bool      `json:"sparse,omitempty"`
	MetaData    *MetaData  `json:"MetaData,omitempty"`
}

// CreateClassRequest represents the QuickBooks create class payload.
type CreateClassRequest struct {
	Name      string     `json:"Name"`
	SubClass  bool       `json:"SubClass,omitempty"`
	ParentRef *Reference `json:"ParentRef,omitempty"`
	Active    *bool      `json:"Active,omitempty"`
}

// SparseUpdateClassRequest represents the QuickBooks sparse class update payload.
type SparseUpdateClassRequest struct {
	ID        string     `json:"Id"`
	SyncToken string     `json:"SyncToken"`
	Sparse    bool       `json:"sparse"`
	Name      string     `json:"Name,omitempty"`
	SubClass  bool       `json:"SubClass,omitempty"`
	ParentRef *Reference `json:"ParentRef,omitempty"`
	Active    *bool      `json:"Active,omitempty"`
	Domain    string     `json:"domain,omitempty"`
}
