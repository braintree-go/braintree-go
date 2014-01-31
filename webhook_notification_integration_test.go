package braintree

import (
	"encoding/base64"
	"testing"
)

func TestWebhookParseMerchantAccountAccepted(t *testing.T) {
	webhookGateway := testGateway.WebhookNotification()

	payload := base64.StdEncoding.EncodeToString([]byte(`
		<notification>
	    <timestamp type="datetime">2014-01-26T10:32:28+00:00</timestamp>
	    <kind>sub_merchant_account_approved</kind>
	    <subject>
	      <merchant_account>
          <id>123</id>
          <master_merchant_account>
            <id>master_ma_for_123</id>
            <status>active</status>
          </master_merchant_account>
          <status>active</status>
        </merchant_account>
	    </subject>
	  </notification>
  `))
	signature := webhookGateway.PublicKey + "|" + Hexdigest(webhookGateway.PrivateKey, payload)

	notification, err := webhookGateway.Parse(signature, payload)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if notification.Kind != SubMerchantAccountApproved {
		t.Log("Incorrect Notification kind, expected sub_merchant_account_approved got", notification.Kind)
		t.Fail()
	}

	if notification.Subject.MerchantAccount == nil {
		t.Log("Notification should have a merchant account")
		t.FailNow()
	}

	if notification.Subject.MerchantAccount.Id != "123" {
		t.Log("Incorrect Merchant Id, expected '123' got", notification.Subject.MerchantAccount.Id)
		t.Fail()
	}

	if notification.Subject.MerchantAccount.Status != "active" {
		t.Log("Incorrect Merchant Status, expected 'active' got", notification.Subject.MerchantAccount.Status)
		t.Fail()
	}
}

func TestWebhookParseMerchantAccountDeclined(t *testing.T) {
	webhookGateway := testGateway.WebhookNotification()

	payload := base64.StdEncoding.EncodeToString([]byte(`
		<notification>
	    <timestamp type="datetime">2014-01-26T10:32:28+00:00</timestamp>
	    <kind>sub_merchant_account_declined</kind>
	    <subject>
	      <api-error-response>
          <message>Credit score is too low</message>
          <errors>
            <errors type="array"/>
              <merchant-account>
                <errors type="array">
                  <error>
                    <code>82621</code>
                    <message>Credit score is too low</message>
                    <attribute type="symbol">base</attribute>
                  </error>
                </errors>
              </merchant-account>
            </errors>
            <merchant-account>
              <id>123</id>
              <status>suspended</status>
              <master-merchant-account>
                <id>master_ma_for_123</id>
                <status>suspended</status>
              </master-merchant-account>
            </merchant-account>
        </api-error-response>
	    </subject>
	  </notification>
  `))
	signature := webhookGateway.PublicKey + "|" + Hexdigest(webhookGateway.PrivateKey, payload)

	notification, err := webhookGateway.Parse(signature, payload)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if notification.Kind != SubMerchantAccountDeclined {
		t.Log("Incorrect Notification kind, expected sub_merchant_account_declined got", notification.Kind)
		t.Fail()
	}

	if notification.Subject.APIErrorResponse == nil {
		t.Log("Notification should have an error response")
		t.FailNow()
	}

	if notification.Subject.APIErrorResponse.ErrorMessage != "Credit score is too low" {
		t.Log("Incorrect Error Message, expected 'Credit score is too low' got", notification.Subject.APIErrorResponse.ErrorMessage)
		t.Fail()
	}

	if notification.Subject.MerchantAccount == nil {
		t.Log("Notification should have a merchant account")
		t.FailNow()
	}

	if notification.Subject.MerchantAccount.Id != "123" {
		t.Log("Incorrect Merchant Id, expected '123' got", notification.Subject.MerchantAccount.Id)
		t.Fail()
	}

	if notification.Subject.MerchantAccount.Status != "suspended" {
		t.Log("Incorrect Merchant Status, expected 'suspended' got", notification.Subject.MerchantAccount.Status)
		t.Fail()
	}
}
