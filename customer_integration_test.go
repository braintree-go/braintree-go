package braintree

import "testing"

func TestCustomerCreate(t *testing.T) {
	customer := Customer{
		FirstName: "Lionel",
		LastName:  "Barrow",
		Company:   "Braintree",
		Email:     "lionel.barrow@example.com",
		Phone:     "312.555.1234",
		Fax:       "614.555.5678",
		Website:   "http://www.example.com",
		CreditCard: &CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}

	result, err := customerGateway.Create(customer)

	if err != nil {
		t.Errorf(err.Error())
	} else if !result.Success() {
		t.Errorf("Customer create response was unsuccessful")
	} else if result.Customer().Id == "" {
		t.Errorf("Customer did not receive an ID")
	}
}

/* This test will fail unless the account under test has CVV rules set up in the
Braintree gateway. See https://www.braintreepayments.com/docs/ruby/card_verifications/overview 
for more details. */
func TestCustomerCreateWithErrors(t *testing.T) {
	customer := Customer{
		FirstName: "Lionel",
		LastName:  "Barrow",
		Company:   "Braintree",
		Email:     "lionel.barrow@example.com",
		Phone:     "312.555.1234",
		Fax:       "614.555.5678",
		Website:   "http://www.example.com",
		CreditCard: &CreditCard{
			Number:         TestCreditCards["visa"].Number,
			ExpirationDate: "05/14",
			CVV:            "200",
			Options: &CreditCardOptions{
				VerifyCard: true,
			},
		},
	}

	result, err := customerGateway.Create(customer)

	if err != nil {
		t.Errorf(err.Error())
	} else if result.Success() {
		t.Errorf("Customer created succeeded with bad CVV and verify card true")
	}
}

func TestCustomerFind(t *testing.T) {
	sentCustomer := Customer{
		FirstName: "Lionel",
		LastName:  "Barrow",
	}

	createResult, err := customerGateway.Create(sentCustomer)

	if err != nil {
		t.Errorf(err.Error())
	} else if !createResult.Success() {
		t.Errorf("Customer create response was unsuccessful")
	} else if createResult.Customer().Id == "" {
		t.Errorf("Customer did not receive an ID")
	}

	findResult, err := customerGateway.Find(createResult.Customer().Id)

	if err != nil {
		t.Errorf(err.Error())
	} else if !findResult.Success() {
		t.Errorf("Could not find the customer we just created")
	} else if findResult.Customer().Id == "" {
		t.Errorf("The old customer's ID and the new one did not match")
	}
}

func TestFindNonExistantCustomer(t *testing.T) {
	result, err := customerGateway.Find("bad_customer_id")

	if err == nil {
		t.Errorf("Did not receive an error when trying to find a non-existant customer")
	} else if result.Success() {
		t.Errorf("Customer find response was successful on bad data")
	} else if err.Error() != "A customer with that ID could not be found" {
		t.Errorf("Got the wrong error message when finding a non-existant customer.")
	}
}
