package braintree

import (
	"encoding/xml"
	"github.com/CompleteSet/braintree-go/date"
)

type Disbursement struct {
	XMLName            xml.Name         `xml:"disbursement" json:"disbursement" bson:"disbursement"`
	Id                 string           `xml:"id" json:"id" bson:"id"`
	ExceptionMessage   string           `xml:"exception-message" json:"exceptionMessage" bson:"exceptionMessage"`
	DisbursementDate   *date.Date       `xml:"disbursement-date" json:"disbursementDate" bson:"disbursementDate"`
	FollowUpAction     string           `xml:"follow-up-action" json:"followUpAction" bson:"followUpAction"`
	Success            bool             `xml:"success" json:"success" bson:"success"`
	Retry              bool             `xml:"retry" json:"retry" bson:"retry"`
	Amount             *Decimal         `xml:"amount" json:"amount" bson:"amount"`
	MerchantAccount    *MerchantAccount `xml:"merchant-account" json:"merchantAccount" bson:"merchantAccount"`
	CurrencyIsoCode    string           `xml:"currency-iso-code" json:"currencyIsoCode" bson:"currencyIsoCode"`
	SubmerchantAccount bool             `xml:"sub-merchant-account" json:"subMerchantAccount" bson:"subMerchantAccount"`
	Status             string           `xml:"status" json:"status" bson:"status"`
	TransactionIds     []string         `xml:"transaction-ids>item" json:"transactionIds.item" bson:"transactionIds.item"`
}

const (
	// Exception messages
	BankRejected         = "bank_rejected"
	InsufficientFunds    = "insuffient_funds"
	AccountNotAuthorized = "account_not_authorized"

	// Followup actions
	ContactUs                = "contact_us"
	UpdateFundingInformation = "update_funding_information"
	None                     = "none"
)

func (d *Disbursement) Transactions(g *TransactionGateway) (*TransactionSearchResult, error) {
	query := new(SearchQuery)
	f := query.AddMultiField("ids")
	f.Items = d.TransactionIds

	result, err := g.Search(query)
	if err != nil {
		return nil, err
	}

	return result, err
}
