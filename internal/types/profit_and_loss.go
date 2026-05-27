// Package types contains transport types for external integrations.
package types

// ProfitAndLossSummarizeColumnBy is an alias for SummarizeColumnBy.
type ProfitAndLossSummarizeColumnBy = SummarizeColumnBy

const (
	ProfitAndLossSummarizeColumnByTotal               = SummarizeColumnByTotal
	ProfitAndLossSummarizeColumnByMonth               = SummarizeColumnByMonth
	ProfitAndLossSummarizeColumnByWeek                = SummarizeColumnByWeek
	ProfitAndLossSummarizeColumnByDays                = SummarizeColumnByDays
	ProfitAndLossSummarizeColumnByQuarter             = SummarizeColumnByQuarter
	ProfitAndLossSummarizeColumnByYear                = SummarizeColumnByYear
	ProfitAndLossSummarizeColumnByCustomers           = SummarizeColumnByCustomers
	ProfitAndLossSummarizeColumnByVendors             = SummarizeColumnByVendors
	ProfitAndLossSummarizeColumnByClasses             = SummarizeColumnByClasses
	ProfitAndLossSummarizeColumnByDepartments         = SummarizeColumnByDepartments
	ProfitAndLossSummarizeColumnByEmployees           = SummarizeColumnByEmployees
	ProfitAndLossSummarizeColumnByProductsAndServices = SummarizeColumnByProductsAndServices
)

// ProfitAndLossReport represents the QuickBooks profit and loss report response.
type ProfitAndLossReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// ProfitAndLossQuery represents the supported query parameters for the profit and loss report.
type ProfitAndLossQuery struct {
	StartDate         *Date
	EndDate           *Date
	DateMacro         ReportDateMacro
	AccountingMethod  ReportBasis
	SummarizeColumnBy SummarizeColumnBy
	SortOrder         ReportSortOrder
	Customer          string
	Vendor            string
	Class             string
	Department        string
	Item              string
	QZURL             *bool
	AdjustedGainLoss  *bool
}
