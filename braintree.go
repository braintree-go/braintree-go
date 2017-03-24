package braintree

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

const LibraryVersion = "0.9.0"

type ApiVersion int

const (
	ApiVersion3 ApiVersion = 3
	ApiVersion4            = 4
)

func New(env Environment, merchId, pubKey, privKey string) *Braintree {
	return &Braintree{
		Environment: env,
		MerchantId:  merchId,
		PublicKey:   pubKey,
		PrivateKey:  privKey,
	}
}

func NewWithHttpClient(env Environment, merchantId, publicKey, privateKey string, client *http.Client) *Braintree {
	return &Braintree{
		Environment: env,
		MerchantId:  merchantId,
		PublicKey:   publicKey,
		PrivateKey:  privateKey,
		HttpClient:  client,
	}
}

type Braintree struct {
	Environment Environment
	MerchantId  string
	PublicKey   string
	PrivateKey  string
	Logger      *log.Logger
	HttpClient  *http.Client
}

func (g *Braintree) MerchantURL() string {
	return g.Environment.BaseURL + "/merchants/" + g.MerchantId
}

func (g *Braintree) execute(method, path string, xmlObj interface{}) (*Response, error) {
	return g.executeVersion(method, path, xmlObj, ApiVersion3)
}

func (g *Braintree) executeVersion(method, path string, xmlObj interface{}, apiVersion ApiVersion) (*Response, error) {
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
	req.Header.Set("X-ApiVersion", fmt.Sprintf("%d", apiVersion))
	req.SetBasicAuth(g.PublicKey, g.PrivateKey)

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
	return &WebhookNotificationGateway{g}
}

func (g *Braintree) Settlement() *SettlementGateway {
	return &SettlementGateway{g}
}
