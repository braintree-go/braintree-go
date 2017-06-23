package braintree

import (
	"time"

	"github.com/lionelbarrow/braintree-go/customfields"
	"github.com/lionelbarrow/braintree-go/nullable"
)

type TransactionStatus string

const (
	TransactionStatusAuthorizationExpired   TransactionStatus = "authorization_expired"
	TransactionStatusAuthorizing            TransactionStatus = "authorizing"
	TransactionStatusAuthorized             TransactionStatus = "authorized"
	TransactionStatusGatewayRejected        TransactionStatus = "gateway_rejected"
	TransactionStatusFailed                 TransactionStatus = "failed"
	TransactionStatusProcessorDeclined      TransactionStatus = "processor_declined"
	TransactionStatusSettled                TransactionStatus = "settled"
	TransactionStatusSettlementConfirmed    TransactionStatus = "settlement_confirmed"
	TransactionStatusSettlementDeclined     TransactionStatus = "settlement_declined"
	TransactionStatusSettlementPending      TransactionStatus = "settlement_pending"
	TransactionStatusSettling               TransactionStatus = "settling"
	TransactionStatusSubmittedForSettlement TransactionStatus = "submitted_for_settlement"
	TransactionStatusVoided                 TransactionStatus = "voided"
	TransactionStatusUnrecognized           TransactionStatus = "unrecognized"
)

type Transaction struct {
	XMLName                      string                    `xml:"transaction"`
	Id                           string                    `xml:"id,omitempty"`
	Status                       TransactionStatus         `xml:"status,omitempty"`
	Type                         string                    `xml:"type,omitempty"`
	CurrencyISOCode              string                    `xml:"currency-iso-code,omitempty"`
	Amount                       *Decimal                  `xml:"amount"`
	OrderId                      string                    `xml:"order-id,omitempty"`
	PaymentMethodToken           string                    `xml:"payment-method-token,omitempty"`
	PaymentMethodNonce           string                    `xml:"payment-method-nonce,omitempty"`
	MerchantAccountId            string                    `xml:"merchant-account-id,omitempty"`
	PlanId                       string                    `xml:"plan-id,omitempty"`
	SubscriptionId               string                    `xml:"subscription-id,omitempty"`
	CreditCard                   *CreditCard               `xml:"credit-card,omitempty"`
	Customer                     *Customer                 `xml:"customer,omitempty"`
	BillingAddress               *Address                  `xml:"billing,omitempty"`
	ShippingAddress              *Address                  `xml:"shipping,omitempty"`
	DeviceData                   string                    `xml:"device-data,omitempty"`
	ServiceFeeAmount             *Decimal                  `xml:"service-fee-amount,attr,omitempty"`
	CreatedAt                    *time.Time                `xml:"created-at,omitempty"`
	UpdatedAt                    *time.Time                `xml:"updated-at,omitempty"`
	DisbursementDetails          *DisbursementDetails      `xml:"disbursement-details,omitempty"`
	RefundId                     string                    `xml:"refund-id,omitempty"`
	RefundIds                    *[]string                 `xml:"refund-ids>item,omitempty"`
	RefundedTransactionId        *string                   `xml:"refunded-transaction-id,omitempty"`
	ProcessorResponseCode        ProcessorResponseCode     `xml:"processor-response-code,omitempty"`
	ProcessorResponseText        string                    `xml:"processor-response-text,omitempty"`
	ProcessorAuthorizationCode   string                    `xml:"processor-authorization-code,omitempty"`
	SettlementBatchId            string                    `xml:"settlement-batch-id,omitempty"`
	PaymentInstrumentType        string                    `xml:"payment-instrument-type,omitempty"`
	PayPalDetails                *PayPalDetails            `xml:"paypal,omitempty"`
	VenmoAccountDetails          *VenmoAccountDetails      `xml:"venmo-account,omitempty"`
	AndroidPayDetails            *AndroidPayDetails        `xml:"android-pay-card,omitempty"`
	ApplePayDetails              *ApplePayDetails          `xml:"apple-pay,omitempty"`
	AdditionalProcessorResponse  string                    `xml:"additional-processor-response,omitempty"`
	RiskData                     *RiskData                 `xml:"risk-data,omitempty"`
	Descriptor                   *Descriptor               `xml:"descriptor,omitempty"`
	CustomFields                 customfields.CustomFields `xml:"custom-fields,omitempty"`
	AVSErrorResponseCode         AVSResponseCode           `xml:"avs-error-response-code,omitempty"`
	AVSPostalCodeResponseCode    AVSResponseCode           `xml:"avs-postal-code-response-code,omitempty"`
	AVSStreetAddressResponseCode AVSResponseCode           `xml:"avs-street-address-response-code,omitempty"`
	CVVResponseCode              CVVResponseCode           `xml:"cvv-response-code,omitempty"`
	GatewayRejectionReason       GatewayRejectionReason    `xml:"gateway-rejection-reason,omitempty"`
}

type TransactionRequest struct {
	XMLName            string                    `xml:"transaction"`
	CustomerID         string                    `xml:"customer-id,omitempty"`
	Type               string                    `xml:"type,omitempty"`
	Amount             *Decimal                  `xml:"amount"`
	OrderId            string                    `xml:"order-id,omitempty"`
	PaymentMethodToken string                    `xml:"payment-method-token,omitempty"`
	PaymentMethodNonce string                    `xml:"payment-method-nonce,omitempty"`
	MerchantAccountId  string                    `xml:"merchant-account-id,omitempty"`
	PlanId             string                    `xml:"plan-id,omitempty"`
	CreditCard         *CreditCard               `xml:"credit-card,omitempty"`
	Customer           *Customer                 `xml:"customer,omitempty"`
	BillingAddress     *Address                  `xml:"billing,omitempty"`
	ShippingAddress    *Address                  `xml:"shipping,omitempty"`
	DeviceData         string                    `xml:"device-data,omitempty"`
	Options            *TransactionOptions       `xml:"options,omitempty"`
	ServiceFeeAmount   *Decimal                  `xml:"service-fee-amount,attr,omitempty"`
	RiskData           *RiskDataRequest          `xml:"risk-data,omitempty"`
	Descriptor         *Descriptor               `xml:"descriptor,omitempty"`
	CustomFields       customfields.CustomFields `xml:"custom-fields,omitempty"`
}

// TODO: not all transaction fields are implemented yet, here are the missing fields (add on demand)
//
// <transaction>
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
	StoreInVault                     bool `xml:"store-in-vault,omitempty"`
	AddBillingAddressToPaymentMethod bool `xml:"add-billing-address-to-payment-method,omitempty"`
	StoreShippingAddressInVault      bool `xml:"store-shipping-address-in-vault,omitempty"`
}

type TransactionSearchResult struct {
	XMLName           string              `xml:"credit-card-transactions"`
	CurrentPageNumber *nullable.NullInt64 `xml:"current-page-number"`
	PageSize          *nullable.NullInt64 `xml:"page-size"`
	TotalItems        *nullable.NullInt64 `xml:"total-items"`
	Transactions      []*Transaction      `xml:"transaction"`
}

type RiskData struct {
	ID       string `xml:"id"`
	Decision string `xml:"decision"`
}

type RiskDataRequest struct {
	CustomerBrowser string `xml:"customer-browser"`
	CustomerIP      string `xml:"customer-ip"`
}
