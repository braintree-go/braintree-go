package braintree

type DiscountList struct {
	XMLName   string     `json:"discounts" xml:"discounts"`
	Discounts []Discount `json:"discount" xml:"discount"`
}

type Discount struct {
	XMLName string `json:"discount" xml:"discount"`
	Modification
}
