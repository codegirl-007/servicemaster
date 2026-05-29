// Package qbo contains transport types for the QuickBooks Online API.
package qbo

import "time"

// EstimateTxnStatus represents the documented estimate status values.
type EstimateTxnStatus string

const (
	EstimateTxnStatusAccepted  EstimateTxnStatus = "Accepted"
	EstimateTxnStatusClosed    EstimateTxnStatus = "Closed"
	EstimateTxnStatusPending   EstimateTxnStatus = "Pending"
	EstimateTxnStatusRejected  EstimateTxnStatus = "Rejected"
	EstimateTxnStatusConverted EstimateTxnStatus = "Converted"
)

// EstimateLineDetailType represents the documented estimate line detail types.
type EstimateLineDetailType string

const (
	EstimateLineDetailTypeSalesItem       EstimateLineDetailType = "SalesItemLineDetail"
	EstimateLineDetailTypeGroup           EstimateLineDetailType = "GroupLineDetail"
	EstimateLineDetailTypeDescriptionOnly EstimateLineDetailType = "DescriptionOnly"
	EstimateLineDetailTypeSubTotal        EstimateLineDetailType = "SubTotalLineDetail"
	EstimateLineDetailTypeTaxLine         EstimateLineDetailType = "TaxLineDetail"
	EstimateLineDetailTypeDiscountLine    EstimateLineDetailType = "DiscountLineDetail"
)

// EstimateResponse represents the QuickBooks estimate response envelope.
type EstimateResponse struct {
	Estimate Estimate  `json:"Estimate"`
	Time     time.Time `json:"time"`
}

// Estimate represents a QuickBooks estimate object.
type Estimate struct {
	ID                    string               `json:"Id"`
	SyncToken             string               `json:"SyncToken"`
	CustomerRef           *Reference           `json:"CustomerRef,omitempty"`
	Line                  []EstimateLine       `json:"Line,omitempty"`
	TxnDate               *Date                `json:"TxnDate,omitempty"`
	ExpirationDate        *Date                `json:"ExpirationDate,omitempty"`
	AcceptedDate          *Date                `json:"AcceptedDate,omitempty"`
	TxnStatus             EstimateTxnStatus    `json:"TxnStatus,omitempty"`
	DocNumber             string               `json:"DocNumber,omitempty"`
	CustomerMemo          *Memo                `json:"CustomerMemo,omitempty"`
	PrivateNote           string               `json:"PrivateNote,omitempty"`
	BillEmail             *EmailAddress        `json:"BillEmail,omitempty"`
	EmailStatus           EmailStatus          `json:"EmailStatus,omitempty"`
	PrintStatus           PrintStatus          `json:"PrintStatus,omitempty"`
	CurrencyRef           *Reference           `json:"CurrencyRef,omitempty"`
	ExchangeRate          float64              `json:"ExchangeRate,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail          *TxnTaxDetail        `json:"TxnTaxDetail,omitempty"`
	ApplyTaxAfterDiscount *bool                `json:"ApplyTaxAfterDiscount,omitempty"`
	TotalAmt              float64              `json:"TotalAmt,omitempty"`
	HomeTotalAmt          float64              `json:"HomeTotalAmt,omitempty"`
	SalesTermRef          *Reference           `json:"SalesTermRef,omitempty"`
	ClassRef              *Reference           `json:"ClassRef,omitempty"`
	DepartmentRef         *Reference           `json:"DepartmentRef,omitempty"`
	ProjectRef            *Reference           `json:"ProjectRef,omitempty"`
	BillAddr              *PhysicalAddress     `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress     `json:"ShipAddr,omitempty"`
	ShipFromAddr          *PhysicalAddress     `json:"ShipFromAddr,omitempty"`
	ShipDate              *Date                `json:"ShipDate,omitempty"`
	ShipMethodRef         *Reference           `json:"ShipMethodRef,omitempty"`
	LinkedTxn             []LinkedTxn          `json:"LinkedTxn,omitempty"`
	MetaData              *MetaData            `json:"MetaData,omitempty"`
	Domain                string               `json:"domain,omitempty"`
	Sparse                *bool                `json:"sparse,omitempty"`
}

