package braintree

type CustomerGateway struct {
	gateway Gateway
}

func (this CustomerGateway) Create(customer Customer) (Response, error) {
	_, err := customer.ToXML()
	if err != nil {
		return ErrorResponse{}, nil
	}
	return ErrorResponse{}, nil
}
