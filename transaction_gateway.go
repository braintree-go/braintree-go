package braintree

import (
	"bytes"
	"errors"
)

type TransactionGateway struct {
	gateway Gateway
}

func (this TransactionGateway) Sale(tx Transaction) (Response, error) {
	transactionXML, err := tx.ToXML()
	if err != nil {
		return ErrorResponse{}, errors.New("Error encoding transaction as XML: " + err.Error())
	}

	requestBody := bytes.NewBuffer(transactionXML)
	responseBody, responseCode, err := this.gateway.Execute("POST", "/transactions", requestBody)
	if err != nil {
		return ErrorResponse{}, err
	}

	if responseCode == 201 {
		txResponse, err := ParseTransactionResponse(responseBody)
		if err != nil {
			return ErrorResponse{}, errors.New("Error decoding transaction response XML: " + err.Error())
		}
		return txResponse, nil
	}

	return ParseErrorResponse(responseBody)
}

func (this TransactionGateway) Find(txId string) (Response, error) {
	responseBody, _, err := this.gateway.Execute("GET", "/transactions/"+txId, bytes.NewBuffer([]byte{}))
	if err != nil {
		if err.Error() == "Got unexpected response from Braintree: 404 Not Found" {
			return ErrorResponse{}, errors.New("A transaction with that ID could not be found")
		} else {
			return ErrorResponse{}, err
		}
	}

	return ParseTransactionResponse(responseBody)
}
