package braintree

import (
	"encoding/xml"
	"github.com/brianpowell/braintree-go/nullable"
)

type DisbursementDetails struct {
	XMLName                        xml.Name           `xml:"disbursement-details" json:"disbursementDetails" bson:"disbursementDetails"`
	DisbursementDate               string             `xml:"disbursement-date" json:"disbursementDate" bson:"disbursementDate"`
	SettlementAmount               *Decimal           `xml:"settlement-amount" json:"settlementAmount" bson:"settlementAmount"`
	SettlementCurrencyIsoCode      string             `xml:"settlement-currency-iso-code" json:"settlementCurrencyIsoCode" bson:"settlementCurrencyIsoCode"`
	SettlementCurrencyExchangeRate *Decimal           `xml:"settlement-currency-exchange-rate" json:"settlementCurrencyExchangeRate" bson:"settlementCurrencyExchangeRate"`
	FundsHeld                      *nullable.NullBool `xml:"funds-held" json:"fundsHeld" bson:"fundsHeld"`
	Success                        *nullable.NullBool `xml:"success" json:"success" bson:"success"`
}
