package braintree

type Configuration struct {
  environment string
  merchant_id string
  public_key string
  private_key string
}

type Gateway struct {
  config Configuration
}

func (this Gateway) Transaction() TransactionGateway {
  return TransactionGateway{}
}

func NewGateway(config Configuration) Gateway {
  return Gateway{config}
}

type TransactionResponse struct {}

func (this TransactionResponse) IsValid() bool { return false }

type TransactionGateway struct {}

func (this TransactionGateway) Sale(request TransactionRequest) TransactionResponse {
  return TransactionResponse{}
}
