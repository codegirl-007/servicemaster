// Package types contains transport types for external integrations.
package types

import "time"

// BillLineDetailType represents the documented QuickBooks Bill line detail types.
type BillLineDetailType string

const (
	BillLineDetailTypeAccountBasedExpense BillLineDetailType = "AccountBasedExpenseLineDetail"
	BillLineDetailTypeItemBasedExpense    BillLineDetailType = "ItemBasedExpenseLineDetail"
	BillLineDetailTypeTaxLine             BillLineDetailType = "TaxLineDetail"
)

// BillableStatus represents the documented billable status values.
type BillableStatus string

const (
	BillableStatusBillable      BillableStatus = "Billable"
	BillableStatusNotBillable   BillableStatus = "NotBillable"
	BillableStatusHasBeenBilled BillableStatus = "HasBeenBilled"
)

// GlobalTaxCalculation represents the documented tax application methods.
type GlobalTaxCalculation string

const (
	GlobalTaxCalculationTaxExcluded   GlobalTaxCalculation = "TaxExcluded"
	GlobalTaxCalculationTaxInclusive  GlobalTaxCalculation = "TaxInclusive"
	GlobalTaxCalculationNotApplicable GlobalTaxCalculation = "NotApplicable"
)

// BillLinkedTxnType represents the documented linked transaction types for bills.
type BillLinkedTxnType string

const (
	BillLinkedTxnTypePurchaseOrder    BillLinkedTxnType = "PurchaseOrder"
	BillLinkedTxnTypeBillPaymentCheck BillLinkedTxnType = "BillPaymentCheck"
	BillLinkedTxnTypeReimburseCharge  BillLinkedTxnType = "ReimburseCharge"
)

// TransactionLocationType represents the FR-only transaction location values.
type TransactionLocationType string

const (
	TransactionLocationTypeWithinFrance        TransactionLocationType = "WithinFrance"
	TransactionLocationTypeFranceOverseas      TransactionLocationType = "FranceOverseas"
	TransactionLocationTypeOutsideFranceWithEU TransactionLocationType = "OutsideFranceWithEU"
	TransactionLocationTypeOutsideEU           TransactionLocationType = "OutsideEU"
)

// BillResponse represents the QuickBooks bill response envelope.
type BillResponse struct {
	Bill Bill      `json:"Bill"`
	Time time.Time `json:"time"`
}

// Bill represents a QuickBooks bill object.
type Bill struct {
	ID                      string                  `json:"Id"`
	VendorRef               Reference               `json:"VendorRef"`
	Line                    []BillLine              `json:"Line"`
	SyncToken               string                  `json:"SyncToken"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	HomeBalance             float64                 `json:"HomeBalance,omitempty"`
	RecurDataRef            *Reference              `json:"RecurDataRef,omitempty"`
	Balance                 float64                 `json:"Balance,omitempty"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	APAccountRef            *Reference              `json:"APAccountRef,omitempty"`
	SalesTermRef            *Reference              `json:"SalesTermRef,omitempty"`
	LinkedTxn               []LinkedTxn             `json:"LinkedTxn,omitempty"`
	TotalAmt                float64                 `json:"TotalAmt,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	DueDate                 *Date                   `json:"DueDate,omitempty"`
	MetaData                *MetaData               `json:"MetaData,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	TxnTaxDetail            *BillTxnTaxDetail       `json:"TxnTaxDetail,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	IncludeInAnnualTPAR     *bool                   `json:"IncludeInAnnualTPAR,omitempty"`
}

