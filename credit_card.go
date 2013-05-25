package braintree

type CreditCard struct {
	Number         string `xml:"number"`
	ExpirationDate string `xml:"expiration-date"`
}
