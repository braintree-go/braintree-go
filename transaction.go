package braintree

import (
	"time"

	"github.com/processout/braintree-go/nullable"
)

type Transaction struct {
	XMLName    string `xml:"transaction"`
	Id         string `xml:"id,omitempty"`
	CustomerID string `xml:"customer-id,omitempty"`

	Status        string         `xml:"status,omitempty"`
	StatusHistory *StatusHistory `xml:"status-history,omitempty"`

	Type                 string                   `xml:"type,omitempty"`
	Amount               *Decimal                 `xml:"amount"`
	CurrencyISOCode      string                   `xml:"currency-iso-code,omitempty"`
	Channel              string                   `xml:"channel,omitempty"`
	Tax                  *Decimal                 `xml:"tax-amount"`
	OrderId              string                   `xml:"order-id,omitempty"`
	PaymentMethodToken   string                   `xml:"payment-method-token,omitempty"`
	PaymentMethodNonce   string                   `xml:"payment-method-nonce,omitempty"`
	MerchantAccountId    string                   `xml:"merchant-account-id,omitempty"`
	PlanId               string                   `xml:"plan-id,omitempty"`
	CreditCard           *CreditCard              `xml:"credit-card,omitempty"`
	Customer             *Customer                `xml:"customer,omitempty"`
	BillingAddress       *Address                 `xml:"billing,omitempty"`
	ShippingAddress      *Address                 `xml:"shipping,omitempty"`
	Options              *TransactionOptions      `xml:"options,omitempty"`
	ThreeDSecurePassThru *ThreeDSecurePassThrough `xml:"three-d-secure-pass-thru,omitempty"`
	ExternalVault        *ExternalVault           `xml:"external-vault,omitempty"`
	FeeAmount            *Decimal                 `xml:"transaction-fee-amount"`

	ServiceFeeAmount      *Decimal `xml:"service-fee-amount,omitempty"`
	ServiceFeeTransaction string   `xml:"transaction-fee-currency-iso-code,omitempty"`

	CreatedAt           *time.Time           `xml:"created-at,omitempty"`
	UpdatedAt           *time.Time           `xml:"updated-at,omitempty"`
	DisbursementDetails *DisbursementDetails `xml:"disbursement-details,omitempty"`

	RefundId  string  `xml:"refund-id,omitempty"`
	RefundIds *Refund `xml:"refund-ids,omitempty"`

	Disputes []Dispute `xml:"disputes,omitempty"`

	AVSErrorCode                 string `xml:"avs-error-response-code,omitempty"`
	AVSPostalCodeResponseCode    string `xml:"avs-postal-code-response-code,omitempty"`
	AVSStreetAddressResponseCode string `xml:"avs-street-address-response-code,omitempty"`
	CVVResponseCode              string `xml:"cvv-response-code,omitempty"`

	RefundedTransactionId      *string `xml:"refunded-transaction-id,omitempty"`
	ProcessorResponseCode      string  `xml:"processor-response-code,omitempty"`
	ProcessorResponseText      string  `xml:"processor-response-text,omitempty"`
	ProcessorAuthorizationCode string  `xml:"processor-authorization-code,omitempty"`
	SettlementBatchId          string  `xml:"settlement-batch-id,omitempty"`
	NetworkTransactionID       string  `xml:"network-transaction-id,omitempty"`

	SubscriptionID        string        `xml:"subscription-id,omitempty"`
	Descriptor            *Descriptor   `xml:"descriptor,omitempty"`
	PaymentInstrumentType string        `xml:"payment-instrument-type,omitempty"`
	TransactionSource     string        `xml:"transaction-source,omitempty"`
	Recurring             bool          `xml:"recurring,omitempty"`
	ARN                   string        `xml:"acquirer-reference-number,omitempty" json:"acquirer_reference_number,omitempty"`
	ApplePayCard          *ApplePayCard `xml:"apple_pay_card,omitempty"`
}

type Refund struct {
	IDs []string `xml:"item"`
}

type StatusHistory struct {
	Events []StatusEvent `xml:"status-event"`
}

type ExternalVault struct {
	PreviousNetworkTransactionID string `xml:"previous-network-transaction-id,omitempty"`
	Status                       string `xml:"status,omitempty"`
}

