package braintree

import "encoding/xml"

type PaymentMethodGateway struct {
	*Braintree
}

type PaymentMethodRequest struct {
	XMLName            xml.Name `xml:"payment-method"`
	CustomerId         string   `xml:"customer-id,omitempty"`
	PaymentMethodNonce string   `xml:"payment-method-nonce,omitempty"`
}

func (g *PaymentMethodGateway) Create(paymentMethodRequest *PaymentMethodRequest) (PaymentMethod, error) {
	resp, err := g.execute("POST", "payment_methods", paymentMethodRequest)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}
