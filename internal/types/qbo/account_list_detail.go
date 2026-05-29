// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// ReportBasis represents a QuickBooks report accounting basis.
type ReportBasis string

const (
	ReportBasisCash    ReportBasis = "Cash"
	ReportBasisAccrual ReportBasis = "Accrual"
)

// ReportRowType represents a QuickBooks report row type.
type ReportRowType string

const (
	ReportRowTypeData    ReportRowType = "Data"
	ReportRowTypeSection ReportRowType = "Section"
)

// ReportColumnType represents a QuickBooks report column type.
type ReportColumnType string

const (
	ReportColumnTypeAccount ReportColumnType = "Account"
	ReportColumnTypeMoney   ReportColumnType = "Money"
)

// ReportSortOrder represents the report row sort order.
type ReportSortOrder string

const (
	ReportSortOrderAscend  ReportSortOrder = "ascend"
	ReportSortOrderDescend ReportSortOrder = "descend"
)

// ReportDateMacro represents a predefined report date range.
type ReportDateMacro string

const (
	ReportDateMacroToday                   ReportDateMacro = "Today"
	ReportDateMacroYesterday               ReportDateMacro = "Yesterday"
	ReportDateMacroThisWeek                ReportDateMacro = "This Week"
	ReportDateMacroLastWeek                ReportDateMacro = "Last Week"
	ReportDateMacroThisWeekToDate          ReportDateMacro = "This Week-to-date"
	ReportDateMacroLastWeekToDate          ReportDateMacro = "Last Week-to-date"
	ReportDateMacroNextWeek                ReportDateMacro = "Next Week"
	ReportDateMacroNext4Weeks              ReportDateMacro = "Next 4 Weeks"
	ReportDateMacroThisMonth               ReportDateMacro = "This Month"
	ReportDateMacroLastMonth               ReportDateMacro = "Last Month"
	ReportDateMacroThisMonthToDate         ReportDateMacro = "This Month-to-date"
	ReportDateMacroLastMonthToDate         ReportDateMacro = "Last Month-to-date"
	ReportDateMacroNextMonth               ReportDateMacro = "Next Month"
	ReportDateMacroThisFiscalQuarter       ReportDateMacro = "This Fiscal Quarter"
	ReportDateMacroLastFiscalQuarter       ReportDateMacro = "Last Fiscal Quarter"
	ReportDateMacroThisFiscalQuarterToDate ReportDateMacro = "This Fiscal Quarter-to-date"
	ReportDateMacroLastFiscalQuarterToDate ReportDateMacro = "Last Fiscal Quarter-to-date"
	ReportDateMacroNextFiscalQuarter       ReportDateMacro = "Next Fiscal Quarter"
	ReportDateMacroThisFiscalYear          ReportDateMacro = "This Fiscal Year"
	ReportDateMacroLastFiscalYear          ReportDateMacro = "Last Fiscal Year"
	ReportDateMacroThisFiscalYearToDate    ReportDateMacro = "This Fiscal Year-to-date"
	ReportDateMacroLastFiscalYearToDate    ReportDateMacro = "Last Fiscal Year-to-date"
	ReportDateMacroNextFiscalYear          ReportDateMacro = "Next Fiscal Year"
)

// ReportAccountStatus represents the account_status query parameter.
type ReportAccountStatus string

const (
	ReportAccountStatusDeleted    ReportAccountStatus = "Deleted"
	ReportAccountStatusNotDeleted ReportAccountStatus = "Not_Deleted"
)

// ReportAccountType represents the account_type query parameter.
type ReportAccountType string

const (
	ReportAccountTypeAccountsPayable       ReportAccountType = "AccountsPayable"
	ReportAccountTypeAccountsReceivable    ReportAccountType = "AccountsReceivable"
	ReportAccountTypeBank                  ReportAccountType = "Bank"
	ReportAccountTypeCostOfGoodsSold       ReportAccountType = "CostOfGoodsSold"
	ReportAccountTypeCreditCard            ReportAccountType = "CreditCard"
	ReportAccountTypeEquity                ReportAccountType = "Equity"
	ReportAccountTypeExpense               ReportAccountType = "Expense"
	ReportAccountTypeFixedAsset            ReportAccountType = "FixedAsset"
	ReportAccountTypeIncome                ReportAccountType = "Income"
	ReportAccountTypeLongTermLiability     ReportAccountType = "LongTermLiability"
	ReportAccountTypeNonPosting            ReportAccountType = "NonPosting"
	ReportAccountTypeOtherAsset            ReportAccountType = "OtherAsset"
	ReportAccountTypeOtherCurrentAsset     ReportAccountType = "OtherCurrentAsset"
	ReportAccountTypeOtherCurrentLiability ReportAccountType = "OtherCurrentLiability"
	ReportAccountTypeOtherExpense          ReportAccountType = "OtherExpense"
	ReportAccountTypeOtherIncome           ReportAccountType = "OtherIncome"
)

