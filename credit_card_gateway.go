package braintree

import (
	"context"
	"encoding/xml"
	"net/url"
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

// ExpiredIDs finds IDs of credit cards that have expired, returning the IDs
// only. Use Expired and ExpiredNext to get pages of credit cards.
func (g *CreditCardGateway) ExpiredIDs(ctx context.Context) (*SearchResult, error) {
	resp, err := g.execute(ctx, "POST", "/payment_methods/all/expired_ids", nil)
	if err != nil {
		return nil, err
	}

	var searchResult struct {
		PageSize int `xml:"page-size"`
		Ids      struct {
			Item []string `xml:"item"`
		} `xml:"ids"`
	}
	err = xml.Unmarshal(resp.Body, &searchResult)
	if err != nil {
		return nil, err
	}

	return &SearchResult{
		PageSize: searchResult.PageSize,
		IDs:      searchResult.Ids.Item,
	}, nil
}

// ExpiringBetweenIDs finds IDs of credit cards that expire between the
// specified dates, returning the IDs only. Use ExpiringBetween and
// ExpiringBetweenNext to get pages of credit cards.
func (g *CreditCardGateway) ExpiringBetweenIDs(ctx context.Context, fromDate, toDate time.Time) (*SearchResult, error) {
	qs := url.Values{}
	qs.Set("start", fromDate.UTC().Format("012006"))
	qs.Set("end", toDate.UTC().Format("012006"))
	resp, err := g.execute(ctx, "POST", "/payment_methods/all/expiring_ids?"+qs.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var searchResult struct {
		PageSize int `xml:"page-size"`
		Ids      struct {
			Item []string `xml:"item"`
		} `xml:"ids"`
	}
	err = xml.Unmarshal(resp.Body, &searchResult)
	if err != nil {
		return nil, err
	}

	return &SearchResult{
		PageSize: searchResult.PageSize,
		IDs:      searchResult.Ids.Item,
	}, nil
}

// ExpiringBetween finds credit cards that expire between the specified dates,
// returning the first page of results. Use ExpiringBetweenNext to get
// subsequent pages.
func (g *CreditCardGateway) ExpiringBetween(ctx context.Context, fromDate, toDate time.Time) (*CreditCardSearchResult, error) {
	searchResult, err := g.ExpiringBetweenIDs(ctx, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	pageSize := searchResult.PageSize
	ids := searchResult.IDs

	endOffset := pageSize
	if endOffset > len(ids) {
		endOffset = len(ids)
	}

	firstPageQuery := &SearchQuery{}
	firstPageQuery.AddMultiField("ids").Items = ids[:endOffset]
	firstPageCreditCards, err := g.fetchExpiringBetween(ctx, firstPageQuery, fromDate, toDate)

	firstPageResult := &CreditCardSearchResult{
		TotalItems:        len(ids),
		TotalIDs:          ids,
		CurrentPageNumber: 1,
		PageSize:          pageSize,
		CreditCards:       firstPageCreditCards,
	}

	return firstPageResult, err
}

// ExpiringBetweenNext finds the next page of credit cards that expire between
// the specified dates. Use ExpiringBetween to start and get the first page of
// results. Returns a nil result and nil error when no more results are
// available.
func (g *CreditCardGateway) ExpiringBetweenNext(ctx context.Context, fromDate, toDate time.Time, prevResult *CreditCardSearchResult) (*CreditCardSearchResult, error) {
	startOffset := prevResult.CurrentPageNumber * prevResult.PageSize
	endOffset := startOffset + prevResult.PageSize
	if endOffset > len(prevResult.TotalIDs) {
		endOffset = len(prevResult.TotalIDs)
	}
	if startOffset >= endOffset {
		return nil, nil
	}

	nextPageQuery := &SearchQuery{}
	nextPageQuery.AddMultiField("ids").Items = prevResult.TotalIDs[startOffset:endOffset]
	nextPageCreditCard, err := g.fetchExpiringBetween(ctx, nextPageQuery, fromDate, toDate)

	nextPageResult := &CreditCardSearchResult{
		TotalItems:        prevResult.TotalItems,
		TotalIDs:          prevResult.TotalIDs,
		CurrentPageNumber: prevResult.CurrentPageNumber + 1,
		PageSize:          prevResult.PageSize,
		CreditCards:       nextPageCreditCard,
	}

	return nextPageResult, err
}

func (g *CreditCardGateway) fetchExpiringBetween(ctx context.Context, query *SearchQuery, fromDate, toDate time.Time) ([]*CreditCard, error) {
	qs := url.Values{}
	qs.Set("start", fromDate.UTC().Format("012006"))
	qs.Set("end", toDate.UTC().Format("012006"))
	resp, err := g.execute(ctx, "POST", "/payment_methods/all/expiring?"+qs.Encode(), query)
	if err != nil {
		return nil, err
	}

	var v struct {
		CreditCards []*CreditCard `xml:"credit-card"`
	}

	err = xml.Unmarshal(resp.Body, &v)
	if err != nil {
		return nil, err
	}

	return v.CreditCards, nil
}
