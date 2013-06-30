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

		// Enable for debug logging
		// fmt.Println("RESP:", string(r.Body))
	}
	return nil
}

func (r *Response) apiError() error {
	var b BraintreeError
	xml.Unmarshal(r.Body, &b)
	if b.ErrorMessage != "" {
		return &b
	}
	if r.StatusCode > 299 {
		return fmt.Errorf("%s (%d)", http.StatusText(r.StatusCode), r.StatusCode)
	}
	return nil
}

type BraintreeError struct {
	ErrorMessage string `xml:"message"`
}

func (e *BraintreeError) Error() string {
	return e.ErrorMessage
}

type InvalidResponseError struct {
	*Response
}

func (e *InvalidResponseError) Error() string {
	return fmt.Sprintf("braintree returned invalid response (%d)", e.StatusCode)
}
