// Package types contains transport types for external integrations.
package types

// VendorBalanceReport represents the QuickBooks vendor balance report response.
type VendorBalanceReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// VendorBalanceQuery represents the supported query parameters for the vendor balance report.
type VendorBalanceQuery struct {
	QZURL             *bool
	AccountingMethod  ReportBasis
	DateMacro         ReportDateMacro
	APPaid            string
	ReportDate        *Date
	SortOrder         ReportSortOrder
	SummarizeColumnBy ProfitAndLossSummarizeColumnBy
	Department        string
	Vendor            string
}
