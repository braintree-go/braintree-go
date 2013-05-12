package braintree

type Configuration struct {
	environment Environment
	merchantId  string
	publicKey   string
	privateKey  string
}

func (this Configuration) BaseURL() string {
	return this.environment.baseURL + "/merchant/" + this.merchantId
}

type Environment struct {
	name    string
	baseURL string
}

var Sandbox = Environment{name: "sandbox", baseURL: "https://sandbox.braintreegateway.com"}
