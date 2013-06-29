package braintree

import (
	"os"
)

var (
	TestCreditCards = map[string]CreditCard{
		"visa":       CreditCard{Number: "4111111111111111"},
		"mastercard": CreditCard{Number: "5555555555554444"},
		"discover":   CreditCard{Number: "6011111111111117"},
	}

	testConfiguration = Configuration{
		environment: Sandbox,
		merchantId:  os.Getenv("BRAINTREE_MERCH_ID"),
		publicKey:   os.Getenv("BRAINTREE_PUB_KEY"),
		privateKey:  os.Getenv("BRAINTREE_PRIV_KEY"),
	}

	gateway = NewGateway(testConfiguration)
)

type blowUpGateway struct{}

func (this blowUpGateway) Execute(method, url string, body []byte) (*Response, error) {
	return &Response{StatusCode: 500, Status: "500 Internal Server Error"}, nil
}

type badInputGateway struct{}

func (this badInputGateway) Execute(method, url string, body []byte) (*Response, error) {
	xml := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><api-error-response><errors><errors type=\"array\"/></errors><message>Card Issuer Declined CVV</message></api-error-response>"
	return &Response{StatusCode: 422, Body: []byte(xml)}, nil
}

type notFoundGateway struct{}

func (this notFoundGateway) Execute(method, url string, body []byte) (*Response, error) {
	return &Response{StatusCode: 404}, nil
}
