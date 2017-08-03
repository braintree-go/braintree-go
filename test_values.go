package braintree

import (
	"fmt"
	"os"
	"time"
)

var testCreditCards = map[string]CreditCard{
	"visa":       CreditCard{Number: "4111111111111111"},
	"mastercard": CreditCard{Number: "5555555555554444"},
	"discover":   CreditCard{Number: "6011111111111117"},
	// See https://developers.braintreepayments.com/reference/general/testing/python#3d-secure
	"visa_3ds_success":       CreditCard{Number: "4000000000000002"},
	"mastercard_3ds_success": CreditCard{Number: "5200000000000007"},
	"visa_3ds_fail":          CreditCard{Number: "4000000000000028"},
	"mastercard_3ds_fail":    CreditCard{Number: "5200000000000023"},
}

var testGateway = New(
	Sandbox,
	os.Getenv("BRAINTREE_MERCH_ID"),
	os.Getenv("BRAINTREE_PUB_KEY"),
	os.Getenv("BRAINTREE_PRIV_KEY"),
)

var testTimeZone = func() *time.Location {
	tzName := os.Getenv("BRAINTREE_TIMEZONE")
	if tzName == "" {
		return time.UTC
	}
	tz, err := time.LoadLocation(tzName)
	if err != nil {
		panic(fmt.Errorf("Error loading time zone location %s: %s", tzName, err))
	}
	return tz
}()

var testMerchantAccountId = os.Getenv("BRAINTREE_MERCH_ACCT_ID")

// Merchant Account which has AVS and CVV checking turned on.
var avsAndCVVTestMerchantAccountId = os.Getenv("BRAINTREE_MERCH_ACCT_ID_FOR_AVS_CVV")
