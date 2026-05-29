// Package types contains transport types for external integrations.
package types

// FECReport represents the QuickBooks FEC report response.
type FECReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// FECReportQuery represents the supported query parameters for the FEC report.
type FECReportQuery struct {
	StartDate  *Date
	EndDate    *Date
	AddDueDate *bool
}
