// +build unit

package braintree

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestWebhookParseRequest(t *testing.T) {
	t.Parallel()

	gateway := New(Sandbox, "mid", "sz9g7zhxz8838v7h", "0c809a2d2e8f4e4c817900ff441c9554")
	webhookGateway := gateway.WebhookNotification()

	body := strings.NewReader("bt_signature=sz9g7zhxz8838v7h%7C4b532339b3107eae876d7637d59217858f320098&bt_payload=PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPG5vdGlm%0AaWNhdGlvbj4KICA8a2luZD5jaGVjazwva2luZD4KICA8dGltZXN0YW1wIHR5%0AcGU9ImRhdGV0aW1lIj4yMDE3LTA0LTI2VDA3OjEyOjI0WjwvdGltZXN0YW1w%0APgogIDxzdWJqZWN0PgogICAgPGNoZWNrIHR5cGU9ImJvb2xlYW4iPnRydWU8%0AL2NoZWNrPgogIDwvc3ViamVjdD4KPC9ub3RpZmljYXRpb24%2BCg%3D%3D%0A")
	r := &http.Request{
		Method:        "POST",
		Header:        http.Header{"Content-Type": {"application/x-www-form-urlencoded"}},
		ContentLength: int64(body.Len()),
		Body:          ioutil.NopCloser(body),
	}

	notification, err := webhookGateway.ParseRequest(r)

	if err != nil {
		t.Fatal(err)
	} else if notification.Kind != CheckWebhook {
		t.Fatal("Incorrect Notification kind, expected check got", notification.Kind)
	}
}

