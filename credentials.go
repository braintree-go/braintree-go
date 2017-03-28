package braintree

type Credentials interface {
	Environment() Environment
	MerchantID() string
	AuthorizationHeader() string
}
