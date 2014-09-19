package braintree

import (
	"reflect"
	"testing"
)

func TestTransactionResponse(t *testing.T) {

	txXml := []byte(`<transaction service-fee-amount="2.10">
		<id>txid</id>
		<customer-id>customerid</customer-id>
		<status>status</status>
		<type>type</type>
		<amount>4.1</amount>
		<order-id>orderid</order-id>
		<payment-method-token>token</payment-method-token>
		<merchant-account-id>accountid</merchant-account-id>
		<plan-id>planid</plan-id>
		<created-at>createdat</created-at>
		<updated-at>updatedat</updated-at>
		<processor-authorization-code>authcode</processor-authorization-code>
	</transaction>`)

	expectedTx := &Transaction{
		Id:                  "txid",
		CustomerID:          "customerid",
		Status:              "status",
		Type:                "type",
		Amount:              4.1,
		AmountStr:           "4.1",
		OrderId:             "orderid",
		PaymentMethodToken:  "token",
		MerchantAccountId:   "accountid",
		PlanId:              "planid",
		ServiceFeeAmount:    2.1,
		ServiceFeeAmountStr: "2.10",
		CreatedAt:           "createdat",
		UpdatedAt:           "updatedat",
		AuthCode:            "authcode",
	}

	r := &Response{
		Body: txXml,
	}

	tx, err := r.transaction()
	if err != nil {
		t.Fatalf("failed to parse transaction: %v", err)
	}
	if !reflect.DeepEqual(tx, expectedTx) {
		t.Error("actual and expected transactions differ")
		t.Logf("actual: %+v", tx)
		t.Logf("expected: %+v", expectedTx)
	}
}
