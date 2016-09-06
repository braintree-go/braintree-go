package braintree

import (
	"encoding/xml"
	"time"
)

type Address struct {
	XMLName            xml.Name
	Id                 string     `xml:"id,omitempty" json:"id,omitempty" bson:"id,omitempty"`
	CustomerId         string     `xml:"customer-id,omitempty" json:"customerId,omitempty" bson:"customerId,omitempty"`
	FirstName          string     `xml:"first-name,omitempty" json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName           string     `xml:"last-name,omitempty" json:"lastName,omitempty" bson:"lastName,omitempty"`
	Company            string     `xml:"company,omitempty" json:"company,omitempty" bson:"company,omitempty"`
	StreetAddress      string     `xml:"street-address,omitempty" json:"streetAddress,omitempty" bson:"streetAddress,omitempty"`
	ExtendedAddress    string     `xml:"extended-address,omitempty" json:"extendedAddress,omitempty" bson:"extendedAddress,omitempty"`
	Locality           string     `xml:"locality,omitempty" json:"locality,omitempty" bson:"locality,omitempty"`
	Region             string     `xml:"region,omitempty" json:"region,omitempty" bson:"region,omitempty"`
	PostalCode         string     `xml:"postal-code,omitempty" json:"postalCode,omitempty" bson:"postalCode,omitempty"`
	CountryCodeAlpha2  string     `xml:"country-code-alpha2,omitempty" json:"countryCodeAlpha2,omitempty" bson:"countryCodeAlpha2,omitempty"`
	CountryCodeAlpha3  string     `xml:"country-code-alpha3,omitempty" json:"countryCodeAlpha3,omitempty" bson:"countryCodeAlpha3,omitempty"`
	CountryCodeNumeric string     `xml:"country-code-numeric,omitempty" json:"countryCodeNumeric,omitempty" bson:"countryCodeNumeric,omitempty"`
	CountryName        string     `xml:"country-name,omitempty" json:"countryName,omitempty" bson:"countryName,omitempty"`
	CreatedAt          *time.Time `xml:"created-at,omitempty" json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt          *time.Time `xml:"updated-at,omitempty" json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Addresses struct {
	XMLName string     `xml:"addresses" json:"addresses" bson:"addresses"`
	Address []*Address `xml:"address" json:"address" bson:"address"`
}
