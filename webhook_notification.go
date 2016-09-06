package braintree

import (
	"encoding/xml"
	"time"
)

const (
	DisbursementWebhook                      = "disbursement"
	DisbursementExceptionWebhook             = "disbursement_exception"
	SubscriptionCanceledWebhook              = "subscription_canceled"
	SubscriptionChargedSuccessfullyWebhook   = "subscription_charged_successfully"
	SubscriptionChargedUnsuccessfullyWebhook = "subscription_charged_unsuccessfully"
	SubscriptionExpiredWebhook               = "subscription_expired"
	SubscriptionTrialEndedWebhook            = "subscription_trial_ended"
	SubscriptionWentActiveWebhook            = "subscription_went_active"
	SubscriptionWentPastDueWebhook           = "subscription_went_past_due"

	SubMerchantAccountApprovedWebhook  = "sub_merchant_account_approved"
	SubMerchantAccountDeclinedWebhook  = "sub_merchant_account_declined"
	TransactionDisbursedWebhook        = "transaction_disbursed"
	PartnerMerchantConnectedWebhook    = "partner_merchant_connected"
	PartnerMerchantDisconnectedWebhook = "partner_merchant_disconnected"
	PartnerMerchantDeclinedWebhook     = "partner_merchant_declined"
)

type WebhookNotification struct {
	XMLName   xml.Name        `xml:"notification" json:"notification" bson:"notification"`
	Timestamp time.Time       `xml:"timestamp" json:"timestamp" bson:"timestamp"`
	Kind      string          `xml:"kind" json:"kind" bson:"kind"`
	Subject   *webhookSubject `xml:"subject" json:"subject" bson:"subject"`
}

func (n *WebhookNotification) MerchantAccount() *MerchantAccount {
	if n.Subject.APIErrorResponse != nil && n.Subject.APIErrorResponse.MerchantAccount != nil {
		return n.Subject.APIErrorResponse.MerchantAccount
	} else if n.Subject.MerchantAccount != nil {
		return n.Subject.MerchantAccount
	}
	return nil
}

func (n *WebhookNotification) Disbursement() *Disbursement {
	if n.Subject.Disbursement != nil {
		return n.Subject.Disbursement
	} else {
		return nil
	}
}

type webhookSubject struct {
	XMLName          xml.Name         `xml:"subject" json:"subject" bson:"subject"`
	APIErrorResponse *BraintreeError  `xml:",omitempty" json:",omitempty" bson:",omitempty"`
	Disbursement     *Disbursement    `xml:"disbursement,omitempty" json:"disbursement,omitempty" bson:"disbursement,omitempty"`
	Subscription     *Subscription    `xml:",omitempty" json:",omitempty" bson:",omitempty"`
	MerchantAccount  *MerchantAccount `xml:"merchant-account,omitempty" json:"merchantAccount,omitempty" bson:"merchantAccount,omitempty"`
	Transaction      *Transaction     `xml:",omitempty" json:",omitempty" bson:",omitempty"`

	// Remaining Fields:
	// partner_merchant
}