// CreateEstimateRequest represents the documented create estimate payload.
type CreateEstimateRequest struct {
	// CustomerRef is required.
	CustomerRef Reference `json:"CustomerRef"`
	// Line is required.
	Line                  []EstimateLine       `json:"Line"`
	TxnDate               *Date                `json:"TxnDate,omitempty"`
	ExpirationDate        *Date                `json:"ExpirationDate,omitempty"`
	AcceptedDate          *Date                `json:"AcceptedDate,omitempty"`
	TxnStatus             EstimateTxnStatus    `json:"TxnStatus,omitempty"`
	DocNumber             string               `json:"DocNumber,omitempty"`
	CustomerMemo          *Memo                `json:"CustomerMemo,omitempty"`
	PrivateNote           string               `json:"PrivateNote,omitempty"`
	BillEmail             *EmailAddress        `json:"BillEmail,omitempty"`
	EmailStatus           EmailStatus          `json:"EmailStatus,omitempty"`
	PrintStatus           PrintStatus          `json:"PrintStatus,omitempty"`
	CurrencyRef           *Reference           `json:"CurrencyRef,omitempty"`
	ExchangeRate          float64              `json:"ExchangeRate,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail          *TxnTaxDetail        `json:"TxnTaxDetail,omitempty"`
	ApplyTaxAfterDiscount *bool                `json:"ApplyTaxAfterDiscount,omitempty"`
	SalesTermRef          *Reference           `json:"SalesTermRef,omitempty"`
	ClassRef              *Reference           `json:"ClassRef,omitempty"`
	DepartmentRef         *Reference           `json:"DepartmentRef,omitempty"`
	ProjectRef            *Reference           `json:"ProjectRef,omitempty"`
	BillAddr              *PhysicalAddress     `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress     `json:"ShipAddr,omitempty"`
	ShipFromAddr          *PhysicalAddress     `json:"ShipFromAddr,omitempty"`
	ShipDate              *Date                `json:"ShipDate,omitempty"`
	ShipMethodRef         *Reference           `json:"ShipMethodRef,omitempty"`
	LinkedTxn             []LinkedTxn          `json:"LinkedTxn,omitempty"`
}

// SparseUpdateEstimateRequest represents the documented sparse update payload.
type SparseUpdateEstimateRequest struct {
	ID                    string               `json:"Id"`
	SyncToken             string               `json:"SyncToken"`
	Sparse                bool                 `json:"sparse"`
	CustomerRef           *Reference           `json:"CustomerRef,omitempty"`
	Line                  []EstimateLine       `json:"Line,omitempty"`
	TxnDate               *Date                `json:"TxnDate,omitempty"`
	ExpirationDate        *Date                `json:"ExpirationDate,omitempty"`
	AcceptedDate          *Date                `json:"AcceptedDate,omitempty"`
	TxnStatus             EstimateTxnStatus    `json:"TxnStatus,omitempty"`
	DocNumber             string               `json:"DocNumber,omitempty"`
	CustomerMemo          *Memo                `json:"CustomerMemo,omitempty"`
	PrivateNote           string               `json:"PrivateNote,omitempty"`
	BillEmail             *EmailAddress        `json:"BillEmail,omitempty"`
	EmailStatus           EmailStatus          `json:"EmailStatus,omitempty"`
	PrintStatus           PrintStatus          `json:"PrintStatus,omitempty"`
	CurrencyRef           *Reference           `json:"CurrencyRef,omitempty"`
	ExchangeRate          float64              `json:"ExchangeRate,omitempty"`
	GlobalTaxCalculation  GlobalTaxCalculation `json:"GlobalTaxCalculation,omitempty"`
	TxnTaxDetail          *TxnTaxDetail        `json:"TxnTaxDetail,omitempty"`
	ApplyTaxAfterDiscount *bool                `json:"ApplyTaxAfterDiscount,omitempty"`
	SalesTermRef          *Reference           `json:"SalesTermRef,omitempty"`
	ClassRef              *Reference           `json:"ClassRef,omitempty"`
	DepartmentRef         *Reference           `json:"DepartmentRef,omitempty"`
	ProjectRef            *Reference           `json:"ProjectRef,omitempty"`
	BillAddr              *PhysicalAddress     `json:"BillAddr,omitempty"`
	ShipAddr              *PhysicalAddress     `json:"ShipAddr,omitempty"`
	ShipFromAddr          *PhysicalAddress     `json:"ShipFromAddr,omitempty"`
	ShipDate              *Date                `json:"ShipDate,omitempty"`
	ShipMethodRef         *Reference           `json:"ShipMethodRef,omitempty"`
	LinkedTxn             []LinkedTxn          `json:"LinkedTxn,omitempty"`
	Domain                string               `json:"domain,omitempty"`
}

// EstimateLine represents a QuickBooks estimate line.
type EstimateLine struct {
	ID                    string                   `json:"Id,omitempty"`
	LineNum               float64                  `json:"LineNum,omitempty"`
	Description           string                   `json:"Description,omitempty"`
	Amount                float64                  `json:"Amount,omitempty"`
	DetailType            EstimateLineDetailType   `json:"DetailType,omitempty"`
	SalesItemLineDetail   *SalesItemLineDetail     `json:"SalesItemLineDetail,omitempty"`
	GroupLineDetail       *EstimateGroupLineDetail `json:"GroupLineDetail,omitempty"`
	DescriptionLineDetail *DescriptionLineDetail   `json:"DescriptionLineDetail,omitempty"`
	SubTotalLineDetail    *SubTotalLineDetail      `json:"SubTotalLineDetail,omitempty"`
	TaxLineDetail         *TaxLineDetail           `json:"TaxLineDetail,omitempty"`
}

// EstimateGroupLineDetail represents grouped estimate line details.
type EstimateGroupLineDetail struct {
	GroupItemRef *Reference     `json:"GroupItemRef,omitempty"`
	Quantity     float64        `json:"Quantity,omitempty"`
	Line         []EstimateLine `json:"Line,omitempty"`
}
