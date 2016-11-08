package braintree

type PaymentMethod struct {
	CustomerId         string                `xml:"customer-id,omitempty" json:"customerId,omitempty" bson:"customerId,omitempty"`
	Token              string                `xml:"token,omitempty" json:"token,omitempty" bson:"token,omitempty"`
	PaymentMethodNonce string                `xml:"payment-method-nonce,omitempty" json:"paymentMethodNonce,omitempty" bson:"paymentMethodNonce,omitempty"`
	Number             string                `xml:"number,omitempty" json:"number,omitempty" bson:"number,omitempty"`
	ExpirationDate     string                `xml:"expiration-date,omitempty" json:"expirationDate,omitempty" bson:"expirationDate,omitempty"`
	ExpirationMonth    string                `xml:"expiration-month,omitempty" json:"expirationMonth,omitempty" bson:"expirationMonth,omitempty"`
	ExpirationYear     string                `xml:"expiration-year,omitempty" json:"expirationYear,omitempty" bson:"expirationYear,omitempty"`
	CVV                string                `xml:"cvv,omitempty" json:"cvv,omitempty" bson:"cvv,omitempty"`
	Options            *PaymentMethodOptions `xml:"options,omitempty" json:"options,omitempty" bson:"options,omitempty"`
	CardholderName     string                `xml:"cardholder-name,omitempty" json:"cardholderName,omitempty" bson:"cardholderName,omitempty"`
	DeviceData         string                `xml:"device-data,omitempty" json:"deviceData,omitempty" bson:"deviceData,omitempty"`
	BillingId          string                `xml:"billing-id,omitempty" json:"billingId,omitempty" bson:"billingId,omitempty"`
	BillingAddress     *Address              `xml:"billing-address,omitempty" json:"billingAddress,omitempty" bson:"billingAddress,omitempty"`
}

type PaymentMethodOptions struct {
	VerifyCard                    bool   `xml:"verify-card,omitempty" json:"verifyCard,omitempty" bson:"verifyCard,omitempty"`
	MakeDefault                   bool   `xml:"make-default,omitempty" json:"makeDefault,omitempty" bson:"makeDefault,omitempty"`
	FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty" json:"failOnDuplicatePaymentMethod,omitempty" bson:"failOnDuplicatePaymentMethod,omitempty"`
	VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty" json:"verificationMerchantAccountId,omitempty" bson:"verificationMerchantAccountId,omitempty"`
}

type PaymentMethods struct {
	PaymentMethod []*PaymentMethod `xml:"payment-method" json:"paymentMethod" bson:"paymentMethod"`
}
