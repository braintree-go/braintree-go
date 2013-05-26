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
	if responseCode == 201 {
		response, err = ParseTransactionResponse(responseBody)
	} else if responseCode == 422 {
		response, err = ParseErrorResponse(responseBody)
	} else {
		err = errors.New("Unknown response code: " + string(responseCode))
	}

	if err != nil {
		return ErrorResponse{}, err
	}
	return response, nil
}
