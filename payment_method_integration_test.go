// +build integration

package braintree

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/hellofresh/braintree-go/testhelpers"
)

func TestPaymentMethod(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cust, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.PaymentMethod()

	// Create using credit card
	addr := &AddressRequest{
		FirstName:          "Robert",
		LastName:           "Smith",
		Company:            "The Cure",
		StreetAddress:      "39 Acacia Avenue",
		ExtendedAddress:    "SAV Studios",
		Locality:           "North End",
		Region:             "London",
		PostalCode:         "SW1A 0AA",
		CountryCodeAlpha2:  "GB",
		CountryCodeAlpha3:  "GBR",
		CountryCodeNumeric: "826",
		CountryName:        "United Kingdom",
	}

	paymentMethod, err := g.Create(ctx, &PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: FakeNonceTransactableVisa,
		BillingAddress:     addr,
	})
	if err != nil {
		t.Fatal(err)
	}

	if paymentMethod.GetCustomerId() != cust.Id {
		t.Errorf("Got paymentMethod customer Id %#v, want %#v", paymentMethod.GetCustomerId(), cust.Id)
	}
	if paymentMethod.GetToken() == "" {
		t.Errorf("Got paymentMethod token %#v, want a value", paymentMethod.GetToken())
	}

	if card, ok := paymentMethod.(*CreditCard); ok {
		ba := card.BillingAddress
		if ba.FirstName != addr.FirstName {
			t.Errorf("Got paymentMethod billing adress first name %#v, want %#v", ba.FirstName, addr.FirstName)
		}
		if ba.LastName != addr.LastName {
			t.Errorf("Got paymentMethod billing adress last name %#v, want %#v", ba.LastName, addr.LastName)
		}
		if ba.Company != addr.Company {
			t.Errorf("Got paymentMethod billing adress company %#v, want %#v", ba.Company, addr.Company)
		}
		if ba.StreetAddress != addr.StreetAddress {
			t.Errorf("Got paymentMethod billing adress street address %#v, want %#v", ba.StreetAddress, addr.StreetAddress)
		}
		if ba.ExtendedAddress != addr.ExtendedAddress {
			t.Errorf("Got paymentMethod billing adress extended address %#v, want %#v", ba.ExtendedAddress, addr.ExtendedAddress)
		}
		if ba.Locality != addr.Locality {
			t.Errorf("Got paymentMethod billing adress locality %#v, want %#v", ba.Locality, addr.Locality)
		}
		if ba.Region != addr.Region {
			t.Errorf("Got paymentMethod billing adress region %#v, want %#v", ba.Region, addr.Region)
		}
		if ba.PostalCode != addr.PostalCode {
			t.Errorf("Got paymentMethod billing adress postal code %#v, want %#v", ba.PostalCode, addr.PostalCode)
		}
		if ba.CountryCodeAlpha2 != addr.CountryCodeAlpha2 {
			t.Errorf("Got paymentMethod billing adress country alpha2 %#v, want %#v", ba.CountryCodeAlpha2, addr.CountryCodeAlpha2)
		}
		if ba.CountryCodeAlpha3 != addr.CountryCodeAlpha3 {
			t.Errorf("Got paymentMethod billing adress country alpha3 %#v, want %#v", ba.CountryCodeAlpha3, addr.CountryCodeAlpha3)
		}
		if ba.CountryCodeNumeric != addr.CountryCodeNumeric {
			t.Errorf("Got paymentMethod billing adress country numeric %#v, want %#v", ba.CountryCodeNumeric, addr.CountryCodeNumeric)
		}
		if ba.CountryName != addr.CountryName {
			t.Errorf("Got paymentMethod billing adress country name %#v, want %#v", ba.CountryName, addr.CountryName)
		}
	} else {
		t.Error("paymentMethod should have been a credit card")
	}

	// Update using different credit card
	rand.Seed(time.Now().UTC().UnixNano())
	token := fmt.Sprintf("btgo_test_token_%d", rand.Int()+1)
	paymentMethod, err = g.Update(ctx, paymentMethod.GetToken(), &PaymentMethodRequest{
		PaymentMethodNonce: FakeNonceTransactableMasterCard,
		Token:              token,
	})
	if err != nil {
		t.Fatal(err)
	}

	if paymentMethod.GetToken() != token {
		t.Errorf("Got paymentMethod token %#v, want %#v", paymentMethod.GetToken(), token)
	}

	// Updating with different payment method type should fail
	if _, err = g.Update(ctx, token, &PaymentMethodRequest{PaymentMethodNonce: FakeNoncePayPalBillingAgreement}); err == nil {
		t.Errorf("Updating with a different payment method type should have failed")
	}

	// Find credit card
	paymentMethod, err = g.Find(ctx, token)
	if err != nil {
		t.Fatal(err)
	}

	if paymentMethod.GetCustomerId() != cust.Id {
		t.Errorf("Got paymentMethod customer Id %#v, want %#v", paymentMethod.GetCustomerId(), cust.Id)
	}
	if paymentMethod.GetToken() != token {
		t.Errorf("Got paymentMethod token %#v, want %#v", paymentMethod.GetToken(), token)
	}

	// Delete credit card
	if err := g.Delete(ctx, token); err != nil {
		t.Fatal(err)
	}

	// Create using PayPal
	paymentMethod, err = g.Create(ctx, &PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: FakeNoncePayPalBillingAgreement,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Find PayPal
	_, err = g.Find(ctx, paymentMethod.GetToken())
	if err != nil {
		t.Fatal(err)
	}

	// Updating a PayPal account with a different payment method nonce of any kind should fail
	if _, err = g.Update(ctx, paymentMethod.GetToken(), &PaymentMethodRequest{PaymentMethodNonce: FakeNoncePayPalOneTimePayment}); err == nil {
		t.Errorf("Updating a PayPal account with a different nonce should have failed")
	}

	// Delete PayPal
	if err := g.Delete(ctx, paymentMethod.GetToken()); err != nil {
		t.Fatal(err)
	}

	// Cleanup
	if err := testGateway.Customer().Delete(ctx, cust.Id); err != nil {
		t.Fatal(err)
	}
}

func TestPaymentMethodFailedAutoVerification(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cust, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.PaymentMethod()
	pm, err := g.Create(ctx, &PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: FakeNonceProcessorDeclinedVisa,
	})
	if err == nil {
		t.Fatal("Got no error, want error")
	}
	if g, w := err.(*BraintreeError).ErrorMessage, "Do Not Honor"; g != w {
		t.Fatalf("Got error %q, want error %q", g, w)
	}

	t.Logf("%#v\n", err)
	t.Logf("%#v\n", pm)
}

func TestPaymentMethodForceNotVerified(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cust, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.PaymentMethod()
	pm, err := g.Create(ctx, &PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: FakeNonceProcessorDeclinedVisa,
		Options: &PaymentMethodRequestOptions{
			VerifyCard: testhelpers.BoolPtr(false),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", pm)
}
