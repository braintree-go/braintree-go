//+build integration

package braintree

import (
	"context"
	"testing"
	"time"
)

func TestDisputeSearch(t *testing.T) {
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

	dg := testGateway.Dispute()
	dispute := tx.Disputes[0]

	if dispute.AmountDisputed.Cmp(NewDecimal(100, 2)) != 0 {
		t.Errorf("got Dispute AmountDisputed %s, want %s", dispute.AmountDisputed, "1.00")
	}

	query := new(SearchQuery)
	f := query.AddTextField("id")
	f.Is = dispute.ID

	result, err := dg.Search(ctx, query)
	if err != nil {
		t.Fatalf("failed to search dispute: %v", err)
	}

	if len(result.Disputes) != 1 {
		t.Fatalf("expected 1 dispute, but got %d", len(result.Disputes))
	}

	if result.Disputes[0].ID != dispute.ID {
		t.Fatalf("expected transaction with id %s, but got %s", result.Disputes[0].ID, dispute.ID)
	}

	err = dg.Finalize(ctx, dispute.ID)

	if err != nil {
		t.Fatalf("failed to finalize dispute: %v", err)
	}
}

func TestDisputeSearchPage(t *testing.T) {
	ctx := context.Background()
	txg := testGateway.Transaction()
	dg := testGateway.Dispute()
	cg := testGateway.Customer()

	customer, err := cg.Create(ctx, &CustomerRequest{
		FirstName: "John",
		LastName:  "Smith",
	})

	if err != nil {
		t.Fatalf("failed to created a customer: %v", err)
	}

	const transactionCount = 51
	createdDisputeIDs := map[string]bool{}
	for i := 0; i < transactionCount; i++ {
		tx, err := txg.Create(ctx, &TransactionRequest{
			Type:       "sale",
			Amount:     NewDecimal(100, 2),
			CustomerID: customer.Id,
			CreditCard: &CreditCard{
				Number:         "4023898493988028",
				ExpirationDate: "12/" + time.Now().Format("2006"),
			},
			Options: &TransactionOptions{
				SubmitForSettlement: true,
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		createdDisputeIDs[tx.Disputes[0].ID] = true
	}

	query := new(SearchQuery)
	mf := query.AddMultiField("kind")
	mf.Items = []string{string(DisputeKindChargeback)}
	mf = query.AddMultiField("status")
	mf.Items = []string{string(DisputeStatusOpen)}
	tf := query.AddTextField("customer-id")
	tf.Is = customer.Id

	var result *DisputeSearchResult
	var matchedDisputeIDs []string
	var page int = 1

	for {
		result, err = dg.SearchPage(ctx, query, result, page)
		if err != nil {
			t.Fatalf("failed to search dispute: %v; page %d", err, page)
		}

		if result.PageCount != 2 {
			t.Fatalf("expected 2 pages of disputes, but got %d", result.PageCount)
		}

		if result.TotalItems != transactionCount {
			t.Fatalf("expected %d disputes, but got %d", transactionCount, result.TotalItems)
		}

		for _, dispute := range result.Disputes {
			matchedDisputeIDs = append(matchedDisputeIDs, dispute.ID)
		}

		page++
		if page > result.PageCount {
			break
		}
	}

	for _, disputeID := range matchedDisputeIDs {
		err = dg.Finalize(ctx, disputeID)
		if err != nil {
			t.Fatalf("failed to finalize dispute: %v", err)
		}
		delete(createdDisputeIDs, disputeID)
	}

	if len(createdDisputeIDs) > 0 {
		t.Fatalf("disputes not returned = %v", createdDisputeIDs)
	}
}

func TestDisputeSearchNext(t *testing.T) {
	ctx := context.Background()
	txg := testGateway.Transaction()
	dg := testGateway.Dispute()
	cg := testGateway.Customer()

	customer, err := cg.Create(ctx, &CustomerRequest{
		FirstName: "John",
		LastName:  "Smith",
	})

	if err != nil {
		t.Fatalf("failed to created a customer: %v", err)
	}

	const transactionCount = 51
	createdDisputeIDs := map[string]bool{}
	for i := 0; i < transactionCount; i++ {
		tx, err := txg.Create(ctx, &TransactionRequest{
			Type:       "sale",
			Amount:     NewDecimal(100, 2),
			CustomerID: customer.Id,
			CreditCard: &CreditCard{
				Number:         "4023898493988028",
				ExpirationDate: "12/" + time.Now().Format("2006"),
			},
			Options: &TransactionOptions{
				SubmitForSettlement: true,
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		createdDisputeIDs[tx.Disputes[0].ID] = true
	}

	query := new(SearchQuery)
	mf := query.AddMultiField("kind")
	mf.Items = []string{string(DisputeKindChargeback)}
	mf = query.AddMultiField("status")
	mf.Items = []string{string(DisputeStatusOpen)}
	tf := query.AddTextField("customer-id")
	tf.Is = customer.Id

	var index int = 0
	var matchedDisputeIDs []string
	var result *DisputeSearchResult

	for {
		if index == 0 {
			result, err = dg.Search(ctx, query)
		} else {
			result, err = dg.SearchNext(ctx, query, result)
		}

		if err != nil {
			t.Fatalf("failed to search dispute: %v", err)
		}

		if result == nil && err == nil {
			break
		}

		if result.PageCount != 2 {
			t.Fatalf("expected 2 pages of disputes, but got %d", result.PageCount)
		}

		if result.TotalItems != transactionCount {
			t.Fatalf("expected %d disputes, but got %d", transactionCount, result.TotalItems)
		}

		for _, dispute := range result.Disputes {
			matchedDisputeIDs = append(matchedDisputeIDs, dispute.ID)
		}
		index++
	}

	for _, disputeID := range matchedDisputeIDs {
		err = dg.Finalize(ctx, disputeID)
		if err != nil {
			t.Fatalf("failed to finalize dispute: %v", err)
		}
		delete(createdDisputeIDs, disputeID)
	}

	if len(createdDisputeIDs) > 0 {
		t.Fatalf("disputes not returned = %v", createdDisputeIDs)
	}
}

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
