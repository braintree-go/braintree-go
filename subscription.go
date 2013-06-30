package braintree

type Subscription struct {
	XMLName            string `xml:"subscription"`
	Id                 string `xml:"id,omitempty"`
	PaymentMethodToken string `xml:"payment-method-token"`
	PlanId             string `xml:"plan-id"`
}
