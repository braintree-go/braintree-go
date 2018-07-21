package braintree

import "encoding/xml"

type PaymentMethodNonce struct {
	XMLName     xml.Name `xml:"payment-method-nonce"`
	Type        string   `xml:"type"`
	Nonce       string   `xml:"nonce"`
	Description string   `xml:"description"`
}
