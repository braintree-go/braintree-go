package braintree

import (
	"context"
	"encoding/xml"
)

type SubscriptionGateway struct {
	*Braintree
}

func (g *SubscriptionGateway) Create(ctx context.Context, sub *SubscriptionRequest) (*Subscription, error) {
	resp, err := g.execute(ctx, "POST", "subscriptions", sub)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.subscription()
	}
	return nil, &invalidResponseError{resp}
}

func (g *SubscriptionGateway) Update(ctx context.Context, sub *SubscriptionRequest) (*Subscription, error) {
	resp, err := g.execute(ctx, "PUT", "subscriptions/"+sub.Id, sub)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.subscription()
	}
	return nil, &invalidResponseError{resp}
}

func (g *SubscriptionGateway) Find(ctx context.Context, subId string) (*Subscription, error) {
	resp, err := g.execute(ctx, "GET", "subscriptions/"+subId, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.subscription()
	}
	return nil, &invalidResponseError{resp}
}

func (g *SubscriptionGateway) Cancel(ctx context.Context, subId string) (*Subscription, error) {
	resp, err := g.execute(ctx, "PUT", "subscriptions/"+subId+"/cancel", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.subscription()
	}
	return nil, &invalidResponseError{resp}
}

// RetryCharge retries to charge for a Subscription. All options,
// including the Subscription ID, are to be provided by the
// SubscriptionTransactionRequest passed as an argument. Note that the
// Amount has to be > 0.
func (g *SubscriptionGateway) RetryCharge(ctx context.Context, txReq *SubscriptionTransactionRequest) error {
	resp, err := g.execute(ctx, "POST", "transactions", txReq)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 201:
		return nil
	}
	return &invalidResponseError{resp}
}

// SearchIDs finds subscriptions matching the search query, returning the IDs
// only. Use Search and SearchNext to get pages of subscriptions.
func (g *SubscriptionGateway) SearchIDs(ctx context.Context, query *SearchQuery) (*SearchResult, error) {
	resp, err := g.execute(ctx, "POST", "subscriptions/advanced_search_ids", query)
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

// Search finds subscriptions matching the search query, returning the first
// page of results. Use SearchNext to get subsequent pages.
func (g *SubscriptionGateway) Search(ctx context.Context, query *SearchQuery) (*SubscriptionSearchResult, error) {
	searchResult, err := g.SearchIDs(ctx, query)
	if err != nil {
		return nil, err
	}

	pageSize := searchResult.PageSize
	ids := searchResult.IDs

	endOffset := pageSize
	if endOffset > len(ids) {
		endOffset = len(ids)
	}

	firstPageQuery := query.shallowCopy()
	firstPageQuery.AddMultiField("ids").Items = ids[:endOffset]
	firstPageSubscriptions, err := g.fetchSubscriptions(ctx, firstPageQuery)

	firstPageResult := &SubscriptionSearchResult{
		TotalItems:        len(ids),
		TotalIDs:          ids,
		CurrentPageNumber: 1,
		PageSize:          pageSize,
		Subscriptions:     firstPageSubscriptions,
	}

	return firstPageResult, err
}

// SearchNext finds the next page of Subscriptions matching the search query.
// Use Search to start a search and get the first page of results.
// Returns a nil result and nil error when no more results are available.
func (g *SubscriptionGateway) SearchNext(ctx context.Context, query *SearchQuery, prevResult *SubscriptionSearchResult) (*SubscriptionSearchResult, error) {
	startOffset := prevResult.CurrentPageNumber * prevResult.PageSize
	endOffset := startOffset + prevResult.PageSize
	if endOffset > len(prevResult.TotalIDs) {
		endOffset = len(prevResult.TotalIDs)
	}
	if startOffset >= endOffset {
		return nil, nil
	}

	nextPageQuery := query.shallowCopy()
	nextPageQuery.AddMultiField("ids").Items = prevResult.TotalIDs[startOffset:endOffset]
	nextPageSubscriptions, err := g.fetchSubscriptions(ctx, nextPageQuery)

	nextPageResult := &SubscriptionSearchResult{
		TotalItems:        prevResult.TotalItems,
		TotalIDs:          prevResult.TotalIDs,
		CurrentPageNumber: prevResult.CurrentPageNumber + 1,
		PageSize:          prevResult.PageSize,
		Subscriptions:     nextPageSubscriptions,
	}

	return nextPageResult, err
}

func (g *SubscriptionGateway) fetchSubscriptions(ctx context.Context, query *SearchQuery) ([]*Subscription, error) {
	resp, err := g.execute(ctx, "POST", "subscriptions/advanced_search", query)
	if err != nil {
		return nil, err
	}
	var v struct {
		XMLName       string          `xml:"subscriptions"`
		Subscriptions []*Subscription `xml:"subscription"`
	}
	err = xml.Unmarshal(resp.Body, &v)
	if err != nil {
		return nil, err
	}
	return v.Subscriptions, err
}
