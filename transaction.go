package braintree

import (
	"github.com/CompleteSet/braintree-go/nullable"
	"time"
)

type Transaction struct {
	XMLName                    string               `xml:"transaction" json:"transaction" bson:"transaction"`
	Id                         string               `xml:"id,omitempty" json:"id,omitempty" bson:"id,omitempty"`
	CustomerID                 string               `xml:"customer-id,omitempty" json:"customerId,omitempty" bson:"customerId,omitempty"`
	Status                     string               `xml:"status,omitempty" json:"status,omitempty" bson:"status,omitempty"`
	Type                       string               `xml:"type,omitempty" json:"type,omitempty" bson:"type,omitempty"`
	Amount                     *Decimal             `xml:"amount" json:"amount" bson:"amount"`
	OrderId                    string               `xml:"order-id,omitempty" json:"orderId,omitempty" bson:"orderId,omitempty"`
	PaymentMethodToken         string               `xml:"payment-method-token,omitempty" json:"paymentMethodToken,omitempty" bson:"paymentMethodToken,omitempty"`
	PaymentMethodNonce         string               `xml:"payment-method-nonce,omitempty" json:"paymentMethodNonce,omitempty" bson:"paymentMethodNonce,omitempty"`
	MerchantAccountId          string               `xml:"merchant-account-id,omitempty" json:"merchantAccountId,omitempty" bson:"merchantAccountId,omitempty"`
	PlanId                     string               `xml:"plan-id,omitempty" json:"planId,omitempty" bson:"planId,omitempty"`
	CreditCard                 *CreditCard          `xml:"credit-card,omitempty" json:"creditCard,omitempty" bson:"creditCard,omitempty"`
	Customer                   *Customer            `xml:"customer,omitempty" json:"customer,omitempty" bson:"customer,omitempty"`
	BillingAddress             *Address             `xml:"billing,omitempty" json:"billing,omitempty" bson:"billing,omitempty"`
	ShippingAddress            *Address             `xml:"shipping,omitempty" json:"shipping,omitempty" bson:"shipping,omitempty"`
	Options                    *TransactionOptions  `xml:"options,omitempty" json:"options,omitempty" json:"options,omitempty"`
	ServiceFeeAmount           *Decimal             `xml:"service-fee-amount,attr,omitempty" json:"serviceFeeAmount,attr,omitempty" bson:"serviceFeeAmount,attr,omitempty"`
	CreatedAt                  *time.Time           `xml:"created-at,omitempty" json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt                  *time.Time           `xml:"updated-at,omitempty" json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DisbursementDetails        *DisbursementDetails `xml:"disbursement-details,omitempty" json:"disbursementDetails,omitempty" bson:"disbursementDetails,omitempty"`
	RefundId                   string               `xml:"refund-id,omitempty" json:"refundId,omitempty" bson:"refundId,omitempty"`
	RefundIds                  *[]string            `xml:"refund-ids>item,omitempty" json:"refundIds.item,omitempty" bson:"refundIds.item,omitempty"`
	RefundedTransactionId      *string              `xml:"refunded-transaction-id,omitempty" json:"refundedTransactionId,omitempty" bson:"refundedTransactionId,omitempty"`
	ProcessorResponseCode      int                  `xml:"processor-response-code,omitempty" json:"processorResponseCode,omitempty" bson:"processorResponseCode,omitempty"`
	ProcessorResponseText      string               `xml:"processor-response-text,omitempty" json:"processorResponseText,omitempty" bson:"processorResponseText,omitempty"`
	ProcessorAuthorizationCode string               `xml:"processor-authorization-code,omitempty" json:"processorAuthorizationCode,omitempty" bson:"processorAuthorizationCode,omitempty"`
	SettlementBatchId          string               `xml:"settlement-batch-id,omitempty" json:"settlementBatchId,omitempty" bson:"settlementBatchId,omitempty"`
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
