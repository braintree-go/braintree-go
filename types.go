package braintree

type Address struct {
	StreetAddress string `xml:"street-address,omitempty"`
	Locality      string `xml:"locality,omitempty"`
	Region        string `xml:"region,omitempty"`
	PostalCode    string `xml:"postal-code,omitempty"`
}
