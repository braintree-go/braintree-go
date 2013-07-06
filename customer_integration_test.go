package braintree

import (
	"testing"
)

func TestCustomerCreateAndUpdate(t *testing.T) {
	customer, err := testGateway.Customer().Create(&Customer{
		FirstName: "Lionel",
		LastName:  "Barrow",
		Company:   "Braintree",
		Email:     "lionel.barrow@example.com",
		Phone:     "312.555.1234",
		Fax:       "614.555.5678",
		Website:   "http://www.example.com",
		CreditCard: &CreditCard{
			Number:         testCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	})

	t.Log(customer)

	if err != nil {
		t.Fatal(err)
	}
	if customer.Id == "" {
		t.Fatal("invalid customer id")
	}

	c2, err := testGateway.Customer().Update(&Customer{
		Id:        customer.Id,
		FirstName: "John",
	})

	t.Log(c2)

	if err != nil {
		t.Fatal(err)
	}
	if c2.FirstName != "John" {
		t.Fatal("first name not changed")
	}
}

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestCustomerCreateWithErrors(t *testing.T) {
	_, err := testGateway.Customer().Create(&Customer{
		FirstName: "Lionel",
		LastName:  "Barrow",
		Company:   "Braintree",
		Email:     "lionel.barrow@example.com",
		Phone:     "312.555.1234",
		Fax:       "614.555.5678",
		Website:   "http://www.example.com",
		CreditCard: &CreditCard{
			Number:         testCreditCards["visa"].Number,
			ExpirationDate: "05/14",
			CVV:            "200",
			Options: &CreditCardOptions{
				VerifyCard: true,
			},
		},
	})

	if err == nil {
		t.Fatal("Did not receive error when creating invalid customer")
	}
}

func TestCustomerFind(t *testing.T) {
	customer, err := testGateway.Customer().Create(&Customer{
		FirstName: "Lionel",
		LastName:  "Barrow",
	})

	t.Log(customer)

	if err != nil {
		t.Fatal(err)
	}

	customer2, err := testGateway.Customer().Find(customer.Id)

	t.Log(customer2)

	if err != nil {
		t.Fatal(err)
	}
	if customer2.Id != customer.Id {
		t.Fatal("ids do not match")
	}
}

func TestFindNonExistantCustomer(t *testing.T) {
	customer, err := testGateway.Customer().Find("bad_customer_id")

	t.Log(customer)

	if err == nil {
		t.Fatal("should report error, invalid customer id")
	}
	if err.Error() != "Not Found (404)" {
		t.Fatal(err)
	}
}
