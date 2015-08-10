package braintree

import "testing"

func TestPaypalAccount(t *testing.T) {
	_, err := testGateway.Customer().Create(&Customer{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.PaypalAccount()

	if testPaypalAccounts["test"].Token == "" {
		t.Skip("No paypal token for test account, skipping integration test")
	}

	// Find
	paypalAccount, err := g.Find(testPaypalAccounts["test"].Token)
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
		Email: testPaypalAccounts["example"].Email,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(paypalAccount2)

	if paypalAccount2.Token != paypalAccount.Token {
		t.Fatal("tokens do not match")
	}
	if paypalAccount2.Email != testPaypalAccounts["example"].Email {
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
