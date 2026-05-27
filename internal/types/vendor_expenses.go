// Package types contains transport types for external integrations.
package types

// VendorExpensesReport represents the QuickBooks vendor expenses report response.
type VendorExpensesReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// VendorExpensesQuery represents the supported query parameters for the vendor expenses report.
type VendorExpensesQuery struct {
	Customer          string
	Vendor            string
	EndDate           *Date
	DateMacro         ReportDateMacro
	Class             string
	SortOrder         ReportSortOrder
	SummarizeColumnBy SummarizeColumnBy
	Department        string
	AccountingMethod  ReportBasis
	StartDate         *Date
}
