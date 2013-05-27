package braintree

import "testing"

func TestTransactionCreate(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 100.00,
		CreditCard: &CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
		Options: &TransactionOptions{
			SubmitForSettlement: true,
		},
	}

	result, err := txGateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if !result.Success() {
		t.Errorf("Transaction create response was unsuccessful")
	} else if result.Transaction().Id == "" {
		t.Errorf("Transaction did not receive an ID")
	} else if result.Transaction().Status != "submitted_for_settlement" {
		t.Errorf("Transaction was not submitted for settlement")
	}
}

func TestTransactionCreateWhenGatewayRejected(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 2010.00,
		CreditCard: &CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}

	result, err := txGateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if result.Success() {
		t.Errorf("Transaction create response was successful, expected failure")
	} else if result.Message() != "Card Issuer Declined CVV" {
		t.Errorf("Got wrong error message. Got: " + result.Message())
	}
}

func TestFindTransaction(t *testing.T) {
	tx := Transaction{
		Type:   "sale",
		Amount: 100.00,
		CreditCard: &CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}

	saleResult, err := txGateway.Sale(tx)

	if err != nil {
		t.Errorf(err.Error())
	} else if !saleResult.Success() {
		t.Errorf("Transaction create response was unsuccessful")
	}

	txId := saleResult.Transaction().Id

	findResult, err := txGateway.Find(txId)

	if err != nil {
		t.Errorf(err.Error())
	} else if !findResult.Success() {
		t.Errorf("Transaction find response was unsuccessful")
	} else if findResult.Transaction().Id != txId {
		t.Errorf("Transaction find came back with the wrong transaction!")
	}
}

func TestFindNonExistantTransaction(t *testing.T) {
	result, err := txGateway.Find("bad transaction ID")

	if err == nil {
		t.Errorf("Did not receive an error when trying to find a non-existant transaction")
	} else if result.Success() {
		t.Errorf("Transaction find response was successful on bad data")
	} else if err.Error() != "A transaction with that ID could not be found" {
		t.Errorf("Got the wrong error message when finding a non-existant transaction. Got: " + err.Error())
	}
}

/* This test will fail unless the account under test is set up with a merchant account with
the ID "my_euro_ma", which presents and settles in Euros. */
func TestAllTransactionFields(t *testing.T) {
	sentTx := Transaction{
		Type:    "sale",
		Amount:  100.00,
		OrderId: "my_custom_order",
		CreditCard: &CreditCard{
			Number:         TestCreditCards["visa"].Number,
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
			StoreInVault:                     true,
			AddBillingAddressToPaymentMethod: true,
			StoreShippingAddressInVault:      true,
		},
	}

	result, err := txGateway.Sale(sentTx)

	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	} else if !result.Success() {
		t.Errorf("Transaction create response was unsuccessful. Message: " + result.Message())
		t.FailNow()
	}

	receivedTx := result.Transaction()

	if receivedTx.Type != sentTx.Type {
		t.Errorf("Type was wrong")
	} else if receivedTx.Amount != sentTx.Amount {
		t.Errorf("Amount was wrong")
	} else if receivedTx.OrderId != sentTx.OrderId {
		t.Errorf("OrderID was wrong")
	} else if receivedTx.Customer.FirstName != sentTx.Customer.FirstName {
		t.Errorf("Customer name was wrong")
	} else if receivedTx.BillingAddress.StreetAddress != sentTx.BillingAddress.StreetAddress {
		t.Errorf("Billing street address was wrong")
	} else if receivedTx.BillingAddress.Locality != sentTx.BillingAddress.Locality {
		t.Errorf("Billing locality was wrong")
	} else if receivedTx.BillingAddress.Region != sentTx.BillingAddress.Region {
		t.Errorf("Billing region was wrong")
	} else if receivedTx.BillingAddress.PostalCode != sentTx.BillingAddress.PostalCode {
		t.Errorf("Billing postal code was wrong")
	} else if receivedTx.ShippingAddress.StreetAddress != sentTx.ShippingAddress.StreetAddress {
		t.Errorf("Shipping street address was wrong")
	} else if receivedTx.ShippingAddress.Locality != sentTx.ShippingAddress.Locality {
		t.Errorf("Shipping locality was wrong")
	} else if receivedTx.ShippingAddress.Region != sentTx.ShippingAddress.Region {
		t.Errorf("Shipping region was wrong")
	} else if receivedTx.ShippingAddress.PostalCode != sentTx.ShippingAddress.PostalCode {
		t.Errorf("Shipping postal code was wrong")
	}

	if receivedTx.CreditCard.Token == "" {
		t.Errorf("Should have received a token when credit card is stored in vault")
	} else if receivedTx.Customer.Id == "" {
		t.Errorf("Should have received a customer ID when providing customer details on a transaction")
	}
}
