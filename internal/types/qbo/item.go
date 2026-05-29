// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// ItemType represents the documented QuickBooks item types.
type ItemType string

const (
	ItemTypeInventory    ItemType = "Inventory"
	ItemTypeService      ItemType = "Service"
	ItemTypeNonInventory ItemType = "NonInventory"
	ItemTypeCategory     ItemType = "Category"
	ItemTypeGroup        ItemType = "Group"
)

// ItemCategoryType represents the documented France item category types.
type ItemCategoryType string

const (
	ItemCategoryTypeProduct ItemCategoryType = "Product"
	ItemCategoryTypeService ItemCategoryType = "Service"
)

// ItemSource represents the documented QuickBooks Commerce source value.
type ItemSource string

const (
	ItemSourceQBCommerce ItemSource = "QBCommerce"
)

// ItemResponse represents the QuickBooks item response envelope.
type ItemResponse struct {
	Item Item      `json:"Item"`
	Time time.Time `json:"time"`
}

// Item represents a QuickBooks item object.
type Item struct {
	ID                  string           `json:"Id"`
	SyncToken           string           `json:"SyncToken"`
	Name                string           `json:"Name,omitempty"`
	Active              *bool            `json:"Active,omitempty"`
	Type                ItemType         `json:"Type,omitempty"`
	Description         string           `json:"Description,omitempty"`
	UnitPrice           float64          `json:"UnitPrice,omitempty"`
	PurchaseDesc        string           `json:"PurchaseDesc,omitempty"`
	PurchaseCost        float64          `json:"PurchaseCost,omitempty"`
	SKU                 string           `json:"Sku,omitempty"`
	SubItem             *bool            `json:"SubItem,omitempty"`
	ParentRef           *Reference       `json:"ParentRef,omitempty"`
	FullyQualifiedName  string           `json:"FullyQualifiedName,omitempty"`
	Level               int              `json:"Level,omitempty"`
	IncomeAccountRef    *Reference       `json:"IncomeAccountRef,omitempty"`
	ExpenseAccountRef   *Reference       `json:"ExpenseAccountRef,omitempty"`
	AssetAccountRef     *Reference       `json:"AssetAccountRef,omitempty"`
	TrackQtyOnHand      *bool            `json:"TrackQtyOnHand,omitempty"`
	QtyOnHand           float64          `json:"QtyOnHand,omitempty"`
	InvStartDate        *Date            `json:"InvStartDate,omitempty"`
	ReorderPoint        float64          `json:"ReorderPoint,omitempty"`
	SalesTaxCodeRef     *Reference       `json:"SalesTaxCodeRef,omitempty"`
	PurchaseTaxCodeRef  *Reference       `json:"PurchaseTaxCodeRef,omitempty"`
	Taxable             *bool            `json:"Taxable,omitempty"`
	SalesTaxIncluded    *bool            `json:"SalesTaxIncluded,omitempty"`
	PurchaseTaxIncluded *bool            `json:"PurchaseTaxIncluded,omitempty"`
	ItemCategoryType    ItemCategoryType `json:"ItemCategoryType,omitempty"`
	Source              ItemSource       `json:"Source,omitempty"`
	MetaData            *MetaData        `json:"MetaData,omitempty"`
	Domain              string           `json:"domain,omitempty"`
	Sparse              *bool            `json:"sparse,omitempty"`
}

// CreateItemRequest represents the documented create item payload.
type CreateItemRequest struct {
	// Name is required.
	Name                string           `json:"Name"`
	Active              *bool            `json:"Active,omitempty"`
	Type                ItemType         `json:"Type,omitempty"`
	Description         string           `json:"Description,omitempty"`
	UnitPrice           float64          `json:"UnitPrice,omitempty"`
	PurchaseDesc        string           `json:"PurchaseDesc,omitempty"`
	PurchaseCost        float64          `json:"PurchaseCost,omitempty"`
	SKU                 string           `json:"Sku,omitempty"`
	SubItem             *bool            `json:"SubItem,omitempty"`
	ParentRef           *Reference       `json:"ParentRef,omitempty"`
	IncomeAccountRef    *Reference       `json:"IncomeAccountRef,omitempty"`
	ExpenseAccountRef   *Reference       `json:"ExpenseAccountRef,omitempty"`
	AssetAccountRef     *Reference       `json:"AssetAccountRef,omitempty"`
	TrackQtyOnHand      *bool            `json:"TrackQtyOnHand,omitempty"`
	QtyOnHand           float64          `json:"QtyOnHand,omitempty"`
	InvStartDate        *Date            `json:"InvStartDate,omitempty"`
	ReorderPoint        float64          `json:"ReorderPoint,omitempty"`
	SalesTaxCodeRef     *Reference       `json:"SalesTaxCodeRef,omitempty"`
	PurchaseTaxCodeRef  *Reference       `json:"PurchaseTaxCodeRef,omitempty"`
	Taxable             *bool            `json:"Taxable,omitempty"`
	SalesTaxIncluded    *bool            `json:"SalesTaxIncluded,omitempty"`
	PurchaseTaxIncluded *bool            `json:"PurchaseTaxIncluded,omitempty"`
	ItemCategoryType    ItemCategoryType `json:"ItemCategoryType,omitempty"`
	Source              ItemSource       `json:"Source,omitempty"`
}

// SparseUpdateItemRequest represents the documented sparse update payload.
type SparseUpdateItemRequest struct {
	// ID is required for update.
	ID string `json:"Id"`
	// SyncToken is required for update.
	SyncToken string `json:"SyncToken"`
	// Sparse must be true for sparse update.
	Sparse              bool             `json:"sparse"`
	Name                string           `json:"Name,omitempty"`
	Active              *bool            `json:"Active,omitempty"`
	Type                ItemType         `json:"Type,omitempty"`
	Description         string           `json:"Description,omitempty"`
	UnitPrice           float64          `json:"UnitPrice,omitempty"`
	PurchaseDesc        string           `json:"PurchaseDesc,omitempty"`
	PurchaseCost        float64          `json:"PurchaseCost,omitempty"`
	SKU                 string           `json:"Sku,omitempty"`
	SubItem             *bool            `json:"SubItem,omitempty"`
	ParentRef           *Reference       `json:"ParentRef,omitempty"`
	IncomeAccountRef    *Reference       `json:"IncomeAccountRef,omitempty"`
	ExpenseAccountRef   *Reference       `json:"ExpenseAccountRef,omitempty"`
	AssetAccountRef     *Reference       `json:"AssetAccountRef,omitempty"`
	TrackQtyOnHand      *bool            `json:"TrackQtyOnHand,omitempty"`
	QtyOnHand           float64          `json:"QtyOnHand,omitempty"`
	InvStartDate        *Date            `json:"InvStartDate,omitempty"`
	ReorderPoint        float64          `json:"ReorderPoint,omitempty"`
	SalesTaxCodeRef     *Reference       `json:"SalesTaxCodeRef,omitempty"`
	PurchaseTaxCodeRef  *Reference       `json:"PurchaseTaxCodeRef,omitempty"`
	Taxable             *bool            `json:"Taxable,omitempty"`
	SalesTaxIncluded    *bool            `json:"SalesTaxIncluded,omitempty"`
	PurchaseTaxIncluded *bool            `json:"PurchaseTaxIncluded,omitempty"`
	ItemCategoryType    ItemCategoryType `json:"ItemCategoryType,omitempty"`
	Source              ItemSource       `json:"Source,omitempty"`
	Domain              string           `json:"domain,omitempty"`
}
