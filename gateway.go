package braintree

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io/ioutil"
	"net/http"
)

type Gateway interface {
	Execute(method, urlExtension string, body []byte) (*Response, error)
}

func NewGateway(config Configuration) BraintreeGateway {
	return BraintreeGateway{
		config: config,
		client: &http.Client{},
	}
}

type BraintreeGateway struct {
	config Configuration
	client *http.Client
}

func (this BraintreeGateway) Execute(method, urlExtension string, body []byte) (*Response, error) {
	bodyBuffer := bytes.NewBuffer(body)

	request, err := http.NewRequest(method, this.config.BaseURL()+urlExtension, bodyBuffer)
	if err != nil {
		return nil, errors.New("Error creating HTTP request: " + err.Error())
	}

	request.Header.Set("Content-Type", "application/xml")
	request.Header.Set("Accept", "application/xml")
	request.Header.Set("Accept-Encoding", "gzip")
	request.Header.Set("User-Agent", "Braintree-Go")
	request.Header.Set("X-ApiVersion", "3")
	request.SetBasicAuth(this.config.publicKey, this.config.privateKey)

	response, err := this.client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, errors.New("Error sending request to Braintree: " + err.Error())
	}

	gzipBody, err := gzip.NewReader(response.Body)
	defer gzipBody.Close()
	if err != nil {
		return nil, errors.New("Error reading gzipped response from Braintree: " + err.Error())
	}

	contents, err := ioutil.ReadAll(gzipBody)
	if err != nil {
		return nil, errors.New("Error reading response from Braintree: " + err.Error())
	}

	return &Response{StatusCode: response.StatusCode, Status: response.Status, Body: contents}, nil
}

func (this BraintreeGateway) Transaction() TransactionGateway {
	return TransactionGateway{this}
}

func (this BraintreeGateway) CreditCard() CreditCardGateway {
	return CreditCardGateway{this}
}

func (this BraintreeGateway) Customer() CustomerGateway {
	return CustomerGateway{this}
}
