package braintree

import (
	"encoding/xml"
)

type Address struct {
	XMLName            xml.Name
	Id                 string `xml:"id,omitempty" json:"-"`
	CustomerId         string `xml:"customer-id,omitempty" json:"-"`
	FirstName          string `xml:"first-name,omitempty" json:"-"`
	LastName           string `xml:"last-name,omitempty" json:"-"`
	Company            string `xml:"company,omitempty" json:"-"`
	StreetAddress      string `xml:"street-address,omitempty" json:"street_address"`
	ExtendedAddress    string `xml:"extended-address,omitempty" json:"-"`
	Locality           string `xml:"locality,omitempty" json:"locality"`
	Region             string `xml:"region,omitempty" json:"region"`
	PostalCode         string `xml:"postal-code,omitempty" json:"postal_code"`
	CountryCodeAlpha2  string `xml:"country-code-alpha2,omitempty" json:"-"`
	CountryCodeAlpha3  string `xml:"country-code-alpha3,omitempty" json:"-"`
	CountryCodeNumeric string `xml:"country-code-numeric,omitempty" json:"-"`
	CountryName        string `xml:"country-name,omitempty" json:"-"`
	CreatedAt          string `xml:"created-at,omitempty" json:"-"`
	UpdatedAt          string `xml:"updated-at,omitempty" json:"-"`
}

type Addresses struct {
	XMLName string     `xml:"addresses"`
	Address []*Address `xml:"address"`
}
