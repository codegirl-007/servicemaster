// Package qbo contains transport types for the QuickBooks Online API.
package qbo

// APAgingDetailReport represents the QuickBooks AP aging detail report response.
type APAgingDetailReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// APAgingDetailQueryColumn represents the columns query values for the AP aging detail report.
type APAgingDetailQueryColumn string

const (
	APAgingDetailQueryColumnCreateBy           APAgingDetailQueryColumn = "create_by"
	APAgingDetailQueryColumnCreateDate         APAgingDetailQueryColumn = "create_date"
	APAgingDetailQueryColumnDocNum             APAgingDetailQueryColumn = "doc_num"
	APAgingDetailQueryColumnDueDate            APAgingDetailQueryColumn = "due_date"
	APAgingDetailQueryColumnLastModBy          APAgingDetailQueryColumn = "last_mod_by"
	APAgingDetailQueryColumnLastModDate        APAgingDetailQueryColumn = "last_mod_date"
	APAgingDetailQueryColumnMemo               APAgingDetailQueryColumn = "memo"
	APAgingDetailQueryColumnPastDue            APAgingDetailQueryColumn = "past_due"
	APAgingDetailQueryColumnTermName           APAgingDetailQueryColumn = "term_name"
	APAgingDetailQueryColumnTxDate             APAgingDetailQueryColumn = "tx_date"
	APAgingDetailQueryColumnTxnType            APAgingDetailQueryColumn = "txn_type"
	APAgingDetailQueryColumnVendBillAddr       APAgingDetailQueryColumn = "vend_bill_addr"
	APAgingDetailQueryColumnVendCompName       APAgingDetailQueryColumn = "vend_comp_name"
	APAgingDetailQueryColumnVendName           APAgingDetailQueryColumn = "vend_name"
	APAgingDetailQueryColumnVendPriCont        APAgingDetailQueryColumn = "vend_pri_cont"
	APAgingDetailQueryColumnVendPriEmail       APAgingDetailQueryColumn = "vend_pri_email"
	APAgingDetailQueryColumnVendPriTel         APAgingDetailQueryColumn = "vend_pri_tel"
	APAgingDetailQueryColumnDeptName           APAgingDetailQueryColumn = "dept_name"
	APAgingDetailQueryColumnCurrency           APAgingDetailQueryColumn = "currency"
	APAgingDetailQueryColumnExchangeRate       APAgingDetailQueryColumn = "exch_rate"
	APAgingDetailQueryColumnNegForeignOpenBal  APAgingDetailQueryColumn = "neg_foreign_open_bal"
	APAgingDetailQueryColumnSubtNegHomeOpenBal APAgingDetailQueryColumn = "subt_neg_home_open_bal"
	APAgingDetailQueryColumnNegForeignAmount   APAgingDetailQueryColumn = "neg_foreign_amount"
	APAgingDetailQueryColumnSubtNegHomeAmount  APAgingDetailQueryColumn = "subt_neg_home_amount"
	APAgingDetailQueryColumnSubtNegOpenBal     APAgingDetailQueryColumn = "subt_neg_open_bal"
	APAgingDetailQueryColumnSubtNegAmount      APAgingDetailQueryColumn = "subt_neg_amount"
)

// APAgingDetailQuery represents the supported query parameters for the AP aging detail report.
type APAgingDetailQuery struct {
	ShipVia          string
	Term             string
	StartDueDate     *Date
	EndDueDate       *Date
	AccountingMethod ReportBasis
	Custom1          string
	Custom2          string
	Custom3          string
	ReportDate       *Date
	NumPeriods       int
	Vendor           string
	PastDue          int
	AgingPeriod      float64
	Columns          []APAgingDetailQueryColumn
}
