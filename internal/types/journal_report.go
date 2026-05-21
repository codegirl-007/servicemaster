// Package types contains transport types for external integrations.
package types

// JournalReportQueryColumn represents the columns query values for the journal report.
type JournalReportQueryColumn string

const (
	JournalReportQueryColumnAcctNumWithExtn JournalReportQueryColumn = "acct_num_with_extn"
	JournalReportQueryColumnAccountName     JournalReportQueryColumn = "account_name"
	JournalReportQueryColumnCreditAmt       JournalReportQueryColumn = "credit_amt"
	JournalReportQueryColumnCreateBy        JournalReportQueryColumn = "create_by"
	JournalReportQueryColumnCreateDate      JournalReportQueryColumn = "create_date"
	JournalReportQueryColumnDebtAmt         JournalReportQueryColumn = "debt_amt"
	JournalReportQueryColumnDocNum          JournalReportQueryColumn = "doc_num"
	JournalReportQueryColumnDueDate         JournalReportQueryColumn = "due_date"
	JournalReportQueryColumnIsArPaid        JournalReportQueryColumn = "is_ar_paid"
	JournalReportQueryColumnIsApPaid        JournalReportQueryColumn = "is_ap_paid"
	JournalReportQueryColumnItemName        JournalReportQueryColumn = "item_name"
	JournalReportQueryColumnJournalCodeName JournalReportQueryColumn = "journal_code_name"
	JournalReportQueryColumnLastModBy       JournalReportQueryColumn = "last_mod_by"
	JournalReportQueryColumnLastModDate     JournalReportQueryColumn = "last_mod_date"
	JournalReportQueryColumnMemo            JournalReportQueryColumn = "memo"
	JournalReportQueryColumnName            JournalReportQueryColumn = "name"
	JournalReportQueryColumnNegOpenBal      JournalReportQueryColumn = "neg_open_bal"
	JournalReportQueryColumnPaidDate        JournalReportQueryColumn = "paid_date"
	JournalReportQueryColumnPmtMthd         JournalReportQueryColumn = "pmt_mthd"
	JournalReportQueryColumnQuantity        JournalReportQueryColumn = "quantity"
	JournalReportQueryColumnRate            JournalReportQueryColumn = "rate"
	JournalReportQueryColumnTxDate          JournalReportQueryColumn = "tx_date"
	JournalReportQueryColumnTxnNum          JournalReportQueryColumn = "txn_num"
	JournalReportQueryColumnTxnType         JournalReportQueryColumn = "txn_type"
)

// JournalReport represents the QuickBooks journal report response.
type JournalReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// JournalReportQuery represents the supported query parameters for the journal report.
type JournalReportQuery struct {
	EndDate   *Date
	DateMacro ReportDateMacro
	SortBy    JournalReportQueryColumn
	SortOrder ReportSortOrder
	StartDate *Date
	Columns   []JournalReportQueryColumn
}
