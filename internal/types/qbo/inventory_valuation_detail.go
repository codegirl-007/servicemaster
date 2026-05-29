// Package qbo contains transport types for the QuickBooks Online API.
package qbo

// InventoryValuationDetailQueryColumn represents the columns query values for the inventory valuation detail report.
type InventoryValuationDetailQueryColumn string

const (
	InventoryValuationDetailQueryColumnTxDate        InventoryValuationDetailQueryColumn = "tx_date"
	InventoryValuationDetailQueryColumnTxnId         InventoryValuationDetailQueryColumn = "txn_id"
	InventoryValuationDetailQueryColumnTxnType       InventoryValuationDetailQueryColumn = "txn_type"
	InventoryValuationDetailQueryColumnDocNum        InventoryValuationDetailQueryColumn = "doc_num"
	InventoryValuationDetailQueryColumnName          InventoryValuationDetailQueryColumn = "name"
	InventoryValuationDetailQueryColumnQuantity      InventoryValuationDetailQueryColumn = "quantity"
	InventoryValuationDetailQueryColumnRate          InventoryValuationDetailQueryColumn = "rate"
	InventoryValuationDetailQueryColumnHomeAmount    InventoryValuationDetailQueryColumn = "home_amount"
	InventoryValuationDetailQueryColumnQtyOnHand     InventoryValuationDetailQueryColumn = "qty_on_hand"
	InventoryValuationDetailQueryColumnAssetValue    InventoryValuationDetailQueryColumn = "asset_value"
	InventoryValuationDetailQueryColumnCreateDate    InventoryValuationDetailQueryColumn = "create_date"
	InventoryValuationDetailQueryColumnCreateBy      InventoryValuationDetailQueryColumn = "create_by"
	InventoryValuationDetailQueryColumnLastModDate   InventoryValuationDetailQueryColumn = "last_mod_date"
	InventoryValuationDetailQueryColumnLastModBy     InventoryValuationDetailQueryColumn = "last_mod_by"
	InventoryValuationDetailQueryColumnItemSku       InventoryValuationDetailQueryColumn = "item_sku"
	InventoryValuationDetailQueryColumnMemo          InventoryValuationDetailQueryColumn = "memo"
	InventoryValuationDetailQueryColumnExchRate      InventoryValuationDetailQueryColumn = "exch_rate"
	InventoryValuationDetailQueryColumnAccountName   InventoryValuationDetailQueryColumn = "account_name"
	InventoryValuationDetailQueryColumnServiceDate   InventoryValuationDetailQueryColumn = "service_date"
	InventoryValuationDetailQueryColumnRateInventory InventoryValuationDetailQueryColumn = "rate_inventory"
	InventoryValuationDetailQueryColumnAssetValueNt  InventoryValuationDetailQueryColumn = "asset_value_nt"
	InventoryValuationDetailQueryColumnTrackingNum   InventoryValuationDetailQueryColumn = "tracking_num"
)

// InventoryValuationDetailReport represents the QuickBooks inventory valuation detail report response.
type InventoryValuationDetailReport struct {
	Header  ReportHeader  `json:"Header"`
	Rows    ReportRows    `json:"Rows"`
	Columns ReportColumns `json:"Columns"`
}

// InventoryValuationDetailQuery represents the supported query parameters for the inventory valuation detail report.
type InventoryValuationDetailQuery struct {
	EndDate      *Date
	EndSvcDate   *Date
	DateMacro    ReportDateMacro
	SvcDateMacro ReportDateMacro
	StartSvcDate *Date
	GroupBy      string
	StartDate    *Date
	Columns      []InventoryValuationDetailQueryColumn
}
