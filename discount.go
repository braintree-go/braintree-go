package braintree

type DiscountList struct {
	XMLName   string     `xml:"discounts" json:"discounts" bson:"discounts"`
	Discounts []Discount `xml:"discount" json:"discount" bson:"discount"`
}

type Discount struct {
	XMLName string `xml:"discount" json:"discount" bson:"discount"`
	Modification
}
