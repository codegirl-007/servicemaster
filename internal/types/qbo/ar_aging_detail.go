// Package qbo contains transport types for the QuickBooks Online API.
package qbo

// ARAgingDetailQueryColumn represents the columns query values for the AR aging detail report.
type ARAgingDetailQueryColumn string

const (
	ARAgingDetailQueryColumnBillAddr       ARAgingDetailQueryColumn = "bill_addr"
	ARAgingDetailQueryColumnCreateBy       ARAgingDetailQueryColumn = "create_by"
	ARAgingDetailQueryColumnCreateDate     ARAgingDetailQueryColumn = "create_date"
	ARAgingDetailQueryColumnCustBillEmail  ARAgingDetailQueryColumn = "cust_bill_email"
	ARAgingDetailQueryColumnCustCompName   ARAgingDetailQueryColumn = "cust_comp_name"
	ARAgingDetailQueryColumnCustMsg        ARAgingDetailQueryColumn = "cust_msg"
	ARAgingDetailQueryColumnCustName       ARAgingDetailQueryColumn = "cust_name"
	ARAgingDetailQueryColumnDelivAddr      ARAgingDetailQueryColumn = "deliv_addr"
	ARAgingDetailQueryColumnDocNum         ARAgingDetailQueryColumn = "doc_num"
	ARAgingDetailQueryColumnDueDate        ARAgingDetailQueryColumn = "due_date"
	ARAgingDetailQueryColumnLastModBy      ARAgingDetailQueryColumn = "last_mod_by"
	ARAgingDetailQueryColumnLastModDate    ARAgingDetailQueryColumn = "last_mod_date"
	ARAgingDetailQueryColumnMemo           ARAgingDetailQueryColumn = "memo"
	ARAgingDetailQueryColumnPastDue        ARAgingDetailQueryColumn = "past_due"
	ARAgingDetailQueryColumnSaleSentState  ARAgingDetailQueryColumn = "sale_sent_state"
	ARAgingDetailQueryColumnShipAddr       ARAgingDetailQueryColumn = "ship_addr"
	ARAgingDetailQueryColumnTermName       ARAgingDetailQueryColumn = "term_name"
	ARAgingDetailQueryColumnTxDate         ARAgingDetailQueryColumn = "tx_date"
	ARAgingDetailQueryColumnTxnType        ARAgingDetailQueryColumn = "txn_type"
	ARAgingDetailQueryColumnSalesCust1     ARAgingDetailQueryColumn = "sales_cust1"
	ARAgingDetailQueryColumnSalesCust2     ARAgingDetailQueryColumn = "sales_cust2"
	ARAgingDetailQueryColumnSalesCust3     ARAgingDetailQueryColumn = "sales_cust3"
	ARAgingDetailQueryColumnDeptName       ARAgingDetailQueryColumn = "dept_name"
	ARAgingDetailQueryColumnSubtOpenBal    ARAgingDetailQueryColumn = "subt_open_bal"
	ARAgingDetailQueryColumnSubtAmount     ARAgingDetailQueryColumn = "subt_amount"
	ARAgingDetailQueryColumnCurrency       ARAgingDetailQueryColumn = "currency"
	ARAgingDetailQueryColumnExchRate       ARAgingDetailQueryColumn = "exch_rate"
	ARAgingDetailQueryColumnForeignOpenBal ARAgingDetailQueryColumn = "foreign_open_bal"
	ARAgingDetailQueryColumnForeignAmount  ARAgingDetailQueryColumn = "foreign_amount"
)

// ARAgingDetailReport represents the QuickBooks AR aging detail report response.
type ARAgingDetailReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// ARAgingDetailQuery represents the supported query parameters for the AR aging detail report.
type ARAgingDetailQuery struct {
	Customer     string
	ShipVia      string
	Term         string
	EndDueDate   *Date
	StartDueDate *Date
	Custom1      string
	Custom2      string
	Custom3      string
	ReportDate   *Date
	NumPeriods   int
	AgingMethod  string
	PastDue      int
	AgingPeriod  float64
	Columns      []ARAgingDetailQueryColumn
}
