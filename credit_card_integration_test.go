// +build integration

package braintree

import (
	"context"
	"testing"

	"time"

	"github.com/lionelbarrow/braintree-go/testhelpers"
)

func TestCreditCard(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cust, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.CreditCard()
	card, err := g.Create(ctx, &CreditCard{
		CustomerId:     cust.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: testhelpers.BoolPtr(true),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(card)

	if card.Token == "" {
		t.Fatal("invalid token")
	}

	// Update
	card2, err := g.Update(ctx, &CreditCard{
		Token:          card.Token,
		Number:         testCreditCards["mastercard"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: testhelpers.BoolPtr(true),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(card2)

	if card2.Token != card.Token {
		t.Fatal("tokens do not match")
	}
	if card2.CardType != "MasterCard" {
		t.Fatal("card type does not match")
	}

	// Delete
	err = g.Delete(ctx, card2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreditCardFailedAutoVerification(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cust, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.CreditCard()
	card, err := g.Create(ctx, &CreditCard{
		CustomerId:         cust.Id,
		PaymentMethodNonce: FakeNonceProcessorDeclinedVisa,
	})
	if err == nil {
		t.Fatal("Got no error, want error")
	}
	if g, w := err.(*BraintreeError).ErrorMessage, "Do Not Honor"; g != w {
		t.Fatalf("Got error %q, want error %q", g, w)
	}

	t.Logf("%#v\n", err)
	t.Logf("%#v\n", card)
}

func TestCreditCardForceNotVerified(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cust, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.CreditCard()
	card, err := g.Create(ctx, &CreditCard{
		CustomerId:         cust.Id,
		PaymentMethodNonce: FakeNonceProcessorDeclinedVisa,
		Options: &CreditCardOptions{
			VerifyCard: testhelpers.BoolPtr(false),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", card)
}

func TestCreateCreditCardWithExpirationMonthAndYear(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	card, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:      customer.Id,
		Number:          testCreditCards["visa"].Number,
		ExpirationMonth: "05",
		ExpirationYear:  "2014",
		CVV:             "100",
	})

	if err != nil {
		t.Fatal(err)
	}
	if card.Token == "" {
		t.Fatal("invalid token")
	}
}

func TestCreateCreditCardInvalidInput(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	card, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
	})

	t.Log(card)

	// This test should fail because customer id is required
	if err == nil {
		t.Fatal("expected to get error creating card because of required fields, but did not")
	}

	// TODO: validate fields
}

func TestFindCreditCard(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	card, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: testhelpers.BoolPtr(true),
		},
	})

	t.Log(card)

	if err != nil {
		t.Fatal(err)
	}
	if card.Token == "" {
		t.Fatal("invalid token")
	}

	card2, err := testGateway.CreditCard().Find(ctx, card.Token)

	t.Log(card2)

	if err != nil {
		t.Fatal(err)
	}
	if card2.Token != card.Token {
		t.Fatal("tokens do not match")
	}
}

func TestFindCreditCardBadData(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	card, err := testGateway.CreditCard().Find(ctx, "invalid_token")

	t.Log(card)

	if err == nil {
		t.Fatal("expected to get error because the token is invalid")
	}
}

func TestSaveCreditCardWithVenmoSDKPaymentMethodCode(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	card, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:                customer.Id,
		VenmoSDKPaymentMethodCode: "stub-" + testCreditCards["visa"].Number,
	})
	if err != nil {
		t.Fatal(err)
	}
	if card.VenmoSDK {
		t.Fatal("venmo card marked")
	}
}

func TestSaveCreditCardWithVenmoSDKSession(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	card, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		Options: &CreditCardOptions{
			VenmoSDKSession: "stub-session",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if card.VenmoSDK {
		t.Fatal("venmo card marked")
	}
}

func TestGetExpiredCards(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	card1, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/18",
	})
	if err != nil {
		t.Fatal(err)
	}

	card2, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "07/20",
	})
	if err != nil {
		t.Fatal(err)
	}

	expiredCards, err := testGateway.CreditCard().Expired(ctx)
	if err != nil {
		t.Fatal(err)
	}
	isCard1Expired := false
	isCard2Expired := false
	for _, card := range expiredCards {
		if card.Token == card1.Token {
			isCard1Expired = true
		}
		if card.Token == card2.Token {
			isCard2Expired = true
		}
	}

	if !isCard1Expired || isCard2Expired {
		t.Fatal("Fail")
	}
}

func TestGetExpiringBetweenCards(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	card1, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/18",
	})
	if err != nil {
		t.Fatal(err)
	}

	card2, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "07/20",
	})
	if err != nil {
		t.Fatal(err)
	}

	fromTime := time.Unix(1524740667, 0) // 04/18
	toTime := time.Unix(1530011067, 0)   // 06/18

	expiredCards, err := testGateway.CreditCard().ExpiringBetween(ctx, fromTime, toTime)
	if err != nil {
		t.Fatal(err)
	}
	isCard1Expired := false
	isCard2Expired := false
	for _, card := range expiredCards {
		if card.Token == card1.Token {
			isCard1Expired = true
		}
		if card.Token == card2.Token {
			isCard2Expired = true
		}
	}

	if !isCard1Expired || isCard2Expired {
		t.Fatal("Fail")
	}
}
