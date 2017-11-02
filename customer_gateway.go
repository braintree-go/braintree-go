package braintree

import (
	"context"
	"encoding/xml"
)

type CustomerGateway struct {
	*Braintree
}

// Create creates a new customer from the passed in customer object.
// If no Id is set, Braintree will assign one.
func (g *CustomerGateway) Create(ctx context.Context, c *Customer) (*Customer, error) {
	resp, err := g.execute(ctx, "POST", "customers", c)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.customer()
	}
	return nil, &invalidResponseError{resp}
}

// Update updates any field that is set in the passed customer object.
// The Id field is mandatory.
func (g *CustomerGateway) Update(ctx context.Context, c *Customer) (*Customer, error) {
	resp, err := g.execute(ctx, "PUT", "customers/"+c.Id, c)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.customer()
	}
	return nil, &invalidResponseError{resp}
}

// Find finds the customer with the given id.
func (g *CustomerGateway) Find(ctx context.Context, id string) (*Customer, error) {
	resp, err := g.execute(ctx, "GET", "customers/"+id, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.customer()
	}
	return nil, &invalidResponseError{resp}
}

func (g *CustomerGateway) Search(ctx context.Context, query *SearchQuery) (*CustomerSearchResult, error) {
	resp, err := g.execute(ctx, "POST", "customers/advanced_search", query)
	if err != nil {
		return nil, err
	}
	var v CustomerSearchResult
	err = xml.Unmarshal(resp.Body, &v)
	if err != nil {
		return nil, err
	}
	return &v, err
}

// Delete deletes the customer with the given id.
func (g *CustomerGateway) Delete(ctx context.Context, id string) error {
	resp, err := g.execute(ctx, "DELETE", "customers/"+id, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}
