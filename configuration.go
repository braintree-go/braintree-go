package braintree

type Configuration struct {
	environment Environment
	merchantId  string
	publicKey   string
	privateKey  string
}

func (this Configuration) BaseURL() string {
	return this.environment.BaseURL + "/merchants/" + this.merchantId
}

type Environment struct {
	Name    string
	BaseURL string
}

var (
	Development = Environment{Name: "development", BaseURL: "http://localhost:3000"}
	Sandbox     = Environment{Name: "sandbox", BaseURL: "https://sandbox.braintreegateway.com"}
)
