package braintree

type CustomerGateway struct {
	*Braintree
}

func (g *CustomerGateway) Create(c *Customer) (*Customer, error) {
	resp, err := g.execute("POST", "customers", c)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.customer()
	}
	return nil, &InvalidResponseError{resp}
}

func (g *CustomerGateway) Find(id string) (*Customer, error) {
	resp, err := g.execute("GET", "customers/"+id, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.customer()
	}
	return nil, &InvalidResponseError{resp}
}
