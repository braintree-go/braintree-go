package braintree

type AddOnList struct {
	XMLName string  `xml:"add-ons" json:"addOns" bson:"addOns"`
	AddOns  []AddOn `xml:"add-on" json:"addOn" bson:"addOn"`
}

type AddOn struct {
	XMLName string `xml:"add-on" json:"addOn" bson:"addOn"`
	Modification
}
