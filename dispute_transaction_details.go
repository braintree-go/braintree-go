package braintree

import "time"

type DisputeTransactionDetails struct {
	XMLName                  string     `xml:"transaction"`
	Amount                   *Decimal   `xml:"amount"`
	CreatedAt                *time.Time `xml:"created-at"`
	Id                       string     `xml:"id"`
	OrderId                  string     `xml:"order-id"`
	PaymentInstrumentSubtype string     `xml:"payment-instrument-subtype"`
	PurchaseOrderNumber      string     `xml:"purchase-order-number"`
}
