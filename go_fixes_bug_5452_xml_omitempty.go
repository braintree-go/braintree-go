// +build !go1.8

package braintree

import "encoding/xml"

// The functions in this file are required because of a bug in versions of go prior to go1.8.
// The bug was reported at https://github.com/golang/go/issues/5452 and fixed in
// https://github.com/golang/go/commit/daa121167b6ce630aba00195f1c3872cda39a50c.
//
// In versions prior to go1.8 the XML encoder did not include pointer fields that were non-nil
// if the field pointed to a value that was the default value for the pointed to type.
//
// To serialize the bool false value when it is set on `VerifyCard`, we must manually control
// if it is serialized or not.

// MarshalXML custom serialization for CreditCardOptions.
func (cco *CreditCardOptions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if cco.VerifyCard == nil {
		type excludeVerifyCard struct {
			VerifyCard                    *bool  `xml:"-"`
			VenmoSDKSession               string `xml:"venmo-sdk-session,omitempty"`
			MakeDefault                   bool   `xml:"make-default,omitempty"`
			FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty"`
			VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty"`
			UpdateExistingToken           string `xml:"update-existing-token,omitempty"`
		}
		return e.EncodeElement(
			excludeVerifyCard{
				VerifyCard:                    cco.VerifyCard,
				VenmoSDKSession:               cco.VenmoSDKSession,
				MakeDefault:                   cco.MakeDefault,
				FailOnDuplicatePaymentMethod:  cco.FailOnDuplicatePaymentMethod,
				VerificationMerchantAccountId: cco.VerificationMerchantAccountId,
				UpdateExistingToken:           cco.UpdateExistingToken,
			},
			start,
		)
	} else {
		type includeVerifyCard struct {
			VerifyCard                    *bool  `xml:"verify-card"`
			VenmoSDKSession               string `xml:"venmo-sdk-session,omitempty"`
			MakeDefault                   bool   `xml:"make-default,omitempty"`
			FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty"`
			VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty"`
			UpdateExistingToken           string `xml:"update-existing-token,omitempty"`
		}
		return e.EncodeElement(
			includeVerifyCard{
				VerifyCard:                    cco.VerifyCard,
				VenmoSDKSession:               cco.VenmoSDKSession,
				MakeDefault:                   cco.MakeDefault,
				FailOnDuplicatePaymentMethod:  cco.FailOnDuplicatePaymentMethod,
				VerificationMerchantAccountId: cco.VerificationMerchantAccountId,
				UpdateExistingToken:           cco.UpdateExistingToken,
			},
			start,
		)
	}
}

// MarshalXML custom serialization for PaymentMethodRequestOptions.
func (pmo *PaymentMethodRequestOptions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if pmo.VerifyCard == nil {
		type excludeVerifyCard struct {
			MakeDefault                   bool   `xml:"make-default,omitempty"`
			FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty"`
			VerifyCard                    *bool  `xml:"-"`
			VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty"`
		}
		return e.EncodeElement(
			excludeVerifyCard{
				MakeDefault:                   pmo.MakeDefault,
				FailOnDuplicatePaymentMethod:  pmo.FailOnDuplicatePaymentMethod,
				VerifyCard:                    pmo.VerifyCard,
				VerificationMerchantAccountId: pmo.VerificationMerchantAccountId,
			},
			start,
		)
	} else {
		type includeVerifyCard struct {
			MakeDefault                   bool   `xml:"make-default,omitempty"`
			FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty"`
			VerifyCard                    *bool  `xml:"verify-card"`
			VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty"`
		}
		return e.EncodeElement(
			includeVerifyCard{
				MakeDefault:                   pmo.MakeDefault,
				FailOnDuplicatePaymentMethod:  pmo.FailOnDuplicatePaymentMethod,
				VerifyCard:                    pmo.VerifyCard,
				VerificationMerchantAccountId: pmo.VerificationMerchantAccountId,
			},
			start,
		)
	}
}

// MarshalXML custom serialization for ClientTokenRequestOptions.
func (ctro *ClientTokenRequestOptions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ctro.VerifyCard == nil {
		type excludeVerifyCard struct {
			FailOnDuplicatePaymentMethod bool  `xml:"fail-on-duplicate-payment-method,omitempty"`
			MakeDefault                  bool  `xml:"make-default,omitempty"`
			VerifyCard                   *bool `xml:"-"`
		}
		return e.EncodeElement(
			excludeVerifyCard{
				FailOnDuplicatePaymentMethod: ctro.FailOnDuplicatePaymentMethod,
				MakeDefault:                  ctro.MakeDefault,
				VerifyCard:                   ctro.VerifyCard,
			},
			start,
		)
	}
	type includeVerifyCard struct {
		FailOnDuplicatePaymentMethod bool  `xml:"fail-on-duplicate-payment-method,omitempty"`
		MakeDefault                  bool  `xml:"make-default,omitempty"`
		VerifyCard                   *bool `xml:"verify-card"`
	}
	return e.EncodeElement(
		includeVerifyCard{
			FailOnDuplicatePaymentMethod: ctro.FailOnDuplicatePaymentMethod,
			MakeDefault:                  ctro.MakeDefault,
			VerifyCard:                   ctro.VerifyCard,
		},
		start,
	)
}
