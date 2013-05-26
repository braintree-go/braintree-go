package braintree

type CreditCard struct {
	Number         string             `xml:"number"`
	ExpirationDate string             `xml:"expiration-date"`
	CVV            string             `xml:"cvv,omitempty"`
	Options        *CreditCardOptions `xml:"options,omitempty"`
}

type CreditCardOptions struct {
	VerifyCard bool `xml:"verify-card,omitempty"`
}
