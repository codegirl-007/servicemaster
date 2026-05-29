// Package types contains transport types for external integrations.
package types

// CashFlowReport represents the QuickBooks cash flow report response.
type CashFlowReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// CashFlowQuery represents the supported query parameters for the cash flow report.
type CashFlowQuery struct {
	Customer          string
	Vendor            string
	EndDate           *Date
	DateMacro         ReportDateMacro
	Class             string
	Item              string
	SortOrder         ReportSortOrder
	SummarizeColumnBy SummarizeColumnBy
	Department        string
	StartDate         *Date
}
