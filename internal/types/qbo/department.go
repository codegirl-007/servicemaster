// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// DepartmentResponse represents the QuickBooks department response envelope.
type DepartmentResponse struct {
	Department Department `json:"Department"`
	Time       time.Time  `json:"time"`
}

// Department represents a QuickBooks department object.
type Department struct {
	ID                 string     `json:"Id"`
	Name               string     `json:"Name,omitempty"`
	SyncToken          string     `json:"SyncToken,omitempty"`
	ParentRef          *Reference `json:"ParentRef,omitempty"`
	FullyQualifiedName string     `json:"FullyQualifiedName,omitempty"`
	SubDepartment      *bool      `json:"SubDepartment,omitempty"`
	Active             *bool      `json:"Active,omitempty"`
	MetaData           *MetaData  `json:"MetaData,omitempty"`
	Domain             string     `json:"domain,omitempty"`
	Sparse             *bool      `json:"sparse,omitempty"`
}

// CreateDepartmentRequest represents the documented create department payload.
type CreateDepartmentRequest struct {
	// Name is required.
	Name      string     `json:"Name"`
	ParentRef *Reference `json:"ParentRef,omitempty"`
}
