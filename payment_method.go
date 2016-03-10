package braintree

type PaymentMethod struct {
	CustomerId         string `xml:"customer-id,omitempty"`
	Token              string `xml:"token,omitempty"`
	PaymentMethodNonce string `xml:"payment-method-nonce,omitempty"`
}
