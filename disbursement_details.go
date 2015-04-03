package braintree

import (
	"encoding/xml"
	"github.com/lionelbarrow/braintree-go/nullable"
)

type DisbursementDetails struct {
	XMLName                        xml.Name           `xml:"disbursement-details"`
	DisbursementDate               string             `xml:"disbursement-date"`
	SettlementAmount               *Decimal           `xml:"settlement-amount"`
	SettlementCurrencyIsoCode      string             `xml:"settlement-currency-iso-code"`
	SettlementCurrencyExchangeRate *Decimal           `xml:"settlement-currency-exchange-rate"`
	FundsHeld                      *nullable.NullBool `xml:"funds-held"`
	Success                        *nullable.NullBool `xml:"success"`
}
