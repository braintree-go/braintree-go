package braintree

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type OauthGateway struct {
	*Braintree
}

type CreateTokenFromCodeInput struct {
	Code string
}

type CreateTokenFromRefreshTokenInput struct {
	RefreshToken string
}

type ConnectURLInput struct {
	MerchantId     string
	RedirectURI    string
	Scope          string
	State          string
	User           map[string]string
	Business       map[string]string
	PaymentMethods []string
}

func (g *OauthGateway) CreateTokenFromCode(ctx context.Context, input *CreateTokenFromCodeInput) (*oauth2.Token, error) {
	return g.Config().Exchange(ctx, input.Code)
}

func (g *OauthGateway) CreateTokenFromRefreshToken(ctx context.Context, input *CreateTokenFromRefreshTokenInput) (*oauth2.Token, error) {
	ts := g.Config().TokenSource(ctx, &oauth2.Token{RefreshToken: input.RefreshToken})
	return ts.Token()
}

func (g *OauthGateway) ConnectURL(input *ConnectURLInput) string {
	params := []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("merchant_id", input.MerchantId),
		oauth2.SetAuthURLParam("redirect_uri", input.RedirectURI),
		oauth2.SetAuthURLParam("scope", input.Scope),
	}
	for k, s := range input.User {
		params = append(params, oauth2.SetAuthURLParam(fmt.Sprintf("user[%v]", k), s))
	}
	for k, s := range input.Business {
		params = append(params, oauth2.SetAuthURLParam(fmt.Sprintf("business[%v]", k), s))
	}
	for i, s := range input.PaymentMethods {
		params = append(params, oauth2.SetAuthURLParam(fmt.Sprintf("payment_methods[%v]", i), s))
	}

	// url
	var buf bytes.Buffer
	uri := g.Config().AuthCodeURL(input.State, params...)
	buf.WriteString(uri)

	// signature
	digest := sha256.Sum256([]byte(g.Config().ClientSecret))
	mac := hmac.New(sha256.New, digest[:])
	mac.Write([]byte(uri))
	v := url.Values{
		"signature": {string(mac.Sum(nil))},
		"algorithm": {"SHA256"},
	}

	// add signature to url
	if strings.Contains(uri, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}

func (g *OauthGateway) Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     g.Braintree.ClientId,
		ClientSecret: g.Braintree.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("%v/oauth/connect", g.Braintree.Environment.BaseURL()),
			TokenURL: fmt.Sprintf("%v/oauth/access_tokens", g.Braintree.Environment.BaseURL()),
		},
	}
}
