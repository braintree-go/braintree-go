package braintree

import "encoding/xml"

type Transaction struct {
	XMLName           string              `xml:"transaction"`
	Id                string              `xml:"id,omitempty"`
	Status            string              `xml:"status,omitempty"`
	Type              string              `xml:"type,omitempty"`
	Amount            float64             `xml:"amount"`
	OrderId           string              `xml:"order-id,omitempty"`
	MerchantAccountId string              `xml:"merchant-account-id,omitempty"`
	CreditCard        *CreditCard         `xml:"credit-card,omitempty"`
	Customer          *Customer           `xml:"customer,omitempty"`
	BillingAddress    *Address            `xml:"billing,omitempty"`
	ShippingAddress   *Address            `xml:"shipping,omitempty"`
	Options           *TransactionOptions `xml:"options,omitempty"`
}

func (this Transaction) ToXML() ([]byte, error) {
	xml, err := xml.Marshal(this)
	if err != nil {
		return []byte{}, err
	}
	return xml, nil
}

type TransactionOptions struct {
	SubmitForSettlement              bool `xml:"submit-for-settlement,omitempty"`
	StoreInVault                     bool `xml:"store-in-vault,omitempty"`
	AddBillingAddressToPaymentMethod bool `xml:"add-billing-address-to-payment-method,omitempty"`
	StoreShippingAddressInVault      bool `xml:"store-shipping-address-in-vault,omitempty"`
}
