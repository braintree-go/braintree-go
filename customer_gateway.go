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

func (this CustomerGateway) Find(id string) (CustomerResult, error) {
	response, err := this.gateway.Execute("GET", "/customers/"+id, []byte{})
	if err != nil {
		return ErrorResult{}, err
	} else if response.StatusCode == 200 {
		return response.CustomerResult()
	} else if response.StatusCode == 404 {
		return ErrorResult{}, errors.New("A customer with that ID could not be found")
	}
	return ErrorResult{}, errors.New("Unexpected response from server: " + string(response.Status))
}
