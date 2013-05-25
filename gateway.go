package braintree

import (
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

type Gateway interface {
	Execute(method, urlExtension string, body io.Reader) ([]byte, int, error)
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

func (this BraintreeGateway) Execute(method, urlExtension string, body io.Reader) ([]byte, int, error) {
	request, err := http.NewRequest(method, this.config.BaseURL()+urlExtension, body)
	if err != nil {
		return []byte{}, 0, errors.New("Error creating HTTP request: " + err.Error())
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
		return []byte{}, 0, errors.New("Error sending request to Braintree: " + err.Error())
	}

	gzipBody, err := gzip.NewReader(response.Body)
	defer gzipBody.Close()
	if err != nil {
		return []byte{}, 0, errors.New("Error reading gzipped response from Braintree: " + err.Error())
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
