// Package qbo contains transport types for the QuickBooks Online API.
package qbo

// TrialBalanceReport represents the QuickBooks trial balance report response.
type TrialBalanceReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// TrialBalanceQuery represents the supported query parameters for the trial balance report.
type TrialBalanceQuery struct {
	AccountingMethod  ReportBasis
	EndDate           *Date
	DateMacro         ReportDateMacro
	SortOrder         ReportSortOrder
	SummarizeColumnBy SummarizeColumnBy
	StartDate         *Date
}
