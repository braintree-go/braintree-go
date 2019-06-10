package braintree

import (
	"context"
	"encoding/xml"
	"fmt"
	"math"
)

type DisputeGateway struct {
	*Braintree
}

func (g *DisputeGateway) fetchDisputes(ctx context.Context, query *SearchQuery, page int) (*DisputeSearchResult, error) {
	resp, err := g.executeVersion(ctx, "POST", fmt.Sprintf("disputes/advanced_search?page=%d", page), query, apiVersion4)
	if err != nil {
		return nil, err
	}
	var v struct {
		CurrentPageNumber int        `xml:"current-page-number"`
		PageSize          int        `xml:"page-size"`
		TotalItems        int        `xml:"total-items"`
		XMLName           string     `xml:"disputes"`
		Disputes          []*Dispute `xml:"dispute"`
	}
	err = xml.Unmarshal(resp.Body, &v)
	if err != nil {
		return nil, err
	}
	pageCount := float64(v.TotalItems) / float64(v.PageSize)
	if math.Mod(pageCount, 1) != 0 {
		pageCount++
	}
	result := &DisputeSearchResult{
		PageCount:         int(math.Trunc(pageCount)),
		PageSize:          v.PageSize,
		TotalItems:        v.TotalItems,
		CurrentPageNumber: v.CurrentPageNumber,
		Disputes:          v.Disputes,
	}
	return result, nil
}

func (g *DisputeGateway) Search(ctx context.Context, query *SearchQuery) (*DisputeSearchResult, error) {
	return g.fetchDisputes(ctx, query, 1)
}

func (g *DisputeGateway) SearchPage(ctx context.Context, query *SearchQuery, searchResult *DisputeSearchResult, page int) (*DisputeSearchResult, error) {
	if searchResult == nil {
		page = 1
	} else if page > searchResult.PageCount {
		return nil, nil
	}
	return g.fetchDisputes(ctx, query, page)
}

func (g *DisputeGateway) SearchNext(ctx context.Context, query *SearchQuery, searchResult *DisputeSearchResult) (*DisputeSearchResult, error) {
	if searchResult == nil {
		return nil, nil
	}
	nextPage := searchResult.CurrentPageNumber + 1
	if nextPage > searchResult.PageCount {
		return nil, nil
	}
	return g.fetchDisputes(ctx, query, nextPage)
}

func (g *DisputeGateway) Find(ctx context.Context, disputeID string) (*Dispute, error) {
	resp, err := g.executeVersion(ctx, "GET", "disputes/"+disputeID, nil, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.dispute()
	}
	return nil, &invalidResponseError{resp}
}

func (g *DisputeGateway) AddTextEvidence(ctx context.Context, disputeID string, textEvidenceRequest *DisputeTextEvidenceRequest) (*DisputeEvidence, error) {
	resp, err := g.executeVersion(ctx, "POST", "disputes/"+disputeID+"/evidence", textEvidenceRequest, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.disputeEvidence()
	}
	return nil, &invalidResponseError{resp}
}

func (g *DisputeGateway) RemoveEvidence(ctx context.Context, disputeID string, evidenceId string) error {
	resp, err := g.executeVersion(ctx, "DELETE", "disputes/"+disputeID+"/evidence/"+evidenceId, nil, apiVersion4)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}

func (g *DisputeGateway) Accept(ctx context.Context, disputeID string) error {
	resp, err := g.executeVersion(ctx, "PUT", "disputes/"+disputeID+"/accept", nil, apiVersion4)
	if err != nil {
		return nil
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}

func (g *DisputeGateway) Finalize(ctx context.Context, disputeID string) error {
	resp, err := g.executeVersion(ctx, "PUT", "disputes/"+disputeID+"/finalize", nil, apiVersion4)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}
