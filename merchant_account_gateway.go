package braintree

type MerchantAccountGateway struct {
	*Braintree
}

// Create a sub merchant account.
func (g *MerchantAccountGateway) Create(ma *MerchantAccount) (*MerchantAccount, error) {
	resp, err := g.execute("POST", "merchant_accounts/create_via_api", ma)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.merchantAccount()
	}
	return nil, &invalidResponseError{resp}
}

// Find finds the merchant account with the specified id.
func (g *MerchantAccountGateway) Find(id string) (*MerchantAccount, error) {
	resp, err := g.execute("GET", "merchant_accounts/"+id, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.merchantAccount()
	}
	return nil, &invalidResponseError{resp}
}

// Update a sub merchant account.
func (g *MerchantAccountGateway) Update(ma *MerchantAccount) (*MerchantAccount, error) {
	resp, err := g.execute("PUT", "merchant_acocunts/"+ma.Id+"/update_via_api", ma)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.merchantAccount()
	}
	return nil, &invalidResponseError{resp}
}
