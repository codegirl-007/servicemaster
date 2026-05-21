// Package types contains transport types for external integrations.
package types

// TransactionListByVendorQueryColumn represents the columns query values for the transaction list by vendor report.
type TransactionListByVendorQueryColumn string

const (
	TransactionListByVendorQueryColumnAccountName  TransactionListByVendorQueryColumn = "account_name"
	TransactionListByVendorQueryColumnCreateBy     TransactionListByVendorQueryColumn = "create_by"
	TransactionListByVendorQueryColumnCreateDate   TransactionListByVendorQueryColumn = "create_date"
	TransactionListByVendorQueryColumnCustMsg      TransactionListByVendorQueryColumn = "cust_msg"
	TransactionListByVendorQueryColumnDueDate      TransactionListByVendorQueryColumn = "due_date"
	TransactionListByVendorQueryColumnDocNum       TransactionListByVendorQueryColumn = "doc_num"
	TransactionListByVendorQueryColumnInvDate      TransactionListByVendorQueryColumn = "inv_date"
	TransactionListByVendorQueryColumnIsApPaid     TransactionListByVendorQueryColumn = "is_ap_paid"
	TransactionListByVendorQueryColumnIsCleared    TransactionListByVendorQueryColumn = "is_cleared"
	TransactionListByVendorQueryColumnIsNoPost     TransactionListByVendorQueryColumn = "is_no_post"
	TransactionListByVendorQueryColumnLastModBy    TransactionListByVendorQueryColumn = "last_mod_by"
	TransactionListByVendorQueryColumnMemo         TransactionListByVendorQueryColumn = "memo"
	TransactionListByVendorQueryColumnName         TransactionListByVendorQueryColumn = "name"
	TransactionListByVendorQueryColumnOtherAccount TransactionListByVendorQueryColumn = "other_account"
	TransactionListByVendorQueryColumnPmtMthd      TransactionListByVendorQueryColumn = "pmt_mthd"
	TransactionListByVendorQueryColumnPrinted      TransactionListByVendorQueryColumn = "printed"
	TransactionListByVendorQueryColumnSalesCust1   TransactionListByVendorQueryColumn = "sales_cust1"
	TransactionListByVendorQueryColumnSalesCust2   TransactionListByVendorQueryColumn = "sales_cust2"
	TransactionListByVendorQueryColumnSalesCust3   TransactionListByVendorQueryColumn = "sales_cust3"
	TransactionListByVendorQueryColumnTermName     TransactionListByVendorQueryColumn = "term_name"
	TransactionListByVendorQueryColumnTrackingNum  TransactionListByVendorQueryColumn = "tracking_num"
	TransactionListByVendorQueryColumnTxDate       TransactionListByVendorQueryColumn = "tx_date"
	TransactionListByVendorQueryColumnTxnType      TransactionListByVendorQueryColumn = "txn_type"
	TransactionListByVendorQueryColumnLastModDate  TransactionListByVendorQueryColumn = "last_mod_date"
	TransactionListByVendorQueryColumnPoStatus     TransactionListByVendorQueryColumn = "po_status"
	TransactionListByVendorQueryColumnShipVia      TransactionListByVendorQueryColumn = "ship_via"
	TransactionListByVendorQueryColumnOlbStatus    TransactionListByVendorQueryColumn = "olb_status"
	TransactionListByVendorQueryColumnVendorName   TransactionListByVendorQueryColumn = "vendor_name"
	TransactionListByVendorQueryColumnDeptName     TransactionListByVendorQueryColumn = "dept_name"
)

// TransactionListByVendorReport represents the QuickBooks transaction list by vendor report response.
type TransactionListByVendorReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// TransactionListByVendorQuery represents the supported query parameters for the transaction list by vendor report.
type TransactionListByVendorQuery struct {
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
	Columns           []TransactionListByVendorQueryColumn
	EndDueDate        *Date
	Vendor            string
	EndDate           *Date
	Memo              string
	APPaid            string
	ModDateMacro      ReportDateMacro
	Printed           string
	CreateDateMacro   ReportDateMacro
	Cleared           string
	QZURL             *bool
	Term              string
	EndCreateDate     *Date
	Name              string
	SortBy            TransactionListByVendorQueryColumn
	SortOrder         ReportSortOrder
	StartCreateDate   *Date
	EndModDate        *Date
}
