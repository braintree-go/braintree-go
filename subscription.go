package braintree

type Subscription struct {
	XMLName            string `xml:"subscription"`
	PaymentMethodToken string `xml:"payment-method-token"`
	PlanId             string `xml:"plan-id"`
	TrialPeriod        bool   `xml:"trial-period,omitempty"`
	TrialDuration      int    `xml:"trial-duration,omitempty"`
	TrialDurationUnit  string `xml:"trial-duration-unit,omitempty"`
}
