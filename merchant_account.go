package braintree

type MerchantAccount struct {
	XMLName                 string                         `xml:"merchant-account,omitempty" json:"-"`
	Id                      string                         `xml:"id,omitempty" json:"-"`
	MasterMerchantAccountId string                         `xml:"master-merchant-account-id,omitempty" json:"-"`
	TOSAccepted             bool                           `xml:"tos_accepted,omitempty" json: "-"`
	Individual              *MerchantAccountPerson         `xml:"individual,omitempty" json:"individual"`
	Business                *MerchantAccountBusiness       `xml:"business,omitempty" json:"business"`
	FundingOptions          *MerchantAccountFundingOptions `xml:"funding,omitempty" json:"funding"`
	Status                  string                         `xml:"status,omitempty"  json:"status,omitempty"`
}

type MerchantAccountPerson struct {
	FirstName   string   `xml:"first-name,omitempty" json:"first_name,omitempty"`
	LastName    string   `xml:"last-name,omitempty" json:"last_name,omitempty"`
	Email       string   `xml:"email,omitempty" json:"email,omitempty"`
	Phone       string   `xml:"phone,omitempty" json:"phone,omitempty"`
	DateOfBirth string   `xml:"date-of-birth,omitempty" json:"dob,omitempty"`
	SSN         string   `xml:"ssn,omitempty" json:"ssn,omitempty"`
	Address     *Address `xml:"address,omitempty" json:"address,omitempty"`
}

type MerchantAccountBusiness struct {
	LegalName string   `xml:"legal-name,omitempty" json:"legal_name,omitempty"`
	DbaName   string   `xml:"dba-name,omitempty" json:"dba_name,omitempty"`
	TaxId     string   `xml:"tax-id,omitempty" json:"tax_id,omitempty"`
	Address   *Address `xml:"address,omitempty" json:"address,omitempty"`
}

type MerchantAccountFundingOptions struct {
	Destination   string `xml:"destination,omitempty" json:"destination,omitempty"`
	Email         string `xml:"email,omitempty" json:"email,omitempty"`
	MobilePhone   string `xml:"mobile-phone,omitempty" json:"mobile_phone,omitempty"`
	AccountNumber string `xml:"account-number,omitempty" json:"account_number,omitempty"`
	RoutingNumber string `xml:"routing-number,omitempty" json:"routing_number,omitempty"`
}

const (
	FUNDING_DEST_BANK         = "bank"
	FUNDING_DEST_MOBILE_PHONE = "mobile_phone"
	FUNDING_DEST_EMAIL        = "email"
)
