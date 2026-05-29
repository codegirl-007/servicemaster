// Package types contains transport types for external integrations.
package types

// TransactionListByCustomerQueryColumn represents the columns query values for the transaction list by customer report.
type TransactionListByCustomerQueryColumn string

const (
	TransactionListByCustomerQueryColumnAccountName  TransactionListByCustomerQueryColumn = "account_name"
	TransactionListByCustomerQueryColumnAmount       TransactionListByCustomerQueryColumn = "amount"
	TransactionListByCustomerQueryColumnCreateBy     TransactionListByCustomerQueryColumn = "create_by"
	TransactionListByCustomerQueryColumnCreateDate   TransactionListByCustomerQueryColumn = "create_date"
	TransactionListByCustomerQueryColumnCustMsg      TransactionListByCustomerQueryColumn = "cust_msg"
	TransactionListByCustomerQueryColumnDueDate      TransactionListByCustomerQueryColumn = "due_date"
	TransactionListByCustomerQueryColumnDocNum       TransactionListByCustomerQueryColumn = "doc_num"
	TransactionListByCustomerQueryColumnInvDate      TransactionListByCustomerQueryColumn = "inv_date"
	TransactionListByCustomerQueryColumnIsApPaid     TransactionListByCustomerQueryColumn = "is_ap_paid"
	TransactionListByCustomerQueryColumnIsCleared    TransactionListByCustomerQueryColumn = "is_cleared"
	TransactionListByCustomerQueryColumnIsNoPost     TransactionListByCustomerQueryColumn = "is_no_post"
	TransactionListByCustomerQueryColumnLastModBy    TransactionListByCustomerQueryColumn = "last_mod_by"
	TransactionListByCustomerQueryColumnMemo         TransactionListByCustomerQueryColumn = "memo"
	TransactionListByCustomerQueryColumnName         TransactionListByCustomerQueryColumn = "name"
	TransactionListByCustomerQueryColumnOtherAccount TransactionListByCustomerQueryColumn = "other_account"
	TransactionListByCustomerQueryColumnPmtMthd      TransactionListByCustomerQueryColumn = "pmt_mthd"
	TransactionListByCustomerQueryColumnPrinted      TransactionListByCustomerQueryColumn = "printed"
	TransactionListByCustomerQueryColumnSalesCust1   TransactionListByCustomerQueryColumn = "sales_cust1"
	TransactionListByCustomerQueryColumnSalesCust2   TransactionListByCustomerQueryColumn = "sales_cust2"
	TransactionListByCustomerQueryColumnSalesCust3   TransactionListByCustomerQueryColumn = "sales_cust3"
	TransactionListByCustomerQueryColumnTermName     TransactionListByCustomerQueryColumn = "term_name"
	TransactionListByCustomerQueryColumnTrackingNum  TransactionListByCustomerQueryColumn = "tracking_num"
	TransactionListByCustomerQueryColumnTxDate       TransactionListByCustomerQueryColumn = "tx_date"
	TransactionListByCustomerQueryColumnTxnType      TransactionListByCustomerQueryColumn = "txn_type"
	TransactionListByCustomerQueryColumnLastModDate  TransactionListByCustomerQueryColumn = "last_mod_date"
	TransactionListByCustomerQueryColumnShipVia      TransactionListByCustomerQueryColumn = "ship_via"
	TransactionListByCustomerQueryColumnOlbStatus    TransactionListByCustomerQueryColumn = "olb_status"
	TransactionListByCustomerQueryColumnIsArPaid     TransactionListByCustomerQueryColumn = "is_ar_paid"
	TransactionListByCustomerQueryColumnExtraDocNum  TransactionListByCustomerQueryColumn = "extra_doc_num"
	TransactionListByCustomerQueryColumnCustName     TransactionListByCustomerQueryColumn = "cust_name"
	TransactionListByCustomerQueryColumnDeptName     TransactionListByCustomerQueryColumn = "dept_name"
)

// TransactionListByCustomerReport represents the QuickBooks transaction list by customer report response.
type TransactionListByCustomerReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// TransactionListByCustomerQuery represents the supported query parameters for the transaction list by customer report.
type TransactionListByCustomerQuery struct {
	DateMacro         ReportDateMacro
	PaymentMethod     string
	DueDateMacro      ReportDateMacro
	ARPaid            string
	BothAmount        string
	TransactionType   string
	DocNum            string
	StartModDate      *Date
	SourceAccountType string
	GroupBy           string
	StartDate         *Date
	Department        string
	StartDueDate      *Date
	Columns           []TransactionListByCustomerQueryColumn
	EndDueDate        *Date
	EndDate           *Date
	Memo              string
	APPaid            string
	ModDateMacro      ReportDateMacro
	Printed           string
	CreateDateMacro   ReportDateMacro
	Cleared           string
	Customer          string
	QZURL             *bool
	Term              string
	EndCreateDate     *Date
	Name              string
	SortBy            TransactionListByCustomerQueryColumn
	SortOrder         ReportSortOrder
	StartCreateDate   *Date
	EndModDate        *Date
}
