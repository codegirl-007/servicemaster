// Package types contains transport types for external integrations.
package types

// ARAgingSummaryReport represents the QuickBooks AR aging summary report response.
type ARAgingSummaryReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// ARAgingSummaryQuery represents the supported query parameters for the AR aging summary report.
type ARAgingSummaryQuery struct {
	Customer    string
	QZURL       *bool
	DateMacro   ReportDateMacro
	AgingMethod string
	ReportDate  *Date
	SortOrder   ReportSortOrder
	Department  string
}
