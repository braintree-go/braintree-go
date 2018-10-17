// +build integration

package braintree

import (
	"context"
	"encoding/xml"
	"testing"

	"github.com/braintree-go/braintree-go/testhelpers"
)

func TestMerchantAccountCreate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	acct := MerchantAccount{
		MasterMerchantAccountId: testMerchantAccountId,
		TOSAccepted:             true,
		Id:                      testhelpers.RandomString(),
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

	merchantAccount, err := testGateway.MerchantAccount().Create(ctx, &acct)

	t.Log(merchantAccount)

	if err != nil {
		t.Fatal(err)
	}

	if merchantAccount.Id == "" {
		t.Fatal("invalid merchant account id")
	}

	ma2, err := testGateway.MerchantAccount().Find(ctx, merchantAccount.Id)

	t.Log(ma2)

	if err != nil {
		t.Fatal(err)
	}

	if ma2.Id != merchantAccount.Id {
		t.Fatal("ids do not match")
	}

}
