package braintree

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
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
	request.Header.Set("X-ApiVersion", "3")
	request.SetBasicAuth(this.config.publicKey, this.config.privateKey)

	return request, nil
}

func (this Gateway) executeRequest(method, urlExtension string, body io.Reader) ([]byte, int, error) {
	request, err := this.newRequest(method, urlExtension, body)
	if err != nil {
		return []byte{}, 0, errors.New("Error creating HTTP request: " + err.Error())
	}

	response, err := this.client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return []byte{}, 0, errors.New("Error sending request to Braintree: " + err.Error())
	}

	gzipBody, err := gzip.NewReader(response.Body)
	defer gzipBody.Close()
	if err != nil {
		return []byte{}, 0, err
	}
	contents, err := ioutil.ReadAll(gzipBody)
	if err != nil {
		return []byte{}, 0, errors.New("Error reading response from Braintree: " + err.Error())
	}

	if response.StatusCode == 201 || response.StatusCode == 422 {
		return contents, response.StatusCode, nil
	}

	return []byte{}, response.StatusCode, errors.New("Got unexpected response from Braintree: " + response.Status)
}

func (this Gateway) ExecuteTransactionRequest(tx TransactionRequest) (TransactionResponse, error) {
	requestBytes, err := tx.ToXML()
	if err != nil {
		return TransactionResponse{}, err
	}
	requestBody := bytes.NewBuffer(requestBytes)
	responseBody, responseCode, err := this.executeRequest("POST", "/transactions", requestBody)
	if err != nil {
		return TransactionResponse{false, ""}, err
	}

	switch responseCode {
	case 201:
		return TransactionResponse{true, ""}, nil
	case 422:
		return TransactionResponse{false, string(responseBody)}, nil
	}
	return TransactionResponse{false, ""}, errors.New("Should never get here")
}

type TransactionResponse struct {
	Success bool
	Message string
}
