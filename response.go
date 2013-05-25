package braintree

import (
	"encoding/xml"
)

type Response interface {
	IsSuccess() bool
	GetMessage() string
	Transaction() Transaction
}

func ParseErrorResponse(content []byte) (ErrorResponse, error) {
	var response ErrorResponse
	err := xml.Unmarshal(content, &response)
	if err != nil {
		return ErrorResponse{}, err
	}
	return response, nil
}

type ErrorResponse struct {
	XMLName string `xml:"api-error-response"`
	Message string `xml:"message"`
}

func (this ErrorResponse) IsSuccess() bool {
	return false
}

func (this ErrorResponse) GetMessage() string {
	return this.Message
}

func (this ErrorResponse) Transaction() Transaction {
	panic("Transaction() called on ErrorResponse")
}

func ParseTransactionResponse(content []byte) (TransactionResponse, error) {
	var tx Transaction
	err := xml.Unmarshal(content, &tx)
	if err != nil {
		return TransactionResponse{}, err
	}
	return TransactionResponse{tx}, nil
}

type TransactionResponse struct {
	Tx Transaction
}

func (this TransactionResponse) IsSuccess() bool {
	return true
}

func (this TransactionResponse) GetMessage() string {
	return ""
}

func (this TransactionResponse) Transaction() Transaction {
	return this.Tx
}
