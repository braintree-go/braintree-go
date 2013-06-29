package braintree

type Address struct {
	StreetAddress string `xml:"street-address,omitempty"`
	Locality      string `xml:"locality,omitempty"`
	Region        string `xml:"region,omitempty"`
	PostalCode    string `xml:"postal-code,omitempty"`
}

type CreditCard struct {
	CustomerId                string             `xml:"customer-id,omitempty"`
	Token                     string             `xml:"token,omitempty"`
	Number                    string             `xml:"number,omitempty"`
	ExpirationDate            string             `xml:"expiration-date,omitempty"`
	CVV                       string             `xml:"cvv,omitempty"`
	VenmoSDKPaymentMethodCode string             `xml:"venmo-sdk-payment-method-code,omitempty"`
	VenmoSDK                  bool               `xml:"venmo-sdk,omitempty"`
	Options                   *CreditCardOptions `xml:"options,omitempty"`
}

type CreditCardOptions struct {
	VerifyCard      bool   `xml:"verify-card,omitempty"`
	VenmoSDKSession string `xml:"venmo-sdk-session,omitempty"`
}

type Customer struct {
	XMLName     string       `xml:"customer"`
	Id          string       `xml:"id,omitempty"`
	FirstName   string       `xml:"first-name,omitempty"`
	LastName    string       `xml:"last-name,omitempty"`
	Company     string       `xml:"company,omitempty"`
	Email       string       `xml:"email,omitempty"`
	Phone       string       `xml:"phone,omitempty"`
	Fax         string       `xml:"fax,omitempty"`
	Website     string       `xml:"website,omitempty"`
	CreditCard  *CreditCard  `xml:"credit-card,omitempty"`
	CreditCards []CreditCard `xml:"credit-cards,omitempty"`
}

type Subscription struct {
	XMLName            string `xml:"subscription"`
	PaymentMethodToken string `xml:"payment-method-token"`
	PlanId             string `xml:"plan-id"`
	TrialPeriod        bool   `xml:"trial-period,omitempty"`
	TrialDuration      int    `xml:"trial-duration,omitempty"`
	TrialDurationUnit  string `xml:"trial-duration-unit,omitempty"`
}

type Transaction struct {
	XMLName            string              `xml:"transaction"`
	Id                 string              `xml:"id,omitempty"`
	CustomerID         string              `xml:"customer-id,omitempty"`
	Status             string              `xml:"status,omitempty"`
	Type               string              `xml:"type,omitempty"`
	Amount             float64             `xml:"amount"`
	OrderId            string              `xml:"order-id,omitempty"`
	PaymentMethodToken string              `xml:"payment-method-token,omitempty"`
	MerchantAccountId  string              `xml:"merchant-account-id,omitempty"`
	CreditCard         *CreditCard         `xml:"credit-card,omitempty"`
	Customer           *Customer           `xml:"customer,omitempty"`
	BillingAddress     *Address            `xml:"billing,omitempty"`
	ShippingAddress    *Address            `xml:"shipping,omitempty"`
	Options            *TransactionOptions `xml:"options,omitempty"`
}

type TransactionOptions struct {
	SubmitForSettlement              bool `xml:"submit-for-settlement,omitempty"`
	StoreInVault                     bool `xml:"store-in-vault,omitempty"`
	AddBillingAddressToPaymentMethod bool `xml:"add-billing-address-to-payment-method,omitempty"`
	StoreShippingAddressInVault      bool `xml:"store-shipping-address-in-vault,omitempty"`
}
