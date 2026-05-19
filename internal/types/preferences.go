// Package types contains transport types for external integrations.
package types

import "time"

// PreferencesWeekDay represents a documented week day value.
type PreferencesWeekDay string

const (
	PreferencesWeekDaySunday    PreferencesWeekDay = "Sunday"
	PreferencesWeekDayMonday    PreferencesWeekDay = "Monday"
	PreferencesWeekDayTuesday   PreferencesWeekDay = "Tuesday"
	PreferencesWeekDayWednesday PreferencesWeekDay = "Wednesday"
	PreferencesWeekDayThursday  PreferencesWeekDay = "Thursday"
	PreferencesWeekDayFriday    PreferencesWeekDay = "Friday"
	PreferencesWeekDaySaturday  PreferencesWeekDay = "Saturday"
)

// CustomFieldType represents documented custom field types.
type CustomFieldType string

const (
	CustomFieldTypeString  CustomFieldType = "StringType"
	CustomFieldTypeBoolean CustomFieldType = "BooleanType"
)

// RecognitionFrequencyType represents documented recognition frequencies.
type RecognitionFrequencyType string

const (
	RecognitionFrequencyTypeDaily   RecognitionFrequencyType = "Daily"
	RecognitionFrequencyTypeWeekly  RecognitionFrequencyType = "Weekly"
	RecognitionFrequencyTypeMonthly RecognitionFrequencyType = "Monthly"
)

// PreferencesResponse represents the QuickBooks preferences response envelope.
type PreferencesResponse struct {
	Preferences Preferences `json:"Preferences"`
	Time        time.Time   `json:"time"`
}

// Preferences represents a QuickBooks preferences object.
type Preferences struct {
	ID                      string                   `json:"Id"`
	SyncToken               string                   `json:"SyncToken"`
	SalesFormsPrefs         *SalesFormsPrefs         `json:"SalesFormsPrefs,omitempty"`
	VendorAndPurchasesPrefs *VendorAndPurchasesPrefs `json:"VendorAndPurchasesPrefs,omitempty"`
	AccountingInfoPrefs     *AccountingInfoPrefs     `json:"AccountingInfoPrefs,omitempty"`
	TaxPrefs                *TaxPrefs                `json:"TaxPrefs,omitempty"`
	TimeTrackingPrefs       *TimeTrackingPrefs       `json:"TimeTrackingPrefs,omitempty"`
	CurrencyPrefs           *CurrencyPrefs           `json:"CurrencyPrefs,omitempty"`
	ReportPrefs             *ReportPrefs             `json:"ReportPrefs,omitempty"`
	ProductAndServicesPrefs *ProductAndServicesPrefs `json:"ProductAndServicesPrefs,omitempty"`
	EmailMessagesPrefs      *EmailMessagesPrefs      `json:"EmailMessagesPrefs,omitempty"`
	OtherPrefs              []NameValue              `json:"OtherPrefs,omitempty"`
	MetaData                *MetaData                `json:"MetaData,omitempty"`
	Domain                  string                   `json:"domain,omitempty"`
	Sparse                  *bool                    `json:"sparse,omitempty"`
}

// UpdatePreferencesRequest represents the documented update preferences payload.
type UpdatePreferencesRequest struct {
	ID                      string                   `json:"Id"`
	SyncToken               string                   `json:"SyncToken"`
	SalesFormsPrefs         *SalesFormsPrefs         `json:"SalesFormsPrefs,omitempty"`
	VendorAndPurchasesPrefs *VendorAndPurchasesPrefs `json:"VendorAndPurchasesPrefs,omitempty"`
	AccountingInfoPrefs     *AccountingInfoPrefs     `json:"AccountingInfoPrefs,omitempty"`
	TaxPrefs                *TaxPrefs                `json:"TaxPrefs,omitempty"`
	TimeTrackingPrefs       *TimeTrackingPrefs       `json:"TimeTrackingPrefs,omitempty"`
	CurrencyPrefs           *CurrencyPrefs           `json:"CurrencyPrefs,omitempty"`
	ReportPrefs             *ReportPrefs             `json:"ReportPrefs,omitempty"`
	ProductAndServicesPrefs *ProductAndServicesPrefs `json:"ProductAndServicesPrefs,omitempty"`
	EmailMessagesPrefs      *EmailMessagesPrefs      `json:"EmailMessagesPrefs,omitempty"`
	OtherPrefs              []NameValue              `json:"OtherPrefs,omitempty"`
	Domain                  string                   `json:"domain,omitempty"`
	Sparse                  *bool                    `json:"sparse,omitempty"`
}

// SalesFormsPrefs represents sales form preferences.
type SalesFormsPrefs struct {
	CustomField                  []CustomFieldDefinition `json:"CustomField,omitempty"`
	AllowDeposit                 *bool                   `json:"AllowDeposit,omitempty"`
	AllowDiscount                *bool                   `json:"AllowDiscount,omitempty"`
	AllowEstimates               *bool                   `json:"AllowEstimates,omitempty"`
	AllowServiceDate             *bool                   `json:"AllowServiceDate,omitempty"`
	DefaultCustomerMessage       *PreferenceMessage      `json:"DefaultCustomerMessage,omitempty"`
	DefaultItem                  *Reference              `json:"DefaultItem,omitempty"`
	DefaultTerms                 *Reference              `json:"DefaultTerms,omitempty"`
	ETransactionEnabledStatus    string                  `json:"ETransactionEnabledStatus,omitempty"`
	EmailCopyToCompany           *bool                   `json:"EmailCopyToCompany,omitempty"`
	CustomTxnNumbers             *bool                   `json:"CustomTxnNumbers,omitempty"`
	AllowShipping                *bool                   `json:"AllowShipping,omitempty"`
	DefaultDiscountAccount       *Reference              `json:"DefaultDiscountAccount,omitempty"`
	AllowPriceRules              *bool                   `json:"AllowPriceRules,omitempty"`
}

