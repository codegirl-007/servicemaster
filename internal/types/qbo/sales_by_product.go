// Package qbo contains transport types for the QuickBooks Online API.
package qbo

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
	SummarizeColumnBy SummarizeColumnBy
	Department        string
	StartDate         *Date
}
