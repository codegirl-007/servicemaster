// Package types contains transport types for external integrations.
package types

// InventoryValuationSummaryReport represents the QuickBooks inventory valuation summary report response.
type InventoryValuationSummaryReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// InventoryValuationSummaryQuery represents the supported query parameters for the inventory valuation summary report.
type InventoryValuationSummaryQuery struct {
	QZURL             *bool
	DateMacro         ReportDateMacro
	Item              string
	ReportDate        *Date
	SortOrder         ReportSortOrder
	SummarizeColumnBy SummarizeColumnBy
}
