package braintree

import (
	"os"
)

var testCreditCards = map[string]CreditCard{
	"visa":       CreditCard{Number: "4111111111111111"},
	"mastercard": CreditCard{Number: "5555555555554444"},
	"discover":   CreditCard{Number: "6011111111111117"},
}

var testGateway = NewBraintree(Config{
	Environment: Sandbox,
	MerchantId:  os.Getenv("BRAINTREE_MERCH_ID"),
	PublicKey:   os.Getenv("BRAINTREE_PUB_KEY"),
	PrivateKey:  os.Getenv("BRAINTREE_PRIV_KEY"),
})
