package braintree

type Plan struct {
	XMLName               string  `xml:"plan"`
	Id                    string  `xml:"id"`
	MerchantId            string  `xml:"merchant-id"`
	BillingDayOfMonth     string  `xml:"billing-day-of-month"` // int
	BillingFrequency      string  `xml:"billing-frequency"`    // int
	CurrencyISOCode       string  `xml:"currency-iso-code"`
	Description           string  `xml:"description"`
	Name                  string  `xml:"name"`
	NumberOfBillingCycles string  `xml:"number-of-billing-cycles"` // int
	Price                 float64 `xml:"price"`
	TrialDuration         string  `xml:"trial-duration"` // int
	TrialDurationUnit     string  `xml:"trial-duration-unit"`
	TrialPeriod           string  `xml:"trial-period"` // bool
	CreatedAt             string  `xml:"created-at"`
	UpdatedAt             string  `xml:"updated-at"`
	// AddOns                []interface{} `xml:"add-ons"`
	// Discounts             []interface{} `xml:"discounts"`
}

// TODO(eaigner): it is suboptimal that we use string instead of int/bool types here,
// but I see no way around this atm to avoid integer parse errors if the field is empty.
//
// If there is a better method, and it can be unmarshalled directly to the correct type
// without errors, this needs to be changed.

type Plans struct {
	XMLName string  `xml:"plans"`
	Plan    []*Plan `xml:"plan"`
}
