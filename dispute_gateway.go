package braintree

import (
	"context"
	"encoding/xml"
)

type DisputeGateway struct {
	*Braintree
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

func (g *DisputeGateway) Search(ctx context.Context, query *SearchQuery) ([]*Dispute, error) {
	resp, err := g.executeVersion(ctx, "POST", "disputes/advanced_search", query, apiVersion4)
	if err != nil {
		return nil, err
	}
	var v struct {
		XMLName  string     `xml:"disputes"`
		Disputes []*Dispute `xml:"dispute"`
	}
	err = xml.Unmarshal(resp.Body, &v)
	if err != nil {
		return nil, err
	}
	return v.Disputes, err
}
