package braintree

type PaymentMethodGateway interface {
	Create(*PaymentMethod) (string, interface{}, error)
	Delete(string) error
	Find(string) (interface{}, error)
	Update(string, *PaymentMethod) (interface{}, error)
}

type PaymentMethodGatewayImpl struct {
	*Braintree
}

func (g *PaymentMethodGatewayImpl) Create(paymentMethod *PaymentMethod) (string, interface{}, error) {
	resp, err := g.execute("POST", "payment_methods", paymentMethod)
	if err != nil {
		return "", nil, err
	}

	if resp.StatusCode != 201 {
		return "", nil, &invalidResponseError{resp}
	}

	if paypalAccount, err := resp.paypalAccount(); err == nil {
		return paypalAccount.Token, paypalAccount, nil
	}

	if creditCard, err := resp.creditCard(); err == nil {
		return creditCard.Token, creditCard, nil
	}

	return "", nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGatewayImpl) Update(token string, paymentMethod *PaymentMethod) (interface{}, error) {
	resp, err := g.execute("PUT", "payment_methods/any/"+token, paymentMethod)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, &invalidResponseError{resp}
	}

	if paypalAccount, err := resp.paypalAccount(); err == nil {
		return paypalAccount, nil
	}

	if creditCard, err := resp.creditCard(); err == nil {
		return creditCard, nil
	}

	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGatewayImpl) Find(token string) (interface{}, error) {
	resp, err := g.execute("GET", "payment_methods/any/"+token, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, &invalidResponseError{resp}
	}

	if paypalAccount, err := resp.paypalAccount(); err == nil {
		return paypalAccount, nil
	}

	if creditCard, err := resp.creditCard(); err == nil {
		return creditCard, nil
	}

	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGatewayImpl) Delete(token string) error {
	resp, err := g.execute("DELETE", "payment_methods/any/"+token, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return &invalidResponseError{resp}
	}
	return nil
}
