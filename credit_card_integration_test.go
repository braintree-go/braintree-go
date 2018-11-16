// +build integration

package braintree

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"time"

	"github.com/braintree-go/braintree-go/testhelpers"
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
		Number:         testCardVisa,
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
	if card.ProductID != "Unknown" { // Unknown appears to be reported from the Braintree Sandbox environment for all cards.
		t.Errorf("got product id %q, want %q", card.ProductID, "Unknown")
	}

	// Update
	card2, err := g.Update(ctx, &CreditCard{
		Token:          card.Token,
		Number:         testCardMastercard,
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
		Number:          testCardVisa,
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
		Number:         testCardVisa,
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
		Number:         testCardVisa,
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
		VenmoSDKPaymentMethodCode: "stub-" + testCardVisa,
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
		Number:         testCardVisa,
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
		Number:         testCardVisa,
		ExpirationDate: now.AddDate(0, -2, 0).Format("01/2006"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("card1", card1.Token)

	card2, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCardVisa,
		ExpirationDate: now.Format("01/2006"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("card2", card2.Token)

	card3, err := testGateway.CreditCard().Create(ctx, &CreditCard{
		CustomerId:     customer.Id,
		Number:         testCardVisa,
		ExpirationDate: now.AddDate(0, 2, 0).Format("01/2006"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("card3", card3.Token)

	fromDate := now.AddDate(0, -1, 0)
	toDate := now.AddDate(0, 1, 0)

	expiringCards := map[string]bool{}
	results, err := testGateway.CreditCard().ExpiringBetweenIDs(ctx, fromDate, toDate)
	if err != nil {
		t.Fatal(err)
	}
	for page := 1; page <= results.PageCount; page++ {
		results, err := testGateway.CreditCard().ExpiringBetweenPage(ctx, fromDate, toDate, results, page)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Iterating page %d (page size: %d, total items: %d)", results.CurrentPageNumber, results.PageSize, results.TotalItems)
		for _, card := range results.CreditCards {
			expiringCards[card.Token] = true
		}
	}
	if expiringCards[card1.Token] {
		t.Fatalf("expiringCards contains card1 (%s), it shouldn't be returned in expiring cards results", card1.Token)
	}
	if !expiringCards[card2.Token] {
		t.Fatalf("expiringCards does not contain card2 (%s), it should be returned in expiring cards results", card2.Token)
	}
	if expiringCards[card3.Token] {
		t.Fatalf("expiringCards contains card3 (%s), it shouldn't be returned in expiring cards results", card3.Token)
	}

	_, err = testGateway.CreditCard().ExpiringBetweenPage(ctx, fromDate, toDate, results, 0)
	t.Logf("%#v", err)
	if err == nil || !strings.Contains(err.Error(), "page 0 out of bounds") {
		t.Errorf("requesting page 0 should result in out of bounds error, but got %#v", err)
	}

	_, err = testGateway.CreditCard().ExpiringBetweenPage(ctx, fromDate, toDate, results, results.PageCount+1)
	t.Logf("%#v", err)
	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf("page %d out of bounds", results.PageCount+1)) {
		t.Errorf("requesting page %d should result in out of bounds error, but got %v", results.PageCount+1, err)
	}
}
