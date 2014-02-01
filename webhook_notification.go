package braintree

import (
	"encoding/xml"
	"time"
)

const (
	SubscriptionCanceled              = "subscription_canceled"
	SubscriptionChargedSuccessfully   = "subscription_charged_successfully"
	SubscriptionChargedUnsuccessfully = "subscription_charged_unsuccessfully"
	SubscriptionExpired               = "subscription_expired"
	SubscriptionTrialEnded            = "subscription_trial_ended"
	SubscriptionWentActive            = "subscription_went_active"
	SubscriptionWentPastDue           = "subscription_went_past_due"

	SubMerchantAccountApproved  = "sub_merchant_account_approved"
	SubMerchantAccountDeclined  = "sub_merchant_account_declined"
	TransactionDisbursed        = "transaction_disbursed"
	PartnerMerchantConnected    = "partner_merchant_connected"
	PartnerMerchantDisconnected = "partner_merchant_disconnected"
	PartnerMerchantDeclined     = "partner_merchant_declined"
)

type WebhookNotification struct {
	XMLName   xml.Name        `xml:"notification"`
	Timestamp time.Time       `xml:"timestamp"`
	Kind      string          `xml:"kind"`
	Subject   *webhookSubject `xml:"subject"`
}

func (n *WebhookNotification) MerchantAccount() *MerchantAccount {
	if n.Subject.APIErrorResponse != nil && n.Subject.APIErrorResponse.MerchantAccount != nil {
		return n.Subject.APIErrorResponse.MerchantAccount
	} else if n.Subject.MerchantAccount != nil {
		return n.Subject.MerchantAccount
	}
	return nil
}

type webhookSubject struct {
	XMLName          xml.Name         `xml:"subject"`
	APIErrorResponse *braintreeError  `xml:,omitempty"`
	Subscription     *Subscription    `xml:",omitempty"`
	MerchantAccount  *MerchantAccount `xml:"merchant-account,omitempty"`
	Transaction      *Transaction     `xml:",omitempty"`

	// Remaining Fields:
	// partner_merchant
}
