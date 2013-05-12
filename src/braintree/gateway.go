package braintree

type Gateway struct {
	config Configuration
}

func (this Gateway) Transaction() TransactionGateway {
	return TransactionGateway{this}
}

func NewGateway(config Configuration) Gateway {
	return Gateway{config}
}

type TransactionGateway struct {
	parent Gateway
}

func (this TransactionGateway) Sale(request TransactionRequest) TransactionResponse {
	return TransactionResponse{}
}

type TransactionResponse struct{}

func (this TransactionResponse) IsSuccess() bool { return false }
