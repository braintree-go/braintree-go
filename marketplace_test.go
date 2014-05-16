package braintree

import (
	"encoding/xml"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

var mmId string
var txId string

func TestMarketplaceMerchantCreate(t *testing.T) {
	mmId = fmt.Sprintf("TMM%d", time.Now().Unix())

	acct := MerchantAccount{
		MasterMerchantAccountId: testMerchantAccountId,
		TOSAccepted:             true,
		Id:                      mmId,
		Individual: &MerchantAccountPerson{
			FirstName:   braintree.MerchantAccountApprove,
			LastName:    "Lastname",
			Email:       "merchant@example.com",
			Phone:       "5558675309",
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
			MobilePhone: "5558675309",
		},
	}

	x, _ := xml.Marshal(&acct)
	t.Log(string(x))

	merchantAccount, err := testGateway.MerchantAccount().Create(&acct)

	t.Log(merchantAccount)

	if err != nil {
		t.Fatal(err)
	}

	if merchantAccount.Status != braintree.MerchantAccountStatusPending {
		t.Fatal(merchantAccount.Status)
	}
}

func TestMarketplaceMerchantActive(t *testing.T) {
	ma2, err := testGateway.MerchantAccount().Find(mmId)

	if err != nil {
		t.Fatal(err)
	}

	if ma2.Status != braintree.MerchantAccountStatusActive {
		t.Fatal("not active yet")
	}
}

func TestMarketplaceMerchantTx(t *testing.T) {
	tx := &Transaction{
		Type:              braintree.TxSale,
		MerchantAccountId: mmId,
		Amount:            100.00 + math.Ceil(rand.Float64()*100.0),
		ServiceFeeAmount:  10.00,
		OrderId:           "test_escrow_order",
		CreditCard: &CreditCard{
			Number:         testCreditCards["visa"].Number,
			ExpirationDate: "05/14",
			CVV:            "100",
		},
		Customer: &Customer{
			FirstName: "Lionel",
		},
		BillingAddress: &Address{
			StreetAddress: "1 E Main St",
			Locality:      "Chicago",
			Region:        "IL",
			PostalCode:    "60637",
		},
		Options: &TransactionOptions{
			SubmitForSettlement: true,
			HoldInEscrow:        true,
		},
	}

	tx2, err := testGateway.Transaction().Create(tx)
	if err != nil {
		t.Fatal(err)
	}

	if tx2.EscrowStatus != braintree.TxEscrowHoldPending {
		t.Fatal(tx2.EscrowStatus)
	}

	txId = tx2.Id
	t.Log(txId)
}

func TestMarketplaceMerchantRelease(t *testing.T) {
	tx, err := testGateway.Transaction().Find(txId)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(txId)
	t.Log(tx.Status)
	t.Log(tx.EscrowStatus)

}
