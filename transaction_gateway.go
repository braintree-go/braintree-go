package braintree

type TransactionGateway struct {
	*Braintree
}

// Create initiates a transaction.
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

// SubmitForSettlement submits the transaction with the specified id for settlement.
// If the amount is omitted, the full amount is settled.
func (g *TransactionGateway) SubmitForSettlement(id string, amount ...float64) (*Transaction, error) {
	var tx *Transaction
	if len(amount) > 0 {
		tx = &Transaction{
			Amount: amount[0],
		}
	}
	resp, err := g.execute("PUT", "transactions/"+id+"/submit_for_settlement", tx)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Find finds the transaction with the specified id.
func (g *TransactionGateway) Find(id string) (*Transaction, error) {
	resp, err := g.execute("GET", "transactions/"+id, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}
