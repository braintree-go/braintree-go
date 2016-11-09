package braintree

type PaymentMethodGateway struct {
	*Braintree
}

func (g *PaymentMethodGateway) Create(paymentmethod *PaymentMethod) (*PaymentMethod, error) {
	resp, err := g.execute("POST", "payment_methods", paymentmethod)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGateway) Update(paymentmethod *PaymentMethod) (*PaymentMethod, error) {
	resp, err := g.execute("PUT", "payment_methods/any/"+paymentmethod.Token, paymentmethod)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGateway) Find(token string) (*PaymentMethod, error) {
	resp, err := g.execute("GET", "payment_methods/"+token, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGateway) Delete(paymentmethod *PaymentMethod) error {
	resp, err := g.execute("DELETE", "payment_methods/any/"+paymentmethod.Token, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}
