package braintree

import "github.com/lionelbarrow/braintree-go/nullable"

type Customer struct {
	XMLName        string          `xml:"customer"`
	Id             string          `xml:"id,omitempty"`
	FirstName      string          `xml:"first-name,omitempty"`
	LastName       string          `xml:"last-name,omitempty"`
	Company        string          `xml:"company,omitempty"`
	Email          string          `xml:"email,omitempty"`
	Phone          string          `xml:"phone,omitempty"`
	Fax            string          `xml:"fax,omitempty"`
	Website        string          `xml:"website,omitempty"`
	CreditCard     *CreditCard     `xml:"credit-card,omitempty"`
	CreditCards    *CreditCards    `xml:"credit-cards,omitempty"`
	PaypalAccount  *PaypalAccount  `xml:"paypal-account,omitempty"`
	PaypalAccounts *PaypalAccounts `xml:"paypal-accounts,omitempty"`
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

type CustomerSearchResult struct {
	XMLName           string              `xml:"customers"`
	CurrentPageNumber *nullable.NullInt64 `xml:"current-page-number"`
	PageSize          *nullable.NullInt64 `xml:"page-size"`
	TotalItems        *nullable.NullInt64 `xml:"total-items"`
	Customers         []*Customer         `xml:"customer"`
}
