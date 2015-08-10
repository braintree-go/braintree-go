package braintree

import (
	"os"
)

var testCreditCards = map[string]CreditCard{
	"visa":       CreditCard{Number: "4111111111111111"},
	"mastercard": CreditCard{Number: "5555555555554444"},
	"discover":   CreditCard{Number: "6011111111111117"},
}

var testPaypalAccounts = map[string]PaypalAccount{
	"example": PaypalAccount{Email: "test@example.com"},
	"test":    PaypalAccount{Email: "jane.doe@example.com", Token: os.Getenv("BRAINTREE_PAYPAL_ACCOUNT_TOKEN")},
}

var testGateway = New(
	Sandbox,
	os.Getenv("BRAINTREE_MERCH_ID"),
	os.Getenv("BRAINTREE_PUB_KEY"),
	os.Getenv("BRAINTREE_PRIV_KEY"),
)

var testMerchantAccountId = os.Getenv("BRAINTREE_MERCH_ACCT_ID")
