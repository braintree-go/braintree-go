package braintree

import "github.com/lionelbarrow/braintree-go/nullable"

const (
	SubscriptionStatusActive       = "Active"
	SubscriptionStatusCanceled     = "Canceled"
	SubscriptionStatusExpired      = "Expired"
	SubscriptionStatusPastDue      = "Past Due"
	SubscriptionStatusPending      = "Pending"
	SubscriptionStatusUnrecognized = "Unrecognized"
)

type Subscription struct {
	XMLName                 string               `xml:"subscription"`
	Id                      string               `xml:"id,omitempty"`
	Balance                 *Decimal             `xml:"balance,omitempty"`
	BillingDayOfMonth       string               `xml:"billing-day-of-month,omitempty"`
	BillingPeriodEndDate    string               `xml:"billing-period-end-date,omitempty"`
	BillingPeriodStartDate  string               `xml:"billing-period-start-date,omitempty"`
	CurrentBillingCycle     string               `xml:"current-billing-cycle,omitempty"`
	DaysPastDue             string               `xml:"days-past-due,omitempty"`
	Discounts               []interface{}        `xml:"discounts,omitempty"`
	FailureCount            string               `xml:"failure-count,omitempty"`
	FirstBillingDate        string               `xml:"first-billing-date,omitempty"`
	MerchantAccountId       string               `xml:"merchant-account-id,omitempty"`
	NeverExpires            *nullable.NullBool   `xml:"never-expires,omitempty"`
	NextBillAmount          *Decimal             `xml:"next-bill-amount,omitempty"`
	NextBillingPeriodAmount *Decimal             `xml:"next-billing-period-amount,omitempty"`
	NextBillingDate         string               `xml:"next-billing-date,omitempty"`
	NumberOfBillingCycles   *nullable.NullInt64  `xml:"number-of-billing-cycles,omitempty"`
	PaidThroughDate         string               `xml:"paid-through-date,omitempty"`
	PaymentMethodToken      string               `xml:"payment-method-token,omitempty"`
	PlanId                  string               `xml:"plan-id,omitempty"`
	Price                   *Decimal             `xml:"price,omitempty"`
	Status                  string               `xml:"status,omitempty"`
	TrialDuration           string               `xml:"trial-duration,omitempty"`
	TrialDurationUnit       string               `xml:"trial-duration-unit,omitempty"`
	TrialPeriod             *nullable.NullBool   `xml:"trial-period,omitempty"`
	Transactions            *Transactions        `xml:"transactions,omitempty"`
	Options                 *SubscriptionOptions `xml:"options,omitempty"`
	// AddOns                  []interface{} `xml:"add-ons,omitempty"`
	// Descriptor              interface{}   `xml:"descriptor,omitempty"`   // struct with name, phone
}

type Subscriptions struct {
	Subscription []*Subscription `xml:"subscription"`
}

type SubscriptionOptions struct {
	DoNotInheritAddOnsOrDiscounts        bool `xml:"do-not-inherit-add-ons-or-discounts,omitempty"`
	ProrateCharges                       bool `xml:"prorate-charges,omitempty"`
	ReplaceAllAddOnsAndDiscounts         bool `xml:"replace-all-add-ons-and-discounts,omitempty"`
	RevertSubscriptionOnProrationFailure bool `xml:"revert-subscription-on-proration-failure,omitempty"`
	StartImmediately                     bool `xml:"start-immediately,omitempty"`
}
