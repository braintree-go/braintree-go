package braintree

import "context"

type CreditCardGateway struct {
	*Braintree
}

// Create creates a new credit card.
func (g *CreditCardGateway) Create(ctx context.Context, card *CreditCard) (*CreditCard, error) {
	resp, err := g.execute(ctx, "POST", "payment_methods", card)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.creditCard()
	}
	return nil, &invalidResponseError{resp}
}

// Update updates a credit card.
func (g *CreditCardGateway) Update(ctx context.Context, card *CreditCard) (*CreditCard, error) {
	resp, err := g.execute(ctx, "PUT", "payment_methods/"+card.Token, card)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.creditCard()
	}
	return nil, &invalidResponseError{resp}
}

// Find finds a credit card by payment method token.
func (g *CreditCardGateway) Find(ctx context.Context, token string) (*CreditCard, error) {
	resp, err := g.execute(ctx, "GET", "payment_methods/"+token, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.creditCard()
	}
	return nil, &invalidResponseError{resp}
}

// Delete deletes a credit card.
func (g *CreditCardGateway) Delete(ctx context.Context, card *CreditCard) error {
	resp, err := g.execute(ctx, "DELETE", "payment_methods/"+card.Token, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}
