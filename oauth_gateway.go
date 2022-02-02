package braintree

import (
	"context"
	"encoding/xml"
)

type OauthGateway struct {
	*Braintree
}

type OAuthCredentialRequest struct {
	XMLName      xml.Name `xml:"credentials"`
	Code         string   `xml:"code"`
	Scope        string   `xml:"scope"`
	GrantType    string   `xml:"grantType"`
	RefreshToken string   `xml:"refreshToken"`
}

type OAuthCredentials struct {
	AccessToken  string `xml:"access-token"`
	RefreshToken string `xml:"refresh-token"`
	TokenType    string `xml:"token-type"`
	Scope        string `xml:"scope"`
	ExpiresAt    string `xml:"expires-at"`
}

func (g *OauthGateway) CreateTokenFromCode(ctx context.Context, oAuthCredentialRequest *OAuthCredentialRequest) (*OAuthCredentials, error) {
	oAuthCredentialRequest.GrantType = "authorization_code"
	resp, err := g.executeVersion(ctx, "POST", "oauth/access_tokens", oAuthCredentialRequest, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.oauth()
	}
	return nil, &invalidResponseError{resp}
}
