package braintree

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

const (
	cardToUse = "Discover"
)

func TestSettlementBatch(t *testing.T) {
	// Get current batch summary
	y, m, d := time.Now().Date()
	date := fmt.Sprintf("%d-%d-%d", y, m, d)
	batchSummary, err := testGateway.Settlement().Generate(&Settlement{Date: date})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(batchSummary)

	// Get the card types
	cardTypes := []string{}
	for _, record := range batchSummary.Records.Type {
		cardTypes = append(cardTypes, record.CardType)
	}

	// Create a new transaction to add 12.34 to the summary
	tx, err := testGateway.Transaction().Create(&Transaction{
		Type:   "sale",
		Amount: NewDecimal(1234, 2),
		CreditCard: &CreditCard{
			Number:         testCreditCards[strings.ToLower(cardToUse)].Number,
			ExpirationDate: "05/14",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx)
	if tx.Id == "" {
		t.Fatal("Received invalid ID on new transaction")
	}
	if tx.Status != "authorized" {
		t.Fatal(tx.Status)
	}

	// Submit for settlement
	ten := NewDecimal(1234, 2)
	tx2, err := testGateway.Transaction().SubmitForSettlement(tx.Id, ten)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx2)
	if x := tx2.Status; x != "submitted_for_settlement" {
		t.Fatal(x)
	}
	if amount := tx2.Amount; amount.Cmp(ten) != 0 {
		t.Fatalf("transaction settlement amount (%s) did not equal amount requested (%s)", amount, ten)
	}

	// Settle
	tx3, err := testGateway.Transaction().Settle(tx.Id)
	t.Log(tx3)
	if err != nil {
		t.Fatal(err)
	}
	if x := tx3.Status; x != "settled" {
		t.Fatal(x)
	}

	// Generate Settlement Batch Summary which will include new transaction
	batchSummary, err = testGateway.Settlement().Generate(&Settlement{Date: date})
	if err != nil {
		t.Fatal(fmt.Sprintf("Unable to get settlement batch: err is %s", err.Error()))
	}
	t.Log(batchSummary)

	// Since these tests are run concurrently, we will not test the amount  only the card types.
	foundTypes := []string{}
	for _, record := range batchSummary.Records.Type {
		foundTypes = append(foundTypes, record.CardType)
	}
	if !reflect.DeepEqual(cardTypes, foundTypes) {
		t.Fatal(fmt.Sprintf("Expected card types: %s, got: %s", cardTypes, foundTypes))
	}
}
