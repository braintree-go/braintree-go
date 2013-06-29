package braintree

type SubscriptionGateway struct {
	*Braintree
}

func (g *SubscriptionGateway) Create(sub *Subscription) (*Subscription, error) {
	resp, err := g.Execute("POST", "subscriptions", sub)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.Subscription()
	}
	return nil, &InvalidResponseError{resp}
}

func (g *SubscriptionGateway) Find(subId string) (*Subscription, error) {
	resp, err := g.Execute("GET", "subscriptions/"+subId, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.Subscription()
	}
	return nil, &InvalidResponseError{resp}
}

func (g *SubscriptionGateway) Cancel(subId string) (*Subscription, error) {
	resp, err := g.Execute("PUT", "subscriptions/"+subId+"/cancel", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.Subscription()
	}
	return nil, &InvalidResponseError{resp}
}