type StatusEvent struct {
	Amount            string `xml:"amount"`
	User              string `xml:"user"`
	TransactionSource string `xml:"transaction-source"`
	Timestamp         string `xml:"timestamp"`
	Status            string `xml:"status"`
}

type Descriptor struct {
	Name  string `xml:"name"`
	Phone string `xml:"phone"`
	URL   string `xml:"url"`
}

// TODO: not all transaction fields are implemented yet, here are the missing fields (add on demand)
//
// <transaction>
//   <currency-iso-code>USD</currency-iso-code>
//   <custom-fields>
//   </custom-fields>
//   <avs-error-response-code nil="true"></avs-error-response-code>
//   <avs-postal-code-response-code>I</avs-postal-code-response-code>
//   <avs-street-address-response-code>I</avs-street-address-response-code>
//   <cvv-response-code>I</cvv-response-code>
//   <gateway-rejection-reason nil="true"></gateway-rejection-reason>
//   <voice-referral-number nil="true"></voice-referral-number>
//   <purchase-order-number nil="true"></purchase-order-number>
//   <tax-amount nil="true"></tax-amount>
//   <tax-exempt type="boolean">false</tax-exempt>
//   <status-history type="array">
//     <status-event>
//       <timestamp type="datetime">2013-10-07T17:26:14Z</timestamp>
//       <status>authorized</status>
//       <amount>7.00</amount>
//       <user>eaigner</user>
//       <transaction-source>Recurring</transaction-source>
//     </status-event>
//     <status-event>
//       <timestamp type="datetime">2013-10-07T17:26:14Z</timestamp>
//       <status>submitted_for_settlement</status>
//       <amount>7.00</amount>
//       <user>eaigner</user>
//       <transaction-source>Recurring</transaction-source>
//     </status-event>
//     <status-event>
//       <timestamp type="datetime">2013-10-08T07:06:38Z</timestamp>
//       <status>settled</status>
//       <amount>7.00</amount>
//       <user nil="true"></user>
//       <transaction-source></transaction-source>
//     </status-event>
//   </status-history>
//   <plan-id>bronze</plan-id>
//   <subscription-id>jqsydb</subscription-id>
//   <subscription>
//     <billing-period-end-date type="date">2013-11-06</billing-period-end-date>
//     <billing-period-start-date type="date">2013-10-07</billing-period-start-date>
//   </subscription>
//   <add-ons type="array"/>
//   <discounts type="array"/>
//   <descriptor>
//     <name nil="true"></name>
//     <phone nil="true"></phone>
//   </descriptor>
//   <recurring type="boolean">true</recurring>
//   <channel nil="true"></channel>
//   <escrow-status nil="true"></escrow-status>
// </transaction>

type Transactions struct {
	Transaction []*Transaction `xml:"transaction"`
}

type TransactionOptions struct {
	SubmitForSettlement              bool `xml:"submit-for-settlement,omitempty"`
	StoreInVaultOnSuccess            bool `xml:"store-in-vault-on-success,omitempty"`
	AddBillingAddressToPaymentMethod bool `xml:"add-billing-address-to-payment-method,omitempty"`
	StoreShippingAddressInVault      bool `xml:"store-shipping-address-in-vault,omitempty"`
	SkipCVV                          bool `xml:"skip-cvv,omitempty"`
}

type ThreeDSecurePassThrough struct {
	CAVV                string `xml:"cavv"`
	DSTransactionID     string `xml:"ds-transaction-id"`
	ECIFlag             string `xml:"eci-flag"`
	ThreeDSecureVersion string `xml:"three-d-secure-version"`
	XID                 string `xml:"xid"`
}

type TransactionSearchResult struct {
	XMLName        string              `xml:"search-results"`
	PageSize       *nullable.NullInt64 `xml:"page-size"`
	TransactionIDs []string            `xml:"ids>item"`
}

type IDs struct {
	Type string `xml:"type,attr"`
}

type ApplePayCard struct {
	Number          string `xml:"number,omitempty"`
	CardHolderName  string `xml:"cardholder_name,omitempty"`
	Cryptogram      string `xml:"cryptogram,omitempty"`
	ExpirationMonth string `xml:"expiration_month,omitempty"`
	ExpirationYear  string `xml:"expiration_year,omitempty"`
	EciIndicator    string `xml:"eci_indicator,omitempty"`
}
