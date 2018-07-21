package braintree

import (
	"context"
)

type PaymentMethodGateway struct {
	*Braintree
}

func (g *PaymentMethodGateway) Create(ctx context.Context, paymentMethodRequest *PaymentMethodRequest) (PaymentMethod, error) {
	resp, err := g.executeVersion(ctx, "POST", "payment_methods", paymentMethodRequest, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGateway) Update(ctx context.Context, token string, paymentMethod *PaymentMethodRequest) (PaymentMethod, error) {
	resp, err := g.executeVersion(ctx, "PUT", "payment_methods/any/"+token, paymentMethod, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGateway) Find(ctx context.Context, token string) (PaymentMethod, error) {
	resp, err := g.executeVersion(ctx, "GET", "payment_methods/any/"+token, nil, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGateway) Delete(ctx context.Context, token string) error {
	resp, err := g.executeVersion(ctx, "DELETE", "payment_methods/any/"+token, nil, apiVersion4)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}
