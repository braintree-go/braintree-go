package braintree

import "golang.org/x/net/context"

type AddOnGateway struct {
	*Braintree
}

func (g *AddOnGateway) All() ([]AddOn, error) {
	return g.AllContext(context.Background())
}

func (g *AddOnGateway) AllContext(ctx context.Context) ([]AddOn, error) {
	resp, err := g.executeContext(ctx, "GET", "add_ons", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.addOns()
	}
	return nil, &invalidResponseError{resp}
}
