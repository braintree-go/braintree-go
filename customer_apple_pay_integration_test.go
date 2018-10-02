// +build integration

package braintree

import (
	"context"
	"reflect"
	"regexp"
	"testing"
)

var isValidBIN = regexp.MustCompile(`^\d{6}$`).MatchString
var isValidLast4 = regexp.MustCompile(`^\d{4}$`).MatchString
var isValidExpiryMonth = regexp.MustCompile(`^\d{2}$`).MatchString
var isValidExpiryYear = regexp.MustCompile(`^\d{4}$`).MatchString

func TestCustomerApplePayCard(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	nonceCardType := "Apple Pay - Visa"

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

	if applePayCard.CardType != nonceCardType {
		t.Errorf("Got ApplePayCard.CardType %v, want %v", applePayCard.CardType, nonceCardType)
	}

	if !isValidExpiryMonth(applePayCard.ExpirationMonth) {
		t.Errorf("ApplePayCard.ExpirationMonth (%s) does not conform expected value", applePayCard.ExpirationMonth)
	}

	if !isValidExpiryYear(applePayCard.ExpirationYear) {
		t.Errorf("ApplePayCard.ExpirationYear (%s) does not conform expected value", applePayCard.ExpirationYear)
	}

	if !isValidBIN(applePayCard.BIN) {
		t.Errorf("ApplePayCard.BIN (%s) does not conform expected value", applePayCard.BIN)
	}

	if !isValidLast4(applePayCard.Last4) {
		t.Errorf("ApplePayCard.Last4 (%s) does not conform expected value", applePayCard.Last4)
	}
}
