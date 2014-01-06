package braintree

import (
	"encoding/xml"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestMerchantAccountCreate(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	acct := MerchantAccount{
		MasterMerchantAccountId: os.Getenv("BRAINTREE_MERCH_ACCT_ID"),
		TOSAccepted:             true,
		Id:                      strconv.Itoa(rand.Int()),
		Individual: &MerchantAccountPerson{
			FirstName:   "Kayle",
			LastName:    "Gishen",
			Email:       "kayle.gishen@example.com",
			Phone:       "5556789012",
			DateOfBirth: "1-1-1989",
			Address: &Address{
				StreetAddress:   "1 E Main St",
				ExtendedAddress: "Suite 404",
				Locality:        "Chicago",
				Region:          "IL",
				PostalCode:      "60622",
			},
		},
		FundingOptions: &MerchantAccountFundingOptions{
			Destination: FUNDING_DEST_MOBILE_PHONE,
			MobilePhone: "5552344567",
		},
	}

	x, _ := xml.Marshal(&acct)
	t.Log(string(x))

	merchantAccount, err := testGateway.MerchantAccount().Create(&acct)

	t.Log(merchantAccount)

	if err != nil {
		t.Fatal(err)
	}

	if merchantAccount.Id == "" {
		t.Fatal("invalid merchant account id")
	}

	ma2, err := testGateway.MerchantAccount().Find(merchantAccount.Id)

	t.Log(ma2)

	if err != nil {
		t.Fatal(err)
	}

	if ma2.Id != merchantAccount.Id {
		t.Fatal("ids do not match")
	}

}
