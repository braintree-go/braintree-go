package braintree

import "errors"

type CreditCardGateway struct {
	gateway Gateway
}

func (this CreditCardGateway) Create(card CreditCard) (CreditCardResult, error) {
	cardXML, err := card.ToXML()
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
