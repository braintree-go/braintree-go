package braintree

type Environment struct {
	BaseURL string
}

var (
	Development = Environment{BaseURL: "http://localhost:3000"}
	Sandbox     = Environment{BaseURL: "https://sandbox.braintreegateway.com"}
	Production  = Environment{BaseURL: "https://www.braintreegateway.com"}
)
