package braintree

import (
	"testing"
)

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
