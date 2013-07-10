package braintree

import (
	"testing"
)

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
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

	g := testGateway.Subscription()

	token := customer.CreditCards.CreditCard[0].Token
	if token == "" {
		t.Fatal("invalid payment method token")
	} else {
		t.Log(token)
	}

	// Create
	sub, err := g.Create(&Subscription{
		PaymentMethodToken: token,
		PlanId:             "test_plan",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	})

	t.Log("sub1", sub)

	if err != nil {
		t.Fatal(err)
	}
	if sub.Id == "" {
		t.Fatal("invalid subscription id")
	}

	// Update
	sub2, err := g.Update(&Subscription{
		Id:     sub.Id,
		PlanId: "test_plan_2",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	})

	t.Log("sub2", sub2)

	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != sub.Id {
		t.Fatal(sub2.Id)
	}
	if x := sub2.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}

	// Find
	sub3, err := g.Find(sub.Id)
	if err != nil {
		t.Fatal(err)
	}
	if sub3.Id != sub2.Id {
		t.Fatal(sub3.Id)
	}

	// Cancel
	_, err = g.Cancel(sub2.Id)
	if err != nil {
		t.Fatal(err)
	}
}
