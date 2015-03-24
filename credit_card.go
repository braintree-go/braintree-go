package braintree

type CreditCard struct {
	CustomerId                string             `json:"customer-id,omitempty" xml:"customer-id,omitempty"`
	Token                     string             `json:"token,omitempty" xml:"token,omitempty"`
	Number                    string             `json:"number,omitempty" xml:"number,omitempty"`
	ExpirationDate            string             `json:"expiration-date,omitempty" xml:"expiration-date,omitempty"`
	ExpirationMonth           string             `json:"expiration-month,omitempty" xml:"expiration-month,omitempty"`
	ExpirationYear            string             `json:"expiration-year,omitempty" xml:"expiration-year,omitempty"`
	CVV                       string             `json:"cvv,omitempty" xml:"cvv,omitempty"`
	VenmoSDKPaymentMethodCode string             `json:"venmo-sdk-payment-method-code,omitempty" xml:"venmo-sdk-payment-method-code,omitempty"`
	VenmoSDK                  bool               `json:"venmo-sdk,omitempty" xml:"venmo-sdk,omitempty"`
	Options                   *CreditCardOptions `json:"options,omitempty" xml:"options,omitempty"`
	CreatedAt                 string             `json:"created-at,omitempty" xml:"created-at,omitempty"`
	UpdatedAt                 string             `json:"updated-at,omitempty" xml:"updated-at,omitempty"`
	Bin                       string             `json:"bin,omitempty" xml:"bin,omitempty"`
	CardType                  string             `json:"card-type,omitempty" xml:"card-type,omitempty"`
	CardholderName            string             `json:"cardholder-name,omitempty" xml:"cardholder-name,omitempty"`
	CustomerLocation          string             `json:"customer-location,omitempty" xml:"customer-location,omitempty"`
	ImageURL                  string             `json:"image-url,omitempty" xml:"image-url,omitempty"`
	Default                   bool               `json:"default,omitempty" xml:"default,omitempty"`
	Expired                   bool               `json:"expired,omitempty" xml:"expired,omitempty"`
	Last4                     string             `json:"last-4,omitempty" xml:"last-4,omitempty"`
	Commercial                string             `json:"commercial,omitempty" xml:"commercial,omitempty"`
	Debit                     string             `json:"debit,omitempty" xml:"debit,omitempty"`
	DurbinRegulated           string             `json:"durbin-regulated,omitempty" xml:"durbin-regulated,omitempty"`
	Healthcare                string             `json:"healthcare,omitempty" xml:"healthcare,omitempty"`
	Payroll                   string             `json:"payroll,omitempty" xml:"payroll,omitempty"`
	Prepaid                   string             `json:"prepaid,omitempty" xml:"prepaid,omitempty"`
	CountryOfIssuance         string             `json:"country-of-issuance,omitempty" xml:"country-of-issuance,omitempty"`
	IssuingBank               string             `json:"issuing-bank,omitempty" xml:"issuing-bank,omitempty"`
	UniqueNumberIdentifier    string             `json:"unique-number-identifier,omitempty" xml:"unique-number-identifier,omitempty"`
	BillingAddress            *Address           `json:"billing-address,omitempty" xml:"billing-address,omitempty"`
	Subscriptions             *Subscriptions     `json:"subscriptions,omitempty" xml:"subscriptions,omitempty"`
}

type CreditCards struct {
	CreditCard []*CreditCard `json:"credit-card" xml:"credit-card"`
}

type CreditCardOptions struct {
	VerifyCard                    bool   `json:"verify-card,omitempty" xml:"verify-card,omitempty"`
	VenmoSDKSession               string `json:"venmo-sdk-session,omitempty" xml:"venmo-sdk-session,omitempty"`
	MakeDefault                   bool   `json:"make-default,omitempty" xml:"make-default,omitempty"`
	FailOnDuplicatePaymentMethod  bool   `json:"fail-on-duplicate-payment-method,omitempty" xml:"fail-on-duplicate-payment-method,omitempty"`
	VerificationMerchantAccountId string `json:"verification-merchant-account-id,omitempty" xml:"verification-merchant-account-id,omitempty"`
	UpdateExistingToken           string `json:"update-existing-token,omitempty" xml:"update-existing-token,omitempty"`
}

// AllSubscriptions returns all subscriptions for this card, or nil if none present.
func (card *CreditCard) AllSubscriptions() []*Subscription {
	if card.Subscriptions != nil {
		subs := card.Subscriptions.Subscription
		if len(subs) > 0 {
			a := make([]*Subscription, 0, len(subs))
			for _, s := range subs {
				a = append(a, s)
			}
			return a
		}
	}
	return nil
}
