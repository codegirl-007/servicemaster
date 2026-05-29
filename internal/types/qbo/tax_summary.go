// Package qbo contains transport types for the QuickBooks Online API.
package qbo

// TaxSummaryReport represents the QuickBooks tax summary report response.
type TaxSummaryReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// TaxSummaryQuery represents the supported query parameters for the tax summary report.
type TaxSummaryQuery struct {
	AgencyID         string
	AccountingMethod ReportBasis
	EndDate          *Date
	DateMacro        ReportDateMacro
	SortOrder        ReportSortOrder
	StartDate        *Date
}