// VendorAndPurchasesPrefs represents vendor and purchase preferences.
type VendorAndPurchasesPrefs struct {
	TrackingByCustomer          *bool      `json:"TrackingByCustomer,omitempty"`
	BillableExpenseTracking     *bool      `json:"BillableExpenseTracking,omitempty"`
	TaxIncluded                 *bool      `json:"TaxIncluded,omitempty"`
	DefaultTerms                *Reference `json:"DefaultTerms,omitempty"`
	DefaultMarkup               float64    `json:"DefaultMarkup,omitempty"`
	DefaultExpenseAccount       *Reference `json:"DefaultExpenseAccount,omitempty"`
	ETransactionEnabledStatus   string     `json:"ETransactionEnabledStatus,omitempty"`
}

// AccountingInfoPrefs represents accounting information preferences.
type AccountingInfoPrefs struct {
	ClassTrackingPerTxn     *bool                  `json:"ClassTrackingPerTxn,omitempty"`
	ClassTrackingPerTxnLine *bool                  `json:"ClassTrackingPerTxnLine,omitempty"`
	TrackDepartments        *bool                  `json:"TrackDepartments,omitempty"`
	DepartmentTerminology   string                 `json:"DepartmentTerminology,omitempty"`
	CustomerTerminology     string                 `json:"CustomerTerminology,omitempty"`
	BookCloseDate           *Date                  `json:"BookCloseDate,omitempty"`
	FirstMonthOfFiscalYear  FiscalYearStartMonth   `json:"FirstMonthOfFiscalYear,omitempty"`
}

// TaxPrefs represents tax preferences.
type TaxPrefs struct {
	UsingSalesTax       *bool      `json:"UsingSalesTax,omitempty"`
	PartnerTaxEnabled   *bool      `json:"PartnerTaxEnabled,omitempty"`
	TaxGroupCodeRef     *Reference `json:"TaxGroupCodeRef,omitempty"`
	DefaultTaxCodeRef   *Reference `json:"DefaultTaxCodeRef,omitempty"`
}

// TimeTrackingPrefs represents time tracking preferences.
type TimeTrackingPrefs struct {
	UsingTimeTracking        *bool                  `json:"UsingTimeTracking,omitempty"`
	BillCustomers            *bool                  `json:"BillCustomers,omitempty"`
	ShowBillRateToAll        *bool                  `json:"ShowBillRateToAll,omitempty"`
	MarkTimeEntriesBillable  *bool                  `json:"MarkTimeEntriesBillable,omitempty"`
	FirstDayOfWeek           PreferencesWeekDay     `json:"FirstDayOfWeek,omitempty"`
	WorkWeekStartDate        *Date                  `json:"WorkWeekStartDate,omitempty"`
	FrequencyType            RecognitionFrequencyType `json:"FrequencyType,omitempty"`
}

// CurrencyPrefs represents currency preferences.
type CurrencyPrefs struct {
	HomeCurrency         *Reference `json:"HomeCurrency,omitempty"`
	MultiCurrencyEnabled *bool      `json:"MultiCurrencyEnabled,omitempty"`
}

// ReportPrefs represents report preferences.
type ReportPrefs struct {
	ReportBasis ReportBasis `json:"ReportBasis,omitempty"`
}

// ProductAndServicesPrefs represents product and service preferences.
type ProductAndServicesPrefs struct {
	ForSales             *bool `json:"ForSales,omitempty"`
	ForPurchase          *bool `json:"ForPurchase,omitempty"`
	QuantityOnHand       *bool `json:"QuantityOnHand,omitempty"`
	QuantityWithPriceAndRate *bool `json:"QuantityWithPriceAndRate,omitempty"`
	StockKeepingUnit     *bool `json:"StockKeepingUnit,omitempty"`
	CategoriesEnabled    *bool `json:"CategoriesEnabled,omitempty"`
}

// EmailMessagesPrefs represents default email message preferences.
type EmailMessagesPrefs struct {
	EstimateMessage     *PreferenceMessage `json:"EstimateMessage,omitempty"`
	InvoiceMessage      *PreferenceMessage `json:"InvoiceMessage,omitempty"`
	SalesReceiptMessage *PreferenceMessage `json:"SalesReceiptMessage,omitempty"`
	StatementMessage    *PreferenceMessage `json:"StatementMessage,omitempty"`
}

// PreferenceMessage represents a message subject/body pair.
type PreferenceMessage struct {
	Subject string `json:"Subject,omitempty"`
	Message string `json:"Message,omitempty"`
}

// CustomFieldDefinition represents a custom field definition in preferences.
type CustomFieldDefinition struct {
	Name     string          `json:"Name,omitempty"`
	Type     CustomFieldType `json:"Type,omitempty"`
	BooleanValue *bool       `json:"BooleanValue,omitempty"`
	StringValue  string      `json:"StringValue,omitempty"`
}
