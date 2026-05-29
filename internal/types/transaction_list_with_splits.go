// Package types contains transport types for external integrations.
package types

// TransactionListWithSplitsQueryColumn represents the columns query values for the transaction list with splits report.
type TransactionListWithSplitsQueryColumn string

const (
	TransactionListWithSplitsQueryColumnTxDate      TransactionListWithSplitsQueryColumn = "tx_date"
	TransactionListWithSplitsQueryColumnTxnType     TransactionListWithSplitsQueryColumn = "txn_type"
	TransactionListWithSplitsQueryColumnDocNum      TransactionListWithSplitsQueryColumn = "doc_num"
	TransactionListWithSplitsQueryColumnIsNoPost    TransactionListWithSplitsQueryColumn = "is_no_post"
	TransactionListWithSplitsQueryColumnAccountName TransactionListWithSplitsQueryColumn = "account_name"
	TransactionListWithSplitsQueryColumnMemo        TransactionListWithSplitsQueryColumn = "memo"
	TransactionListWithSplitsQueryColumnAmount      TransactionListWithSplitsQueryColumn = "amount"
	TransactionListWithSplitsQueryColumnIsAdj       TransactionListWithSplitsQueryColumn = "is_adj"
	TransactionListWithSplitsQueryColumnCreateBy    TransactionListWithSplitsQueryColumn = "create_by"
	TransactionListWithSplitsQueryColumnCreateDate  TransactionListWithSplitsQueryColumn = "create_date"
	TransactionListWithSplitsQueryColumnLastModDate TransactionListWithSplitsQueryColumn = "last_mod_date"
	TransactionListWithSplitsQueryColumnLastModBy   TransactionListWithSplitsQueryColumn = "last_mod_by"
	TransactionListWithSplitsQueryColumnCustName    TransactionListWithSplitsQueryColumn = "cust_name"
	TransactionListWithSplitsQueryColumnVendName    TransactionListWithSplitsQueryColumn = "vend_name"
	TransactionListWithSplitsQueryColumnRate        TransactionListWithSplitsQueryColumn = "rate"
	TransactionListWithSplitsQueryColumnQuantity    TransactionListWithSplitsQueryColumn = "quantity"
	TransactionListWithSplitsQueryColumnItemName    TransactionListWithSplitsQueryColumn = "item_name"
	TransactionListWithSplitsQueryColumnEmpName     TransactionListWithSplitsQueryColumn = "emp_name"
	TransactionListWithSplitsQueryColumnPmtMthd     TransactionListWithSplitsQueryColumn = "pmt_mthd"
	TransactionListWithSplitsQueryColumnNatOpenBal  TransactionListWithSplitsQueryColumn = "nat_open_bal"
	TransactionListWithSplitsQueryColumnTaxType     TransactionListWithSplitsQueryColumn = "tax_type"
	TransactionListWithSplitsQueryColumnIsBillable  TransactionListWithSplitsQueryColumn = "is_billable"
	TransactionListWithSplitsQueryColumnDebitAmt    TransactionListWithSplitsQueryColumn = "debt_amt"
	TransactionListWithSplitsQueryColumnCreditAmt   TransactionListWithSplitsQueryColumn = "credit_amt"
	TransactionListWithSplitsQueryColumnIsCleared   TransactionListWithSplitsQueryColumn = "is_cleared"
	TransactionListWithSplitsQueryColumnOlbStatus   TransactionListWithSplitsQueryColumn = "olb_status"
	TransactionListWithSplitsQueryColumnDeptName    TransactionListWithSplitsQueryColumn = "dept_name"
)

// TransactionListWithSplitsReport represents the QuickBooks transaction list with splits report response.
type TransactionListWithSplitsReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// TransactionListWithSplitsQuery represents the supported query parameters for the transaction list with splits report.
type TransactionListWithSplitsQuery struct {
	DocNum            string
	Name              string
	EndDate           *Date
	DateMacro         ReportDateMacro
	PaymentMethod     string
	SourceAccountType string
	TransactionType   string
	GroupBy           string
	SortBy            TransactionListWithSplitsQueryColumn
	SortOrder         ReportSortOrder
	StartDate         *Date
	Columns           []TransactionListWithSplitsQueryColumn
}
