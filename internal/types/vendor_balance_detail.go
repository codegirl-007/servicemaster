// Package types contains transport types for external integrations.
package types

// VendorBalanceDetailQueryColumn represents the columns query values for the vendor balance detail report.
type VendorBalanceDetailQueryColumn string

const (
	VendorBalanceDetailQueryColumnCreateBy          VendorBalanceDetailQueryColumn = "create_by"
	VendorBalanceDetailQueryColumnCreateDate        VendorBalanceDetailQueryColumn = "create_date"
	VendorBalanceDetailQueryColumnDocNum            VendorBalanceDetailQueryColumn = "doc_num"
	VendorBalanceDetailQueryColumnDueDate           VendorBalanceDetailQueryColumn = "due_date"
	VendorBalanceDetailQueryColumnLastModBy         VendorBalanceDetailQueryColumn = "last_mod_by"
	VendorBalanceDetailQueryColumnLastModDate       VendorBalanceDetailQueryColumn = "last_mod_date"
	VendorBalanceDetailQueryColumnMemo              VendorBalanceDetailQueryColumn = "memo"
	VendorBalanceDetailQueryColumnTermName          VendorBalanceDetailQueryColumn = "term_name"
	VendorBalanceDetailQueryColumnTxDate            VendorBalanceDetailQueryColumn = "tx_date"
	VendorBalanceDetailQueryColumnTxnType           VendorBalanceDetailQueryColumn = "txn_type"
	VendorBalanceDetailQueryColumnVendBillAddr      VendorBalanceDetailQueryColumn = "vend_bill_addr"
	VendorBalanceDetailQueryColumnVendCompName      VendorBalanceDetailQueryColumn = "vend_comp_name"
	VendorBalanceDetailQueryColumnVendName          VendorBalanceDetailQueryColumn = "vend_name"
	VendorBalanceDetailQueryColumnVendPriCont       VendorBalanceDetailQueryColumn = "vend_pri_cont"
	VendorBalanceDetailQueryColumnVendPriEmail      VendorBalanceDetailQueryColumn = "vend_pri_email"
	VendorBalanceDetailQueryColumnVendPriTel        VendorBalanceDetailQueryColumn = "vend_pri_tel"
	VendorBalanceDetailQueryColumnDeptName          VendorBalanceDetailQueryColumn = "dept_name"
	VendorBalanceDetailQueryColumnSubtNegOpenBal    VendorBalanceDetailQueryColumn = "subt_neg_open_bal"
	VendorBalanceDetailQueryColumnSubtNegAmount     VendorBalanceDetailQueryColumn = "subt_neg_amount"
	VendorBalanceDetailQueryColumnCurrency          VendorBalanceDetailQueryColumn = "currency"
	VendorBalanceDetailQueryColumnExchRate          VendorBalanceDetailQueryColumn = "exch_rate"
	VendorBalanceDetailQueryColumnNegForeignOpenBal  VendorBalanceDetailQueryColumn = "neg_foreign_open_bal"
	VendorBalanceDetailQueryColumnSubtNegHomeOpenBal VendorBalanceDetailQueryColumn = "subt_neg_home_open_bal"
	VendorBalanceDetailQueryColumnNegForeignAmount   VendorBalanceDetailQueryColumn = "neg_foreign_amount"
	VendorBalanceDetailQueryColumnSubtNegHomeAmount  VendorBalanceDetailQueryColumn = "subt_neg_home_amount"
)

// VendorBalanceDetailReport represents the QuickBooks vendor balance detail report response.
type VendorBalanceDetailReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// VendorBalanceDetailQuery represents the supported query parameters for the vendor balance detail report.
type VendorBalanceDetailQuery struct {
	Term             string
	EndDueDate       *Date
	AccountingMethod ReportBasis
	DateMacro        ReportDateMacro
	StartDueDate     *Date
	DueDateMacro     ReportDateMacro
	SortBy           VendorBalanceDetailQueryColumn
	ReportDate       *Date
	SortOrder        ReportSortOrder
	APPaid           string
	Department       string
	Vendor           string
	Columns          []VendorBalanceDetailQueryColumn
}
