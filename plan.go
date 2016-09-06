package braintree

import (
	"github.com/brianpowell/braintree-go/nullable"
	"time"
)

type Plan struct {
	XMLName               string              `xml:"plan" json:"plan" bson:"plan"`
	Id                    string              `xml:"id" json:"id" bson:"id"`
	MerchantId            string              `xml:"merchant-id" json:"merchantId" bson:"merchantId"`
	BillingDayOfMonth     *nullable.NullInt64 `xml:"billing-day-of-month" json:"billingDayOfMonth" bson:"billingDayOfMonth"`
	BillingFrequency      *nullable.NullInt64 `xml:"billing-frequency" json:"billingFrequency" bson:"billingFrequency"`
	CurrencyISOCode       string              `xml:"currency-iso-code" json:"currencyIsoCode" bson:"currencyIsoCode"`
	Description           string              `xml:"description" json:"description" bson:"description"`
	Name                  string              `xml:"name" json:"name" bson:"name"`
	NumberOfBillingCycles *nullable.NullInt64 `xml:"number-of-billing-cycles" json:"numberOfBillingCycles" bson:"numberOfBillingCycles"`
	Price                 *Decimal            `xml:"price" json:"price" bson:"price"`
	TrialDuration         *nullable.NullInt64 `xml:"trial-duration" json:"trialDuration" bson:"trialDuration"`
	TrialDurationUnit     string              `xml:"trial-duration-unit" json:"trialDurationUnit" bson:"trialDurationUnit"`
	TrialPeriod           *nullable.NullBool  `xml:"trial-period" json:"trialPeriod" bson:"trialPeriod"`
	CreatedAt             *time.Time          `xml:"created-at" json:"createdAt" bson:"createdAt"`
	UpdatedAt             *time.Time          `xml:"updated-at" json:"updatedAt" bson:"updatedAt"`
}

type Plans struct {
	XMLName string  `xml:"plans" json:"plans" bson:"plans"`
	Plan    []*Plan `xml:"plan" json:"plan" bson:"plan"`
}
