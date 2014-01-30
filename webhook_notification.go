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

type webhookSubject struct {
	XMLName          xml.Name        `xml:"subject"`
	APIErrorResponse *braintreeError `xml:,omitempty"`
	Subscription     *Subscription   `xml:",omitempty"`

	// Merchant account will be extracted from api error response on error
	MerchantAccount *MerchantAccount `xml:",omitempty"`
	Transaction     *Transaction     `xml:",omitempty"`

	// Remaining Fields:
	// partner_merchant
}

type WebhookNotification struct {
	XMLName   xml.Name        `xml:"notification"`
	Timestamp time.Time       `xml:"timestamp"`
	Kind      string          `xml:"kind"`
	Subject   *webhookSubject `xml:"subject"`
}

func NewWebhookNotification(xmlData []byte) (*WebhookNotification, error) {
	var n WebhookNotification
	if err := xml.Unmarshal(xmlData, &n); err != nil {
		return nil, err
	}

	if n.Subject.APIErrorResponse != nil && n.Subject.APIErrorResponse.MerchantAccount != nil {
		n.Subject.MerchantAccount = n.Subject.APIErrorResponse.MerchantAccount
	}

	return &n, nil
}
