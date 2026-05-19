// Package types contains transport types for external integrations.
package types

// ProfitAndLossSummarizeColumnBy represents the documented summarize_column_by values.
type ProfitAndLossSummarizeColumnBy string

const (
	ProfitAndLossSummarizeColumnByTotal              ProfitAndLossSummarizeColumnBy = "Total"
	ProfitAndLossSummarizeColumnByMonth              ProfitAndLossSummarizeColumnBy = "Month"
	ProfitAndLossSummarizeColumnByWeek               ProfitAndLossSummarizeColumnBy = "Week"
	ProfitAndLossSummarizeColumnByDays               ProfitAndLossSummarizeColumnBy = "Days"
	ProfitAndLossSummarizeColumnByQuarter            ProfitAndLossSummarizeColumnBy = "Quarter"
	ProfitAndLossSummarizeColumnByYear               ProfitAndLossSummarizeColumnBy = "Year"
	ProfitAndLossSummarizeColumnByCustomers          ProfitAndLossSummarizeColumnBy = "Customers"
	ProfitAndLossSummarizeColumnByVendors            ProfitAndLossSummarizeColumnBy = "Vendors"
	ProfitAndLossSummarizeColumnByClasses            ProfitAndLossSummarizeColumnBy = "Classes"
	ProfitAndLossSummarizeColumnByDepartments        ProfitAndLossSummarizeColumnBy = "Departments"
	ProfitAndLossSummarizeColumnByEmployees          ProfitAndLossSummarizeColumnBy = "Employees"
	ProfitAndLossSummarizeColumnByProductsAndServices ProfitAndLossSummarizeColumnBy = "ProductsAndServices"
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
	SummarizeColumnBy ProfitAndLossSummarizeColumnBy
	SortOrder         ReportSortOrder              
	Customer          string                       
	Vendor            string                       
	Class             string                       
	Department        string                       
	Item              string                       
	QZURL             *bool                        
	AdjustedGainLoss  *bool                        
}
