package braintree

type Environment struct {
	GatewayBaseURL string
}

func (e Environment) BaseURL() string {
	return e.GatewayBaseURL
}

var (
	Development = Environment{GatewayBaseURL: "http://localhost:3000"}
	Sandbox     = Environment{GatewayBaseURL: "https://sandbox.braintreegateway.com"}
	Production  = Environment{GatewayBaseURL: "https://www.braintreegateway.com"}
)
