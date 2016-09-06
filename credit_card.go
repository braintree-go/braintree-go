package braintree

import "time"

type CreditCard struct {
	CustomerId                string             `xml:"customer-id,omitempty" json:"customerId,omitempty" bson:"customerId,omitempty"`
	Token                     string             `xml:"token,omitempty" json:"token,omitempty" bson:"token,omitempty"`
	PaymentMethodNonce        string             `xml:"payment-method-nonce,omitempty" json:"paymentMethodNonce,omitempty" bson:"paymentMethodNonce,omitempty"`
	Number                    string             `xml:"number,omitempty" json:"number,omitempty" bson:"number,omitempty"`
	ExpirationDate            string             `xml:"expiration-date,omitempty" json:"expirationDate,omitempty" bson:"expirationDate,omitempty"`
	ExpirationMonth           string             `xml:"expiration-month,omitempty" json:"expirationMonth,omitempty" bson:"expirationMonth,omitempty"`
	ExpirationYear            string             `xml:"expiration-year,omitempty" json:"expirationYear,omitempty" bson:"expirationYear,omitempty"`
	CVV                       string             `xml:"cvv,omitempty" json:"cvv,omitempty" bson:"cvv,omitempty"`
	VenmoSDKPaymentMethodCode string             `xml:"venmo-sdk-payment-method-code,omitempty" json:"venmoSDKPaymentMethodCode,omitempty" bson:"venmoSDKPaymentMethodCode,omitempty"`
	VenmoSDK                  bool               `xml:"venmo-sdk,omitempty" json:"venmoSDK,omitempty" bson:"venmoSDK,omitempty"`
	Options                   *CreditCardOptions `xml:"options,omitempty" json:"options,omitempty" bson:"options,omitempty"`
	CreatedAt                 *time.Time         `xml:"created-at,omitempty" json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt                 *time.Time         `xml:"updated-at,omitempty" json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Bin                       string             `xml:"bin,omitempty" json:"bin,omitempty" bson:"bin,omitempty"`
	CardType                  string             `xml:"card-type,omitempty" json:"cardType,omitempty" bson:"cardType,omitempty"`
	CardholderName            string             `xml:"cardholder-name,omitempty" json:"cardholderName,omitempty" bson:"cardholderName,omitempty"`
	CustomerLocation          string             `xml:"customer-location,omitempty" json:"customerLocation,omitempty" bson:"customerLocation,omitempty"`
	ImageURL                  string             `xml:"image-url,omitempty" json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	Default                   bool               `xml:"default,omitempty" json:"default,omitempty" bson:"default,omitempty"`
	Expired                   bool               `xml:"expired,omitempty" json:"expired,omitempty" bson:"expired,omitempty"`
	Last4                     string             `xml:"last-4,omitempty" json:"last4,omitempty" bson:"last4,omitempty"`
	Commercial                string             `xml:"commercial,omitempty" json:"commercial,omitempty" bson:"commercial,omitempty"`
	Debit                     string             `xml:"debit,omitempty" json:"debit,omitempty" bson:"debit,omitempty"`
	DurbinRegulated           string             `xml:"durbin-regulated,omitempty" json:"durbinRegulated,omitempty" bson:"durbinRegulated,omitempty"`
	Healthcare                string             `xml:"healthcare,omitempty" json:"healthcare,omitempty" bson:"healthcare,omitempty"`
	Payroll                   string             `xml:"payroll,omitempty" json:"payroll,omitempty" bson:"payroll,omitempty"`
	Prepaid                   string             `xml:"prepaid,omitempty" json:"prepaid,omitempty" bson:"prepaid,omitempty"`
	CountryOfIssuance         string             `xml:"country-of-issuance,omitempty" json:"countryOfIssuance,omitempty" bson:"countryOfIssuance,omitempty"`
	IssuingBank               string             `xml:"issuing-bank,omitempty" json:"issuingBank,omitempty" bson:"issuingBank,omitempty"`
	UniqueNumberIdentifier    string             `xml:"unique-number-identifier,omitempty" json:"uniqueNumberIdentifier,omitempty" bson:"uniqueNumberIdentifier,omitempty"`
	BillingAddress            *Address           `xml:"billing-address,omitempty" json:"billingAddress,omitempty" bson:"billingAddress,omitempty"`
	Subscriptions             *Subscriptions     `xml:"subscriptions,omitempty" json:"subscriptions,omitempty" bson:"subscriptions,omitempty"`
}

type CreditCards struct {
	CreditCard []*CreditCard `xml:"credit-card" json:"creditCard" bson:"creditCard"`
}

type CreditCardOptions struct {
	VerifyCard                    bool   `xml:"verify-card,omitempty" json:"verifyCard,omitempty" bson:"verifyCard,omitempty"`
	VenmoSDKSession               string `xml:"venmo-sdk-session,omitempty" json:"venmoSDKSession,omitempty" bson:"venmoSDKSession,omitempty"`
	MakeDefault                   bool   `xml:"make-default,omitempty" json:"makeDefault,omitempty" bson:"makeDefault,omitempty"`
	FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty" json:"failOnDuplicatePaymentMethod,omitempty" bson:"failOnDuplicatePaymentMethod,omitempty"`
	VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty" json:"verificationMerchantAccountId,omitempty" bson:"verificationMerchantAccountId,omitempty"`
	UpdateExistingToken           string `xml:"update-existing-token,omitempty" json:"updateExistingToken,omitempty" bson:"updateExistingToken,omitempty"`
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
