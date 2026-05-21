// Package types contains transport types for external integrations.
package types

// SalesByProductReport represents the QuickBooks sales by product report response.
type SalesByProductReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// SalesByProductQuery represents the supported query parameters for the sales by product report.
type SalesByProductQuery struct {
	Customer          string
	EndDueDate        *Date
	AccountingMethod  ReportBasis
	EndDate           *Date
	DateMacro         ReportDateMacro
	StartDueDate      *Date
	Class             string
	Item              string
	SortOrder         ReportSortOrder
	SummarizeColumnBy ProfitAndLossSummarizeColumnBy
	Department        string
	StartDate         *Date
}
