package braintree

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
