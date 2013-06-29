package braintree

import (
	"encoding/xml"
)

type CreditCardGateway struct {
	*Braintree
}

func (g *CreditCardGateway) Create(card *CreditCard) (*CreditCard, error) {
	xmlBody, err := xml.Marshal(card)
	if err != nil {
		return nil, err
	}
	resp, err := g.Execute("POST", "payment_methods", xmlBody)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		card, err := resp.CreditCard()
		return card, err
	}
	return nil, &InvalidResponseError{resp}
}

func (g *CreditCardGateway) Find(token string) (*CreditCard, error) {
	resp, err := g.Execute("GET", "payment_methods/"+token, []byte{})
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		card, err := resp.CreditCard()
		return card, err
	}
	return nil, &InvalidResponseError{resp}
}
