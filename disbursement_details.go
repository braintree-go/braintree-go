package braintree

import (
	"encoding/xml"
)

type DisbursementDetails struct {
	XMLName                        xml.Name `xml:"disbursement-details"`
	DisbursementDate               string   `xml:"disbursement-date"`
	SettlementAmount               string   `xml:"settlement-amount"` // float64
	SettlementCurrencyIsoCode      string   `xml:"settlement-currency-iso-code"`
	SettlementCurrencyExchangeRate string   `xml:"settlement-currency-exchange-rate"` // float64
	FundsHeld                      string   `xml:"funds-held"`                        // bool
	Success                        string   `xml:"success"`                           // bool
}
