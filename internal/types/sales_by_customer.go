// Package types contains transport types for external integrations.
package types

// SalesByCustomerReport represents the QuickBooks sales by customer report response.
type SalesByCustomerReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// SalesByCustomerQuery represents the supported query parameters for the sales by customer report.
type SalesByCustomerQuery struct {
	Customer          string
	QZURL             *bool
	AccountingMethod  ReportBasis
	EndDate           *Date
	DateMacro         ReportDateMacro
	Class             string
	Item              string
	SortOrder         ReportSortOrder
	SummarizeColumnBy SummarizeColumnBy
	Department        string
	StartDate         *Date
}
