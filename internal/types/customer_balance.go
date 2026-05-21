// Package types contains transport types for external integrations.
package types

// CustomerBalanceReport represents the QuickBooks customer balance report response.
type CustomerBalanceReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// CustomerBalanceQuery represents the supported query parameters for the customer balance report.
type CustomerBalanceQuery struct {
	Customer          string
	AccountingMethod  ReportBasis
	DateMacro         ReportDateMacro
	ARPaid            string
	ReportDate        *Date
	SortOrder         ReportSortOrder
	SummarizeColumnBy ProfitAndLossSummarizeColumnBy
	Department        string
}
