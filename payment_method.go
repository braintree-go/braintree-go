package braintree

type PaymentMethod struct {
	CustomerId         string                `xml:"customer-id,omitempty" json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	Token              string                `xml:"token,omitempty" json:"token,omitempty" bson:"token,omitempty"`
	PaymentMethodNonce string                `xml:"payment-method-nonce,omitempty" json:"payment_method_nonce,omitempty" bson:"payment_method_nonce,omitempty"`
	Number             string                `xml:"number,omitempty" json:"number,omitempty" bson:"number,omitempty"`
	ExpirationDate     string                `xml:"expiration-date,omitempty" json:"expiration_date,omitempty" bson:"expiration_date,omitempty"`
	ExpirationMonth    string                `xml:"expiration-month,omitempty" json:"expiration_month,omitempty" bson:"expiration_month,omitempty"`
	ExpirationYear     string                `xml:"expiration-year,omitempty" json:"expiration_year,omitempty" bson:"expiration_year,omitempty"`
	CVV                string                `xml:"cvv,omitempty" json:"cvv,omitempty" bson:"cvv,omitempty"`
	Options            *PaymentMethodOptions `xml:"options,omitempty" json:"options,omitempty" bson:"options,omitempty"`
	CardholderName     string                `xml:"cardholder-name,omitempty" json:"cardholder_name,omitempty" bson:"cardholder_name,omitempty"`
	DeviceData         string                `xml:"device-data,omitempty" json:"device_data,omitempty" bson:"device_data,omitempty"`
	BillingId          string                `xml:"billing-id,omitempty" json:"billing_id,omitempty" bson:"billing_id,omitempty"`
	BillingAddress     *Address              `xml:"billing-address,omitempty" json:"billing_address,omitempty" bson:"billing_address,omitempty"`
}

type PaymentMethodOptions struct {
	VerifyCard                    bool   `xml:"verify-card,omitempty" json:"verify_card,omitempty" bson:"verify_card,omitempty"`
	MakeDefault                   bool   `xml:"make-default,omitempty" json:"make_default,omitempty" bson:"make_default,omitempty"`
	FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty" json:"fail_on_duplicate_payment_method,omitempty" bson:"fail_on_duplicate_payment_method,omitempty"`
	VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty" json:"verification_merchant_account_id,omitempty" bson:"verification_merchant_account_id,omitempty"`
}

type PaymentMethods struct {
	ID             string           `xml:"id,omitempty" json:"id,omitempty" bson:"_id,omitempty"`
	PaymentMethods []*PaymentMethod `xml:"payment-methods" json:"payment_methods" bson:"payment_methods"`
}
