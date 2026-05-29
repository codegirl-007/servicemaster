// Package types contains transport types for external integrations.
package types

// CustomerIncomeReport represents the QuickBooks customer income report response.
type CustomerIncomeReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// CustomerIncomeQuery represents the supported query parameters for the customer income report.
type CustomerIncomeQuery struct {
	Customer          string
	Term              string
	AccountingMethod  ReportBasis
	EndDate           *Date
	DateMacro         ReportDateMacro
	Class             string
	SortOrder         ReportSortOrder
	SummarizeColumnBy SummarizeColumnBy
	Department        string
	Vendor            string
	StartDate         *Date
}
