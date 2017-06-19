package braintree

import (
	"github.com/lionelbarrow/braintree-go/customfields"
	"github.com/lionelbarrow/braintree-go/nullable"
)

type Customer struct {
	XMLName            string                    `xml:"customer"`
	Id                 string                    `xml:"id,omitempty"`
	FirstName          string                    `xml:"first-name,omitempty"`
	LastName           string                    `xml:"last-name,omitempty"`
	Company            string                    `xml:"company,omitempty"`
	Email              string                    `xml:"email,omitempty"`
	Phone              string                    `xml:"phone,omitempty"`
	Fax                string                    `xml:"fax,omitempty"`
	Website            string                    `xml:"website,omitempty"`
	CustomFields       customfields.CustomFields `xml:"custom-fields,omitempty"`
	CreditCard         *CreditCard               `xml:"credit-card,omitempty"`
	CreditCards        *CreditCards              `xml:"credit-cards,omitempty"`
	PayPalAccounts     *PayPalAccounts           `xml:"paypal-accounts,omitempty"`
	VenmoAccounts      *VenmoAccounts            `xml:"venmo-accounts,omitempty"`
	ApplePayCards      *ApplePayCards            `xml:"apple-pay-cards,omitempty"`
	PaymentMethodNonce string                    `xml:"payment-method-nonce,omitempty"`
}

// PaymentMethods returns a slice of all PaymentMethods this customer has
func (c *Customer) PaymentMethods() []PaymentMethod {
	var paymentMethods []PaymentMethod
	if c.CreditCards != nil {
		for _, cc := range c.CreditCards.CreditCard {
			paymentMethods = append(paymentMethods, cc)
		}
	}
	if c.PayPalAccounts != nil {
		for _, pp := range c.PayPalAccounts.PayPalAccount {
			paymentMethods = append(paymentMethods, pp)
		}
	}
	if c.VenmoAccounts != nil {
		for _, v := range c.VenmoAccounts.VenmoAccount {
			paymentMethods = append(paymentMethods, v)
		}
	}
	if c.ApplePayCards != nil {
		for _, a := range c.ApplePayCards.ApplePayCard {
			paymentMethods = append(paymentMethods, a)
		}
	}
	return paymentMethods
}

// DefaultCreditCard returns the default credit card, or nil
func (c *Customer) DefaultCreditCard() *CreditCard {
	for _, card := range c.CreditCards.CreditCard {
		if card.Default {
			return card
		}
	}
	return nil
}

// DefaultPaymentMethod returns the default payment method, or nil
func (c *Customer) DefaultPaymentMethod() PaymentMethod {
	if c.CreditCards != nil {
		for _, cc := range c.CreditCards.CreditCard {
			if cc.IsDefault() {
				return cc
			}
		}
	}
	if c.PayPalAccounts != nil {
		for _, pp := range c.PayPalAccounts.PayPalAccount {
			if pp.IsDefault() {
				return pp
			}
		}
	}
	if c.VenmoAccounts != nil {
		for _, v := range c.VenmoAccounts.VenmoAccount {
			if v.IsDefault() {
				return v
			}
		}
	}
	if c.ApplePayCards != nil {
		for _, a := range c.ApplePayCards.ApplePayCard {
			if a.IsDefault() {
				return a
			}
		}
	}
	return nil
}

type CustomerSearchResult struct {
	XMLName           string              `xml:"customers"`
	CurrentPageNumber *nullable.NullInt64 `xml:"current-page-number"`
	PageSize          *nullable.NullInt64 `xml:"page-size"`
	TotalItems        *nullable.NullInt64 `xml:"total-items"`
	Customers         []*Customer         `xml:"customer"`
}