// CreateBillRequest represents the documented create bill payload.
type CreateBillRequest struct {
	// VendorRef is required.
	VendorRef Reference `json:"VendorRef"`
	// Line is required.
	Line                    []BillLine              `json:"Line"`
	TxnDate                 *Date                   `json:"TxnDate,omitempty"`
	DueDate                 *Date                   `json:"DueDate,omitempty"`
	APAccountRef            *Reference              `json:"APAccountRef,omitempty"`
	SalesTermRef            *Reference              `json:"SalesTermRef,omitempty"`
	DocNumber               string                  `json:"DocNumber,omitempty"`
	PrivateNote             string                  `json:"PrivateNote,omitempty"`
	CurrencyRef             *Reference              `json:"CurrencyRef,omitempty"`
	ExchangeRate            float64                 `json:"ExchangeRate,omitempty"`
	DepartmentRef           *Reference              `json:"DepartmentRef,omitempty"`
	GlobalTaxCalculation    GlobalTaxCalculation    `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail            *BillTxnTaxDetail       `json:"TxnTaxDetail,omitempty"`
	TransactionLocationType TransactionLocationType `json:"TransactionLocationType,omitempty"`
	IncludeInAnnualTPAR     *bool                   `json:"IncludeInAnnualTPAR,omitempty"`
}

// BillLine represents a bill line.
type BillLine struct {
	ID                            string                             `json:"Id,omitempty"`
	Amount                        float64                            `json:"Amount"`
	DetailType                    BillLineDetailType                 `json:"DetailType"`
	LinkedTxn                     []LinkedTxn                        `json:"LinkedTxn,omitempty"`
	Description                   string                             `json:"Description,omitempty"`
	LineNum                       float64                            `json:"LineNum,omitempty"`
	AccountBasedExpenseLineDetail *BillAccountBasedExpenseLineDetail `json:"AccountBasedExpenseLineDetail,omitempty"`
	ItemBasedExpenseLineDetail    *BillItemBasedExpenseLineDetail    `json:"ItemBasedExpenseLineDetail,omitempty"`
	TaxLineDetail                 *BillTaxLineDetail                 `json:"TaxLineDetail,omitempty"`
}

// BillAccountBasedExpenseLineDetail represents account-based bill line details.
type BillAccountBasedExpenseLineDetail struct {
	AccountRef      Reference       `json:"AccountRef"`
	TaxAmount       float64         `json:"TaxAmount,omitempty"`
	TaxInclusiveAmt float64         `json:"TaxInclusiveAmt,omitempty"`
	ClassRef        *Reference      `json:"ClassRef,omitempty"`
	TaxCodeRef      *Reference      `json:"TaxCodeRef,omitempty"`
	MarkupInfo      *BillMarkupInfo `json:"MarkupInfo,omitempty"`
	BillableStatus  BillableStatus  `json:"BillableStatus,omitempty"`
	CustomerRef     *Reference      `json:"CustomerRef,omitempty"`
}

// BillItemBasedExpenseLineDetail represents item-based bill line details.
type BillItemBasedExpenseLineDetail struct {
	TaxInclusiveAmt float64         `json:"TaxInclusiveAmt,omitempty"`
	ItemRef         *Reference      `json:"ItemRef,omitempty"`
	CustomerRef     *Reference      `json:"CustomerRef,omitempty"`
	PriceLevelRef   *Reference      `json:"PriceLevelRef,omitempty"`
	ClassRef        *Reference      `json:"ClassRef,omitempty"`
	TaxCodeRef      *Reference      `json:"TaxCodeRef,omitempty"`
	MarkupInfo      *BillMarkupInfo `json:"MarkupInfo,omitempty"`
	BillableStatus  BillableStatus  `json:"BillableStatus,omitempty"`
	Qty             float64         `json:"Qty,omitempty"`
	UnitPrice       float64         `json:"UnitPrice,omitempty"`
}

// BillMarkupInfo represents bill expense markup information.
type BillMarkupInfo struct {
	PriceLevelRef          *Reference `json:"PriceLevelRef,omitempty"`
	Percent                float64    `json:"Percent,omitempty"`
	MarkUpIncomeAccountRef *Reference `json:"MarkUpIncomeAccountRef,omitempty"`
}

// BillTxnTaxDetail represents bill-level tax details.
type BillTxnTaxDetail struct {
	TxnTaxCodeRef *Reference    `json:"TxnTaxCodeRef,omitempty"`
	TotalTax      float64       `json:"TotalTax,omitempty"`
	TaxLine       []BillTaxLine `json:"TaxLine,omitempty"`
}

// BillTaxLine represents a tax line in bill tax details.
type BillTaxLine struct {
	DetailType    BillLineDetailType `json:"DetailType"`
	TaxLineDetail BillTaxLineDetail  `json:"TaxLineDetail"`
	Amount        float64            `json:"Amount,omitempty"`
}

// BillTaxLineDetail represents tax details for a bill tax line.
type BillTaxLineDetail struct {
	TaxRateRef          Reference `json:"TaxRateRef"`
	NetAmountTaxable    float64   `json:"NetAmountTaxable,omitempty"`
	PercentBased        *bool     `json:"PercentBased,omitempty"`
	TaxInclusiveAmount  float64   `json:"TaxInclusiveAmount,omitempty"`
	OverrideDeltaAmount float64   `json:"OverrideDeltaAmount,omitempty"`
	TaxPercent          float64   `json:"TaxPercent,omitempty"`
}
