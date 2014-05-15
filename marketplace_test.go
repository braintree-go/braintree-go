package braintree

import (
	"testing"
)

func TestTransactionHeldInEscrow(t *testing.T) {
	tx := &Transaction{
		Type:             "sale",
		Amount:           100.00 + offset(),
		ServiceFeeAmount: 10.00,
		OrderId:          "my_escrow_order",
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
		ShippingAddress: &Address{
			StreetAddress: "1 E Main St",
			Locality:      "Chicago",
			Region:        "IL",
			PostalCode:    "60637",
		},
		Options: &TransactionOptions{
			SubmitForSettlement:              true,
			StoreInVault:                     false,
			AddBillingAddressToPaymentMethod: true,
			StoreShippingAddressInVault:      false,
			HoldInEscrow:                     true,
		},
		MerchantAccountId: testSubMerchantAccountId,
	}

	tx2, err := testGateway.Transaction().Create(tx)
	if err != nil {
		t.Fatal(err)
	}

	if tx2.Status != "submitted_for_settlement" {
		t.Fail()
	}

	if tx2.EscrowStatus != "hold_pending" {
		t.Fail()
	}
}
