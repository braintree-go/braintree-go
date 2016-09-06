package braintree

type MerchantAccount struct {
	XMLName                 string                         `xml:"merchant-account,omitempty" json:"merchantAccount,omitempty" bson:"merchantAccount,omitempty"`
	Id                      string                         `xml:"id,omitempty" json:"id,omitempty" bson:"id,omitempty"`
	MasterMerchantAccountId string                         `xml:"master-merchant-account-id,omitempty" json:"masterMerchantAccountId,omitempty" bson:"masterMerchantAccountId,omitempty"`
	TOSAccepted             bool                           `xml:"tos_accepted,omitempty" json:"tos_accepted,omitempty" bson:"tos_accepted,omitempty"`
	Individual              *MerchantAccountPerson         `xml:"individual,omitempty" json:"individual,omitempty" bson:"individual,omitempty"`
	Business                *MerchantAccountBusiness       `xml:"business,omitempty" json:"business,omitempty" bson:"business,omitempty"`
	FundingOptions          *MerchantAccountFundingOptions `xml:"funding,omitempty" json:"funding,omitempty" bson:"funding,omitempty"`
	Status                  string                         `xml:"status,omitempty" json:"status,omitempty" bson:"status,omitempty"`
}

type MerchantAccountPerson struct {
	FirstName   string   `xml:"first-name,omitempty" json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName    string   `xml:"last-name,omitempty" json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email       string   `xml:"email,omitempty" json:"email,omitempty" bson:"email,omitempty"`
	Phone       string   `xml:"phone,omitempty" json:"phone,omitempty" bson:"phone,omitempty"`
	DateOfBirth string   `xml:"date-of-birth,omitempty" json:"dateOfBirth,omitempty" bson:"dateOfBirth,omitempty"`
	SSN         string   `xml:"ssn,omitempty" json:"ssn,omitempty" bson:"ssn,omitempty"`
	Address     *Address `xml:"address,omitempty" json:"address,omitempty" bson:"address,omitempty"`
}

type MerchantAccountBusiness struct {
	LegalName string   `xml:"legal-name,omitempty" json:"legalName,omitempty" bson:"legalName,omitempty"`
	DbaName   string   `xml:"dba-name,omitempty" json:"dbaName,omitempty" bson:"dbaName,omitempty"`
	TaxId     string   `xml:"tax-id,omitempty" json:"taxId,omitempty" bson:"taxId,omitempty"`
	Address   *Address `xml:"address,omitempty" json:"address,omitempty" bson:"address,omitempty"`
}

type MerchantAccountFundingOptions struct {
	Destination   string `xml:"destination,omitempty" json:"destination,omitempty" bson:"destination,omitempty"`
	Email         string `xml:"email,omitempty" json:"email,omitempty" bson:"email,omitempty"`
	MobilePhone   string `xml:"mobile-phone,omitempty" json:"mobilePhone,omitempty" bson:"mobilePhone,omitempty"`
	AccountNumber string `xml:"account-number,omitempty" json:"accountNumber,omitempty" bson:"accountNumber,omitempty"`
	RoutingNumber string `xml:"routing-number,omitempty" json:"routingNumber,omitempty" bson:"routingNumber,omitempty"`
}

const (
	FUNDING_DEST_BANK         = "bank"
	FUNDING_DEST_MOBILE_PHONE = "mobile_phone"
	FUNDING_DEST_EMAIL        = "email"
)
