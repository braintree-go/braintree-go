package braintree

import (
	"encoding/xml"
	"errors"
)

type CreditCardGateway struct {
	gateway Gateway
}

func (this CreditCardGateway) Create(card CreditCard) (CreditCardResult, error) {
	cardXML, err := xml.Marshal(card)
	if err != nil {
		return ErrorResult{}, errors.New("Error encoding credit card as XML: " + err.Error())
	}

	response, err := this.gateway.Execute("POST", "/payment_methods", cardXML)
	if err != nil {
		return ErrorResult{}, err
	}

	if response.StatusCode == 201 {
		return response.CreditCardResult()
	} else if response.StatusCode == 422 {
		return response.ErrorResult()
	}

	return ErrorResult{}, errors.New("Unexpected response from server: " + string(response.Status))
}

func (this CreditCardGateway) Find(token string) (CreditCardResult, error) {
	response, err := this.gateway.Execute("GET", "/payment_methods/"+token, []byte{})
	if err != nil {
		return ErrorResult{}, err
	} else if response.StatusCode == 200 {
		return response.CreditCardResult()
	} else if response.StatusCode == 404 {
		return ErrorResult{}, errors.New("A credit card with that token could not be found")
	}
	return ErrorResult{}, errors.New("Unexpected response from server: " + response.Status)
}
