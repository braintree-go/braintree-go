// +build integration

package braintree

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"fmt"
	"strings"

	"github.com/braintree-go/braintree-go/testhelpers"
)

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestCustomer(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	oc := &CustomerRequest{
		FirstName: "Lionel",
		LastName:  "Barrow",
		Company:   "Braintree",
		Email:     "lionel.barrow@example.com",
		Phone:     "312.555.1234",
		Fax:       "614.555.5678",
		Website:   "http://www.example.com",
		CreditCard: &CreditCard{
			Number:         testCardVisa,
			ExpirationDate: "05/14",
			CVV:            "200",
			Options: &CreditCardOptions{
				VerifyCard: testhelpers.BoolPtr(true),
			},
		},
	}

	// Create with errors
	_, err := testGateway.Customer().Create(ctx, oc)
	if err == nil {
		t.Fatal("Did not receive error when creating invalid customer")
	}

	// Create
	oc.CreditCard.CVV = ""
	oc.CreditCard.Options = nil
	customer, err := testGateway.Customer().Create(ctx, oc)

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
	if customer.CreatedAt == nil {
		t.Fatal("generated created at is empty")
	}
	if customer.UpdatedAt == nil {
		t.Fatal("generated updated at is empty")
	}

	// Update
	unique := testhelpers.RandomString()
	newFirstName := "John" + unique
	c2, err := testGateway.Customer().Update(ctx, &CustomerRequest{
		ID:        customer.Id,
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
	c3, err := testGateway.Customer().Find(ctx, customer.Id)

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
	searchResult, err := testGateway.Customer().Search(ctx, query)
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
	err = testGateway.Customer().Delete(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	// Test customer 404
	c4, err := testGateway.Customer().Find(ctx, customer.Id)
	if err == nil {
		t.Fatal("should return 404")
	}
	if err.Error() != "Not Found (404)" {
		t.Fatal(err)
	}
	if apiErr, ok := err.(APIError); !(ok && apiErr.StatusCode() == http.StatusNotFound) {
		t.Fatal(err)
	}
	if c4 != nil {
		t.Fatal(c4)
	}
}

func TestCustomerSearch(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cg := testGateway.Customer()

	const customersCount = 51
	customerIDs := map[string]bool{}
	prefix := "PaginationTest-" + testhelpers.RandomString()
	for i := 0; i < customersCount; i++ {
		unique := testhelpers.RandomString()
		tx, err := cg.Create(ctx, &CustomerRequest{
			FirstName: "John",
			LastName:  "Smith",
			Company:   prefix + unique,
		})
		if err != nil {
			t.Fatal(err)
		}
		customerIDs[tx.Id] = true
	}

	t.Logf("customerIDs = %v", customerIDs)

	query := new(SearchQuery)
	query.AddTextField("company").StartsWith = prefix

	results, err := cg.SearchIDs(ctx, query)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("results.PageSize = %v", results.PageSize)
	t.Logf("results.PageCount = %v", results.PageCount)
	t.Logf("results.IDs = %d %v", len(results.IDs), results.IDs)

	if len(results.IDs) != customersCount {
		t.Fatalf("results.IDs = %v, want %v", len(results.IDs), customersCount)
	}

	for page := 1; page <= results.PageCount; page++ {
		results, err := cg.SearchPage(ctx, query, results, page)
		if err != nil {
			t.Fatal(err)
		}
		for _, cs := range results.Customers {
			if company := cs.Company; !strings.HasPrefix(company, prefix) {
				t.Fatalf("cs.Company = %q, want prefix of %q", company, prefix)
			}
			if customerIDs[cs.Id] {
				delete(customerIDs, cs.Id)
			} else {
				t.Fatalf("cs.Id = %q, not expected", cs.Id)
			}
		}
	}

	if len(customerIDs) > 0 {
		t.Fatalf("customers not returned = %v", customerIDs)
	}

	_, err = cg.SearchPage(ctx, query, results, 0)
	t.Logf("%#v", err)
	if err == nil || !strings.Contains(err.Error(), "page 0 out of bounds") {
		t.Errorf("requesting page 0 should result in out of bounds error, but got %#v", err)
	}

	_, err = cg.SearchPage(ctx, query, results, results.PageCount+1)
	t.Logf("%#v", err)
	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf("page %d out of bounds", results.PageCount+1)) {
		t.Errorf("requesting page %d should result in out of bounds error, but got %v", results.PageCount+1, err)
	}

}

func TestCustomerWithCustomFields(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customFields := map[string]string{
		"custom_field_1": "custom value",
	}

	c := &CustomerRequest{
		CustomFields: customFields,
	}

	customer, err := testGateway.Customer().Create(ctx, c)
	if err != nil {
		t.Fatal(err)
	}

	if x := map[string]string(customer.CustomFields); !reflect.DeepEqual(x, customFields) {
		t.Fatalf("Returned custom fields doesn't match input, got %q, want %q", x, customFields)
	}

	customer, err = testGateway.Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if x := map[string]string(customer.CustomFields); !reflect.DeepEqual(x, customFields) {
		t.Fatalf("Returned custom fields doesn't match input, got %q, want %q", x, customFields)
	}
}

func TestCustomerPaymentMethods(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	paymentMethod1, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNoncePayPalBillingAgreement,
	})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod2, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	expectedPaymentMethods := []PaymentMethod{
		paymentMethod2,
		paymentMethod1,
	}

	customerFound, err := testGateway.Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(customerFound.PaymentMethods(), expectedPaymentMethods) {
		t.Fatalf("Got Customer %#v PaymentMethods %#v, want %#v", customerFound, customerFound.PaymentMethods(), expectedPaymentMethods)
	}
}

