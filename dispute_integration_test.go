//+build integration

package braintree

import (
	"context"
	"testing"
	"time"
)

func TestDisputeFinalize(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tx, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:   "sale",
		Amount: NewDecimal(100, 2),
		CreditCard: &CreditCard{
			Number:         "4023898493988028",
			ExpirationDate: "12/" + time.Now().Format("2006"),
		},
		Options: &TransactionOptions{
			SubmitForSettlement: true,
		},
	})
	if err != nil {
		t.Fatalf("failed to create disputed transaction: %v", err)
	}

	tx, err = testGateway.Transaction().Find(ctx, tx.Id)
	if err != nil {
		t.Fatalf("failed to find disputed transaction: %v", err)
	}

	if len(tx.Disputes) != 1 {
		t.Fatalf("got Transaction with %d disputes, want 1", len(tx.Disputes))
	}

	dispute := tx.Disputes[0]

	if dispute.AmountDisputed.Cmp(NewDecimal(100, 2)) != 0 {
		t.Errorf("got Dispute AmountDisputed %s, want %s", dispute.AmountDisputed, "1.00")
	}

	err = testGateway.Dispute().Finalize(ctx, dispute.ID)

	if err != nil {
		t.Fatalf("failed to finalize dispute: %v", err)
	}
}

func TestDisputeAccept(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tx, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:   "sale",
		Amount: NewDecimal(100, 2),
		CreditCard: &CreditCard{
			Number:         "4023898493988028",
			ExpirationDate: "12/" + time.Now().Format("2006"),
		},
		Options: &TransactionOptions{
			SubmitForSettlement: true,
		},
	})

	if err != nil {
		t.Fatalf("failed to create disputed transaction: %v", err)
	}

	tx, err = testGateway.Transaction().Find(ctx, tx.Id)
	if err != nil {
		t.Fatalf("failed to find disputed transaction: %v", err)
	}

	if len(tx.Disputes) != 1 {
		t.Fatalf("transaction has %d disputes, want 1", len(tx.Disputes))
	}

	dispute := tx.Disputes[0]

	err = testGateway.Dispute().Accept(ctx, dispute.ID)
	if err != nil {
		t.Fatalf("failed to accept dispute: %v", err)
	}

	dispute, err = testGateway.Dispute().Find(ctx, dispute.ID)
	if err != nil {
		t.Fatalf("failed to find dispute: %v", err)
	}

	if dispute.Status != DisputeStatusAccepted {
		t.Fatalf("got Dispute Status %q, want %q", dispute.Status, DisputeStatusAccepted)
	}
}

func TestDisputeTextEvidence(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tx, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:   "sale",
		Amount: NewDecimal(100, 2),
		CreditCard: &CreditCard{
			Number:         "4023898493988028",
			ExpirationDate: "12/" + time.Now().Format("2006"),
		},
		Options: &TransactionOptions{
			SubmitForSettlement: true,
		},
	})
	if err != nil {
		t.Fatalf("failed to create disputed transaction: %v", err)
	}

	tx, err = testGateway.Transaction().Find(ctx, tx.Id)
	if err != nil {
		t.Fatalf("failed to find disputed transaction: %v", err)
	}

	if len(tx.Disputes) != 1 {
		t.Fatalf("got Transaction with %d disputes, want 1", len(tx.Disputes))
	}

	dispute := tx.Disputes[0]

	textEvidence, err := testGateway.Dispute().AddTextEvidence(ctx, dispute.ID, &DisputeTextEvidenceRequest{
		Content:  "some evidence",
		Category: DisputeEvidenceCategoryDeviceName,
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
