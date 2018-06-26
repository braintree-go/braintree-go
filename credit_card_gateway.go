package braintree

import (
	"context"
	"fmt"
	"time"
	"encoding/xml"
)

type CreditCardGateway struct {
	*Braintree
}

// Create creates a new credit card.
func (g *CreditCardGateway) Create(ctx context.Context, card *CreditCard) (*CreditCard, error) {
	resp, err := g.execute(ctx, "POST", "payment_methods", card)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.creditCard()
	}
	return nil, &invalidResponseError{resp}
}

// Update updates a credit card.
func (g *CreditCardGateway) Update(ctx context.Context, card *CreditCard) (*CreditCard, error) {
	resp, err := g.execute(ctx, "PUT", "payment_methods/"+card.Token, card)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.creditCard()
	}
	return nil, &invalidResponseError{resp}
}

// Find finds a credit card by payment method token.
func (g *CreditCardGateway) Find(ctx context.Context, token string) (*CreditCard, error) {
	resp, err := g.execute(ctx, "GET", "payment_methods/"+token, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.creditCard()
	}
	return nil, &invalidResponseError{resp}
}

// Delete deletes a credit card.
func (g *CreditCardGateway) Delete(ctx context.Context, card *CreditCard) error {
	resp, err := g.execute(ctx, "DELETE", "payment_methods/"+card.Token, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}

//  ExpiringBetween Find list of credit card that expire between the specified dates
func (g *CreditCardGateway) ExpiringBetween(ctx context.Context, fromDate, toDate time.Time) ([]*CreditCard, error) {
	path := fmt.Sprintf("/payment_methods/all/expiring?start=%s&end=%s",
		fromDate.Format("012006"),
		toDate.Format("012006"))
	resp, err := g.execute(ctx, "POST", path, nil)
	if err != nil {
		return nil, err
	}

	type searchResult struct {
		XMLName           xml.Name      `xml:"payment-methods"`
		CurrentPageNumber int64         `xml:"current-page-number"`
		PageSize          int64         `xml:"page-size"`
		TotalItems        int64         `xml:"total-items"`
		CreditCards       []*CreditCard `xml:"credit-card"`
	}

	cc := &searchResult{}
	err = xml.Unmarshal(resp.Body, &cc)
	if err != nil {
		return nil, err
	}

	return cc.CreditCards, nil
}
