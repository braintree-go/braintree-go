package braintree

import (
	"context"
	"encoding/xml"
	"fmt"
	"time"
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

//  Expired finds expired credit cards
func (g *CreditCardGateway) Expired(ctx context.Context) ([]*CreditCard, error) {
	resp, err := g.execute(ctx, "POST", "/payment_methods/all/expired_ids", nil)
	if err != nil {
		return nil, err
	}

	var searchResult struct {
		PageSize int `xml:"page-size"`
		Ids struct {
			Item []string `xml:"item"`
		} `xml:"ids"`
	}
	err = xml.Unmarshal(resp.Body, &searchResult)
	if err != nil {
		return nil, err
	}

	searchQuery := &SearchQuery{}
	multiField := searchQuery.AddMultiField("ids")
	for _, item := range searchResult.Ids.Item {
		multiField.Items = append(multiField.Items, item)
	}

	return g.fetchExpired(ctx, searchQuery)
}

// fetchExpired --
func (g *CreditCardGateway) fetchExpired(ctx context.Context, query *SearchQuery) ([]*CreditCard, error) {
	resp, err := g.execute(ctx, "POST", "/payment_methods/all/expired", query)
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

//  ExpiringBetween Find list of credit card that expire between the specified dates
func (g *CreditCardGateway) ExpiringBetween(ctx context.Context, fromDate, toDate time.Time) ([]*CreditCard, error) {
	path := fmt.Sprintf("/payment_methods/all/expiring_ids?start=%s&end=%s",
		fromDate.UTC().Format("012006"),
		toDate.UTC().Format("012006"))
	resp, err := g.execute(ctx, "POST", path, nil)
	if err != nil {
		return nil, err
	}

	var searchResult struct {
		PageSize int `xml:"page-size"`
		Ids struct {
			Item []string `xml:"item"`
		} `xml:"ids"`
	}
	err = xml.Unmarshal(resp.Body, &searchResult)
	if err != nil {
		return nil, err
	}

	searchQuery := &SearchQuery{}
	multiField := searchQuery.AddMultiField("ids")
	for _, item := range searchResult.Ids.Item {
		multiField.Items = append(multiField.Items, item)
	}

	return g.fetchExpiringCreditCards(ctx, searchQuery, fromDate, toDate)
}

// fetchExpiringCreditCards --
func (g *CreditCardGateway) fetchExpiringCreditCards(ctx context.Context, query *SearchQuery, fromDate, toDate time.Time) ([]*CreditCard, error) {
	path := fmt.Sprintf("/payment_methods/all/expiring?start=%s&end=%s",
		fromDate.UTC().Format("012006"),
		toDate.UTC().Format("012006"))
	resp, err := g.execute(ctx, "POST", path, query)
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
