// Package types contains transport types for external integrations.
package types

// ProfitAndLossDetailQueryColumn represents the columns query values for the profit and loss detail report.
type ProfitAndLossDetailQueryColumn string

const (
	ProfitAndLossDetailQueryColumnCreateBy        ProfitAndLossDetailQueryColumn = "create_by"
	ProfitAndLossDetailQueryColumnCreateDate      ProfitAndLossDetailQueryColumn = "create_date"
	ProfitAndLossDetailQueryColumnDocNum          ProfitAndLossDetailQueryColumn = "doc_num"
	ProfitAndLossDetailQueryColumnLastModBy       ProfitAndLossDetailQueryColumn = "last_mod_by"
	ProfitAndLossDetailQueryColumnLastModDate     ProfitAndLossDetailQueryColumn = "last_mod_date"
	ProfitAndLossDetailQueryColumnMemo            ProfitAndLossDetailQueryColumn = "memo"
	ProfitAndLossDetailQueryColumnName            ProfitAndLossDetailQueryColumn = "name"
	ProfitAndLossDetailQueryColumnPmtMthd         ProfitAndLossDetailQueryColumn = "pmt_mthd"
	ProfitAndLossDetailQueryColumnSplitAcc        ProfitAndLossDetailQueryColumn = "split_acc"
	ProfitAndLossDetailQueryColumnTxDate          ProfitAndLossDetailQueryColumn = "tx_date"
	ProfitAndLossDetailQueryColumnTxnType         ProfitAndLossDetailQueryColumn = "txn_type"
	ProfitAndLossDetailQueryColumnTaxCode         ProfitAndLossDetailQueryColumn = "tax_code"
	ProfitAndLossDetailQueryColumnKlassName       ProfitAndLossDetailQueryColumn = "klass_name"
	ProfitAndLossDetailQueryColumnDeptName        ProfitAndLossDetailQueryColumn = "dept_name"
	ProfitAndLossDetailQueryColumnDebitAmt        ProfitAndLossDetailQueryColumn = "debt_amt"
	ProfitAndLossDetailQueryColumnCreditAmt       ProfitAndLossDetailQueryColumn = "credit_amt"
	ProfitAndLossDetailQueryColumnNatOpenBal      ProfitAndLossDetailQueryColumn = "nat_open_bal"
	ProfitAndLossDetailQueryColumnSubtNatAmount   ProfitAndLossDetailQueryColumn = "subt_nat_amount"
	ProfitAndLossDetailQueryColumnSubtNatAmountNt ProfitAndLossDetailQueryColumn = "subt_nat_amount_nt"
	ProfitAndLossDetailQueryColumnRbalNatAmount   ProfitAndLossDetailQueryColumn = "rbal_nat_amount"
	ProfitAndLossDetailQueryColumnRbalNatAmountNt ProfitAndLossDetailQueryColumn = "rbal_nat_amount_nt"
	ProfitAndLossDetailQueryColumnTaxAmount       ProfitAndLossDetailQueryColumn = "tax_amount"
	ProfitAndLossDetailQueryColumnNetAmount       ProfitAndLossDetailQueryColumn = "net_amount"
)

// ProfitAndLossDetailReport represents the QuickBooks profit and loss detail report response.
type ProfitAndLossDetailReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// ProfitAndLossDetailQuery represents the supported query parameters for the profit and loss detail report.
type ProfitAndLossDetailQuery struct {
	Customer         string
	Account          string
	AccountingMethod ReportBasis
	EndDate          *Date
	DateMacro        ReportDateMacro
	AdjustedGainLoss *bool
	Class            string
	SortBy           ProfitAndLossDetailQueryColumn
	PaymentMethod    string
	SortOrder        ReportSortOrder
	Employee         string
	Department       string
	Vendor           string
	AccountType      ReportAccountType
	StartDate        *Date
	Columns          []ProfitAndLossDetailQueryColumn
}
