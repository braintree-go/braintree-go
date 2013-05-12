package braintree

import "encoding/xml"

type Transaction struct {
	XMLName    string     `xml:"transaction"`
	Amount     int        `xml:"amount"`
	CreditCard CreditCard `xml:"credit-card"`
}

func NewTransactionRequest(tx Transaction) TransactionRequest {
	return TransactionRequest{tx}
}

type TransactionRequest struct {
	tx Transaction
}

func (this TransactionRequest) ToXML() ([]byte, error) {
	xml, err := xml.Marshal(this.tx)
	if err != nil {
		return []byte{}, err
	}
	return xml, nil
}
