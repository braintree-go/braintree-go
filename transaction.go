package braintree

type Transaction struct {
	XMLName            string              `json:"transaction" xml:"transaction"`
	Id                 string              `json:"id,omitempty" xml:"id,omitempty"`
	CustomerID         string              `json:"customer-id,omitempty" xml:"customer-id,omitempty"`
	Status             string              `json:"status,omitempty" xml:"status,omitempty"`
	Type               string              `json:"type,omitempty" xml:"type,omitempty"`
	Amount             float64             `json:"amount" xml:"amount"`
	OrderId            string              `json:"order-id,omitempty" xml:"order-id,omitempty"`
	PaymentMethodToken string              `json:"payment-method-token,omitempty" xml:"payment-method-token,omitempty"`
	MerchantAccountId  string              `json:"merchant-account-id,omitempty" xml:"merchant-account-id,omitempty"`
	PlanId             string              `json:"plan-id,omitempty" xml:"plan-id,omitempty"`
	CreditCard         *CreditCard         `json:"credit-card,omitempty" xml:"credit-card,omitempty"`
	Customer           *Customer           `json:"customer,omitempty" xml:"customer,omitempty"`
	BillingAddress     *Address            `json:"billing,omitempty" xml:"billing,omitempty"`
	ShippingAddress    *Address            `json:"shipping,omitempty" xml:"shipping,omitempty"`
	Options            *TransactionOptions `json:"options,omitempty" xml:"options,omitempty"`
	ServiceFeeAmount   float64             `json:"service-fee-amount,attr,omitempty" xml:"service-fee-amount,attr,omitempty"`
	CreatedAt          string              `json:"created-at,omitempty" xml:"created-at,omitempty"`
	UpdatedAt          string              `json:"updated-at,omitempty" xml:"updated-at,omitempty"`
	AuthCode           string              `json:"processor-authorization-code,omitempty" xml:"processor-authorization-code,omitempty"`
}

// TODO: not all transaction fields are implemented yet, here are the missing fields (add on demand)
//
// <transaction>
//   <currency-iso-code>USD</currency-iso-code>
//   <refund-id nil="true"></refund-id>
//   <refund-ids type="array"/>
//   <refunded-transaction-id nil="true"></refunded-transaction-id>
//   <settlement-batch-id>2013-10-08_49grybq7pbtsnvsr</settlement-batch-id>
//   <custom-fields>
//   </custom-fields>
//   <avs-error-response-code nil="true"></avs-error-response-code>
//   <avs-postal-code-response-code>I</avs-postal-code-response-code>
//   <avs-street-address-response-code>I</avs-street-address-response-code>
//   <cvv-response-code>I</cvv-response-code>
//   <gateway-rejection-reason nil="true"></gateway-rejection-reason>
//   <processor-authorization-code>YCSBWR</processor-authorization-code>
//   <processor-response-code>1000</processor-response-code>
//   <processor-response-text>Approved</processor-response-text>
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
//   <disbursement-details>
//     <disbursement-date type="date">2013-10-08</disbursement-date>
//     <settlement-amount>7.00</settlement-amount>
//     <settlement-currency-iso-code>USD</settlement-currency-iso-code>
//     <settlement-currency-exchange-rate>1</settlement-currency-exchange-rate>
//     <funds-held type="boolean">false</funds-held>
//   </disbursement-details>
// </transaction>

type Transactions struct {
	Transaction []*Transaction `json:"transaction" xml:"transaction"`
}

type TransactionOptions struct {
	SubmitForSettlement              bool `json:"submit-for-settlement,omitempty" xml:"submit-for-settlement,omitempty"`
	StoreInVault                     bool `json:"store-in-vault,omitempty" xml:"store-in-vault,omitempty"`
	AddBillingAddressToPaymentMethod bool `json:"add-billing-address-to-payment-method,omitempty" xml:"add-billing-address-to-payment-method,omitempty"`
	StoreShippingAddressInVault      bool `json:"store-shipping-address-in-vault,omitempty" xml:"store-shipping-address-in-vault,omitempty"`
}

type TransactionSearchResult struct {
	XMLName           string         `json:"credit-card-transactions" xml:"credit-card-transactions"`
	CurrentPageNumber string         `json:"current-page-number" xml:"current-page-number"` // int
	PageSize          string         `json:"page-size" xml:"page-size"`           // int
	TotalItems        string         `json:"total-items" xml:"total-items"`         // int
	Transactions      []*Transaction `json:"transaction" xml:"transaction"`
}

