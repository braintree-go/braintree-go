//+build integration

package braintree

import (
	"context"
	"testing"
)

var (
	disputedTransaction = TransactionRequest{
		Type:   "sale",
		Amount: NewDecimal(10, 2),
		CreditCard: &CreditCard{
			Number:         "4023898493988028",
			ExpirationDate: "01/2020",
		},
		Options: &TransactionOptions{
			SubmitForSettlement: true,
		},
	}
)

func TestProcessAndFinalizeDispute(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tx, err := testGateway.Transaction().Create(ctx, &disputedTransaction)

	if err != nil {
		t.Fatalf("failed to create disputed transaction: %v", err)
	}

	query := new(SearchQuery)
	transactionId := query.AddTextField("transaction_id")
	transactionId.Is = tx.Id

	disputes, err := testGateway.Dispute().Search(ctx, query)

	if err != nil {
		t.Fatalf("failed to search for disputes: %v", err)
	}

	if len(disputes) == 0 {
		t.Fatalf("at least one dispute object should be created")
	}

	dispute := disputes[0]

	if dispute.AmountDisputed.Cmp(disputedTransaction.Amount) != 0 {
		t.Errorf("expected AmountDisputed to be %s, was %s", disputedTransaction.Amount, dispute.AmountDisputed)
	}

	foundDispute, err := testGateway.Dispute().Find(ctx, dispute.ID)

	if foundDispute.AmountDisputed.Cmp(dispute.AmountDisputed) != 0 {
		t.Fatalf("disputes with the same id should have equal amounts")
	}

	textEvidence, err := testGateway.Dispute().AddTextEvidence(ctx, dispute.ID, &DisputeTextEvidenceRequest{
		Content:  "some-id",
		Category: EvidenceCategoryDeviceName,
	})

	if err != nil {
		t.Fatalf("failed to add text evidence: %v", err)
	}

	if textEvidence.ID == "" {
		t.Fatal("text evidence can not have empty id")
	}

	err = testGateway.Dispute().RemoveEvidence(ctx, dispute.ID, textEvidence.ID)

	if err != nil {
		t.Fatalf("failed to remove evidence: %v", err)
	}

	err = testGateway.Dispute().Finalize(ctx, dispute.ID)

	if err != nil {
		t.Fatalf("failed to finalize dispute: %v", err)
	}

}

func TestAcceptDispute(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	disputedTransaction.Amount = NewDecimal(100, 2)
	tx, err := testGateway.Transaction().Create(ctx, &disputedTransaction)

	if err != nil {
		t.Fatalf("failed to create disputed transaction: %v", err)
	}

	query := new(SearchQuery)
	transactionId := query.AddTextField("transaction_id")
	transactionId.Is = tx.Id

	disputes, err := testGateway.Dispute().Search(ctx, query)

	if err != nil {
		t.Fatalf("failed to search for disputes: %v", err)
	}

	if len(disputes) == 0 {
		t.Fatalf("at least one dispute object should be created")
	}

	dispute := disputes[0]

	err = testGateway.Dispute().Accept(ctx, dispute.ID)

	if err != nil {
		t.Fatalf("failed to accept dispute: %v", err)
	}
}
