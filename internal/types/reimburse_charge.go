// Package types contains transport types for external integrations.
package types

import "time"

// ReimburseChargeLineDetailType represents the documented reimburse charge line detail types.
type ReimburseChargeLineDetailType string

const (
	ReimburseChargeLineDetailTypeReimburse ReimburseChargeLineDetailType = "ReimburseLineDetail"
)

// ReimburseChargeResponse represents the QuickBooks reimburse charge response envelope.
type ReimburseChargeResponse struct {
	ReimburseCharge ReimburseCharge `json:"ReimburseCharge"`
	Time            time.Time       `json:"time"`
}

// ReimburseCharge represents a QuickBooks reimburse charge object.
type ReimburseCharge struct {
	ID              string                `json:"Id"`
	SyncToken       string                `json:"SyncToken,omitempty"`
	CustomerRef     *Reference            `json:"CustomerRef,omitempty"`
	Amount          float64               `json:"Amount,omitempty"`
	Line            []ReimburseChargeLine `json:"Line,omitempty"`
	CurrencyRef     *Reference            `json:"CurrencyRef,omitempty"`
	HasBeenInvoiced *bool                 `json:"HasBeenInvoiced,omitempty"`
	HomeTotalAmt    float64               `json:"HomeTotalAmt,omitempty"`
	PrivateNote     string                `json:"PrivateNote,omitempty"`
	LinkedTxn       []LinkedTxn           `json:"LinkedTxn,omitempty"`
	ExchangeRate    float64               `json:"ExchangeRate,omitempty"`
	TxnDate         *Date                 `json:"TxnDate,omitempty"`
	MetaData        *MetaData             `json:"MetaData,omitempty"`
	Domain          string                `json:"domain,omitempty"`
	Sparse          *bool                 `json:"sparse,omitempty"`
}

// ReimburseChargeLine represents a reimburse charge line.
type ReimburseChargeLine struct {
	ID                  string                        `json:"Id,omitempty"`
	LineNum             float64                       `json:"LineNum,omitempty"`
	LineID              string                        `json:"LineId,omitempty"`
	Amount              float64                       `json:"Amount,omitempty"`
	Description         string                        `json:"Description,omitempty"`
	DetailType          ReimburseChargeLineDetailType `json:"DetailType,omitempty"`
	ReimburseLineDetail *ReimburseLineDetail          `json:"ReimburseLineDetail,omitempty"`
	LinkedTxn           []LinkedTxn                   `json:"LinkedTxn,omitempty"`
}

// ReimburseLineDetail represents reimburse charge line details.
type ReimburseLineDetail struct {
	ItemRef            *Reference      `json:"ItemRef,omitempty"`
	Qty                float64         `json:"Qty,omitempty"`
	UnitPrice          float64         `json:"UnitPrice,omitempty"`
	TaxCodeRef         *Reference      `json:"TaxCodeRef,omitempty"`
	ItemAccountRef     *Reference      `json:"ItemAccountRef,omitempty"`
	ClassRef           *Reference      `json:"ClassRef,omitempty"`
	MarkupInfo         *BillMarkupInfo `json:"MarkupInfo,omitempty"`
	DiscountAccountRef *Reference      `json:"DiscountAccountRef,omitempty"`
	PercentBased       *bool           `json:"PercentBased,omitempty"`
	DiscountPercent    float64         `json:"DiscountPercent,omitempty"`
}
