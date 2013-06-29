package braintree

type CustomerGateway struct {
	*Braintree
}

func (g *CustomerGateway) Create(c *Customer) (*Customer, error) {
	resp, err := g.Execute("POST", "customers", c)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.Customer()
	}
	return nil, &InvalidResponseError{resp}
}

func (g *CustomerGateway) Find(id string) (*Customer, error) {
	resp, err := g.Execute("GET", "customers/"+id, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.Customer()
	}
	return nil, &InvalidResponseError{resp}
}
