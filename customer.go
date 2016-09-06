package braintree

import "github.com/CompleteSet/braintree-go/nullable"

type Customer struct {
	XMLName     string       `xml:"customer" json:"customer" bson:"customer"`
	Id          string       `xml:"id,omitempty" json:"id,omitempty" bson:"id,omitempty"`
	FirstName   string       `xml:"first-name,omitempty" json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName    string       `xml:"last-name,omitempty" json:"lastName,omitempty" bson:"lastName,omitempty"`
	Company     string       `xml:"company,omitempty" json:"company,omitempty" bson:"company,omitempty"`
	Email       string       `xml:"email,omitempty" json:"email,omitempty" bson:"email,omitempty"`
	Phone       string       `xml:"phone,omitempty" json:"phone,omitempty" bson:"phone,omitempty"`
	Fax         string       `xml:"fax,omitempty" json:"fax,omitempty" bson:"fax,omitempty"`
	Website     string       `xml:"website,omitempty" json:"website,omitempty" bson:"website,omitempty"`
	CreditCard  *CreditCard  `xml:"credit-card,omitempty" json:"creditCard,omitempty" bson:"creditCard,omitempty"`
	CreditCards *CreditCards `xml:"credit-cards,omitempty" json:"creditCards,omitempty" bson:"creditCards,omitempty"`
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
	XMLName           string              `xml:"customers" json:"customers" bson:"customers"`
	CurrentPageNumber *nullable.NullInt64 `xml:"current-page-number" json:"currentPageNumber" bson:"currentPageNumber"`
	PageSize          *nullable.NullInt64 `xml:"page-size" json:"pageSize" bson:"pageSize"`
	TotalItems        *nullable.NullInt64 `xml:"total-items" json:"totalItems" bson:"totalItems"`
	Customers         []*Customer         `xml:"customer" json:"customer" bson:"customer"`
}
