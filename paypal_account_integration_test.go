package braintree

import "testing"

func TestPaypalAccount(t *testing.T) {
	cust, err := testGateway.Customer().Create(&Customer{})
	if err != nil {
		t.Fatal(err)
	}

	nonce := FakeNoncePayPalFuturePayment

	g := testGateway.PaypalAccount()
	paymentMethod, err := testGateway.PaymentMethod().Create(&PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: nonce,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Find
	paypalAccount, err := g.Find(paymentMethod.GetToken())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(paypalAccount)

	if paypalAccount.Token == "" {
		t.Fatal("invalid token")
	}

	// Update
	paypalAccount2, err := g.Update(&PaypalAccount{
		Token: paypalAccount.Token,
		Email: "new-email@example.com",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(paypalAccount2)

	if paypalAccount2.Token != paypalAccount.Token {
		t.Fatal("tokens do not match")
	}
	if paypalAccount2.Email != "new-email@example.com" {
		t.Fatal("paypalAccount email does not match")
	}

	// Delete
	err = g.Delete(paypalAccount2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindPaypalAccountBadData(t *testing.T) {
	paypalAccount, err := testGateway.PaypalAccount().Find("invalid_token")

	t.Log(paypalAccount)

	if err == nil {
		t.Fatal("expected to get error because the token is invalid")
	}
}
