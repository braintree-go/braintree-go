package braintree

import "time"

type PaymentMethod struct {
	CustomerId         string                `xml:"customer-id,omitempty" json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	Token              string                `xml:"token,omitempty" json:"token,omitempty" bson:"token,omitempty"`
	PaymentMethodNonce string                `xml:"payment-method-nonce,omitempty" json:"payment_method_nonce,omitempty" bson:"payment_method_nonce,omitempty"`
	ExpirationDate     string                `xml:"expiration-date,omitempty" json:"expiration_date,omitempty" bson:"expiration_date,omitempty"`
	ExpirationMonth    string                `xml:"expiration-month,omitempty" json:"expiration_month,omitempty" bson:"expiration_month,omitempty"`
	ExpirationYear     string                `xml:"expiration-year,omitempty" json:"expiration_year,omitempty" bson:"expiration_year,omitempty"`
	Bin                string                `xml:"bin,omitempty" json:"bin,omitempty" bson:"bin,omitempty"`
	BillingAgreementID string                `xml:"billing-agreeement-id,omitempty" json:"billing_agreeement_id,omitempty" bson:"billing_agreeement_id,omitempty"`
	CardType           string                `xml:"card-type,omitempty" json:"card_type,omitempty" bson:"card_type,omitempty"`
	CardholderName     string                `xml:"cardholder-name,omitempty" json:"cardholder_name,omitempty" bson:"cardholder_name,omitempty"`
	Country            string                `xml:"country-of-issuance,omitempty" json:"country_of_issuance,omitempty" bson:"country_of_issuance,omitempty"`
	Default            bool                  `xml:"default,omitempty" json:"default,omitempty" bson:"default,omitempty"`
	Expired            bool                  `xml:"expired,omitempty" json:"expired,omitempty" bson:"expired,omitempty"`
	Email              string                `xml:"email,omitempty" json:"email,omitempty" bson:"email,omitempty"`
	GoogleTransId      string                `xml:"google-transaction-id,omitempty" json:"google_transaction_id,omitempty" bson:"google_transaction_id,omitempty"`
	ImageURL           string                `xml:"image-url,omitempty" json:"image_url,omitempty" bson:"image_url,omitempty"`
	IssuingBank        string                `xml:"issuing-bank,omitempty" json:"issuing_bank,omitempty" bson:"issuing_bank,omitempty"`
	Last4              string                `xml:"last-4,omitempty" json:"last_4,omitempty" bson:"last_4,omitempty"`
	MaskedNumber       string                `xml:"masked-number,omitempty" json:"masked_number,omitempty" bson:"masked_number,omitempty"`
	Options            *PaymentMethodOptions `xml:"options,omitempty" json:"options,omitempty" bson:"options,omitempty"`
	PaymentInstName    string                `xml:"payment-instrument-name,omitempty" json:"payment_instrument_name,omitempty" bson:"payment_instrument_name,omitempty"`
	ProductId          string                `xml:"product-id,omitempty" json:"product_id,omitempty" bson:"product_id,omitempty"`
	SourceDesc         string                `xml:"source-description,omitempty" json:"source_description,omitempty" bson:"source_description,omitempty"`
	SourceLast4        string                `xml:"source-card-last-4,omitempty" json:"source_card_last_4,omitempty" bson:"source_card_last_4,omitempty"`
	SourceType         string                `xml:"source-card-type,omitempty" json:"source_card_type,omitempty" bson:"source_card_type,omitempty"`
	UUID               string                `xml:"unique-nubmer-identifier,omitempty" json:"unique_number_identifier,omitempty" bson:"unique_number_identifier,omitempty"`
	VirtualLast4       string                `xml:"virtual-card-last-4,omitempty" json:"virtual_card_last_4,omitempty" bson:"virtual_card_last_4,omitempty"`
	VirtualType        string                `xml:"virtual-card-type,omitempty" json:"virtual_card_type,omitempty" bson:"virtual_card_type,omitempty"`
	BillingAddress     *Address              `xml:"billing-address,omitempty" json:"billing_address,omitempty" bson:"billing_address,omitempty"`
	Created_At         *time.Time            `xml:"created-at,omitempty" json:"created_at,omitempty" bson:"created_at,omitempty"`
	Updated_At         *time.Time            `xml:"updated-at,omitempty" json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type PaymentMethodOptions struct {
	VerifyCard                    bool   `xml:"verify-card,omitempty" json:"verify_card,omitempty" bson:"verify_card,omitempty"`
	MakeDefault                   bool   `xml:"make-default,omitempty" json:"make_default,omitempty" bson:"make_default,omitempty"`
	FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty" json:"fail_on_duplicate_payment_method,omitempty" bson:"fail_on_duplicate_payment_method,omitempty"`
	VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty" json:"verification_merchant_account_id,omitempty" bson:"verification_merchant_account_id,omitempty"`
}

type PaymentMethods struct {
	ID             string           `xml:"id,omitempty" json:"id,omitempty" bson:"_id,omitempty"`
	PaymentMethods []*PaymentMethod `xml:"payment-methods" json:"payment_methods" bson:"payment_methods"`
}
