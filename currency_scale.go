package braintree

var (
	// DefaultCurrencyScales - default currency decimal scales
	DefaultCurrencyScales uint = 2

	// CurrencyScales are the decimal scales of various currencies
	// currencies with 2 decimal scales stripped out
	CurrencyScales = map[string]uint{
		"BHD": 3,
		"CVE": 0,
		"DJF": 0,
		"GNF": 0,
		"IDR": 0,
		"JOD": 3,
		"JPY": 0,
		"KMF": 0,
		"KRW": 0,
		"KWD": 3,
		"LYD": 3,
		"OMR": 3,
		"PYG": 0,
		"RWF": 0,
		"TND": 3,
		"UGX": 0,
		"VND": 0,
		"VUV": 0,
		"XAF": 0,
		"XOF": 0,
		"XPF": 0,
	}
)
