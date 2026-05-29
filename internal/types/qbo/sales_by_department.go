// Package qbo contains transport types for the QuickBooks Online API.
package qbo

// SalesByDepartmentReport represents the QuickBooks sales by department report response.
type SalesByDepartmentReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// SalesByDepartmentQuery represents the supported query parameters for the sales by department report.
type SalesByDepartmentQuery struct {
	Customer          string
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
