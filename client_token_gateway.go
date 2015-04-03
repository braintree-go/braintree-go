package braintree

import "encoding/xml"

const clientTokenVersion = 2

type ClientTokenGateway struct {
	*Braintree
}

func (g *ClientTokenGateway) Generate() (string, error) {
	resp, err := g.execute("POST", "client_token", ClientToken{Version: clientTokenVersion})
	if err != nil {
		return "", err
	}
	switch resp.StatusCode {
	case 201:
		var b struct {
			ClientToken string `xml:"value"`
		}
		if err := xml.Unmarshal(resp.Body, &b); err != nil {
			return "", err
		}
		return b.ClientToken, nil
	}
	return "", &invalidResponseError{resp}
}
