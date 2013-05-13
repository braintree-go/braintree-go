package braintree

import (
  "bytes"
  "net/http"
  "fmt"
  "io"
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
  request, err := http.NewRequest(method, this.config.BaseURL() + urlExtension, body)
  if err != nil {
    return request, err
  }

  request.Header.Set("Content-Type", "application/xml")
  request.Header.Set("Accept", "application/xml")
  request.Header.Set("Accept-Encoding", "gzip")

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

  fmt.Println(request)
  response, err := this.client.Do(request)
  fmt.Println(response)
  fmt.Println(err)
  return TransactionResponse{}, nil
}

type TransactionResponse struct{}

func (this TransactionResponse) IsSuccess() bool { return false }
