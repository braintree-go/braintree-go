package braintree

import "encoding/xml"

type PaymentMethodNonce struct {
	XMLName     xml.Name `xml:"payment-method-nonce"`
	Type        string   `xml:"type"`
	Nonce       string   `xml:"nonce"`
	Description string   `xml:"description"`
}

type PaymentMethodRequest struct {
	XMLName            xml.Name                     `xml:"payment-method"`
	CustomerId         string                       `xml:"customer-id,omitempty"`
	Token              string                       `xml:"token,omitempty"`
	PaymentMethodNonce string                       `xml:"payment-method-nonce,omitempty"`
	Options            *PaymentMethodRequestOptions `xml:"options,omitempty"`
}

type PaymentMethodRequestOptions struct {
	MakeDefault                   bool   `xml:"make-default,omitempty"`
	FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty"`
	VerifyCard                    *bool  `xml:"verify-card,omitempty"`
	VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty"`
}
