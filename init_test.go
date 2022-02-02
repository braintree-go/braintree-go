// +build unit integration

package braintree

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BoltApp/braintree-go/testhelpers"
)

const (
	testCardVisa       = "4111111111111111"
	testCardMastercard = "5555555555554444"
	testCardDiscover   = "6011111111111117"
)

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

func testSubMerchantAccount() string {
	acct := MerchantAccount{
		MasterMerchantAccountId: testMerchantAccountId,
		TOSAccepted:             true,
		Id:                      testhelpers.RandomString(),
		Individual: &MerchantAccountPerson{
			FirstName:   "First",
			LastName:    "Last",
			Email:       "firstlast@example.com",
			Phone:       "0000000000",
			DateOfBirth: "1-1-1900",
			Address: &Address{
				StreetAddress:   "222 W Merchandise Mart Plaza",
				ExtendedAddress: "Suite 800",
				Locality:        "Chicago",
				Region:          "IL",
				PostalCode:      "00000",
			},
		},
		FundingOptions: &MerchantAccountFundingOptions{
			Destination: FUNDING_DEST_MOBILE_PHONE,
			MobilePhone: "0000000000",
		},
	}

	merchantAccount, err := testGateway.MerchantAccount().Create(context.Background(), &acct)
	if err != nil {
		panic(fmt.Errorf("Error creating test sub merchant account: %s", err))
	}

	return merchantAccount.Id
}

func init() {
	logEnabled := flag.Bool("log", false, "enables logging")
	flag.Parse()

	if *logEnabled {
		testGateway.Logger = log.New(os.Stderr, "", 0)
	}
}
