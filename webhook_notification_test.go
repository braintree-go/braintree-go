package braintree

import (
	"encoding/base64"
	"testing"
)

func TestWebhookParseMerchantAccountAccepted(t *testing.T) {
	webhookGateway := testGateway.WebhookNotification()
	hmacer := newHmacer(testGateway)

	payload := base64.StdEncoding.EncodeToString([]byte(`
<notification>
  <timestamp type="datetime">2014-01-26T10:32:28+00:00</timestamp>
  <kind>sub_merchant_account_approved</kind>
  <subject>
    <merchant-account>
      <id>123</id>
      <master-merchant-account>
        <id>master_ma_for_123</id>
        <status>active</status>
      </master-merchant-account>
      <status>active</status>
    </merchant-account>
  </subject>
</notification>`))
	hmacedPayload, err := hmacer.hmac(payload)
	if err != nil {
		t.Fatal(err)
	}
	signature := webhookGateway.PublicKey + "|" + hmacedPayload

	notification, err := webhookGateway.Parse(signature, payload)

	if err != nil {
		t.Fatal(err)
	} else if notification.Kind != SubMerchantAccountApproved {
		t.Fatal("Incorrect Notification kind, expected sub_merchant_account_approved got", notification.Kind)
	} else if notification.MerchantAccount() == nil {
		t.Log(notification.Subject)
		t.Fatal("Notification should have a merchant account")
	} else if notification.MerchantAccount().Id != "123" {
		t.Fatal("Incorrect Merchant Id, expected '123' got", notification.Subject.MerchantAccount.Id)
	} else if notification.MerchantAccount().Status != "active" {
		t.Fatal("Incorrect Merchant Status, expected 'active' got", notification.Subject.MerchantAccount.Status)
	}
}

func TestWebhookParseMerchantAccountDeclined(t *testing.T) {
	webhookGateway := testGateway.WebhookNotification()
	hmacer := newHmacer(testGateway)

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
</notification>`))
	hmacedPayload, err := hmacer.hmac(payload)
	if err != nil {
		t.Fatal(err)
	}
	signature := webhookGateway.PublicKey + "|" + hmacedPayload

	notification, err := webhookGateway.Parse(signature, payload)

	if err != nil {
		t.Fatal(err)
	} else if notification.Kind != SubMerchantAccountDeclined {
		t.Fatal("Incorrect Notification kind, expected sub_merchant_account_declined got", notification.Kind)
	} else if notification.Subject.APIErrorResponse == nil {
		t.Fatal("Notification should have an error response")
	} else if notification.Subject.APIErrorResponse.ErrorMessage != "Credit score is too low" {
		t.Fatal("Incorrect Error Message, expected 'Credit score is too low' got", notification.Subject.APIErrorResponse.ErrorMessage)
	} else if notification.MerchantAccount() == nil {
		t.Fatal("Notification should have a merchant account")
	} else if notification.MerchantAccount().Id != "123" {
		t.Fatal("Incorrect Merchant Id, expected '123' got", notification.Subject.MerchantAccount.Id)
	} else if notification.MerchantAccount().Status != "suspended" {
		t.Fatal("Incorrect Merchant Status, expected 'suspended' got", notification.Subject.MerchantAccount.Status)
	}
}
