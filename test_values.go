package braintree

import "io"

var (
	TestCreditCards = map[string]CreditCard{"visa": CreditCard{Number: "4111111111111111"}}

	testConfiguration = Configuration{
		environment: Development,
		merchantId:  "integration_merchant_id",
		publicKey:   "b6fkbfmhnjdg7333",
		privateKey:  "37912780851d0f68c267ea049cfa0114",
	}

	baseGateway     = NewGateway(testConfiguration)
	txGateway       = TransactionGateway{baseGateway}
	customerGateway = CustomerGateway{baseGateway}
)

type blowUpGateway struct{}

func (this blowUpGateway) Execute(method, url string, body io.Reader) (*Response, error) {
	return &Response{StatusCode: 500, Status: "500 Internal Server Error"}, nil
}

type badInputGateway struct{}

func (this badInputGateway) Execute(method, url string, body io.Reader) (*Response, error) {
	xml := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><api-error-response><errors><errors type=\"array\"/></errors><message>Card Issuer Declined CVV</message></api-error-response>"
	return &Response{StatusCode: 422, Body: []byte(xml)}, nil
}

type notFoundGateway struct{}

func (this notFoundGateway) Execute(method, url string, body io.Reader) (*Response, error) {
	return &Response{StatusCode: 404}, nil
}
