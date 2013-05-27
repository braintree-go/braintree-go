package braintree

import "testing"

func TestCreateCreditCard(t *testing.T) {
	customer := Customer{}

	customerResult, err := gateway.Customer().Create(customer)

	if err != nil {
		t.Errorf(err.Error())
	} else if !customerResult.Success() {
		t.Errorf("Customer create response was unsuccessful")
	}

	creditCard := CreditCard{
		CustomerId:     customerResult.Customer().Id,
		Number:         TestCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: true,
		},
	}

	createResult, err := gateway.CreditCard().Create(creditCard)

	if err != nil {
		t.Errorf(err.Error())
	} else if !createResult.Success() {
		t.Errorf("Credit card create response was unsuccessful")
	}
}

/* This will fail because a customer ID is required. */
func TestCreateCreditCardInvalidInput(t *testing.T) {
	creditCard := CreditCard{
		Number:         TestCreditCards["visa"].Number,
		ExpirationDate: "05/14",
	}

	result, err := gateway.CreditCard().Create(creditCard)

	if err != nil {
		t.Errorf(err.Error())
	} else if result.Success() {
		t.Errorf("Invaid credit card create returned a successful response")
	}
}

func TestFindCreditCard(t *testing.T) {
	customer := Customer{}

	customerResult, err := gateway.Customer().Create(customer)

	if err != nil {
		t.Errorf(err.Error())
	} else if !customerResult.Success() {
		t.Errorf("Customer create response was unsuccessful")
	}

	creditCard := CreditCard{
		CustomerId:     customerResult.Customer().Id,
		Number:         TestCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: true,
		},
	}

	createResult, err := gateway.CreditCard().Create(creditCard)

	if err != nil {
		t.Errorf(err.Error())
	} else if !createResult.Success() {
		t.Errorf("Credit card create response was unsuccessful")
	} else if createResult.CreditCard().Token == "" {
		t.Errorf("Credit card did not receive an token")
	}

	findResult, err := gateway.CreditCard().Find(createResult.CreditCard().Token)

	if err != nil {
		t.Errorf(err.Error())
	} else if !findResult.Success() {
		t.Errorf("Credit card find response was unsuccessful")
	} else if findResult.CreditCard().Token != createResult.CreditCard().Token {
		t.Errorf("Credit card find returned the wrong card!")
	}
}

func TestFindCreditCardBadData(t *testing.T) {
	result, err := gateway.CreditCard().Find("invalid_token")

	if err == nil {
		t.Errorf("No error was returned when finding an invalid credit card")
	} else if result.Success() {
		t.Errorf("Finding an invalid credit card returned a successful result")
	}
}