func TestWebhookParseMerchantAccountAccepted(t *testing.T) {
	t.Parallel()

	webhookGateway := testGateway.WebhookNotification()
	apiKey := testGateway.credentials.(apiKey)
	hmacer := newHmacer(apiKey.publicKey, apiKey.privateKey)

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
	signature := hmacer.publicKey + "|" + hmacedPayload

	notification, err := webhookGateway.Parse(signature, payload)

	if err != nil {
		t.Fatal(err)
	} else if notification.Kind != SubMerchantAccountApprovedWebhook {
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
	t.Parallel()

	webhookGateway := testGateway.WebhookNotification()
	apiKey := testGateway.credentials.(apiKey)
	hmacer := newHmacer(apiKey.publicKey, apiKey.privateKey)

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
	signature := hmacer.publicKey + "|" + hmacedPayload

	notification, err := webhookGateway.Parse(signature, payload)

	if err != nil {
		t.Fatal(err)
	} else if notification.Kind != SubMerchantAccountDeclinedWebhook {
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

func TestWebhookParseDisbursement(t *testing.T) {
	t.Parallel()

	webhookGateway := testGateway.WebhookNotification()
	apiKey := testGateway.credentials.(apiKey)
	hmacer := newHmacer(apiKey.publicKey, apiKey.privateKey)

	payload := base64.StdEncoding.EncodeToString([]byte(`
<notification>
	<timestamp type="datetime">2014-04-06T10:32:28+00:00</timestamp>
	<kind>disbursement</kind>
	<subject>
		<disbursement>
			<id>456</id>
			<transaction-ids type="array">
				<item>afv56j</item>
				<item>kj8hjk</item>
			</transaction-ids>
			<success type="boolean">true</success>
			<retry type="boolean">false</retry>
			<merchant-account>
				<id>123</id>
				<currency-iso-code>USD</currency-iso-code>
				<sub-merchant-account type="boolean">false</sub-merchant-account>
				<status>active</status>
			</merchant-account>
			<amount>100.00</amount>
			<disbursement-date type="date">2014-02-09</disbursement-date>
			<exception-message nil="true"/>
			<follow-up-action nil="true"/>
		</disbursement>
	</subject>
</notification>`))
	hmacedPayload, err := hmacer.hmac(payload)
	if err != nil {
		t.Fatal(err)
	}
	signature := hmacer.publicKey + "|" + hmacedPayload

	notification, err := webhookGateway.Parse(signature, payload)

	if err != nil {
		t.Fatal(err)
	} else if notification.Kind != DisbursementWebhook {
		t.Fatal("Incorrect Notification kind, expected disbursement got", notification.Kind)
	} else if notification.Disbursement() == nil {
		t.Fatal("Notification should have a disbursement")
	} else if notification.Disbursement().Id != "456" {
		t.Fatal("Incorrect disbursement id, expected 456 got", notification.Subject.MerchantAccount.Status)
	} else if len(notification.Disbursement().TransactionIds) != 2 {
		t.Fatal("Disbursement should have two txns")
	} else if notification.Disbursement().TransactionIds[1] != "kj8hjk" {
		t.Fatal("Incorrect txn id on disbursement, expected kj8hjk got", notification.Disbursement().TransactionIds[1])
	} else if notification.Disbursement().MerchantAccount.Id != "123" {
		t.Fatal("Disbursement not associated with correct merchant account")
	} else if notification.Disbursement().ExceptionMessage != "" {
		t.Fatal("Disbursement should not have an exception message")
	}
}

func TestWebhookParseDisbursementException(t *testing.T) {
	t.Parallel()

	webhookGateway := testGateway.WebhookNotification()
	apiKey := testGateway.credentials.(apiKey)
	hmacer := newHmacer(apiKey.publicKey, apiKey.privateKey)

	payload := base64.StdEncoding.EncodeToString([]byte(`
<notification>
	<timestamp type="datetime">2014-04-06T10:32:28+00:00</timestamp>
	<kind>disbursement_exception</kind>
	<subject>
		<disbursement>
			<id>456</id>
			<transaction-ids type="array">
				<item>afv56j</item>
				<item>kj8hjk</item>
			</transaction-ids>
			<success type="boolean">false</success>
			<retry type="boolean">false</retry>
			<merchant-account>
				<id>123</id>
				<currency-iso-code>USD</currency-iso-code>
				<sub-merchant-account type="boolean">false</sub-merchant-account>
				<status>active</status>
			</merchant-account>
			<amount>100.00</amount>
			<disbursement-date type="date">2014-02-09</disbursement-date>
      <exception-message>bank_rejected</exception-message>
      <follow-up-action>update_funding_information</follow-up-action>
		</disbursement>
	</subject>
</notification>`))
	hmacedPayload, err := hmacer.hmac(payload)
	if err != nil {
		t.Fatal(err)
	}
	signature := hmacer.publicKey + "|" + hmacedPayload

	notification, err := webhookGateway.Parse(signature, payload)

	if err != nil {
		t.Fatal(err)
	} else if notification.Disbursement().ExceptionMessage != BankRejected {
		t.Fatal("Disbursement should have a BankRejected exception message")
	} else if notification.Disbursement().FollowUpAction != UpdateFundingInformation {
		t.Fatal("Disbursement followup action should be UpdateFundingInformation")
	}

}

func TestWebhookParseDispute(t *testing.T) {
	t.Parallel()

	webhookGateway := testGateway.WebhookNotification()
	apiKey := testGateway.credentials.(apiKey)
	hmacer := newHmacer(apiKey.publicKey, apiKey.privateKey)

	payload := base64.StdEncoding.EncodeToString([]byte(`
        <notification>
            <timestamp type="datetime">2014-04-06T10:32:28+00:00</timestamp>
            <kind>dispute_opened</kind>
            <subject>
                <dispute>
                    <id>123456</id>
                    <amount>100.00</amount>
                    <amount-disputed>100.00</amount-disputed>
                    <amount-won>95.00</amount-won>
                    <case-number>CASE-12345</case-number>
                    <created-at type="datetime">2017-06-16T20:44:41Z</created-at>
                    <currency-iso-code>USD</currency-iso-code>
                    <forwarded-comments>Forwarded comments</forwarded-comments>
                    <kind>chargeback</kind>
                    <merchant-account-id>abc123</merchant-account-id>
                    <reason>fraud</reason>
                    <reason-code>83</reason-code>
                    <reason-description>Reason code 83 description</reason-description>
                    <received-date type="date">2016-02-15</received-date>
                    <reference-number>123456</reference-number>
                    <reply-by-date type="date">2016-02-22</reply-by-date>
                    <status>open</status>
                    <updated-at type="datetime">2013-04-10T10:50:39Z</updated-at>
                    <original-dispute-id>original_dispute_id</original-dispute-id>
                    <status-history type="array">
                        <status-history>
                            <status>open</status>
                            <timestamp type="datetime">2013-04-10T10:50:39Z</timestamp>
                            <effective-date type="date">2013-04-10</effective-date>
                        </status-history>
                    </status-history>
                    <evidence type="array">
                        <evidence>
                            <created-at type="datetime">2013-04-11T10:50:39Z</created-at>
                            <id>evidence1</id>
                            <url>url_of_file_evidence</url>
                        </evidence>
                        <evidence>
                            <created-at type="datetime">2013-04-11T10:50:39Z</created-at>
                            <id>evidence2</id>
                            <comment>text evidence</comment>
                            <sent-to-processor-at type="date">2009-04-11</sent-to-processor-at>
                        </evidence>
                    </evidence>
                    <transaction>
                        <id>123456</id>
                        <amount>100.00</amount>
                        <created-at>2017-06-21T20:44:41Z</created-at>
                        <order-id nil="true"/>
                        <purchase-order-number nil="true"/>
                        <payment-instrument-subtype>Visa</payment-instrument-subtype>
                    </transaction>
                    <date-opened type="date">2014-03-28</date-opened>
                    <date-won type="date">2014-04-05</date-won>
                </dispute>
            </subject>
        </notification>`))
	hmacedPayload, err := hmacer.hmac(payload)
	if err != nil {
		t.Fatal(err)
	}
	signature := hmacer.publicKey + "|" + hmacedPayload

	notification, err := webhookGateway.Parse(signature, payload)

	if err != nil {
		t.Fatal(err)
	} else if notification.Kind != DisputeOpenedWebhook {
		t.Fatal("Incorrect Notification kind, expected dispute got", notification.Kind)
	} else if notification.Dispute() == nil {
		t.Fatal("Notification should have a dispute")
	} else if notification.Dispute().Kind != DisputeKindChargeback {
		t.Errorf("Incorrect dispute kind, expected %s got %s", DisputeKindChargeback, notification.Dispute().Kind)
	} else if notification.Dispute().Reason != DisputeReasonFraud {
		t.Errorf("Incorrect dispute reason, expected %s got %s", DisputeReasonFraud, notification.Dispute().Reason)
	} else if notification.Dispute().Status != DisputeStatusOpen {
		t.Errorf("Incorrect dispute status, expected %s got %s", DisputeStatusOpen, notification.Dispute().Reason)
	} else if notification.Dispute().ID != "123456" {
		t.Errorf("Incorrect dispute id, expected 456 got %s", notification.Dispute().ID)
	} else if len(notification.Dispute().StatusHistory) != 1 {
		t.Error("Dispute should have one status history entry")
	} else if len(notification.Dispute().Evidence) != 2 {
		t.Error("Dispute should have two evidence entries")
	} else if notification.Dispute().Transaction == nil {
		t.Error("Dispute shoud have transaction details")
	}
}

func TestWebhookParseAccountUpdaterDailyReport(t *testing.T) {
	t.Parallel()

	webhookGateway := testGateway.WebhookNotification()
	apiKey := testGateway.credentials.(apiKey)
	hmacer := newHmacer(apiKey.publicKey, apiKey.privateKey)

	payload := base64.StdEncoding.EncodeToString([]byte(`
        <notification>
            <timestamp type="datetime">2014-04-06T10:32:28+00:00</timestamp>
            <kind>account_updater_daily_report</kind>
            <subject>
        		<account-updater-daily-report>
					<report-date type="date">2016-01-14</report-date>
					<report-url>link-to-csv-report</report-url>
				</account-updater-daily-report>
            </subject>
        </notification>`))
	hmacedPayload, err := hmacer.hmac(payload)
	if err != nil {
		t.Fatal(err)
	}
	signature := hmacer.publicKey + "|" + hmacedPayload

	notification, err := webhookGateway.Parse(signature, payload)

	if err != nil {
		t.Fatal(err)
	} else if notification.Kind != AccountUpdaterDailyReportWebhook {
		t.Fatal("Incorrect Notification kind, expected account_updater_daily_report got", notification.Kind)
	} else if notification.AccountUpdaterDailyReport() == nil {
		t.Fatal("Notification should have a account updater daily report")
	} else if notification.AccountUpdaterDailyReport().ReportDate != "2016-01-14" {
		t.Errorf("Incorrect report date, expected 2016-01-14, got %s", notification.AccountUpdaterDailyReport().ReportDate)
	} else if notification.AccountUpdaterDailyReport().ReportURL != "link-to-csv-report" {
		t.Errorf("Incorrect report url, expected link-to-csv-report, got %s", notification.AccountUpdaterDailyReport().ReportURL)
	}
}
