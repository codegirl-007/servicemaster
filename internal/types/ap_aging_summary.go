// Package types contains transport types for external integrations.
package types

// APAgingSummaryReport represents the QuickBooks AP aging summary report response.
type APAgingSummaryReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// APAgingSummaryQuery represents the supported query parameters for the AP aging summary report.
type APAgingSummaryQuery struct {
	Customer    string
	QZURL       *bool
	Vendor      string
	DateMacro   ReportDateMacro
	Department  string
	ReportDate  *Date
	SortOrder   ReportSortOrder
	AgingMethod string
}
