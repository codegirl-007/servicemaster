// Package qbo contains transport types for the QuickBooks Online API.
package qbo

// CustomerBalanceDetailQueryColumn represents the columns query values for the customer balance detail report.
type CustomerBalanceDetailQueryColumn string

const (
	CustomerBalanceDetailQueryColumnBillAddr       CustomerBalanceDetailQueryColumn = "bill_addr"
	CustomerBalanceDetailQueryColumnCreateBy       CustomerBalanceDetailQueryColumn = "create_by"
	CustomerBalanceDetailQueryColumnCreateDate     CustomerBalanceDetailQueryColumn = "create_date"
	CustomerBalanceDetailQueryColumnCustBillEmail  CustomerBalanceDetailQueryColumn = "cust_bill_email"
	CustomerBalanceDetailQueryColumnCustCompName   CustomerBalanceDetailQueryColumn = "cust_comp_name"
	CustomerBalanceDetailQueryColumnCustMsg        CustomerBalanceDetailQueryColumn = "cust_msg"
	CustomerBalanceDetailQueryColumnCustPhoneOther CustomerBalanceDetailQueryColumn = "cust_phone_other"
	CustomerBalanceDetailQueryColumnCustTel        CustomerBalanceDetailQueryColumn = "cust_tel"
	CustomerBalanceDetailQueryColumnCustName       CustomerBalanceDetailQueryColumn = "cust_name"
	CustomerBalanceDetailQueryColumnDelivAddr      CustomerBalanceDetailQueryColumn = "deliv_addr"
	CustomerBalanceDetailQueryColumnDocNum         CustomerBalanceDetailQueryColumn = "doc_num"
	CustomerBalanceDetailQueryColumnDueDate        CustomerBalanceDetailQueryColumn = "due_date"
	CustomerBalanceDetailQueryColumnLastModBy      CustomerBalanceDetailQueryColumn = "last_mod_by"
	CustomerBalanceDetailQueryColumnLastModDate    CustomerBalanceDetailQueryColumn = "last_mod_date"
	CustomerBalanceDetailQueryColumnMemo           CustomerBalanceDetailQueryColumn = "memo"
	CustomerBalanceDetailQueryColumnSaleSentState  CustomerBalanceDetailQueryColumn = "sale_sent_state"
	CustomerBalanceDetailQueryColumnShipAddr       CustomerBalanceDetailQueryColumn = "ship_addr"
	CustomerBalanceDetailQueryColumnShipDate       CustomerBalanceDetailQueryColumn = "ship_date"
	CustomerBalanceDetailQueryColumnShipVia        CustomerBalanceDetailQueryColumn = "ship_via"
	CustomerBalanceDetailQueryColumnTermName       CustomerBalanceDetailQueryColumn = "term_name"
	CustomerBalanceDetailQueryColumnTrackingNum    CustomerBalanceDetailQueryColumn = "tracking_num"
	CustomerBalanceDetailQueryColumnTxDate         CustomerBalanceDetailQueryColumn = "tx_date"
	CustomerBalanceDetailQueryColumnTxnType        CustomerBalanceDetailQueryColumn = "txn_type"
	CustomerBalanceDetailQueryColumnSalesCust1     CustomerBalanceDetailQueryColumn = "sales_cust1"
	CustomerBalanceDetailQueryColumnSalesCust2     CustomerBalanceDetailQueryColumn = "sales_cust2"
	CustomerBalanceDetailQueryColumnSalesCust3     CustomerBalanceDetailQueryColumn = "sales_cust3"
	CustomerBalanceDetailQueryColumnDeptName       CustomerBalanceDetailQueryColumn = "dept_name"
	CustomerBalanceDetailQueryColumnSubtOpenBal    CustomerBalanceDetailQueryColumn = "subt_open_bal"
	CustomerBalanceDetailQueryColumnSubtAmount     CustomerBalanceDetailQueryColumn = "subt_amount"
	CustomerBalanceDetailQueryColumnCurrency       CustomerBalanceDetailQueryColumn = "currency"
	CustomerBalanceDetailQueryColumnExchRate       CustomerBalanceDetailQueryColumn = "exch_rate"
	CustomerBalanceDetailQueryColumnForeignOpenBal CustomerBalanceDetailQueryColumn = "foreign_open_bal"
	CustomerBalanceDetailQueryColumnForeignAmount  CustomerBalanceDetailQueryColumn = "foreign_amount"
)

// CustomerBalanceDetailReport represents the QuickBooks customer balance detail report response.
type CustomerBalanceDetailReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// CustomerBalanceDetailQuery represents the supported query parameters for the customer balance detail report.
type CustomerBalanceDetailQuery struct {
	Customer     string
	ShipVia      string
	Term         string
	EndDueDate   *Date
	StartDueDate *Date
	Custom1      string
	SortBy       CustomerBalanceDetailQueryColumn
	ARPaid       string
	ReportDate   *Date
	SortOrder    ReportSortOrder
	AgingMethod  string
	Department   string
	Columns      []CustomerBalanceDetailQueryColumn
}
