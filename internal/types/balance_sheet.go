// Package types contains transport types for external integrations.
package types

// BalanceSheetReport represents the QuickBooks balance sheet report response.
type BalanceSheetReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// BalanceSheetQuery represents the supported query parameters for the balance sheet report.
type BalanceSheetQuery struct {
	Customer          string
	QZURL             *bool
	AccountingMethod  ReportBasis
	EndDate           *Date
	DateMacro         ReportDateMacro
	AdjustedGainLoss  *bool
	Class             string
	Item              string
	SortOrder         ReportSortOrder
	SummarizeColumnBy ProfitAndLossSummarizeColumnBy
	Department        string
	Vendor            string
	StartDate         *Date
}
