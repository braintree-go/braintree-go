package braintree

import (
	"encoding/xml"
)

type CustomerGateway struct {
	*Braintree
}

func (g *CustomerGateway) Create(c *Customer) (*Customer, error) {
	xmlBody, err := xml.Marshal(c)
	if err != nil {
		return nil, err
	}
	resp, err := g.Execute("POST", "customers", xmlBody)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		cust, err := resp.Customer()
		return cust, err
	}
	return nil, &InvalidResponseError{resp}
}

func (g *CustomerGateway) Find(id string) (*Customer, error) {
	resp, err := g.Execute("GET", "customers/"+id, []byte{})
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		cust, err := resp.Customer()
		return cust, err
	}
	return nil, &InvalidResponseError{resp}
}
