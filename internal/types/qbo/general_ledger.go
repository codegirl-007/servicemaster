// Package qbo contains transport types for the QuickBooks Online API.
package qbo

// GeneralLedgerQueryColumn represents the columns query values for the general ledger report.
type GeneralLedgerQueryColumn string

const (
	GeneralLedgerQueryColumnAccountName     GeneralLedgerQueryColumn = "account_name"
	GeneralLedgerQueryColumnChkPrintState   GeneralLedgerQueryColumn = "chk_print_state"
	GeneralLedgerQueryColumnCreateBy        GeneralLedgerQueryColumn = "create_by"
	GeneralLedgerQueryColumnCreateDate      GeneralLedgerQueryColumn = "create_date"
	GeneralLedgerQueryColumnCustName        GeneralLedgerQueryColumn = "cust_name"
	GeneralLedgerQueryColumnDocNum          GeneralLedgerQueryColumn = "doc_num"
	GeneralLedgerQueryColumnEmpName         GeneralLedgerQueryColumn = "emp_name"
	GeneralLedgerQueryColumnInvDate         GeneralLedgerQueryColumn = "inv_date"
	GeneralLedgerQueryColumnIsAdj           GeneralLedgerQueryColumn = "is_adj"
	GeneralLedgerQueryColumnIsApPaid        GeneralLedgerQueryColumn = "is_ap_paid"
	GeneralLedgerQueryColumnIsArPaid        GeneralLedgerQueryColumn = "is_ar_paid"
	GeneralLedgerQueryColumnIsCleared       GeneralLedgerQueryColumn = "is_cleared"
	GeneralLedgerQueryColumnItemName        GeneralLedgerQueryColumn = "item_name"
	GeneralLedgerQueryColumnLastModBy       GeneralLedgerQueryColumn = "last_mod_by"
	GeneralLedgerQueryColumnLastModDate     GeneralLedgerQueryColumn = "last_mod_date"
	GeneralLedgerQueryColumnMemo            GeneralLedgerQueryColumn = "memo"
	GeneralLedgerQueryColumnName            GeneralLedgerQueryColumn = "name"
	GeneralLedgerQueryColumnQuantity        GeneralLedgerQueryColumn = "quantity"
	GeneralLedgerQueryColumnRate            GeneralLedgerQueryColumn = "rate"
	GeneralLedgerQueryColumnSplitAcc        GeneralLedgerQueryColumn = "split_acc"
	GeneralLedgerQueryColumnTxDate          GeneralLedgerQueryColumn = "tx_date"
	GeneralLedgerQueryColumnTxnType         GeneralLedgerQueryColumn = "txn_type"
	GeneralLedgerQueryColumnVendName        GeneralLedgerQueryColumn = "vend_name"
	GeneralLedgerQueryColumnNetAmount       GeneralLedgerQueryColumn = "net_amount"
	GeneralLedgerQueryColumnTaxAmount       GeneralLedgerQueryColumn = "tax_amount"
	GeneralLedgerQueryColumnTaxCode         GeneralLedgerQueryColumn = "tax_code"
	GeneralLedgerQueryColumnAccountNum      GeneralLedgerQueryColumn = "account_num"
	GeneralLedgerQueryColumnKlassName       GeneralLedgerQueryColumn = "klass_name"
	GeneralLedgerQueryColumnDeptName        GeneralLedgerQueryColumn = "dept_name"
	GeneralLedgerQueryColumnDebitAmt        GeneralLedgerQueryColumn = "debt_amt"
	GeneralLedgerQueryColumnCreditAmt       GeneralLedgerQueryColumn = "credit_amt"
	GeneralLedgerQueryColumnNatOpenBal      GeneralLedgerQueryColumn = "nat_open_bal"
	GeneralLedgerQueryColumnSubtNatAmount   GeneralLedgerQueryColumn = "subt_nat_amount"
	GeneralLedgerQueryColumnSubtNatAmountNt GeneralLedgerQueryColumn = "subt_nat_amount_nt"
	GeneralLedgerQueryColumnRbalNatAmount   GeneralLedgerQueryColumn = "rbal_nat_amount"
	GeneralLedgerQueryColumnRbalNatAmountNt GeneralLedgerQueryColumn = "rbal_nat_amount_nt"
)

// GeneralLedgerReport represents the QuickBooks general ledger report response.
type GeneralLedgerReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// GeneralLedgerQuery represents the supported query parameters for the general ledger report.
type GeneralLedgerQuery struct {
	Customer          string
	Account           string
	AccountingMethod  ReportBasis
	SourceAccount     string
	EndDate           *Date
	DateMacro         ReportDateMacro
	AccountType       ReportAccountType
	SortBy            GeneralLedgerQueryColumn
	SortOrder         ReportSortOrder
	StartDate         *Date
	SummarizeColumnBy SummarizeColumnBy
	Department        string
	Vendor            string
	Class             string
	Columns           []GeneralLedgerQueryColumn
}
