package braintree

import (
	"bytes"
	"errors"
)

type TransactionGateway struct {
	gateway Gateway
}

func (this TransactionGateway) Sale(tx Transaction) (TransactionResult, error) {
	transactionXML, err := tx.ToXML()
	if err != nil {
		return ErrorResult{}, errors.New("Error encoding transaction as XML: " + err.Error())
	}

	requestBody := bytes.NewBuffer(transactionXML)
	response, err := this.gateway.Execute("POST", "/transactions", requestBody)
	if err != nil {
		return ErrorResult{}, err
	}

	if response.StatusCode == 201 {
		return response.TransactionResult()
	} else if response.StatusCode == 422 {
		return response.ErrorResult()
	}
	return ErrorResult{}, errors.New("Unexpected response from server: " + string(response.Status))
}

func (this TransactionGateway) Find(txId string) (TransactionResult, error) {
	response, err := this.gateway.Execute("GET", "/transactions/"+txId, bytes.NewBuffer([]byte{}))
	if err != nil {
		return ErrorResult{}, err
	} else if response.StatusCode == 200 {
		return response.TransactionResult()
	} else if response.StatusCode == 404 {
		return ErrorResult{}, errors.New("A transaction with that ID could not be found")
	}
	return ErrorResult{}, errors.New("Unexpected response from server: " + string(response.Status))
}
