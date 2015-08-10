package braintree

type PaypalAccountGateway struct {
	*Braintree
}

func (g *PaypalAccountGateway) Update(paypalAccount *PaypalAccount) (*PaypalAccount, error) {
	resp, err := g.execute("PUT", "payment_methods/paypal_account/"+paypalAccount.Token, paypalAccount)
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
	resp, err := g.execute("GET", "payment_methods/paypal_account/"+token, nil)
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
	resp, err := g.execute("DELETE", "payment_methods/paypal_account/"+paypalAccount.Token, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}
