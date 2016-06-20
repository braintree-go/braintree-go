package braintree

import (
	"testing"

	"github.com/lionelbarrow/braintree-go/testhelpers"
)

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestCustomer(t *testing.T) {
	oc := &Customer{
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
	}

	// Create with errors
	_, err := testGateway.Customer().Create(oc)
	if err == nil {
		t.Fatal("Did not receive error when creating invalid customer")
	}

	// Create
	oc.CreditCard.CVV = ""
	oc.CreditCard.Options = nil
	customer, err := testGateway.Customer().Create(oc)

	t.Log(customer)

	if err != nil {
		t.Fatal(err)
	}
	if customer.Id == "" {
		t.Fatal("invalid customer id")
	}
	if card := customer.DefaultCreditCard(); card == nil {
		t.Fatal("invalid credit card")
	}
	if card := customer.DefaultCreditCard(); card.Token == "" {
		t.Fatal("invalid token")
	}

	// Update
	unique := testhelpers.RandomString()
	newFirstName := "John" + unique
	c2, err := testGateway.Customer().Update(&Customer{
		Id:        customer.Id,
		FirstName: newFirstName,
	})

	t.Log(c2)

	if err != nil {
		t.Fatal(err)
	}
	if c2.FirstName != newFirstName {
		t.Fatal("first name not changed")
	}

	// Find
	c3, err := testGateway.Customer().Find(customer.Id)

	t.Log(c3)

	if err != nil {
		t.Fatal(err)
	}
	if c3.Id != customer.Id {
		t.Fatal("ids do not match")
	}

	// Search
	query := new(SearchQuery)
	f := query.AddTextField("first-name")
	f.Is = newFirstName
	searchResult, err := testGateway.Customer().Search(query)
	if err != nil {
		t.Fatal(err)
	}
	if len(searchResult.Customers) == 0 {
		t.Fatal("could not search for a customer")
	}
	if id := searchResult.Customers[0].Id; id != customer.Id {
		t.Fatalf("id from search does not match: got %s, wanted %s", id, customer.Id)
	}

	// Delete
	err = testGateway.Customer().Delete(customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	// Test customer 404
	c4, err := testGateway.Customer().Find(customer.Id)
	if err == nil {
		t.Fatal("should return 404")
	}
	if err.Error() != "Not Found (404)" {
		t.Fatal(err)
	}
	if c4 != nil {
		t.Fatal(c4)
	}
}
