package braintree

import (
	"encoding/xml"
	"time"
)

type Address struct {
	XMLName            xml.Name
	Id                 string     `xml:"id"`
	CustomerId         string     `xml:"customer-id"`
	FirstName          string     `xml:"first-name"`
	LastName           string     `xml:"last-name"`
	Company            string     `xml:"company"`
	StreetAddress      string     `xml:"street-address"`
	ExtendedAddress    string     `xml:"extended-address"`
	Locality           string     `xml:"locality"`
	Region             string     `xml:"region"`
	PostalCode         string     `xml:"postal-code"`
	CountryCodeAlpha2  string     `xml:"country-code-alpha2"`
	CountryCodeAlpha3  string     `xml:"country-code-alpha3"`
	CountryCodeNumeric string     `xml:"country-code-numeric"`
	CountryName        string     `xml:"country-name"`
	CreatedAt          *time.Time `xml:"created-at"`
	UpdatedAt          *time.Time `xml:"updated-at"`
}

type AddressRequest struct {
	XMLName            xml.Name `xml:"address"`
	FirstName          string   `xml:"first-name,omitempty"`
	LastName           string   `xml:"last-name,omitempty"`
	Company            string   `xml:"company,omitempty"`
	StreetAddress      string   `xml:"street-address,omitempty"`
	ExtendedAddress    string   `xml:"extended-address,omitempty"`
	Locality           string   `xml:"locality,omitempty"`
	Region             string   `xml:"region,omitempty"`
	PostalCode         string   `xml:"postal-code,omitempty"`
	CountryCodeAlpha2  string   `xml:"country-code-alpha2,omitempty"`
	CountryCodeAlpha3  string   `xml:"country-code-alpha3,omitempty"`
	CountryCodeNumeric string   `xml:"country-code-numeric,omitempty"`
	CountryName        string   `xml:"country-name,omitempty"`
}

type Addresses struct {
	XMLName string     `xml:"addresses"`
	Address []*Address `xml:"address"`
}
