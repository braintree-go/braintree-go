package braintree

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const apiVersion5 = "5"

// TransparentRedirectGateway is the gateway for TR.
type TransparentRedirectGateway struct {
	*Braintree
}

// TransactionData creates the signed transaction data to embded in the html form.
func (tr *TransparentRedirectGateway) TransactionData(data *TransparentRedirectData) (string, error) {
	data.Kind = TransparentRedirectKindCreateTransaction
	return tr.generateTrData(data)
}

func (tr *TransparentRedirectGateway) generateTrData(inputData *TransparentRedirectData) (trData string, err error) {
	apiKey, ok := tr.credentials.(apiKey)
	if !ok {
		return "", errors.New("generateTrData can only be used with Braintree Credentials that are API Keys")
	}

	v, _ := query.Values(inputData)
	v.Add("api_version", apiVersion5)
	v.Add("time", time.Now().UTC().Format("20060102150405"))
	v.Add("public_key", apiKey.publicKey)
	hmacer := newHmacer(apiKey.publicKey, apiKey.privateKey)
	hmac, err := hmacer.hmac(v.Encode())
	if err != nil {
		return
	}

	trData = hmac + "|" + v.Encode()
	return
}

// ValidateQueryString validates the signature on the query string passed to the callback by braintree.
func (tr *TransparentRedirectGateway) ValidateQueryString(query string) (bool, error) {
	apiKey, ok := tr.credentials.(apiKey)
	if !ok {
		return false, errors.New("validateQueryString can only be used with Braintree Credentials that are API Keys")
	}

	splitQuery := strings.Split(query, "&hash=")
	if len(splitQuery) != 2 {
		return false, errors.New("query is incorrect and has no hash parameter")
	}

	hmacer := newHmacer(apiKey.publicKey, apiKey.privateKey)
	signature, err := hmacer.hmac(splitQuery[0])
	if err != nil {
		return false, errors.New("unable to hmac query")
	}

	return signature == splitQuery[1], nil
}

// FormURL returns the URL the html form has to POST to.
func (tr *TransparentRedirectGateway) FormURL() string {
	return tr.Environment().baseURL + "/transparent_redirect_requests"
}

// Confirm confirms the transaction.
func (tr *TransparentRedirectGateway) Confirm(ctx context.Context, query string) (t *Transaction, err error) {
	v, err := url.ParseQuery(query)
	if err != nil {
		return nil, errors.New("can't parse query")
	}

	if v.Get("http_status") != "200" {
		return nil, errors.New("expected http status to be 200")
	}

	if ok, err := tr.ValidateQueryString(query); !ok {
		return nil, err
	}

	trID := v.Get("id")
	if trID == "" {
		return nil, errors.New("transparent redirect id is not set")
	}

	resp, err := tr.execute(ctx, "POST", "/transparent_redirect_requests/"+trID+"/confirm", t)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200, 201:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}
