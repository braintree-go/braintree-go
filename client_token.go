package braintree

type ClientTokenRequest struct {
	XMLName    string `xml:"client-token" json:"clientToken" bson:"clientToken"`
	CustomerID string `xml:"customerId,omitempty" json:"customerId,omitempty" bson:"customerId,omitempty"`
	Version    int    `xml:"version" json:"version" bson:"version"`
}

type clientToken struct {
	ClientToken string `xml:"value" json:"value" bson:"value"`
}
