package braintree

type Plan struct {
	XMLName               string     `xml:"plan"`
	Id                    string     `xml:"id"`
	MerchantId            string     `xml:"merchant-id"`
	BillingDayOfMonth     *NullInt64 `xml:"billing-day-of-month"`
	BillingFrequency      *NullInt64 `xml:"billing-frequency"`
	CurrencyISOCode       string     `xml:"currency-iso-code"`
	Description           string     `xml:"description"`
	Name                  string     `xml:"name"`
	NumberOfBillingCycles *NullInt64 `xml:"number-of-billing-cycles"`
	Price                 float64    `xml:"price"`
	TrialDuration         *NullInt64 `xml:"trial-duration"`
	TrialDurationUnit     string     `xml:"trial-duration-unit"`
	TrialPeriod           *NullBool  `xml:"trial-period"`
	CreatedAt             string     `xml:"created-at"`
	UpdatedAt             string     `xml:"updated-at"`
	// AddOns                []interface{} `xml:"add-ons"`
	// Discounts             []interface{} `xml:"discounts"`
}

type Plans struct {
	XMLName string  `xml:"plans"`
	Plan    []*Plan `xml:"plan"`
}
