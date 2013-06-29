package braintree

import (
	"testing"
)

func TestSubscription(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}

	t.Log(customer)

	token := customer.CreditCards.CreditCard[0].Token
	if token == "" {
		t.Fatal("invalid payment method token")
	} else {
		t.Log(token)
	}

	sub, err := testGateway.Subscription().Create(&Subscription{
		PaymentMethodToken: token,
		PlanId:             "test_plan",
	})
	if err != nil {
		t.Fatal(err)
	}
	if sub.Id == "" {
		t.Fatal("invalid subscription id")
	}

	sub2, err := testGateway.Subscription().Find(sub.Id)
	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != sub.Id {
		t.Fatal(sub2.Id)
	}

	_, err = testGateway.Subscription().Cancel(sub.Id)
	if err != nil {
		t.Fatal(err)
	}
}
