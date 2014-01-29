package braintree

type AddOnList struct {
	XMLName string  `json:"add-ons" xml:"add-ons"`
	AddOns  []AddOn `json:"add-on" xml:"add-on"`
}

type AddOn struct {
	XMLName string `json:"add-on" xml:"add-on"`
	Modification
}

