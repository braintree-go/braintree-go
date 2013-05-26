package braintree

import "encoding/xml"

type Transaction struct {
	XMLName    string              `xml:"transaction"`
	Id         string              `xml:"id,omitempty"`
	Status     string              `xml:"status,omitempty"`
	Type       string              `xml:"type,omitempty"`
	Amount     float64             `xml:"amount"`
	CreditCard *CreditCard         `xml:"credit-card,omitempty"`
	Options    *TransactionOptions `xml:"options,omitempty"`
}

type TransactionOptions struct {
	SubmitForSettlement bool `xml:"submit-for-settlement,omitempty"`
}

func (this Transaction) ToXML() ([]byte, error) {
	xml, err := xml.Marshal(this)
	if err != nil {
		return []byte{}, err
	}
	return xml, nil
}
