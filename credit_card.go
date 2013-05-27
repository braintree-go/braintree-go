package braintree

import "encoding/xml"

type CreditCard struct {
	CustomerId     string             `xml:"customer-id,omitempty"`
	Token          string             `xml:"token,omitempty"`
	Number         string             `xml:"number"`
	ExpirationDate string             `xml:"expiration-date"`
	CVV            string             `xml:"cvv,omitempty"`
	Options        *CreditCardOptions `xml:"options,omitempty"`
}

func (this CreditCard) ToXML() ([]byte, error) {
	xml, err := xml.Marshal(this)
	if err != nil {
		return []byte{}, err
	}
	return xml, nil
}

type CreditCardOptions struct {
	VerifyCard bool `xml:"verify-card,omitempty"`
}