// AccountListDetailQueryColumn represents the columns and sort_by query values.
type AccountListDetailQueryColumn string

const (
	AccountListDetailQueryColumnAccountName   AccountListDetailQueryColumn = "account_name"
	AccountListDetailQueryColumnAccountType   AccountListDetailQueryColumn = "account_type"
	AccountListDetailQueryColumnDetailAccType AccountListDetailQueryColumn = "detail_acc_type"
	AccountListDetailQueryColumnCreateDate    AccountListDetailQueryColumn = "create_date"
	AccountListDetailQueryColumnCreateBy      AccountListDetailQueryColumn = "create_by"
	AccountListDetailQueryColumnLastModDate   AccountListDetailQueryColumn = "last_mod_date"
	AccountListDetailQueryColumnLastModBy     AccountListDetailQueryColumn = "last_mod_by"
	AccountListDetailQueryColumnAccountDesc   AccountListDetailQueryColumn = "account_desc"
	AccountListDetailQueryColumnAccountBal    AccountListDetailQueryColumn = "account_bal"
)

// AccountListDetailQuery represents the supported query parameters for the report.
type AccountListDetailQuery struct {
	AccountType     ReportAccountType
	AccountStatus   ReportAccountStatus
	StartDate       *Date
	EndDate         *Date
	StartModDate    *Date
	EndModDate      *Date
	ModDateMacro    ReportDateMacro
	CreateDateMacro ReportDateMacro
	SortBy          AccountListDetailQueryColumn
	SortOrder       ReportSortOrder
	Columns         []AccountListDetailQueryColumn
}

// AccountListDetailReport represents the QuickBooks account list detail report response.
type AccountListDetailReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// ReportHeader represents report metadata.
type ReportHeader struct {
	Customer           string         `json:"Customer,omitempty"`
	ReportName         string         `json:"ReportName"`
	Vendor             string         `json:"Vendor,omitempty"`
	Item               string         `json:"Item,omitempty"`
	Employee           string         `json:"Employee,omitempty"`
	ReportBasis        ReportBasis    `json:"ReportBasis,omitempty"`
	StartPeriod        string         `json:"StartPeriod,omitempty"`
	EndPeriod          string         `json:"EndPeriod,omitempty"`
	Class              string         `json:"Class,omitempty"`
	Department         string         `json:"Department,omitempty"`
	SummarizeColumnsBy string         `json:"SummarizeColumnsBy,omitempty"`
	Currency           string         `json:"Currency"`
	Option             []ReportOption `json:"Option"`
	Time               time.Time      `json:"Time"`
}

// ReportOption represents a report option entry.
type ReportOption struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// ReportRows represents the top-level report row container.
type ReportRows struct {
	Row []ReportRow `json:"Row"`
}

// ReportRow represents a single report row.
type ReportRow struct {
	ColData []ReportColumnData `json:"ColData,omitempty"`
	Type    ReportRowType      `json:"type"`
	Group   string             `json:"group,omitempty"`
	Rows    *ReportRows        `json:"Rows,omitempty"`
	Header  *ReportRowSection  `json:"Header,omitempty"`
	Summary *ReportRowSection  `json:"Summary,omitempty"`
}

// ReportRowSection represents a report section header or summary row.
type ReportRowSection struct {
	ColData []ReportColumnData `json:"ColData,omitempty"`
}

// ReportColumnData represents a single report cell.
type ReportColumnData struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value"`
	Href  string `json:"href,omitempty"`
}

// ReportColumns represents the top-level report column container.
type ReportColumns struct {
	Column []ReportColumn `json:"Column"`
}

// ReportColumn represents a report column definition.
type ReportColumn struct {
	ColType  ReportColumnType `json:"ColType"`
	ColTitle string           `json:"ColTitle"`
}
