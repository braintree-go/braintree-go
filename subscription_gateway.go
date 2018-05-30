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

// RetryCharge retries to charge for a Subscription. All options,
// including the Subscription ID, are to be provided by the
// SubscriptionTransactionRequest passed as an argument. Note that the
// Amount has to be > 0.
func (g *SubscriptionGateway) RetryCharge(ctx context.Context, txReq *SubscriptionTransactionRequest) error {
	resp, err := g.execute(ctx, "POST", "transactions", txReq)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 201:
		return nil
	}
	return &invalidResponseError{resp}
}
