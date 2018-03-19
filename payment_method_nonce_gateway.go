package braintree

import (
	"context"
)

type PaymentMethodNonceGateway struct {
	*Braintree
}

func (g *PaymentMethodNonceGateway) Find(ctx context.Context, nonce *Nonce) (*Nonce, error) {
	resp, err := g.executeVersion(ctx, "GET", g.MerchantURL()+"/payment_method_nonces/"+nonce.Nonce, nil, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.nonce()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodNonceGateway) Create(ctx context.Context, token string) (*Nonce, error) {
	resp, err := g.executeVersion(ctx, "POST", g.MerchantURL()+"/payment_methods/"+token+"/nonces", nil, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.nonce()
	}
	return nil, &invalidResponseError{resp}
}
