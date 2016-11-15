package braintree

type PaymentMethod struct {
	CustomerId         string                `xml:"customer-id,omitempty"`
	Token              string                `xml:"token,omitempty"`
	PaymentMethodNonce string                `xml:"payment-method-nonce,omitempty"`
	Options            *PaymentMethodOptions `xml:"options,omitempty"`
}

type PaymentMethodOptions struct {
	VerifyCard                    bool   `xml:"verify-card,omitempty"`
	MakeDefault                   bool   `xml:"make-default,omitempty"`
	FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty"`
	VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty"`
}
