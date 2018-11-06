package braintree

import (
	"context"
	"encoding/xml"
)

const clientTokenVersion = 2

// ClientTokenGateway represents the provider client token generation.
type ClientTokenGateway struct {
	*Braintree
}

// Generate generates a new client token.
func (g *ClientTokenGateway) Generate(ctx context.Context) (string, error) {
	return g.generate(ctx, &ClientTokenRequest{
		Version: clientTokenVersion,
	})
}

// GenerateWithCustomer generates a new client token for the customer id.
func (g *ClientTokenGateway) GenerateWithCustomer(ctx context.Context, customerID string) (string, error) {
	return g.generate(ctx, &ClientTokenRequest{
		Version:    clientTokenVersion,
		CustomerID: customerID,
	})
}

// GenerateWithRequest generates a new client token using custom request options.
func (g *ClientTokenGateway) GenerateWithRequest(ctx context.Context, req *ClientTokenRequest) (string, error) {
	if req == nil {
		req = &ClientTokenRequest{}
	}
	if req.Version == 0 {
		req.Version = clientTokenVersion
	}
	return g.generate(ctx, req)
}

func (g *ClientTokenGateway) generate(ctx context.Context, req *ClientTokenRequest) (string, error) {
	resp, err := g.execute(ctx, "POST", "client_token", req)
	if err != nil {
		return "", err
	}
	switch resp.StatusCode {
	case 201:
		var b clientToken
		if err := xml.Unmarshal(resp.Body, &b); err != nil {
			return "", err
		}
		return b.ClientToken, nil
	}
	return "", &invalidResponseError{resp}
}
