// +build integration

package braintree

import (
	"context"
	"strconv"
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

func TestGetExpiringBetweenCards(t *testing.T) {
	now := time.Now()

	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	card1, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "01/" + strconv.Itoa(now.Year()-2),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("card1", card1.Token)

	card2, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "12/" + strconv.Itoa(now.Year()),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("card2", card2.Token)

	card3, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "01/" + strconv.Itoa(now.Year()+2),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("card3", card3.Token)

	fromDate := now.AddDate(-1, 0, 0)
	toDate := now.AddDate(1, 0, 0)

	expiringCards := map[string]bool{}
	results, err := testGateway.CreditCard().ExpiringBetween(ctx, fromDate, toDate)
	if err != nil {
		t.Fatal(err)
	}
	for {
		t.Logf("Iterating page %d (page size: %d, total items: %d)", results.CurrentPageNumber, results.PageSize, results.TotalItems)
		for _, card := range results.CreditCards {
			expiringCards[card.Token] = true
		}
		results, err = testGateway.CreditCard().ExpiringBetweenNext(ctx, fromDate, toDate, results)
		if err != nil {
			t.Fatal(err)
		}
		if results == nil {
			break
		}
	}
	if expiringCards[card1.Token] {
		t.Fatalf("expiringCards contains card1 (%s), it shouldn't be expired", card1.Token)
	}
	if !expiringCards[card2.Token] {
		t.Fatalf("expiringCards does not contain card2 (%s), it should be expired", card2.Token)
	}
	if expiringCards[card3.Token] {
		t.Fatalf("expiringCards contains card3 (%s), it shouldn't be expired", card3.Token)
	}
}