func TestCustomerDefaultPaymentMethod(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	defaultPaymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNoncePayPalBillingAgreement,
	})
	if err != nil {
		t.Fatal(err)
	}

	customerFound, err := testGateway.Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(customerFound.DefaultPaymentMethod(), defaultPaymentMethod) {
		t.Fatalf("Got Customer %#v DefaultPaymentMethod %#v, want %#v", customerFound, customerFound.DefaultPaymentMethod(), defaultPaymentMethod)
	}
}

func TestCustomerDefaultPaymentMethodManuallySet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	_, err = testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod2, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNoncePayPalBillingAgreement,
	})
	if err != nil {
		t.Fatal(err)
	}
	paypalAccount, err := testGateway.PayPalAccount().Update(ctx, &PayPalAccount{
		Token: paymentMethod2.GetToken(),
		Options: &PayPalAccountOptions{
			MakeDefault: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	customerFound, err := testGateway.Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(customerFound.DefaultPaymentMethod(), paypalAccount) {
		t.Fatalf("Got Customer %#v DefaultPaymentMethod %#v, want %#v", customerFound, customerFound.DefaultPaymentMethod(), paypalAccount)
	}
}

func TestCustomerPaymentMethodNonce(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{PaymentMethodNonce: FakeNonceTransactable})
	if err != nil {
		t.Fatal(err)
	}

	customerFound, err := testGateway.Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if len(customer.PaymentMethods()) != 1 {
		t.Fatalf("Customer %#v has %#v payment method(s), want 1 payment method", customerFound, len(customer.PaymentMethods()))
	}
}

func TestCustomerAddresses(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{
		FirstName: "Jenna",
		LastName:  "Smith",
	})
	if err != nil {
		t.Fatal(err)
	}
	if customer.Id == "" {
		t.Fatal("invalid customer id")
	}

	addrReqs := []*AddressRequest{
		&AddressRequest{
			FirstName:          "Jenna",
			LastName:           "Smith",
			Company:            "Braintree",
			StreetAddress:      "1 E Main St",
			ExtendedAddress:    "Suite 403",
			Locality:           "Chicago",
			Region:             "Illinois",
			PostalCode:         "60622",
			CountryCodeAlpha2:  "US",
			CountryCodeAlpha3:  "USA",
			CountryCodeNumeric: "840",
			CountryName:        "United States of America",
		},
		&AddressRequest{
			FirstName:          "Bob",
			LastName:           "Rob",
			Company:            "Paypal",
			StreetAddress:      "1 W Main St",
			ExtendedAddress:    "Suite 402",
			Locality:           "Boston",
			Region:             "Massachusetts",
			PostalCode:         "02140",
			CountryCodeAlpha2:  "US",
			CountryCodeAlpha3:  "USA",
			CountryCodeNumeric: "840",
			CountryName:        "United States of America",
		},
	}

	for _, addrReq := range addrReqs {
		_, err = testGateway.Address().Create(ctx, customer.Id, addrReq)
		if err != nil {
			t.Fatal(err)
		}
	}

	customerWithAddrs, err := testGateway.Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if customerWithAddrs.Addresses == nil || len(customerWithAddrs.Addresses.Address) != 2 {
		t.Fatal("wrong number of addresses returned")
	}

	for _, addr := range customerWithAddrs.Addresses.Address {
		if addr.Id == "" {
			t.Fatal("generated id is empty")
		}

		var addrReq *AddressRequest
		for _, ar := range addrReqs {
			if ar.PostalCode == addr.PostalCode {
				addrReq = ar
				break
			}
		}

		if addrReq == nil {
			t.Fatal("did not return sent address")
		}

		t.Logf("%+v\n", addr)
		t.Logf("%+v\n", addrReq)

		if addr.CustomerId != customer.Id {
			t.Errorf("got customer id %s, want %s", addr.CustomerId, customer.Id)
		}
		if addr.FirstName != addrReq.FirstName {
			t.Errorf("got first name %s, want %s", addr.FirstName, addrReq.FirstName)
		}
		if addr.LastName != addrReq.LastName {
			t.Errorf("got last name %s, want %s", addr.LastName, addrReq.LastName)
		}
		if addr.Company != addrReq.Company {
			t.Errorf("got company %s, want %s", addr.Company, addrReq.Company)
		}
		if addr.StreetAddress != addrReq.StreetAddress {
			t.Errorf("got street address %s, want %s", addr.StreetAddress, addrReq.StreetAddress)
		}
		if addr.ExtendedAddress != addrReq.ExtendedAddress {
			t.Errorf("got extended address %s, want %s", addr.ExtendedAddress, addrReq.ExtendedAddress)
		}
		if addr.Locality != addrReq.Locality {
			t.Errorf("got locality %s, want %s", addr.Locality, addrReq.Locality)
		}
		if addr.Region != addrReq.Region {
			t.Errorf("got region %s, want %s", addr.Region, addrReq.Region)
		}
		if addr.CountryCodeAlpha2 != addrReq.CountryCodeAlpha2 {
			t.Errorf("got country code alpha 2 %s, want %s", addr.CountryCodeAlpha2, addrReq.CountryCodeAlpha2)
		}
		if addr.CountryCodeAlpha3 != addrReq.CountryCodeAlpha3 {
			t.Errorf("got country code alpha 3 %s, want %s", addr.CountryCodeAlpha3, addrReq.CountryCodeAlpha3)
		}
		if addr.CountryCodeNumeric != addrReq.CountryCodeNumeric {
			t.Errorf("got country code numeric %s, want %s", addr.CountryCodeNumeric, addrReq.CountryCodeNumeric)
		}
		if addr.CountryName != addrReq.CountryName {
			t.Errorf("got country name %s, want %s", addr.CountryName, addrReq.CountryName)
		}
		if addr.CreatedAt == nil {
			t.Error("got created at nil, want a value")
		}
		if addr.UpdatedAt == nil {
			t.Error("got updated at nil, want a value")
		}
	}

	err = testGateway.Customer().Delete(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}
}
