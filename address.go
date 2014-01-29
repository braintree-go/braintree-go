package braintree

import (
	"encoding/xml"
)

type Address struct {
	XMLName            xml.Name
	Id                 string `json:"id,omitempty" xml:"id,omitempty"`
	CustomerId         string `json:"customer-id,omitempty" xml:"customer-id,omitempty"`
	FirstName          string `json:"first-name,omitempty" xml:"first-name,omitempty"`
	LastName           string `json:"last-name,omitempty" xml:"last-name,omitempty"`
	Company            string `json:"company,omitempty" xml:"company,omitempty"`
	StreetAddress      string `json:"street-address,omitempty" xml:"street-address,omitempty"`
	ExtendedAddress    string `json:"extended-address,omitempty" xml:"extended-address,omitempty"`
	Locality           string `json:"locality,omitempty" xml:"locality,omitempty"`
	Region             string `json:"region,omitempty" xml:"region,omitempty"`
	PostalCode         string `json:"postal-code,omitempty" xml:"postal-code,omitempty"`
	CountryCodeAlpha2  string `json:"country-code-alpha2,omitempty" xml:"country-code-alpha2,omitempty"`
	CountryCodeAlpha3  string `json:"country-code-alpha3,omitempty" xml:"country-code-alpha3,omitempty"`
	CountryCodeNumeric string `json:"country-code-numeric,omitempty" xml:"country-code-numeric,omitempty"`
	CountryName        string `json:"country-name,omitempty" xml:"country-name,omitempty"`
	CreatedAt          string `json:"created-at,omitempty" xml:"created-at,omitempty"`
	UpdatedAt          string `json:"updated-at,omitempty" xml:"updated-at,omitempty"`
}

type Addresses struct {
	XMLName string     `json:"addresses" xml:"addresses"`
	Address []*Address `json:"address" xml:"address"`
}

