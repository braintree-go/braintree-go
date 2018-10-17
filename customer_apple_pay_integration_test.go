// +build integration

package braintree

import (
	"context"
	"reflect"
	"testing"

	"github.com/braintree-go/braintree-go/testhelpers"
)

func TestCustomerApplePayCard(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceApplePayVisa,
	})
	if err != nil {
		t.Fatal(err)
	}
	applePayCard := paymentMethod.(*ApplePayCard)

	customerFound, err := testGateway.Customer().Find(ctx, customer.Id)
	if err != nil {
		t.Fatal(err)
	}

	if customerFound.ApplePayCards == nil || len(customerFound.ApplePayCards.ApplePayCard) != 1 {
		t.Fatalf("Customer %#v expected to have one ApplePayCard", customerFound)
	}
	if !reflect.DeepEqual(customerFound.ApplePayCards.ApplePayCard[0], applePayCard) {
		t.Fatalf("Got Customer %#v ApplePayCard %#v, want %#v", customerFound, customerFound.ApplePayCards.ApplePayCard[0], applePayCard)
	}

	wantNonceCardType := "Apple Pay - Visa"
	if applePayCard.CardType != wantNonceCardType {
		t.Errorf("Got ApplePayCard.CardType %v, want %v", applePayCard.CardType, wantNonceCardType)
	}
	if !testhelpers.ValidExpiryMonth(applePayCard.ExpirationMonth) {
		t.Errorf("ApplePayCard.ExpirationMonth (%s) does not conform expected value", applePayCard.ExpirationMonth)
	}
	if !testhelpers.ValidExpiryYear(applePayCard.ExpirationYear) {
		t.Errorf("ApplePayCard.ExpirationYear (%s) does not conform expected value", applePayCard.ExpirationYear)
	}
	if !testhelpers.ValidBIN(applePayCard.BIN) {
		t.Errorf("ApplePayCard.BIN (%s) does not conform expected value", applePayCard.BIN)
	}
	if !testhelpers.ValidLast4(applePayCard.Last4) {
		t.Errorf("ApplePayCard.Last4 (%s) does not conform expected value", applePayCard.Last4)
	}
}
