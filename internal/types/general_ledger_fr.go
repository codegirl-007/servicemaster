// Package types contains transport types for external integrations.
package types

// GeneralLedgerFR represents the QuickBooks General Ledger FR report response.
type GeneralLedgerFR struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// GeneralLedgerFRQuery represents the supported query parameters for the General Ledger FR report.
type GeneralLedgerFRQuery struct {
	Customer          string
	Account           string
	AccountingMethod  ReportBasis
	SourceAccount     string
	EndDate           *Date
	DateMacro         ReportDateMacro
	AccountType       ReportAccountType
	SortBy            string
	SortOrder         ReportSortOrder
	StartDate         *Date
	SummarizeColumnBy SummarizeColumnBy
	Department        string
	Vendor            string
	Class             string
}
