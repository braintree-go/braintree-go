package braintree

import "context"

type SubscriptionGateway struct {
	*Braintree
}

func (g *SubscriptionGateway) Create(ctx context.Context, sub *SubscriptionRequest) (*Subscription, error) {
	resp, err := g.execute(ctx, "POST", "subscriptions", sub)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.subscription()
	}
	return nil, &invalidResponseError{resp}
}

func (g *SubscriptionGateway) Update(ctx context.Context, sub *SubscriptionRequest) (*Subscription, error) {
	resp, err := g.execute(ctx, "PUT", "subscriptions/"+sub.Id, sub)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.subscription()
	}
	return nil, &invalidResponseError{resp}
}

func (g *SubscriptionGateway) Find(ctx context.Context, subId string) (*Subscription, error) {
	resp, err := g.execute(ctx, "GET", "subscriptions/"+subId, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.subscription()
	}
	return nil, &invalidResponseError{resp}
}

func (g *SubscriptionGateway) Cancel(ctx context.Context, subId string) (*Subscription, error) {
	resp, err := g.execute(ctx, "PUT", "subscriptions/"+subId+"/cancel", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.subscription()
	}
	return nil, &invalidResponseError{resp}
}
