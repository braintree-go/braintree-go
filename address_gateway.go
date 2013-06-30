package braintree

import (
	"encoding/xml"
)

type AddressGateway struct {
	*Braintree
}

func (g *AddressGateway) Create(a *Address) (*Address, error) {
	cid := a.CustomerId
	a.CustomerId = ""
	a.XMLName = xml.Name{Local: "address"}
	resp, err := g.execute("POST", "customers/"+cid+"/addresses", a)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.address()
	}
	return nil, &InvalidResponseError{resp}
}

func (g *AddressGateway) Delete(customerId, addrId string) error {
	resp, err := g.execute("DELETE", "customers/"+customerId+"/addresses/"+addrId, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &InvalidResponseError{resp}
}
