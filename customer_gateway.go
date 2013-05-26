package braintree

import (
	"errors"
)

type CustomerGateway struct {
	gateway Gateway
}

func (this CustomerGateway) Create(customer Customer) (CustomerResult, error) {
	customerXML, err := customer.ToXML()
	if err != nil {
		return ErrorResult{}, errors.New("Error encoding customer as XML: " + err.Error())
	}

	response, err := this.gateway.Execute("POST", "/customers", customerXML)
	if err != nil {
		return ErrorResult{}, err
	}

	if response.StatusCode == 201 {
		return response.CustomerResult()
	} else if response.StatusCode == 422 {
		return response.ErrorResult()
	}

	return ErrorResult{}, errors.New("Unexpected response from server: " + string(response.Status))
}
