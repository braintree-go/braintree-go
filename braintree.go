package braintree

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type apiVersion int

type HTTPClient interface {
	Do(req *http.Request) (res *http.Response, err error)
	Get(url string) (res *http.Response, err error)
	Post(urstring, bodyType string, body io.Reader) (res *http.Response, err error)
	PostForm(url string, values url.Values) (res *http.Response, err error)
	Head(url string) (res *http.Response, err error)
}

const (
	apiVersion3    apiVersion = 3
	apiVersion4               = 4
	LibraryVersion            = "0.10.0"
)

// New creates a Braintree with API Keys.
func New(env Environment, merchId, pubKey, privKey string) *Braintree {
	return &Braintree{credentials: newAPIKey(env, merchId, pubKey, privKey)}
}

// NewWithHttpClient creates a Braintree with API Keys and a HTTP Client.
func NewWithHttpClient(env Environment, merchantId, publicKey, privateKey string, client HTTPClient) *Braintree {
	return &Braintree{credentials: newAPIKey(env, merchantId, publicKey, privateKey), HttpClient: client}
}

// New creates a Braintree with an Access Token.
//
// Note: When using an access token, webhooks are unsupported and the
// WebhookNotification() function will panic.
func NewWithAccessToken(accessToken string) (*Braintree, error) {
	c, err := newAccessToken(accessToken)
	if err != nil {
		return nil, err
	}
	return &Braintree{credentials: c}, nil
}

// Braintree interacts with the Braintree API.
type Braintree struct {
	credentials credentials
	Logger      *log.Logger
	HttpClient  HTTPClient
}

func (g *Braintree) Environment() Environment {
	return g.credentials.Environment()
}

func (g *Braintree) MerchantID() string {
	return g.credentials.MerchantID()
}

func (g *Braintree) MerchantURL() string {
	return g.Environment().BaseURL() + "/merchants/" + g.MerchantID()
}

func (g *Braintree) execute(method, path string, xmlObj interface{}) (*Response, error) {
	return g.executeVersion(method, path, xmlObj, apiVersion3)
}

func (g *Braintree) executeVersion(method, path string, xmlObj interface{}, v apiVersion) (*Response, error) {
	var buf bytes.Buffer
	if xmlObj != nil {
		xmlBody, err := xml.Marshal(xmlObj)
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(xmlBody)
		if err != nil {
			return nil, err
		}
	}

	url := g.MerchantURL() + "/" + path

	if g.Logger != nil {
		g.Logger.Printf("> %s %s\n%s", method, url, buf.String())
	}

	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", fmt.Sprintf("Braintree Go %s", LibraryVersion))
	req.Header.Set("X-ApiVersion", fmt.Sprintf("%d", v))
	req.Header.Set("Authorization", g.credentials.AuthorizationHeader())

	httpClient := g.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	btr := &Response{
		Response: resp,
	}
	err = btr.unpackBody()
	if err != nil {
		return nil, err
	}

	if g.Logger != nil {
		g.Logger.Printf("<\n%s", string(btr.Body))
	}

	err = btr.apiError()
	if err != nil {
		return nil, err
	}
	return btr, nil
}

func (g *Braintree) ClientToken() *ClientTokenGateway {
	return &ClientTokenGateway{g}
}

func (g *Braintree) MerchantAccount() *MerchantAccountGateway {
	return &MerchantAccountGateway{g}
}

func (g *Braintree) Transaction() *TransactionGateway {
	return &TransactionGateway{g}
}

func (g *Braintree) Testing() *TestingGateway {
	return &TestingGateway{g}
}

func (g *Braintree) PaymentMethod() *PaymentMethodGateway {
	return &PaymentMethodGateway{g}
}

func (g *Braintree) CreditCard() *CreditCardGateway {
	return &CreditCardGateway{g}
}

func (g *Braintree) PayPalAccount() *PayPalAccountGateway {
	return &PayPalAccountGateway{g}
}

func (g *Braintree) Customer() *CustomerGateway {
	return &CustomerGateway{g}
}

func (g *Braintree) Subscription() *SubscriptionGateway {
	return &SubscriptionGateway{g}
}

func (g *Braintree) Plan() *PlanGateway {
	return &PlanGateway{g}
}

func (g *Braintree) Address() *AddressGateway {
	return &AddressGateway{g}
}

func (g *Braintree) AddOn() *AddOnGateway {
	return &AddOnGateway{g}
}

func (g *Braintree) Discount() *DiscountGateway {
	return &DiscountGateway{g}
}

func (g *Braintree) WebhookNotification() *WebhookNotificationGateway {
	if apiKey, ok := g.credentials.(apiKey); !ok {
		panic(errors.New("WebhookNotifications can only be used with Braintree Credentials that are API Keys."))
	} else {
		return &WebhookNotificationGateway{Braintree: g, apiKey: apiKey}
	}
}

func (g *Braintree) Settlement() *SettlementGateway {
	return &SettlementGateway{g}
}
