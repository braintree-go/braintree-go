package braintree

import "testing"

func TestCreatePaymentMethod(t *testing.T) {
	cust, err := testGateway.Customer().Create(&Customer{})
	if err != nil {
		t.Fatal(err)
	}

	nonce := FakeNonceTransactable

	paymentMethod, err := testGateway.PaymentMethod().Create(&PaymentMethodRequest{
		CustomerId:         cust.Id,
		PaymentMethodNonce: nonce,
	})
	if err != nil {
		t.Fatal(err)
	}

	if paymentMethod.GetCustomerId() != cust.Id {
		t.Errorf("Got paymentMethod customer Id %#v, want %#v", paymentMethod.GetCustomerId(), cust.Id)
	}
	if paymentMethod.GetToken() == "" {
		t.Errorf("Got paymentMethod token %#v, want a value", paymentMethod.GetToken())
	}
}
