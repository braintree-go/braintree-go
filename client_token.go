package braintree

// ClientTokenRequest represents the parameters for the client token request.
type ClientTokenRequest struct {
	XMLName           string                     `xml:"client-token"`
	CustomerID        string                     `xml:"customer-id,omitempty"`
	MerchantAccountID string                     `xml:"merchant-account-id,omitempty"`
	Options           *ClientTokenRequestOptions `xml:"options,omitempty"`
	Version           int                        `xml:"version"`
}

type clientToken struct {
	ClientToken string `xml:"value"`
}

// ClientTokenRequestOptions represents options map for the client token request.
type ClientTokenRequestOptions struct {
	FailOnDuplicatePaymentMethod bool  `xml:"fail-on-duplicate-payment-method,omitempty"`
	MakeDefault                  bool  `xml:"make-default,omitempty"`
	VerifyCard                   *bool `xml:"verify-card,omitempty"`
}
