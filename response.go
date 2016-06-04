package braintree

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
	Body []byte
}

// TODO: remove dedicated unmarshal methods (redundant)

func (r *Response) merchantAccount() (*MerchantAccount, error) {
	var b MerchantAccount
	if err := xml.Unmarshal(r.Body, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *Response) transaction() (*Transaction, error) {
	var b Transaction
	if err := xml.Unmarshal(r.Body, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *Response) creditCard() (*CreditCard, error) {
	var b CreditCard
	if err := xml.Unmarshal(r.Body, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *Response) customer() (*Customer, error) {
	var b Customer
	if err := xml.Unmarshal(r.Body, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *Response) subscription() (*Subscription, error) {
	var b Subscription
	if err := xml.Unmarshal(r.Body, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *Response) settlement() (*SettlementBatchSummary, error) {
	var b SettlementBatchSummary
	if err := xml.Unmarshal(r.Body, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *Response) address() (*Address, error) {
	var b Address
	if err := xml.Unmarshal(r.Body, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *Response) addOns() ([]AddOn, error) {
	var b AddOnList
	if err := xml.Unmarshal(r.Body, &b); err != nil {
		return nil, err
	}
	return b.AddOns, nil
}

func (r *Response) discounts() ([]Discount, error) {
	var b DiscountList
	if err := xml.Unmarshal(r.Body, &b); err != nil {
		return nil, err
	}
	return b.Discounts, nil
}

func (r *Response) unpackBody() error {
	if len(r.Body) == 0 {
		b, err := gzip.NewReader(r.Response.Body)
		if err != nil {
			return err
		}
		defer r.Response.Body.Close()

		buf, err := ioutil.ReadAll(b)
		if err != nil {
			return err
		}
		r.Body = buf
	}
	return nil
}

func (r *Response) apiError() error {
	var b BraintreeError
	xml.Unmarshal(r.Body, &b)
	if b.ErrorMessage != "" {
		b.statusCode = r.StatusCode
		return &b
	}
	if r.StatusCode > 299 {
		return fmt.Errorf("%s (%d)", http.StatusText(r.StatusCode), r.StatusCode)
	}
	return nil
}

type APIError interface {
	error
	StatusCode() int
}

type invalidResponseError struct {
	resp *Response
}

type InvalidResponseError interface {
	error
	Response() *Response
}

func (e *invalidResponseError) Error() string {
	return fmt.Sprintf("braintree returned invalid response - status(%d) body(%s)", e.resp.StatusCode, string(e.resp.Body))
}

func (e *invalidResponseError) Response() *Response {
	return e.resp
}
