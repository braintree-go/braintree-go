package braintree

import "encoding/xml"

type Transaction struct {
	XMLName    string     `xml:"transaction"`
	Id         string     `xml:"id,omitempty"`
	OrderId    string     `xml:"order_id,omitempty"`
	Type       string     `xml:"type"`
	Amount     float64    `xml:"amount"`
	CreditCard CreditCard `xml:"credit-card"`
}

func (this Transaction) ToXML() ([]byte, error) {
	xml, err := xml.Marshal(this)
	if err != nil {
		return []byte{}, err
	}
	return xml, nil
}
