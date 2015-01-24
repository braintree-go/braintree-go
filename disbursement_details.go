package braintree

import (
	"encoding/xml"
)

type DisbursementDetails struct {
	XMLName                        xml.Name     `xml:"disbursement-details"`
	DisbursementDate               string       `xml:"disbursement-date"`
	SettlementAmount               *NullFloat64 `xml:"settlement-amount"`
	SettlementCurrencyIsoCode      string       `xml:"settlement-currency-iso-code"`
	SettlementCurrencyExchangeRate *NullFloat64 `xml:"settlement-currency-exchange-rate"`
	FundsHeld                      *NullBool    `xml:"funds-held"`
	Success                        *NullBool    `xml:"success"`
}
