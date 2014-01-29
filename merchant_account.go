package braintree

type MerchantAccount struct {
	XMLName                 string                         `json:"merchant-account,omitempty" xml:"merchant-account,omitempty"`
	Id                      string                         `json:"id,omitempty" xml:"id,omitempty"`
	MasterMerchantAccountId string                         `json:"master-merchant-account-id,omitempty" xml:"master-merchant-account-id,omitempty"`
	TOSAccepted             bool                           `json:"tos_accepted,omitempty" xml:"tos_accepted,omitempty"`
	Individual              *MerchantAccountPerson         `json:"individual,omitempty" xml:"individual,omitempty"`
	Business                *MerchantAccountBusiness       `json:"business,omitempty" xml:"business,omitempty"`
	FundingOptions          *MerchantAccountFundingOptions `json:"funding,omitempty" xml:"funding,omitempty"`
}

type MerchantAccountPerson struct {
	FirstName   string   `json:"first-name,omitempty" xml:"first-name,omitempty"`
	LastName    string   `json:"last-name,omitempty" xml:"last-name,omitempty"`
	Email       string   `json:"email,omitempty" xml:"email,omitempty"`
	Phone       string   `json:"phone,omitempty" xml:"phone,omitempty"`
	DateOfBirth string   `json:"date-of-birth,omitempty" xml:"date-of-birth,omitempty"`
	SSN         string   `json:"ssn,omitempty" xml:"ssn,omitempty"`
	Address     *Address `json:"address,omitempty" xml:"address,omitempty"`
}

type MerchantAccountBusiness struct {
	LegalName string   `json:"legal-name,omitempty" xml:"legal-name,omitempty"`
	DbaName   string   `json:"dba-name,omitempty" xml:"dba-name,omitempty"`
	TaxId     string   `json:"tax-id,omitempty" xml:"tax-id,omitempty"`
	Address   *Address `json:"address,omitempty" xml:"address,omitempty"`
}

type MerchantAccountFundingOptions struct {
	Destination   string `json:"destination,omitempty" xml:"destination,omitempty"`
	Email         string `json:"email,omitempty" xml:"email,omitempty"`
	MobilePhone   string `json:"mobile-phone,omitempty" xml:"mobile-phone,omitempty"`
	AccountNumber string `json:"account-number,omitempty" xml:"account-number,omitempty"`
	RoutingNumber string `json:"routing-number,omitempty" xml:"routing-number,omitempty"`
}

const (
	FUNDING_DEST_BANK         = "bank"
	FUNDING_DEST_MOBILE_PHONE = "mobile_phone"
	FUNDING_DEST_EMAIL        = "email"
)

