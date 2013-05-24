package braintree

import (
	"bytes"
	"io"
	"net/http"
)

func NewGateway(config Configuration) Gateway {
	return Gateway{
		config: config,
		client: &http.Client{},
	}
}

type Gateway struct {
	config Configuration
	client *http.Client
}

func (this Gateway) newRequest(method, urlExtension string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, this.config.BaseURL()+urlExtension, body)
	if err != nil {
		return request, err
	}

	request.Header.Set("Content-Type", "application/xml")
	request.Header.Set("Accept", "application/xml")
	request.Header.Set("Accept-Encoding", "gzip")
	request.Header.Set("User-Agent", "Braintree-Go")
	request.Header.Set("X-ApiVersion", "2.0.0")
	request.SetBasicAuth(this.config.publicKey, this.config.privateKey)

	return request, nil
}

func (this Gateway) ExecuteTransactionRequest(tx TransactionRequest) (TransactionResponse, error) {
	requestBytes, err := tx.ToXML()
	if err != nil {
		return TransactionResponse{}, err
	}
	requestBody := bytes.NewBuffer(requestBytes)
	request, err := this.newRequest("POST", "/transactions", requestBody)
	if err != nil {
		return TransactionResponse{}, err
	}
	response, err := this.client.Do(request)
	if err != nil {
		return TransactionResponse{}, err
	}
	if response.Status == "201 Created" {
		return TransactionResponse{true}, nil
	}
	return TransactionResponse{}, nil
}

type TransactionResponse struct {
	Success bool
}
