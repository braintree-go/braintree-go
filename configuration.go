package braintree

type Configuration struct {
	environment Environment
	merchantId  string
	publicKey   string
	privateKey  string
}

func (this Configuration) BaseURL() string {
	return this.environment.BaseURL + "/merchant/" + this.merchantId
}

type Environment struct {
	Name    string
	BaseURL string
}

var Sandbox = Environment{Name: "sandbox", BaseURL: "https://sandbox.braintreegateway.com"}
