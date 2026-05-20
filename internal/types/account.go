// Package types contains transport types for external integrations.
package types

import "time"

// AccountType represents a QuickBooks top-level account type.
type AccountType string

const (
	AccountTypeBank               AccountType = "Bank"
	AccountTypeOtherCurrentAsset  AccountType = "Other Current Asset"
	AccountTypeFixedAsset         AccountType = "Fixed Asset"
	AccountTypeOtherAsset         AccountType = "Other Asset"
	AccountTypeAccountsReceivable AccountType = "Accounts Receivable"
)

// AccountSubType represents a QuickBooks account subtype.
type AccountSubType string

const (
	AccountSubTypeCashOnHand                           AccountSubType = "CashOnHand"
	AccountSubTypeChecking                             AccountSubType = "Checking"
	AccountSubTypeMoneyMarket                          AccountSubType = "MoneyMarket"
	AccountSubTypeRentsHeldInTrust                     AccountSubType = "RentsHeldInTrust"
	AccountSubTypeSavings                              AccountSubType = "Savings"
	AccountSubTypeTrustAccounts                        AccountSubType = "TrustAccounts"
	AccountSubTypeCashAndCashEquivalents               AccountSubType = "CashAndCashEquivalents"
	AccountSubTypeOtherEarMarkedBankAccounts           AccountSubType = "OtherEarMarkedBankAccounts"
	AccountSubTypeAllowanceForBadDebts                 AccountSubType = "AllowanceForBadDebts"
	AccountSubTypeDevelopmentCosts                     AccountSubType = "DevelopmentCosts"
	AccountSubTypeEmployeeCashAdvances                 AccountSubType = "EmployeeCashAdvances"
	AccountSubTypeOtherCurrentAssets                   AccountSubType = "OtherCurrentAssets"
	AccountSubTypeInventory                            AccountSubType = "Inventory"
	AccountSubTypeInvestmentMortgageRealEstateLoans    AccountSubType = "Investment_MortgageRealEstateLoans"
	AccountSubTypeInvestmentOther                      AccountSubType = "Investment_Other"
	AccountSubTypeInvestmentTaxExemptSecurities        AccountSubType = "Investment_TaxExemptSecurities"
	AccountSubTypeInvestmentUSGovernmentObligations    AccountSubType = "Investment_USGovernmentObligations"
	AccountSubTypeLoansToOfficers                      AccountSubType = "LoansToOfficers"
	AccountSubTypeLoansToOthers                        AccountSubType = "LoansToOthers"
	AccountSubTypeLoansToStockholders                  AccountSubType = "LoansToStockholders"
	AccountSubTypePrepaidExpenses                      AccountSubType = "PrepaidExpenses"
	AccountSubTypeRetainage                            AccountSubType = "Retainage"
	AccountSubTypeUndepositedFunds                     AccountSubType = "UndepositedFunds"
	AccountSubTypeAccumulatedDepletion                 AccountSubType = "AccumulatedDepletion"
	AccountSubTypeAccumulatedDepreciation              AccountSubType = "AccumulatedDepreciation"
	AccountSubTypeDepletableAssets                     AccountSubType = "DepletableAssets"
	AccountSubTypeFixedAssetComputers                  AccountSubType = "FixedAssetComputers"
	AccountSubTypeFixedAssetCopiers                    AccountSubType = "FixedAssetCopiers"
	AccountSubTypeFixedAssetFurniture                  AccountSubType = "FixedAssetFurniture"
	AccountSubTypeFixedAssetPhone                      AccountSubType = "FixedAssetPhone"
	AccountSubTypeFixedAssetPhotoVideo                 AccountSubType = "FixedAssetPhotoVideo"
	AccountSubTypeFixedAssetSoftware                   AccountSubType = "FixedAssetSoftware"
	AccountSubTypeFixedAssetOtherToolsEquipment        AccountSubType = "FixedAssetOtherToolsEquipment"
	AccountSubTypeFurnitureAndFixtures                 AccountSubType = "FurnitureAndFixtures"
	AccountSubTypeLand                                 AccountSubType = "Land"
	AccountSubTypeLeaseholdImprovements                AccountSubType = "LeaseholdImprovements"
	AccountSubTypeOtherFixedAssets                     AccountSubType = "OtherFixedAssets"
	AccountSubTypeAccumulatedAmortization              AccountSubType = "AccumulatedAmortization"
	AccountSubTypeBuildings                            AccountSubType = "Buildings"
	AccountSubTypeIntangibleAssets                     AccountSubType = "IntangibleAssets"
	AccountSubTypeMachineryAndEquipment                AccountSubType = "MachineryAndEquipment"
	AccountSubTypeVehicles                             AccountSubType = "Vehicles"
	AccountSubTypeLeaseBuyout                          AccountSubType = "LeaseBuyout"
	AccountSubTypeOtherLongTermAssets                  AccountSubType = "OtherLongTermAssets"
	AccountSubTypeSecurityDeposits                     AccountSubType = "SecurityDeposits"
	AccountSubTypeAccumulatedAmortizationOfOtherAssets AccountSubType = "AccumulatedAmortizationOfOtherAssets"
	AccountSubTypeGoodwill                             AccountSubType = "Goodwill"
	AccountSubTypeLicenses                             AccountSubType = "Licenses"
	AccountSubTypeOrganizationalCosts                  AccountSubType = "OrganizationalCosts"
	AccountSubTypeAccountsReceivable                   AccountSubType = "AccountsReceivable"
)

// AccountResponse represents the QuickBooks account response envelope.
type AccountResponse struct {
	Account Account   `json:"Account"`
	Time    time.Time `json:"time"`
}

// Account represents a QuickBooks account object.
type Account struct {
	FullyQualifiedName            string         `json:"FullyQualifiedName"`
	Domain                        string         `json:"domain"`
	Name                          string         `json:"Name"`
	Classification                string         `json:"Classification"`
	AccountSubType                AccountSubType `json:"AccountSubType"`
	CurrencyRef                   Reference      `json:"CurrencyRef"`
	CurrentBalanceWithSubAccounts float64        `json:"CurrentBalanceWithSubAccounts"`
	Sparse                        bool           `json:"sparse"`
	MetaData                      MetaData       `json:"MetaData"`
	AccountType                   AccountType    `json:"AccountType"`
	CurrentBalance                float64        `json:"CurrentBalance"`
	Active                        bool           `json:"Active"`
	SyncToken                     string         `json:"SyncToken"`
	ID                            string         `json:"Id"`
	SubAccount                    bool           `json:"SubAccount"`
}

// CreateAccountRequest represents the minimum QuickBooks payload for creating an account.
type CreateAccountRequest struct {
	// Name is required and must be unique.
	Name string `json:"Name"`
	// AccountType is conditionally required when AccountSubType is not provided.
	AccountType AccountType `json:"AccountType,omitempty"`
	// AccountSubType is conditionally required when AccountType is not provided.
	AccountSubType AccountSubType `json:"AccountSubType,omitempty"`
	AcctNum        string         `json:"AcctNum,omitempty"`
	TaxCodeRef     *Reference     `json:"TaxCodeRef,omitempty"`
}

// Reference represents a QuickBooks reference object.
type Reference struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// MetaData represents QuickBooks object metadata timestamps.
type MetaData struct {
	CreateTime      time.Time `json:"CreateTime"`
	LastUpdatedTime time.Time `json:"LastUpdatedTime"`
}
