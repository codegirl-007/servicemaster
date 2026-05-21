// Package types contains transport types for external integrations.
package types

import "time"

// InventoryAdjustmentLineDetailType represents documented inventory adjustment line detail types.
type InventoryAdjustmentLineDetailType string

const (
	InventoryAdjustmentLineDetailTypeItemAdjustment InventoryAdjustmentLineDetailType = "ItemAdjustmentLineDetail"
)

// InventoryAdjustmentResponse represents the QuickBooks inventory adjustment response envelope.
type InventoryAdjustmentResponse struct {
	InventoryAdjustment InventoryAdjustment `json:"InventoryAdjustment"`
	Time                time.Time           `json:"time"`
}

// InventoryAdjustment represents a QuickBooks inventory adjustment object.
type InventoryAdjustment struct {
	ID               string                    `json:"Id"`
	SyncToken        string                    `json:"SyncToken,omitempty"`
	TxnDate          *Date                     `json:"TxnDate,omitempty"`
	DocNumber        string                    `json:"DocNumber,omitempty"`
	PrivateNote      string                    `json:"PrivateNote,omitempty"`
	AdjustAccountRef Reference                 `json:"AdjustAccountRef"`
	Line             []InventoryAdjustmentLine `json:"Line"`
	MetaData         *MetaData                 `json:"MetaData,omitempty"`
	Domain           string                    `json:"domain,omitempty"`
	Sparse           *bool                     `json:"sparse,omitempty"`
}

// CreateInventoryAdjustmentRequest represents the documented create inventory adjustment payload.
type CreateInventoryAdjustmentRequest struct {
	AdjustAccountRef Reference                 `json:"AdjustAccountRef"`
	Line             []InventoryAdjustmentLine `json:"Line"`
	TxnDate          *Date                     `json:"TxnDate,omitempty"`
	DocNumber        string                    `json:"DocNumber,omitempty"`
	PrivateNote      string                    `json:"PrivateNote,omitempty"`
}

// UpdateInventoryAdjustmentRequest represents the documented update inventory adjustment payload.
type UpdateInventoryAdjustmentRequest struct {
	ID               string                    `json:"Id"`
	SyncToken        string                    `json:"SyncToken"`
	AdjustAccountRef Reference                 `json:"AdjustAccountRef"`
	Line             []InventoryAdjustmentLine `json:"Line"`
	TxnDate          *Date                     `json:"TxnDate,omitempty"`
	DocNumber        string                    `json:"DocNumber,omitempty"`
	PrivateNote      string                    `json:"PrivateNote,omitempty"`
	Domain           string                    `json:"domain,omitempty"`
}

// DeleteInventoryAdjustmentRequest represents the documented delete inventory adjustment payload.
type DeleteInventoryAdjustmentRequest struct {
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
}

// InventoryAdjustmentDeleteResponse represents the QuickBooks deleted inventory adjustment response envelope.
type InventoryAdjustmentDeleteResponse struct {
	InventoryAdjustment DeletedEntity `json:"InventoryAdjustment"`
	Time                time.Time     `json:"time"`
}

// InventoryAdjustmentLine represents an inventory adjustment line.
type InventoryAdjustmentLine struct {
	ID                        string                            `json:"Id,omitempty"`
	DetailType                InventoryAdjustmentLineDetailType `json:"DetailType,omitempty"`
	ItemAdjustmentLineDetail  *ItemAdjustmentLineDetail         `json:"ItemAdjustmentLineDetail,omitempty"`
}

// ItemAdjustmentLineDetail represents item adjustment line details.
type ItemAdjustmentLineDetail struct {
	QtyDiff  float64    `json:"QtyDiff,omitempty"`
	ItemRef  *Reference `json:"ItemRef,omitempty"`
}
