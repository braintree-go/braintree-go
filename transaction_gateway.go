package braintree

import (
	"encoding/xml"
	"errors"
)

type TransactionGateway struct {
	gateway Gateway
}

func (this TransactionGateway) Create(tx Transaction) (TransactionResult, error) {
	transactionXML, err := xml.Marshal(tx)
	if err != nil {
		return ErrorResult{}, errors.New("Error encoding transaction as XML: " + err.Error())
	}

	response, err := this.gateway.Execute("POST", "/transactions", transactionXML)
	if err != nil {
		return ErrorResult{}, err
	}

	if response.StatusCode == 201 {
		return response.TransactionResult()
	} else if response.StatusCode == 422 {
		return response.ErrorResult()
	}

	return ErrorResult{}, errors.New("Unexpected response from server: " + response.Status)
}

func (this TransactionGateway) Find(txId string) (TransactionResult, error) {
	response, err := this.gateway.Execute("GET", "/transactions/"+txId, []byte{})
	if err != nil {
		return ErrorResult{}, err
	} else if response.StatusCode == 200 {
		return response.TransactionResult()
	} else if response.StatusCode == 404 {
		return ErrorResult{}, errors.New("A transaction with that ID could not be found")
	}
	return ErrorResult{}, errors.New("Unexpected response from server: " + response.Status)
}
