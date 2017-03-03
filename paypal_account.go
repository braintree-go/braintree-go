package braintree

import "time"

type PayPalAccount struct {
	CustomerId    string         `xml:"customer-id,omitempty"`
	Token         string         `xml:"token,omitempty"`
	Email         string         `xml:"email,omitempty"`
	ImageURL      string         `xml:"image-url,omitempty"`
	CreatedAt     *time.Time     `xml:"created-at,omitempty"`
	UpdatedAt     *time.Time     `xml:"updated-at,omitempty"`
	Subscriptions *Subscriptions `xml:"subscriptions,omitempty"`
}

type PayPalAccounts struct {
	PayPalAccount []*PayPalAccount `xml:"paypal-account"`
}

func (paypalAccount *PayPalAccount) GetCustomerId() string {
	return paypalAccount.CustomerId
}

func (paypalAccount *PayPalAccount) GetToken() string {
	return paypalAccount.Token
}

// AllSubscriptions returns all subscriptions for this paypal account, or nil if none present.
func (paypalAccount *PayPalAccount) AllSubscriptions() []*Subscription {
	if paypalAccount.Subscriptions != nil {
		subs := paypalAccount.Subscriptions.Subscription
		if len(subs) > 0 {
			a := make([]*Subscription, 0, len(subs))
			for _, s := range subs {
				a = append(a, s)
			}
			return a
		}
	}
	return nil
}
