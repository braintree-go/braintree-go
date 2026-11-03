package braintree

import "time"

type CreditCard struct {
	CustomerId                string                         `xml:"customer-id,omitempty"`
	Token                     string                         `xml:"token,omitempty"`
	PaymentMethodNonce        string                         `xml:"payment-method-nonce,omitempty"`
	Number                    string                         `xml:"number,omitempty"`
	ExpirationDate            string                         `xml:"expiration-date,omitempty"`
	ExpirationMonth           string                         `xml:"expiration-month,omitempty"`
	ExpirationYear            string                         `xml:"expiration-year,omitempty"`
	CVV                       string                         `xml:"cvv,omitempty"`
	NetworkTokenAttributes    *NetworkTokenizationAttributes `xml:"network-tokenization-attributes,omitempty"`
	VenmoSDKPaymentMethodCode string                         `xml:"venmo-sdk-payment-method-code,omitempty"`
	VenmoSDK                  bool                           `xml:"venmo-sdk,omitempty"`
	Options                   *CreditCardOptions             `xml:"options,omitempty"`
	CreatedAt                 *time.Time                     `xml:"created-at,omitempty"`
	UpdatedAt                 *time.Time                     `xml:"updated-at,omitempty"`
	Bin                       string                         `xml:"bin,omitempty"`
	CardType                  string                         `xml:"card-type,omitempty"`
	CardholderName            string                         `xml:"cardholder-name,omitempty"`
	CustomerLocation          string                         `xml:"customer-location,omitempty"`
	ImageURL                  string                         `xml:"image-url,omitempty"`
	Default                   bool                           `xml:"default,omitempty"`
	Expired                   bool                           `xml:"expired,omitempty"`
	Last4                     string                         `xml:"last-4,omitempty"`
	Commercial                string                         `xml:"commercial,omitempty"`
	Debit                     string                         `xml:"debit,omitempty"`
	DurbinRegulated           string                         `xml:"durbin-regulated,omitempty"`
	Healthcare                string                         `xml:"healthcare,omitempty"`
	Payroll                   string                         `xml:"payroll,omitempty"`
	Prepaid                   string                         `xml:"prepaid,omitempty"`
	CountryOfIssuance         string                         `xml:"country-of-issuance,omitempty"`
	IssuingBank               string                         `xml:"issuing-bank,omitempty"`
	UniqueNumberIdentifier    string                         `xml:"unique-number-identifier,omitempty"`
	BillingAddress            *Address                       `xml:"billing-address,omitempty"`
	Subscriptions             *Subscriptions                 `xml:"subscriptions,omitempty"`
	Verifications             *CreditCardVerifications       `xml:"verifications,omitempty"`
}

type CreditCards struct {
	CreditCard []*CreditCard `xml:"credit-card"`
}

type CreditCardVerification struct {
	Id                           string      `xml:"id,omitempty"`
	Status                       string      `xml:"status,omitempty"`
	Amount                       string      `xml:"amount,omitempty"`
	CurrencyIsoCode              string      `xml:"currency-iso-code,omitempty"`
	MerchantAccountId            string      `xml:"merchant-account-id,omitempty"`
	ProcessorResponseCode        string      `xml:"processor-response-code,omitempty"`
	ProcessorResponseText        string      `xml:"processor-response-text,omitempty"`
	ProcessorResponseType        string      `xml:"processor-response-type,omitempty"`
	NetworkResponseCode          string      `xml:"network-response-code,omitempty"`
	NetworkResponseText          string      `xml:"network-response-text,omitempty"`
	NetworkTransactionID         string      `xml:"network-transaction-id,omitempty"`
	MastercardTransactionLinkID  string      `xml:"mastercard-transaction-link-id,omitempty"`
	AdditionalProcessorResponse  string      `xml:"additional-processor-response,omitempty"`
	GatewayRejectionReason       string      `xml:"gateway-rejection-reason,omitempty"`
	CvvResponseCode              string      `xml:"cvv-response-code,omitempty"`
	AvsErrorResponseCode         string      `xml:"avs-error-response-code,omitempty"`
	AvsPostalCodeResponseCode    string      `xml:"avs-postal-code-response-code,omitempty"`
	AvsStreetAddressResponseCode string      `xml:"avs-street-address-response-code,omitempty"`
	GraphQLId                    string      `xml:"graphql-id,omitempty"`
	CreatedAt                    *time.Time  `xml:"created-at,omitempty"`
	CreditCard                   *CreditCard `xml:"credit-card,omitempty"`
	Billing                      *Address    `xml:"billing,omitempty"`
	RiskData                     *RiskData   `xml:"risk-data,omitempty"`
}

type CreditCardVerifications struct {
	CreditCardVerification []*CreditCardVerification `xml:"verification"`
}

type RiskData struct {
	Id                   string   `xml:"id,omitempty"`
	Decision             string   `xml:"decision,omitempty"`
	DecisionReasons      []string `xml:"decision-reasons,omitempty"`
	DeviceDataCaptured   bool     `xml:"device-data-captured,omitempty"`
	FraudServiceProvider string   `xml:"fraud-service-provider,omitempty"`
	TransactionRiskScore string   `xml:"transaction-risk-score,omitempty"`
}

type CreditCardOptions struct {
	VerifyCard                    bool   `xml:"verify-card,omitempty"`
	VenmoSDKSession               string `xml:"venmo-sdk-session,omitempty"`
	MakeDefault                   bool   `xml:"make-default,omitempty"`
	FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty"`
	VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty"`
	UpdateExistingToken           string `xml:"update-existing-token,omitempty"`
}

type NetworkTokenizationAttributes struct {
	Cryptogram         string `xml:"cryptogram,omitempty"`
	EcommerceIndicator string `xml:"ecommerce-indicator,omitempty"`
	TokenRequestorID   string `xml:"token-requestor-id,omitempty"`
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

// AllVerifications returns all verifications for this card, or nil if none present.
func (card *CreditCard) AllVerifications() []*CreditCardVerification {
	if card.Verifications != nil {
		verifs := card.Verifications.CreditCardVerification
		if len(verifs) > 0 {
			a := make([]*CreditCardVerification, 0, len(verifs))
			for _, v := range verifs {
				a = append(a, v)
			}
			return a
		}
	}
	return nil
}
