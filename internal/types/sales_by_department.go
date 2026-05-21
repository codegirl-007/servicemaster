// Package types contains transport types for external integrations.
package types

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
	SummarizeColumnBy ProfitAndLossSummarizeColumnBy
	Department        string
	StartDate         *Date
}
