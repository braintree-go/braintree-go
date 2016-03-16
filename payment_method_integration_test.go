package braintree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestPaymentMethod(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	customer, err := testGateway.Customer().Create(&Customer{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.PaymentMethod()

	// Create
	paymentMethod, err := g.Create(&PaymentMethod{
		CustomerId:         customer.Id,
		PaymentMethodNonce: testFakeValidNonce,
	})
	if err != nil {
		t.Fatal(err)
	}
	creditCard := paymentMethod.(*CreditCard)

	t.Log(paymentMethod)

	if creditCard.Token == "" {
		t.Fatal("invalid token")
	}

	// Update
	token := fmt.Sprintf("new_test_token_%d", rand.Int()+1)
	paymentMethod, err = g.Update(creditCard.Token, &PaymentMethod{
		PaymentMethodNonce: testFakeValidMastercardNonce,
		Token:              token,
	})
	if err != nil {
		t.Fatal(err)
	}
	creditCard = paymentMethod.(*CreditCard)

	t.Log(paymentMethod)

	if creditCard.Token == "" {
		t.Fatal("invalid token")
	}

	// Updating with different payment method nonce should fail
	if _, err = g.Update(token, &PaymentMethod{PaymentMethodNonce: testFakePayPalFutureNonce}); err == nil {
		t.Fatal(err)
	}

	// Find
	paymentMethod, err = g.Find(token)
	if err != nil {
		t.Fatal(err)
	}
	creditCard = paymentMethod.(*CreditCard)

	if creditCard.Token != token || creditCard.CustomerId != customer.Id {
		t.Fatal("Unexpected payment method attributes")
	}

	// Delete
	if err := g.Delete(token); err != nil {
		t.Fatal(err)
	}

	// Cleanup
	if err := testGateway.Customer().Delete(customer.Id); err != nil {
		t.Fatal(err)
	}
}
