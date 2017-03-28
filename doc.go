/*
Package braintree is a client library for Braintree.

Initializing

Initialize it with API Keys:

	c := braintree.NewAPIKey(env, merchantID, publicKey, privateKey)
	bt := &braintree.Braintree{Credentials: c}

Initialize it with an Access Token:

	c, _ := braintree.NewAccessToken(accessToken)
	bt := &braintree.Braintree{Credentials: c}

Loggers and HTTP Clients

Optionally configure a logger and HTTP client:

	bt := &braintree.Braintree{
		Credentials: ...,
		Logger: log.New(...),
		HttpClient: ...,
	}

Creating Transactions

Create transactions:

	t, err := bt.Transaction().Create(&braintree.Transaction{
		Type:   "sale",
		Amount: braintree.NewDecimal(100, 2), // $1.00
		PaymentMethodNonce: braintree.FakeNonceTransactable,
	})

API Errors

API errors are intended to be consumed in two ways. One, they can be dealt with as a single unit:

	t, err := bt.Transaction().Create(...)
	err.Error() => "A top level error message"

Second, you can drill down to see specific error messages on a field-by-field basis:

	err.For("Transaction").On("Base")[0].Message => "A more specific error message"
*/
package braintree
