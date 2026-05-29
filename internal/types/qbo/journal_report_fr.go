// Package qbo contains transport types for the QuickBooks Online API.
package qbo

// JournalReportFRQueryColumn represents the columns query values for the Journal Report FR report.
type JournalReportFRQueryColumn string

const (
	JournalReportFRQueryColumnAcctNumWithExtn JournalReportFRQueryColumn = "acct_num_with_extn"
	JournalReportFRQueryColumnAccountName     JournalReportFRQueryColumn = "account_name"
	JournalReportFRQueryColumnCreditAmt       JournalReportFRQueryColumn = "credit_amt"
	JournalReportFRQueryColumnCreateBy        JournalReportFRQueryColumn = "create_by"
	JournalReportFRQueryColumnCreateDate      JournalReportFRQueryColumn = "create_date"
	JournalReportFRQueryColumnDebitAmt        JournalReportFRQueryColumn = "debt_amt"
	JournalReportFRQueryColumnDocNum          JournalReportFRQueryColumn = "doc_num"
	JournalReportFRQueryColumnDueDate         JournalReportFRQueryColumn = "due_date"
	JournalReportFRQueryColumnIsArPaid        JournalReportFRQueryColumn = "is_ar_paid"
	JournalReportFRQueryColumnIsApPaid        JournalReportFRQueryColumn = "is_ap_paid"
	JournalReportFRQueryColumnItemName        JournalReportFRQueryColumn = "item_name"
	JournalReportFRQueryColumnJournalCodeName JournalReportFRQueryColumn = "journal_code_name"
	JournalReportFRQueryColumnLastModBy       JournalReportFRQueryColumn = "last_mod_by"
	JournalReportFRQueryColumnLastModDate     JournalReportFRQueryColumn = "last_mod_date"
	JournalReportFRQueryColumnMemo            JournalReportFRQueryColumn = "memo"
	JournalReportFRQueryColumnName            JournalReportFRQueryColumn = "name"
	JournalReportFRQueryColumnNegOpenBal      JournalReportFRQueryColumn = "neg_open_bal"
	JournalReportFRQueryColumnPaidDate        JournalReportFRQueryColumn = "paid_date"
	JournalReportFRQueryColumnPmtMthd         JournalReportFRQueryColumn = "pmt_mthd"
	JournalReportFRQueryColumnQuantity        JournalReportFRQueryColumn = "quantity"
	JournalReportFRQueryColumnRate            JournalReportFRQueryColumn = "rate"
	JournalReportFRQueryColumnTxDate          JournalReportFRQueryColumn = "tx_date"
	JournalReportFRQueryColumnTxnNum          JournalReportFRQueryColumn = "txn_num"
	JournalReportFRQueryColumnTxnType         JournalReportFRQueryColumn = "txn_type"
)

// JournalReportFR represents the QuickBooks Journal Report FR report response.
type JournalReportFR struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// JournalReportFRQuery represents the supported query parameters for the Journal Report FR report.
type JournalReportFRQuery struct {
	JournalCode string
	EndDate     *Date
	DateMacro   ReportDateMacro
	SortBy      JournalReportFRQueryColumn
	SortOrder   ReportSortOrder
	StartDate   *Date
	Columns     []JournalReportFRQueryColumn
}
