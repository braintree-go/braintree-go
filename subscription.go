package braintree

const (
	SubscriptionStatusActive       = "Active"
	SubscriptionStatusCanceled     = "Canceled"
	SubscriptionStatusExpired      = "Expired"
	SubscriptionStatusPastDue      = "Past Due"
	SubscriptionStatusPending      = "Pending"
	SubscriptionStatusUnrecognized = "Unrecognized"
)

type Subscription struct {
	XMLName                 string               `json:"subscription" xml:"subscription"`
	Id                      string               `json:"id,omitempty" xml:"id,omitempty"`
	Balance                 float64              `json:"balance,omitempty" xml:"balance,omitempty"`
	BillingDayOfMonth       string               `json:"billing-day-of-month,omitempty" xml:"billing-day-of-month,omitempty"`
	BillingPeriodEndDate    string               `json:"billing-period-end-date,omitempty" xml:"billing-period-end-date,omitempty"`
	BillingPeriodStartDate  string               `json:"billing-period-start-date,omitempty" xml:"billing-period-start-date,omitempty"`
	CurrentBillingCycle     string               `json:"current-billing-cycle,omitempty" xml:"current-billing-cycle,omitempty"`
	DaysPastDue             string               `json:"days-past-due,omitempty" xml:"days-past-due,omitempty"`
	Discounts               []interface{}        `json:"discounts,omitempty" xml:"discounts,omitempty"`
	FailureCount            string               `json:"failure-count,omitempty" xml:"failure-count,omitempty"`
	FirstBillingDate        string               `json:"first-billing-date,omitempty" xml:"first-billing-date,omitempty"`
	MerchantAccountId       string               `json:"merchant-account-id,omitempty" xml:"merchant-account-id,omitempty"`
	NeverExpires            string               `json:"never-expires,omitempty" xml:"never-expires,omitempty"` // bool
	NextBillAmount          float64              `json:"next-bill-amount,omitempty" xml:"next-bill-amount,omitempty"`
	NextBillingPeriodAmount float64              `json:"next-billing-period-amount,omitempty" xml:"next-billing-period-amount,omitempty"`
	NextBillingDate         string               `json:"next-billing-date,omitempty" xml:"next-billing-date,omitempty"`
	NumberOfBillingCycles   string               `json:"number-of-billing-cycles,omitempty" xml:"number-of-billing-cycles,omitempty"` // int
	PaidThroughDate         string               `json:"paid-through-date,omitempty" xml:"paid-through-date,omitempty"`
	PaymentMethodToken      string               `json:"payment-method-token,omitempty" xml:"payment-method-token,omitempty"`
	PlanId                  string               `json:"plan-id,omitempty" xml:"plan-id,omitempty"`
	Price                   float64              `json:"price,omitempty" xml:"price,omitempty"`
	Status                  string               `json:"status,omitempty" xml:"status,omitempty"`
	TrialDuration           string               `json:"trial-duration,omitempty" xml:"trial-duration,omitempty"`
	TrialDurationUnit       string               `json:"trial-duration-unit,omitempty" xml:"trial-duration-unit,omitempty"`
	TrialPeriod             string               `json:"trial-period,omitempty" xml:"trial-period,omitempty"` // bool
	Transactions            *Transactions        `json:"transactions,omitempty" xml:"transactions,omitempty"`
	Options                 *SubscriptionOptions `json:"options,omitempty" xml:"options,omitempty"`
	// AddOns                  []interface{} `json:"add-ons,omitempty" xml:"add-ons,omitempty"`
	// Descriptor              interface{}   `json:"descriptor,omitempty" xml:"descriptor,omitempty"`   // struct with name, phone
}

type Subscriptions struct {
	Subscription []*Subscription `json:"subscription" xml:"subscription"`
}

// TODO(eaigner): same considerations apply as with plan type marshalling

type SubscriptionOptions struct {
	DoNotInheritAddOnsOrDiscounts        bool `json:"do-not-inherit-add-ons-or-discounts,omitempty" xml:"do-not-inherit-add-ons-or-discounts,omitempty"`
	ProrateCharges                       bool `json:"prorate-charges,omitempty" xml:"prorate-charges,omitempty"`
	ReplaceAllAddOnsAndDiscounts         bool `json:"replace-all-add-ons-and-discounts,omitempty" xml:"replace-all-add-ons-and-discounts,omitempty"`
	RevertSubscriptionOnProrationFailure bool `json:"revert-subscription-on-proration-failure,omitempty" xml:"revert-subscription-on-proration-failure,omitempty"`
	StartImmediately                     bool `json:"start-immediately,omitempty" xml:"start-immediately,omitempty"`
}

