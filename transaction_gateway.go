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

	var response Response
	switch responseCode {
	case 201:
		response, err = ParseTransactionResponse(responseBody)
	case 422:
		response, err = ParseErrorResponse(responseBody)
	}
	if err != nil {
		return ErrorResponse{}, err
	}
	return response, nil
}
