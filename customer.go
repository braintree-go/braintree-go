package braintree

type Customer struct {
	XMLName     string       `json:"customer" xml:"customer"`
	Id          string       `json:"id,omitempty" xml:"id,omitempty"`
	FirstName   string       `json:"first-name,omitempty" xml:"first-name,omitempty"`
	LastName    string       `json:"last-name,omitempty" xml:"last-name,omitempty"`
	Company     string       `json:"company,omitempty" xml:"company,omitempty"`
	Email       string       `json:"email,omitempty" xml:"email,omitempty"`
	Phone       string       `json:"phone,omitempty" xml:"phone,omitempty"`
	Fax         string       `json:"fax,omitempty" xml:"fax,omitempty"`
	Website     string       `json:"website,omitempty" xml:"website,omitempty"`
	CreditCard  *CreditCard  `json:"credit-card,omitempty" xml:"credit-card,omitempty"`
	CreditCards *CreditCards `json:"credit-cards,omitempty" xml:"credit-cards,omitempty"`
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

