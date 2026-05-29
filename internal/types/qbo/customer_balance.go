// Package qbo contains transport types for the QuickBooks Online API.
package qbo

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
	ARPaid            ReportARPaid
	ReportDate        *Date
	SortOrder         ReportSortOrder
	SummarizeColumnBy SummarizeColumnBy
	Department        string
}
