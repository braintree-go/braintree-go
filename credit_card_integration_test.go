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

	result, err := gateway.CreditCard().Create(creditCard)

	if err != nil {
		t.Errorf(err.Error())
	} else if !result.Success() {
		t.Errorf("Credit card create response was unsuccessful")
		t.Errorf("Message: " + result.Message())
	} else if result.CreditCard().Token == "" {
		t.Errorf("Credit card did not receive an token")
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
