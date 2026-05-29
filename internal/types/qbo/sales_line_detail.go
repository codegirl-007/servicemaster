package qbo

// SalesItemLineDetail represents sales-item line details on QuickBooks sales forms.
type SalesItemLineDetail struct {
	ItemRef         *Reference `json:"ItemRef,omitempty"`
	ClassRef        *Reference `json:"ClassRef,omitempty"`
	TaxCodeRef      *Reference `json:"TaxCodeRef,omitempty"`
	ServiceDate     *Date      `json:"ServiceDate,omitempty"`
	Qty             float64    `json:"Qty,omitempty"`
	UnitPrice       float64    `json:"UnitPrice,omitempty"`
	TaxInclusiveAmt float64    `json:"TaxInclusiveAmt,omitempty"`
	DiscountAmt     float64    `json:"DiscountAmt,omitempty"`
	DiscountRate    float64    `json:"DiscountRate,omitempty"`
	ItemAccountRef  *Reference `json:"ItemAccountRef,omitempty"`
}

// GroupLineDetail represents grouped line details on QuickBooks sales forms.
type GroupLineDetail struct {
	GroupItemRef *Reference    `json:"GroupItemRef,omitempty"`
	Quantity     float64       `json:"Quantity,omitempty"`
	Line         []InvoiceLine `json:"Line,omitempty"`
}

// DescriptionLineDetail represents descriptive line details on QuickBooks sales forms.
type DescriptionLineDetail struct {
	TaxCodeRef  *Reference `json:"TaxCodeRef,omitempty"`
	ServiceDate *Date      `json:"ServiceDate,omitempty"`
}

// SubTotalLineDetail represents subtotal line details on QuickBooks sales forms.
type SubTotalLineDetail struct{}
