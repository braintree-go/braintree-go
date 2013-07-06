package braintree

type TransactionGateway struct {
	*Braintree
}

func (g *TransactionGateway) Create(tx *Transaction) (*Transaction, error) {
	resp, err := g.execute("POST", "transactions", tx)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

func (g *TransactionGateway) Find(txId string) (*Transaction, error) {
	resp, err := g.execute("GET", "transactions/"+txId, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}
