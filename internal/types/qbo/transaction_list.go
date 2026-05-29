// Package qbo contains transport types for the QuickBooks Online API.
package qbo

// TransactionListQueryColumn represents the columns query values for the transaction list report.
type TransactionListQueryColumn string

const (
	TransactionListQueryColumnAccountName  TransactionListQueryColumn = "account_name"
	TransactionListQueryColumnCreateBy     TransactionListQueryColumn = "create_by"
	TransactionListQueryColumnCreateDate   TransactionListQueryColumn = "create_date"
	TransactionListQueryColumnCustMsg      TransactionListQueryColumn = "cust_msg"
	TransactionListQueryColumnDueDate      TransactionListQueryColumn = "due_date"
	TransactionListQueryColumnDocNum       TransactionListQueryColumn = "doc_num"
	TransactionListQueryColumnInvDate      TransactionListQueryColumn = "inv_date"
	TransactionListQueryColumnIsApPaid     TransactionListQueryColumn = "is_ap_paid"
	TransactionListQueryColumnIsCleared    TransactionListQueryColumn = "is_cleared"
	TransactionListQueryColumnIsNoPost     TransactionListQueryColumn = "is_no_post"
	TransactionListQueryColumnLastModBy    TransactionListQueryColumn = "last_mod_by"
	TransactionListQueryColumnMemo         TransactionListQueryColumn = "memo"
	TransactionListQueryColumnName         TransactionListQueryColumn = "name"
	TransactionListQueryColumnOtherAccount TransactionListQueryColumn = "other_account"
	TransactionListQueryColumnPmtMthd      TransactionListQueryColumn = "pmt_mthd"
	TransactionListQueryColumnPrinted      TransactionListQueryColumn = "printed"
	TransactionListQueryColumnSalesCust1   TransactionListQueryColumn = "sales_cust1"
	TransactionListQueryColumnSalesCust2   TransactionListQueryColumn = "sales_cust2"
	TransactionListQueryColumnSalesCust3   TransactionListQueryColumn = "sales_cust3"
	TransactionListQueryColumnTermName     TransactionListQueryColumn = "term_name"
	TransactionListQueryColumnTrackingNum  TransactionListQueryColumn = "tracking_num"
	TransactionListQueryColumnTxDate       TransactionListQueryColumn = "tx_date"
	TransactionListQueryColumnTxnType      TransactionListQueryColumn = "txn_type"
	TransactionListQueryColumnIsAdj        TransactionListQueryColumn = "is_adj"
	TransactionListQueryColumnLastModDate  TransactionListQueryColumn = "last_mod_date"
	TransactionListQueryColumnShipVia      TransactionListQueryColumn = "ship_via"
	TransactionListQueryColumnOlbStatus    TransactionListQueryColumn = "olb_status"
	TransactionListQueryColumnExtraDocNum  TransactionListQueryColumn = "extra_doc_num"
	TransactionListQueryColumnIsArPaid     TransactionListQueryColumn = "is_ar_paid"
	TransactionListQueryColumnDeptName     TransactionListQueryColumn = "dept_name"
)

// TransactionListReport represents the QuickBooks transaction list report response.
type TransactionListReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// TransactionListQuery represents the supported query parameters for the transaction list report.
type TransactionListQuery struct {
	DateMacro         ReportDateMacro
	PaymentMethod     ReportPaymentMethod
	DueDateMacro      ReportDateMacro
	ARPaid            ReportARPaid
	BothAmount        string
	TransactionType   ReportTransactionType
	DocNum            string
	StartModDate      *Date
	SourceAccountType ReportAccountType
	GroupBy           ReportGroupBy
	StartDate         *Date
	Department        string
	StartDueDate      *Date
	Columns           []TransactionListQueryColumn
	EndDueDate        *Date
	Vendor            string
	EndDate           *Date
	Memo              string
	APPaid            ReportAPPaid
	ModDateMacro      ReportDateMacro
	Printed           ReportPrinted
	CreateDateMacro   ReportDateMacro
	Cleared           ReportCleared
	Customer          string
	QZURL             *bool
	Term              string
	EndCreateDate     *Date
	Name              string
	SortBy            TransactionListQueryColumn
	SortOrder         ReportSortOrder
	StartCreateDate   *Date
	EndModDate        *Date
}
