package braintree

import (
	"github.com/lionelbarrow/braintree-go/nullable"
	"time"
)

type Plan struct {
	XMLName               string              `xml:"plan"`
	Id                    string              `xml:"id"`
	MerchantId            string              `xml:"merchant-id"`
	BillingDayOfMonth     *nullable.NullInt64 `xml:"billing-day-of-month"`
	BillingFrequency      *nullable.NullInt64 `xml:"billing-frequency"`
	CurrencyISOCode       string              `xml:"currency-iso-code"`
	Description           string              `xml:"description"`
	Name                  string              `xml:"name"`
	NumberOfBillingCycles *nullable.NullInt64 `xml:"number-of-billing-cycles"`
	Price                 *Decimal            `xml:"price"`
	TrialDuration         *nullable.NullInt64 `xml:"trial-duration"`
	TrialDurationUnit     string              `xml:"trial-duration-unit"`
	TrialPeriod           *nullable.NullBool  `xml:"trial-period"`
	CreatedAt             *time.Time          `xml:"created-at"`
	UpdatedAt             *time.Time          `xml:"updated-at"`
}

type Plans struct {
	XMLName string  `xml:"plans"`
	Plan    []*Plan `xml:"plan"`
}
