package braintree

import (
	"encoding/xml"
	"errors"
)

type Response struct {
	StatusCode int
	Status     string
	Body       []byte
}

func (this Response) TransactionResult() (TransactionResult, error) {
	var tx Transaction
	err := xml.Unmarshal(this.Body, &tx)
	if err != nil {
		return ErrorResult{}, errors.New("Error unmarshalling transaction XML: " + err.Error())
	}
	return SuccessfulTransactionResult{tx}, nil
}

func (this Response) CustomerResult() (CustomerResult, error) {
	var customer Customer
	err := xml.Unmarshal(this.Body, &customer)
	if err != nil {
		return ErrorResult{}, errors.New("Error unmarshalling customer XML: " + err.Error())
	}
	return SuccessfulCustomerResult{customer}, nil
}

func (this Response) ErrorResult() (ErrorResult, error) {
	var result ErrorResult
	err := xml.Unmarshal(this.Body, &result)
	if err != nil {
		return ErrorResult{}, errors.New("Error unmarshalling error XML: " + err.Error())
	}
	return result, nil
}

type TransactionResult interface {
	Transaction() Transaction
	Success() bool
	Message() string
}

type SuccessfulTransactionResult struct {
	tx Transaction
}

func (this SuccessfulTransactionResult) Transaction() Transaction {
	return this.tx
}

func (this SuccessfulTransactionResult) Success() bool {
	return true
}

func (this SuccessfulTransactionResult) Message() string {
	return ""
}

type CustomerResult interface {
	Customer() Customer
	Success() bool
	Message() string
}

type SuccessfulCustomerResult struct {
	customer Customer
}

func (this SuccessfulCustomerResult) Customer() Customer {
	return this.customer
}

func (this SuccessfulCustomerResult) Success() bool {
	return true
}

func (this SuccessfulCustomerResult) Message() string {
	return ""
}

type ErrorResult struct {
	XMLName      string `xml:"api-error-response"`
	ErrorMessage string `xml:"message"`
}

func (this ErrorResult) Success() bool {
	return false
}

func (this ErrorResult) Message() string {
	return this.ErrorMessage
}

func (this ErrorResult) Transaction() Transaction {
	panic("Transaction() called on ErrorResult")
}

func (this ErrorResult) Customer() Customer {
	panic("Customer() called on ErrorResult")
}
