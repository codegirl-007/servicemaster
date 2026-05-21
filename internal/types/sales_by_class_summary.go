// Package types contains transport types for external integrations.
package types

// SalesByClassSummaryReport represents the QuickBooks sales by class summary report response.
type SalesByClassSummaryReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// SalesByClassSummaryQuery represents the supported query parameters for the sales by class summary report.
type SalesByClassSummaryQuery struct {
	Customer          string
	AccountingMethod  ReportBasis
	EndDate           *Date
	DateMacro         ReportDateMacro
	Class             string
	Item              string
	SummarizeColumnBy ProfitAndLossSummarizeColumnBy
	Department        string
	StartDate         *Date
}
