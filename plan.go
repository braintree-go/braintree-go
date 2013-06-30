package braintree

type Plan struct {
	XMLName string `xml:"plan"`
	Id      string `xml:"id"`
}

type Plans struct {
	XMLName string  `xml:"plans"`
	Plan    []*Plan `xml:"plan"`
}
