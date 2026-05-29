// Package types contains transport types for external integrations.
package types

import "time"

// TimeActivityNameOf represents the documented time activity party types.
type TimeActivityNameOf string

const (
	TimeActivityNameOfEmployee TimeActivityNameOf = "Employee"
	TimeActivityNameOfVendor   TimeActivityNameOf = "Vendor"
)

// TimeActivityResponse represents the QuickBooks time activity response envelope.
type TimeActivityResponse struct {
	TimeActivity TimeActivity `json:"TimeActivity"`
	Time         time.Time    `json:"time"`
}

// TimeActivity represents a QuickBooks time activity object.
type TimeActivity struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken,omitempty"`
	NameOf                  TimeActivityNameOf      `json:"NameOf,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	Description             string                  `json:"Description,omitempty"`
	Hours                   float64                 `json:"Hours,omitempty"`
	Minutes                 float64                 `json:"Minutes,omitempty"`
	BreakHours              float64                 `json:"BreakHours,omitempty"`
	BreakMinutes            float64                 `json:"BreakMinutes,omitempty"`
	StartTime               string                  `json:"StartTime,omitempty"`
	EndTime                 string                  `json:"EndTime,omitempty"`
	HourlyRate              float64                 `json:"HourlyRate,omitempty"`
	CostRate                float64                 `json:"CostRate,omitempty"`
	Taxable                 *bool                   `json:"Taxable,omitempty"`
	BillableStatus          BillableStatus          `json:"BillableStatus,omitempty"`
	EmployeeRef             *Reference              `json:"EmployeeRef,omitempty"`
	VendorRef               *Reference              `json:"VendorRef,omitempty"`
	CustomerRef             *Reference              `json:"CustomerRef,omitempty"`
	ProjectRef              *Reference              `json:"ProjectRef,omitempty"`
	ItemRef                 *Reference              `json:"ItemRef,omitempty"`
	PayrollItemRef          *Reference              `json:"PayrollItemRef,omitempty"`
	ClassRef                *Reference              `json:"ClassRef,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	MetaData                *MetaData               `json:"MetaData,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
	Sparse                  *bool                   `json:"sparse,omitempty"`
}

// CreateTimeActivityRequest represents the documented create time activity payload.
type CreateTimeActivityRequest struct {
	NameOf                  TimeActivityNameOf      `json:"NameOf"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	Description             string                  `json:"Description,omitempty"`
	Hours                   float64                 `json:"Hours,omitempty"`
	Minutes                 float64                 `json:"Minutes,omitempty"`
	BreakHours              float64                 `json:"BreakHours,omitempty"`
	BreakMinutes            float64                 `json:"BreakMinutes,omitempty"`
	StartTime               string                  `json:"StartTime,omitempty"`
	EndTime                 string                  `json:"EndTime,omitempty"`
	HourlyRate              float64                 `json:"HourlyRate,omitempty"`
	CostRate                float64                 `json:"CostRate,omitempty"`
	Taxable                 *bool                   `json:"Taxable,omitempty"`
	EmployeeRef             *Reference              `json:"EmployeeRef,omitempty"`
	VendorRef               *Reference              `json:"VendorRef,omitempty"`
	CustomerRef             *Reference              `json:"CustomerRef,omitempty"`
	ProjectRef              *Reference              `json:"ProjectRef,omitempty"`
	ItemRef                 *Reference              `json:"ItemRef,omitempty"`
	PayrollItemRef          *Reference              `json:"PayrollItemRef,omitempty"`
	ClassRef                *Reference              `json:"ClassRef,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
}

// UpdateTimeActivityRequest represents the documented full update time activity payload.
type UpdateTimeActivityRequest struct {
	ID                      string                  `json:"Id"`
	SyncToken               string                  `json:"SyncToken"`
	NameOf                  TimeActivityNameOf      `json:"NameOf,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	Description             string                  `json:"Description,omitempty"`
	Hours                   float64                 `json:"Hours,omitempty"`
	Minutes                 float64                 `json:"Minutes,omitempty"`
	BreakHours              float64                 `json:"BreakHours,omitempty"`
	BreakMinutes            float64                 `json:"BreakMinutes,omitempty"`
	StartTime               string                  `json:"StartTime,omitempty"`
	EndTime                 string                  `json:"EndTime,omitempty"`
	HourlyRate              float64                 `json:"HourlyRate,omitempty"`
	CostRate                float64                 `json:"CostRate,omitempty"`
	Taxable                 *bool                   `json:"Taxable,omitempty"`
	EmployeeRef             *Reference              `json:"EmployeeRef,omitempty"`
	VendorRef               *Reference              `json:"VendorRef,omitempty"`
	CustomerRef             *Reference              `json:"CustomerRef,omitempty"`
	ProjectRef              *Reference              `json:"ProjectRef,omitempty"`
	ItemRef                 *Reference              `json:"ItemRef,omitempty"`
	PayrollItemRef          *Reference              `json:"PayrollItemRef,omitempty"`
	ClassRef                *Reference              `json:"ClassRef,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	Domain                  string                  `json:"domain,omitempty"`
}

// DeleteTimeActivityRequest represents the documented delete time activity payload.
type DeleteTimeActivityRequest struct {
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
}

// TimeActivityDeleteResponse represents the QuickBooks deleted time activity response envelope.
type TimeActivityDeleteResponse struct {
	TimeActivity DeletedEntity `json:"TimeActivity"`
	Time         time.Time     `json:"time"`
}
