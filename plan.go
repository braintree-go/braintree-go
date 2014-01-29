package braintree

type Plan struct {
	XMLName               string  `json:"plan" xml:"plan"`
	Id                    string  `json:"id" xml:"id"`
	MerchantId            string  `json:"merchant-id" xml:"merchant-id"`
	BillingDayOfMonth     string  `json:"billing-day-of-month" xml:"billing-day-of-month"` // int
	BillingFrequency      string  `json:"billing-frequency" xml:"billing-frequency"`    // int
	CurrencyISOCode       string  `json:"currency-iso-code" xml:"currency-iso-code"`
	Description           string  `json:"description" xml:"description"`
	Name                  string  `json:"name" xml:"name"`
	NumberOfBillingCycles string  `json:"number-of-billing-cycles" xml:"number-of-billing-cycles"` // int
	Price                 float64 `json:"price" xml:"price"`
	TrialDuration         string  `json:"trial-duration" xml:"trial-duration"` // int
	TrialDurationUnit     string  `json:"trial-duration-unit" xml:"trial-duration-unit"`
	TrialPeriod           string  `json:"trial-period" xml:"trial-period"` // bool
	CreatedAt             string  `json:"created-at" xml:"created-at"`
	UpdatedAt             string  `json:"updated-at" xml:"updated-at"`
	// AddOns                []interface{} `json:"add-ons" xml:"add-ons"`
	// Discounts             []interface{} `json:"discounts" xml:"discounts"`
}

// TODO(eaigner): it is suboptimal that we use string instead of int/bool types here,
// but I see no way around this atm to avoid integer parse errors if the field is empty.
//
// If there is a better method, and it can be unmarshalled directly to the correct type
// without errors, this needs to be changed.

type Plans struct {
	XMLName string  `json:"plans" xml:"plans"`
	Plan    []*Plan `json:"plan" xml:"plan"`
}

