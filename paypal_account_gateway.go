package braintree

type PaypalAccountGateway struct {
	*Braintree
}

func (g *PaypalAccountGateway) Update(paypalAccount *PaypalAccount) (*PaypalAccount, error) {
	resp, err := g.executeVersion("PUT", "payment_methods/paypal_account/"+paypalAccount.Token, paypalAccount, ApiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paypalAccount()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaypalAccountGateway) Find(token string) (*PaypalAccount, error) {
	resp, err := g.executeVersion("GET", "payment_methods/paypal_account/"+token, nil, ApiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paypalAccount()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaypalAccountGateway) Delete(paypalAccount *PaypalAccount) error {
	resp, err := g.executeVersion("DELETE", "payment_methods/paypal_account/"+paypalAccount.Token, nil, ApiVersion4)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}
