## 0.9.0 (August 5th, 2015)

BACKWARDS INCOMPATIBILITES:

* Time fields such as `CreatedAt`, `UpdatedAt`, and `DisbursementDetails`
  changed to be `time.Time` or `date.Date` types to make their use simpler
  (rather than forcing the user to deserialize)
* Additional currency fields migrated from `float64` to `Decimal` to avoid
  precision loss
* `Null(Int64|Bool)` types added to support empty XML elements. Many of
  the nonstring, `string`, fields were updated to be nullable of their actual
  type.
* `ClientToken` was renamed to `ClientTokenRequest`
* `ClientToken` API changed to allow generation of client tokens with or
  without an associated customer. `NewClientTokenRequest` was removed.

IMPROVEMENTS:

* `CustomerGateway.Search` added to permit advanced searching for customers by
  metadata
* `BraintreeError` type was exposed to make it easier to inspect whether the
  errors returned by the library are network on Braintree Gatway errors
* `ClientTokenGateway.GenerateWithCustomer` added to generate a customer
  specific client token

## 0.8.0 (April 3, 2015)

BACKWARDS INCOMPATIBILITES:

* Webhook constants made more uniform via `Webhook` suffix
* All currency amounts changed from `float` to `Decimal` to remove loss of
  precision

IMPROVEMENTS:

* Specification of a custom `http.Client` to use via `Braintree.HttpClient`.
  This enables `AppEngine` support which required a being able to use a custom
  `http.Client`.
* `DisbursementDetails` added to `Transaction`
* Support for querying disbursement webhooks added via `WebhookNotification.Disbursement`
* `TransactionGateway.Settle` added to automatically settle transactions in
  sandbox (`SubmitForSettlement` should be used in production)
* `PaymentMethodNonce` added to `CreditCard`
* `PaymentMethodNonce` added to `Transaction`
* `Decimal` arbitrary precision numeric type added to be used for currency
  amounts
* `ClientToken` support added via `ClientTokenGateway` to generate new client
  tokens

BUG FIXES:

* Typo in path for merchant account updates (`MerchantAccountGateway.Update`)
  was fixed.

## 0.7.0 (April 3, 2014)

BACKWARDS INCOMPATIBILITES:

* `InvalidResponseError` was unexported to encourage use of the new
  `BraintreeError` type
* `CreditCard.Default` changed from string to bool
* `CreditCard.Expired` changed from string to bool

IMPROVEMENTS:

* `CustomerGateway.Update` added to update metadata about the customer
* `CustomerGateway.Delete` added to allow customers to be deleted
* `Customer.DefaultCreditCard` added to return the default credit card
  associated with the customer
* `BraintreeError` type added to expose metadata about gateway errors in
  a structured manner
* `TransactionGateway.SubmitForSettlement` added to allow transactions to be
  submitted to be settled
* `TransactionGateway.Void` added to allow transactions to be voided
* Additional fields added to `Plan` (all except `Addons` and `Discounts`)
* Additional fields added to `Subscription` (all except `Addons` and `Descriptor`)
* `Subscription.Update` added to allow subscription data to be updated
* Remaining fields added to `CreditCard` and `CreditCardOptions`
* `CreditCardGateway.Update` added to update credit card information
* `CreditCardGateway.Delete` added to allow credit cards to be deleted
* `CreditCard.AllSubscriptions` added to allow subscriptions for a credit card
  to be queried
* `PlanGateway.Find` added to lookup plan by id
* `SubscriptionStatus*` constants were added to make comparisons easier
* `TransactionGateway.Search` added to permit searching for transactions by
  metadata
* `CreatedAt`, `UpdatedAt`, `PlanId` added to `Transaction`
* `ParseDate` added to facilitate parsing the date format returned by Braintree
* Adedd `AddOn` support via `AddOnGateway`
* Adedd `Discount` support via `DiscountGateway`
* Adedd `MerchantAccount` support via `MerchantAccountGateway` for submerchant
  support. Includes addition of `ServiceFeeAmount` to `Transaction`

BUG FIXES:

* `AddressGateway.Create` now copies address for sanitization to avoid
  modifying passed struct
* Errors during failed HTTP requests no longer cause a nil pointer dereference
  (when a `nil` body was `Close`d)

## 0.6.0 (June 30, 2015)

BACKWARDS INCOMPATIBILITES:

* Large scale refactoring from `0.5.0`

IMPROVEMENTS:

* Start of `Subscription` and `Plan` support
* `Address` `Create` and `Delete` support added via `AddressGateway`
* `ExpirationMonth` and `ExpirationYear` added to `CreditCard`

## 0.5.0 (May 27, 2013)

Initial release
