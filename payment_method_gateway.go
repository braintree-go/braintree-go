package braintree

type PaymentMethodGateway interface {
	Create(*PaymentMethod) (*PaymentMethod, error)
	Delete(string) error
	Find(string) (*PaymentMethod, error)
	Update(string, *PaymentMethod) (*PaymentMethod, error)
}

type PaymentMethodGatewayImpl struct {
	*Braintree
}

func (g *PaymentMethodGatewayImpl) Create(paymentMethod *PaymentMethod) (*PaymentMethod, error) {
	resp, err := g.execute("POST", "payment_methods", paymentMethod)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGatewayImpl) Update(token string, paymentMethod *PaymentMethod) (*PaymentMethod, error) {
	resp, err := g.execute("PUT", "payment_methods/any/"+token, paymentMethod)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGatewayImpl) Find(token string) (*PaymentMethod, error) {
	resp, err := g.execute("GET", "payment_methods/any/"+token, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGatewayImpl) Delete(token string) error {
	resp, err := g.execute("DELETE", "payment_methods/any/"+token, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}
