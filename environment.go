package braintree

type Environment struct {
	baseURL string
}

func NewEnvironment(baseURL string) Environment {
	return Environment{baseURL: baseURL}
}

func (e Environment) BaseURL() string {
	return e.baseURL
}

var (
	Development = NewEnvironment("http://localhost:3000")
	Sandbox     = NewEnvironment("https://sandbox.braintreegateway.com")
	Production  = NewEnvironment("https://www.braintreegateway.com")
)
