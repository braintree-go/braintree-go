package braintree

type Config struct {
	Environment Environment
	MerchantId  string
	PublicKey   string
	PrivateKey  string
}

func (c *Config) BaseURL() string {
	return c.Environment.BaseURL + "/merchants/" + c.MerchantId
}

type Environment struct {
	Name    string
	BaseURL string
}

var (
	Development = Environment{
		Name:    "development",
		BaseURL: "http://localhost:3000",
	}
	Sandbox = Environment{
		Name:    "sandbox",
		BaseURL: "https://sandbox.braintreegateway.com",
	}
	Production = Environment{
		Name:    "production",
		BaseURL: "https://www.braintreegateway.com",
	}
)
