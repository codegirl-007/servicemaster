// Package types contains transport types for external integrations.
package types

import "time"

// TaxPaymentResponse represents the QuickBooks tax payment response envelope.
type TaxPaymentResponse struct {
	TaxPayment TaxPayment `json:"TaxPayment"`
	Time       time.Time  `json:"time"`
}

// TaxPayment represents a QuickBooks tax payment object.
type TaxPayment struct {
	ID                string     `json:"Id"`
	SyncToken         string     `json:"SyncToken,omitempty"`
	Refund            string     `json:"Refund,omitempty"`
	PaymentDate       *Date      `json:"PaymentDate,omitempty"`
	PaymentAccountRef *Reference `json:"PaymentAccountRef,omitempty"`
	Description       string     `json:"Description,omitempty"`
	PaymentAmount     string     `json:"PaymentAmount,omitempty"`
	MetaData          *MetaData  `json:"MetaData,omitempty"`
	Domain            string     `json:"domain,omitempty"`
	Sparse            string     `json:"sparse,omitempty"`
}
