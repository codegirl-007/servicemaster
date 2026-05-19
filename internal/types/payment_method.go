// Package types contains transport types for external integrations.
package types

import "time"

// PaymentMethodType represents the documented payment method types.
type PaymentMethodType string

const (
	PaymentMethodTypeCreditCard    PaymentMethodType = "CREDIT_CARD"
	PaymentMethodTypeNonCreditCard PaymentMethodType = "NON_CREDIT_CARD"
)

// PaymentMethodResponse represents the QuickBooks payment method response envelope.
type PaymentMethodResponse struct {
	PaymentMethod PaymentMethod `json:"PaymentMethod"`
	Time          time.Time     `json:"time"`
}

// PaymentMethod represents a QuickBooks payment method object.
type PaymentMethod struct {
	ID        string            `json:"Id"`
	Name      string            `json:"Name,omitempty"`
	SyncToken string            `json:"SyncToken,omitempty"`
	Active    *bool             `json:"Active,omitempty"`
	Type      PaymentMethodType `json:"Type,omitempty"`
	MetaData  *MetaData         `json:"MetaData,omitempty"`
	Domain    string            `json:"domain,omitempty"`
	Sparse    *bool             `json:"sparse,omitempty"`
}

// CreatePaymentMethodRequest represents the documented create payment method payload.
type CreatePaymentMethodRequest struct {
	// Name is required.
	Name string `json:"Name"`
}
