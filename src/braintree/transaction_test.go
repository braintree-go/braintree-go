package braintree

import (
  "testing"
)

func Test_Transaction_Create(t *testing.T) {
  config := Configuration{environment: "sandbox", merchant_id: "integration_merchant_id", public_key: "integration_public_key", private_key: "integration_private_key"}
  gateway := NewGateway(config)

  request := NewTransactionRequest().Amount(100).CreditCard().Number(TestCards["visa"].Number).ExpirationDate("05/2014").Done()
  response := gateway.Transaction().Sale(request)
  if !response.IsValid() {
    t.Log("Transaction create response was invalid")
    t.Fail()
  }
}
