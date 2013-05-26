package braintree

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
