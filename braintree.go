package braintree

import (
	"bytes"
	"net/http"
)

func New(c Config) *Braintree {
	return &Braintree{c}
}

type Braintree struct {
	Config
}

func (g *Braintree) Execute(method, path string, body []byte) (*Response, error) {
	req, err := http.NewRequest(method, g.BaseURL()+"/"+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", "Braintree-Go")
	req.Header.Set("X-ApiVersion", "3")
	req.SetBasicAuth(g.PublicKey, g.PrivateKey)

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	btr := &Response{
		Response: resp,
	}
	err = btr.unpackBody()
	if err != nil {
		return nil, err
	}
	err = btr.apiError()
	if err != nil {
		return nil, err
	}
	return btr, nil
}

func (g *Braintree) Transaction() *TransactionGateway {
	return &TransactionGateway{g}
}

func (g *Braintree) CreditCard() *CreditCardGateway {
	return &CreditCardGateway{g}
}

func (g *Braintree) Customer() *CustomerGateway {
	return &CustomerGateway{g}
}
