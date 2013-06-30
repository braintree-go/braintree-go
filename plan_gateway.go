package braintree

import (
	"encoding/xml"
)

type PlanGateway struct {
	*Braintree
}

func (g *PlanGateway) All() ([]*Plan, error) {
	resp, err := g.Execute("GET", "plans", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		var b Plans
		if err := xml.Unmarshal(resp.Body, &b); err != nil {
			return nil, err
		}
		return b.Plan, nil
	}
	return nil, &InvalidResponseError{resp}
}
