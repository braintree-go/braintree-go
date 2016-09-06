package braintree

import "github.com/brianpowell/braintree-go/nullable"

const (
	SubscriptionStatusActive       = "Active"
	SubscriptionStatusCanceled     = "Canceled"
	SubscriptionStatusExpired      = "Expired"
	SubscriptionStatusPastDue      = "Past Due"
	SubscriptionStatusPending      = "Pending"
	SubscriptionStatusUnrecognized = "Unrecognized"
)

type Subscription struct {
	XMLName                 string               `xml:"subscription" json:"subscription" bson:"subscription"`
	Id                      string               `xml:"id,omitempty" json:"id,omitempty" bson:"id,omitempty"`
	Balance                 *Decimal             `xml:"balance,omitempty" json:"balance,omitempty" bson:"balance,omitempty"`
	BillingDayOfMonth       string               `xml:"billing-day-of-month,omitempty" json:"billingDayOfMonth,omitempty" bson:"billingDayOfMonth,omitempty"`
	BillingPeriodEndDate    string               `xml:"billing-period-end-date,omitempty" json:"billingPeriodEndDate,omitempty" bson:"billingPeriodEndDate,omitempty"`
	BillingPeriodStartDate  string               `xml:"billing-period-start-date,omitempty" json:"billingPeriodStartDate,omitempty" bson:"billingPeriodStartDate,omitempty"`
	CurrentBillingCycle     string               `xml:"current-billing-cycle,omitempty" json:"currentBillingCycle,omitempty" bson:"currentBillingCycle,omitempty"`
	DaysPastDue             string               `xml:"days-past-due,omitempty" json:"daysPastDue,omitempty" bson:"daysPastDue,omitempty"`
	Discounts               []interface{}        `xml:"discounts,omitempty" json:"discounts,omitempty" bson:"discounts,omitempty"`
	FailureCount            string               `xml:"failure-count,omitempty" json:"failureCount,omitempty" bson:"failureCount,omitempty"`
	FirstBillingDate        string               `xml:"first-billing-date,omitempty" json:"firstBillingDate,omitempty" bson:"firstBillingDate,omitempty"`
	MerchantAccountId       string               `xml:"merchant-account-id,omitempty" json:"merchantAccountId,omitempty" bson:"merchantAccountId,omitempty"`
	NeverExpires            *nullable.NullBool   `xml:"never-expires,omitempty" json:"neverExpires,omitempty" bson:"neverExpires,omitempty"`
	NextBillAmount          *Decimal             `xml:"next-bill-amount,omitempty" json:"nextBillAmount,omitempty" bson:"nextBillAmount,omitempty"`
	NextBillingPeriodAmount *Decimal             `xml:"next-billing-period-amount,omitempty" json:"nextBillingPeriodAmount,omitempty" bson:"nextBillingPeriodAmount,omitempty"`
	NextBillingDate         string               `xml:"next-billing-date,omitempty" json:"next-billing-date,omitempty" bson:"next-billing-date,omitempty"`
	NumberOfBillingCycles   *nullable.NullInt64  `xml:"number-of-billing-cycles,omitempty" json:"numberOfBillingCycles,omitempty" bson:"numberOfBillingCycles,omitempty"`
	PaidThroughDate         string               `xml:"paid-through-date,omitempty" json:"paidThroughDate,omitempty" bson:"paidThroughDate,omitempty"`
	PaymentMethodToken      string               `xml:"payment-method-token,omitempty" json:"paymentMethodToken,omitempty" bson:"paymentMethodToken,omitempty"`
	PlanId                  string               `xml:"plan-id,omitempty" json:"planId,omitempty" bson:"planId,omitempty"`
	Price                   *Decimal             `xml:"price,omitempty" json:"price,omitempty" bson:"price,omitempty"`
	Status                  string               `xml:"status,omitempty" json:"status,omitempty" bson:"status,omitempty"`
	TrialDuration           string               `xml:"trial-duration,omitempty" json:"trialDuration,omitempty" bson:"trialDuration,omitempty"`
	TrialDurationUnit       string               `xml:"trial-duration-unit,omitempty" json:"trialDurationUnit,omitempty" bson:"trialDurationUnit,omitempty"`
	TrialPeriod             *nullable.NullBool   `xml:"trial-period,omitempty" json:"trialPeriod,omitempty" bson:"trialPeriod,omitempty"`
	Transactions            *Transactions        `xml:"transactions,omitempty" json:"transactions,omitempty" bson:"transactions,omitempty"`
	Options                 *SubscriptionOptions `xml:"options,omitempty" json:"options,omitempty" bson:"options,omitempty"`
	// AddOns                  []interface{} `xml:"add-ons,omitempty" json:"add-ons,omitempty" bson:"add-ons,omitempty"`
	// Descriptor              interface{}   `xml:"descriptor,omitempty"`   // struct with name, p json:"descriptor,omitempty"`   // struct with name, " bson:"descriptor,omitempty"`   // struct with name, "hone
}

type Subscriptions struct {
	Subscription []*Subscription `xml:"subscription" json:"subscription" bson:"subscription"`
}

type SubscriptionOptions struct {
	DoNotInheritAddOnsOrDiscounts        bool `xml:"do-not-inherit-add-ons-or-discounts,omitempty" json:"doNotInheritAddOnsOrDiscounts,omitempty" bson:"doNotInheritAddOnsOrDiscounts,omitempty"`
	ProrateCharges                       bool `xml:"prorate-charges,omitempty" json:"prorateCharges,omitempty" bson:"prorateCharges,omitempty"`
	ReplaceAllAddOnsAndDiscounts         bool `xml:"replace-all-add-ons-and-discounts,omitempty" json:"replaceAllAddOnsAndDiscounts,omitempty" bson:"replaceAllAddOnsAndDiscounts,omitempty"`
	RevertSubscriptionOnProrationFailure bool `xml:"revert-subscription-on-proration-failure,omitempty" json:"revertSubscriptionOnProrationFailure,omitempty" bson:"revertSubscriptionOnProrationFailure,omitempty"`
	StartImmediately                     bool `xml:"start-immediately,omitempty" json:"startImmediately,omitempty" bson:"startImmediately,omitempty"`
}
