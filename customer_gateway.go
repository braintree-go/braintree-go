package braintree

type CustomerGateway struct {
	gateway Gateway
}

func (this CustomerGateway) Create(customer Customer) (CustomerResult, error) {
	return ErrorResult{}, nil
}
