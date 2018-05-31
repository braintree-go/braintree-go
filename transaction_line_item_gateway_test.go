// +build integration

package braintree

import (
	"context"
	"testing"
)

func TestTransactionWithLineItemsZero(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(txn.Id)

	lineItems, err := testGateway.TransactionLineItem().Find(ctx, txn.Id)

	if err != nil {
		t.Fatal(err)
	}

	if g, w := len(lineItems), 0; g != w {
		t.Fatalf("got %d line items, want %d line items", g, w)
	}
}

func TestTransactionWithLineItemsSingleOnlyRequiredFields(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:        "Name #1",
				Kind:        TransactionLineItemKindDebit,
				Quantity:    NewDecimal(10232, 4),
				UnitAmount:  NewDecimal(451232, 4),
				TotalAmount: NewDecimal(4515, 2),
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(txn.Id)

	lineItems, err := testGateway.TransactionLineItem().Find(ctx, txn.Id)

	if err != nil {
		t.Fatal(err)
	}

	if g, w := len(lineItems), 1; g != w {
		t.Fatalf("got %d line items, want %d line items", g, w)
	}

	l := lineItems[0]
	if g, w := l.Name, "Name #1"; g != w {
		t.Errorf("got name %q, want %q", g, w)
	}
	if g, w := l.Kind, TransactionLineItemKindDebit; g != w {
		t.Errorf("got kind %q, want %q", g, w)
	}
	if g, w := l.Quantity, NewDecimal(10232, 4); g.Cmp(w) != 0 {
		t.Errorf("got quantity %q, want %q", g, w)
	}
	if g, w := l.UnitAmount, NewDecimal(451232, 4); g.Cmp(w) != 0 {
		t.Errorf("got unit amount %q, want %q", g, w)
	}
	if g, w := l.TotalAmount, NewDecimal(4515, 2); g.Cmp(w) != 0 {
		t.Errorf("got total amount %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsSingleZeroAmountFields(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(0, 0),
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(0, 0),
				DiscountAmount: NewDecimal(0, 0),
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(txn.Id)

	lineItems, err := testGateway.TransactionLineItem().Find(ctx, txn.Id)

	if err != nil {
		t.Fatal(err)
	}

	if g, w := len(lineItems), 1; g != w {
		t.Fatalf("got %d line items, want %d line items", g, w)
	}

	l := lineItems[0]
	if g, w := l.Name, "Name #1"; g != w {
		t.Errorf("got name %q, want %q", g, w)
	}
	if g, w := l.Kind, TransactionLineItemKindDebit; g != w {
		t.Errorf("got kind %q, want %q", g, w)
	}
	if g, w := l.Quantity, NewDecimal(10232, 4); g.Cmp(w) != 0 {
		t.Errorf("got quantity %q, want %q", g, w)
	}
	if g, w := l.UnitAmount, NewDecimal(451232, 4); g.Cmp(w) != 0 {
		t.Errorf("got unit amount %q, want %q", g, w)
	}
	if g, w := l.UnitTaxAmount, NewDecimal(0, 0); g.Cmp(w) != 0 {
		t.Errorf("got unit tax amount %q, want %q", g, w)
	}
	if g, w := l.TotalAmount, NewDecimal(4515, 2); g.Cmp(w) != 0 {
		t.Errorf("got total amount %q, want %q", g, w)
	}
	if g, w := l.TaxAmount, NewDecimal(0, 0); g.Cmp(w) != 0 {
		t.Errorf("got tax amount %q, want %q", g, w)
	}
	if g, w := l.DiscountAmount, NewDecimal(0, 0); g.Cmp(w) != 0 {
		t.Errorf("got discount amount %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsSingle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(txn.Id)

	lineItems, err := testGateway.TransactionLineItem().Find(ctx, txn.Id)

	if err != nil {
		t.Fatal(err)
	}

	if g, w := len(lineItems), 1; g != w {
		t.Fatalf("got %d line items, want %d line items", g, w)
	}

	l := lineItems[0]
	if g, w := l.Name, "Name #1"; g != w {
		t.Errorf("got name %q, want %q", g, w)
	}
	if g, w := l.Description, "Description #1"; g != w {
		t.Errorf("got description %q, want %q", g, w)
	}
	if g, w := l.Kind, TransactionLineItemKindDebit; g != w {
		t.Errorf("got kind %q, want %q", g, w)
	}
	if g, w := l.Quantity, NewDecimal(10232, 4); g.Cmp(w) != 0 {
		t.Errorf("got quantity %q, want %q", g, w)
	}
	if g, w := l.UnitAmount, NewDecimal(451232, 4); g.Cmp(w) != 0 {
		t.Errorf("got unit amount %q, want %q", g, w)
	}
	if g, w := l.UnitTaxAmount, NewDecimal(123, 2); g.Cmp(w) != 0 {
		t.Errorf("got unit tax amount %q, want %q", g, w)
	}
	if g, w := l.UnitOfMeasure, "gallon"; g != w {
		t.Errorf("got unit of measure %q, want %q", g, w)
	}
	if g, w := l.TotalAmount, NewDecimal(4515, 2); g.Cmp(w) != 0 {
		t.Errorf("got total amount %q, want %q", g, w)
	}
	if g, w := l.TaxAmount, NewDecimal(455, 2); g.Cmp(w) != 0 {
		t.Errorf("got tax amount %q, want %q", g, w)
	}
	if g, w := l.DiscountAmount, NewDecimal(102, 2); g.Cmp(w) != 0 {
		t.Errorf("got discount amount %q, want %q", g, w)
	}
	if g, w := l.ProductCode, "23434"; g != w {
		t.Errorf("got product code %q, want %q", g, w)
	}
	if g, w := l.CommodityCode, "9SAASSD8724"; g != w {
		t.Errorf("got commodity code %q, want %q", g, w)
	}
	if g, w := l.URL, "https://example.com/products/23434"; g != w {
		t.Errorf("got url %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsMultiple(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:          "Name #2",
				Description:   "Description #2",
				Kind:          TransactionLineItemKindCredit,
				Quantity:      NewDecimal(202, 2),
				UnitAmount:    NewDecimal(5, 0),
				UnitOfMeasure: "gallon",
				TotalAmount:   NewDecimal(4515, 2),
				TaxAmount:     NewDecimal(455, 2),
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(txn.Id)

	lineItems, err := testGateway.TransactionLineItem().Find(ctx, txn.Id)

	if err != nil {
		t.Fatal(err)
	}

	if g, w := len(lineItems), 2; g != w {
		t.Fatalf("got %d line items, want %d line items", g, w)
	}

	{
		l := lineItems[0]
		if g, w := l.Name, "Name #1"; g != w {
			t.Errorf("got name %q, want %q", g, w)
		}
		if g, w := l.Description, "Description #1"; g != w {
			t.Errorf("got description %q, want %q", g, w)
		}
		if g, w := l.Kind, TransactionLineItemKindDebit; g != w {
			t.Errorf("got kind %q, want %q", g, w)
		}
		if g, w := l.Quantity, NewDecimal(10232, 4); g.Cmp(w) != 0 {
			t.Errorf("got quantity %q, want %q", g, w)
		}
		if g, w := l.UnitAmount, NewDecimal(451232, 4); g.Cmp(w) != 0 {
			t.Errorf("got unit amount %q, want %q", g, w)
		}
		if g, w := l.UnitTaxAmount, NewDecimal(123, 2); g.Cmp(w) != 0 {
			t.Errorf("got unit tax amount %q, want %q", g, w)
		}
		if g, w := l.UnitOfMeasure, "gallon"; g != w {
			t.Errorf("got unit of measure %q, want %q", g, w)
		}
		if g, w := l.TotalAmount, NewDecimal(4515, 2); g.Cmp(w) != 0 {
			t.Errorf("got total amount %q, want %q", g, w)
		}
		if g, w := l.TaxAmount, NewDecimal(455, 2); g.Cmp(w) != 0 {
			t.Errorf("got tax amount %q, want %q", g, w)
		}
		if g, w := l.DiscountAmount, NewDecimal(102, 2); g.Cmp(w) != 0 {
			t.Errorf("got discount amount %q, want %q", g, w)
		}
		if g, w := l.ProductCode, "23434"; g != w {
			t.Errorf("got product code %q, want %q", g, w)
		}
		if g, w := l.CommodityCode, "9SAASSD8724"; g != w {
			t.Errorf("got commodity code %q, want %q", g, w)
		}
		if g, w := l.URL, "https://example.com/products/23434"; g != w {
			t.Errorf("got url %q, want %q", g, w)
		}
	}

	{
		l := lineItems[1]
		if g, w := l.Name, "Name #2"; g != w {
			t.Errorf("got name %q, want %q", g, w)
		}
		if g, w := l.Description, "Description #2"; g != w {
			t.Errorf("got description %q, want %q", g, w)
		}
		if g, w := l.Kind, TransactionLineItemKindCredit; g != w {
			t.Errorf("got kind %q, want %q", g, w)
		}
		if g, w := l.Quantity, NewDecimal(202, 2); g.Cmp(w) != 0 {
			t.Errorf("got quantity %q, want %q", g, w)
		}
		if g, w := l.UnitAmount, NewDecimal(5, 0); g.Cmp(w) != 0 {
			t.Errorf("got unit amount %q, want %q", g, w)
		}
		if g, w := l.UnitOfMeasure, "gallon"; g != w {
			t.Errorf("got unit of measure %q, want %q", g, w)
		}
		if g, w := l.TotalAmount, NewDecimal(4515, 2); g.Cmp(w) != 0 {
			t.Errorf("got total amount %q, want %q", g, w)
		}
		if g, w := l.TaxAmount, NewDecimal(455, 2); g.Cmp(w) != 0 {
			t.Errorf("got tax amount %q, want %q", g, w)
		}
		if g, w := l.DiscountAmount, (*Decimal)(nil); g != nil {
			t.Errorf("got discount amount %q, want %q", g, w)
		}
		if g, w := l.ProductCode, ""; g != w {
			t.Errorf("got product code %q, want %q", g, w)
		}
		if g, w := l.CommodityCode, ""; g != w {
			t.Errorf("got commodity code %q, want %q", g, w)
		}
		if g, w := l.URL, ""; g != w {
			t.Errorf("got url %q, want %q", g, w)
		}
	}
}

func TestTransactionWithLineItemsMaxMultiple(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	lineItems := TransactionLineItemRequests{}
	for i := 0; i < 249; i++ {
		lineItems = append(lineItems, &TransactionLineItemRequest{
			Name:           "Name #1",
			Description:    "Description #1",
			Kind:           TransactionLineItemKindDebit,
			Quantity:       NewDecimal(10232, 4),
			UnitAmount:     NewDecimal(451232, 4),
			UnitTaxAmount:  NewDecimal(123, 2),
			UnitOfMeasure:  "gallon",
			TotalAmount:    NewDecimal(4515, 2),
			TaxAmount:      NewDecimal(455, 2),
			DiscountAmount: NewDecimal(102, 2),
			ProductCode:    "23434",
			CommodityCode:  "9SAASSD8724",
			URL:            "https://example.com/products/23434",
		})
	}

	txn, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems:          lineItems,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(txn.Id)

	foundLineItems, err := testGateway.TransactionLineItem().Find(ctx, txn.Id)

	if err != nil {
		t.Fatal(err)
	}

	if g, w := len(foundLineItems), len(lineItems); g != w {
		t.Fatalf("got %d line items, want %d line items", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorCommodityCodeIsTooLong(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "0123456789123",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("CommodityCode")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95801"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "CommodityCode"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Commodity code is too long."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorDescriptionIsTooLong(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "This is a line item description which is far too long. Like, way too long to be practical. We don't like how long this line item description is.",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("Description")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95803"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "Description"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Description is too long."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorDiscountAmountIsTooLarge(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(214748364800, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("DiscountAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95805"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "DiscountAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Discount amount is too large."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorDiscountAmountCannotBeNegative(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(-200, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("DiscountAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95806"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "DiscountAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Discount amount cannot be negative."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorTaxAmountIsTooLarge(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(214748364800, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("TaxAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95828"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "TaxAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Tax amount is too large."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorTaxAmountCannotBeNegative(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(-200, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("TaxAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95829"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "TaxAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Tax amount cannot be negative."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorKindIsRequired(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("Kind")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95808"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "Kind"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Kind is required."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorNameIsRequired(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("Name")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95822"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "Name"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Name is required."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorNameIsTooLong(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "123456789012345678901234567890123456",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("Name")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95823"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "Name"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Name is too long."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorProductCodeIsTooLong(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "123456789012345678901234567890123456",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("ProductCode")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95809"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "ProductCode"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Product code is too long."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorQuantityIsRequired(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("Quantity")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95811"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "Quantity"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Quantity is required."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorQuantityIsTooLarge(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(21474836480000, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("Quantity")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95812"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "Quantity"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Quantity is too large."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorTotalAmountIsRequired(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("TotalAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95814"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "TotalAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Total amount is required."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorTotalAmountIsTooLarge(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(214748364800, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("TotalAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95815"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "TotalAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Total amount is too large."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorTotalAmountMustBeGreaterThanZero(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(-200, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("TotalAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95816"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "TotalAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Total amount must be greater than zero."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorUnitAmountIsRequired(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("UnitAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95818"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "UnitAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Unit amount is required."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorUnitAmountIsTooLarge(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(21474836480000, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("UnitAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95819"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "UnitAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Unit amount is too large."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorUnitAmountMustBeGreaterThanZero(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(-20000, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("UnitAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95820"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "UnitAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Unit amount must be greater than zero."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorUnitOfMeasureIsTooLong(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "1234567890123",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("UnitOfMeasure")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95821"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "UnitOfMeasure"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Unit of measure is too long."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorUnitTaxAmountIsInvalid(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(1234, 3),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("UnitTaxAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95824"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "UnitTaxAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Unit tax amount is an invalid format."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorUnitTaxAmountIsTooLarge(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(214748364800, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("UnitTaxAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95825"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "UnitTaxAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Unit tax amount is too large."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorUnitTaxAmountCannotBeNegative(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems: TransactionLineItemRequests{
			&TransactionLineItemRequest{
				Name:           "Name #1",
				Description:    "Description #1",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(123, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
			&TransactionLineItemRequest{
				Name:           "Name #2",
				Description:    "Description #2",
				Kind:           TransactionLineItemKindDebit,
				Quantity:       NewDecimal(10232, 4),
				UnitAmount:     NewDecimal(451232, 4),
				UnitTaxAmount:  NewDecimal(-200, 2),
				UnitOfMeasure:  "gallon",
				TotalAmount:    NewDecimal(4515, 2),
				TaxAmount:      NewDecimal(455, 2),
				DiscountAmount: NewDecimal(102, 2),
				ProductCode:    "23434",
				CommodityCode:  "9SAASSD8724",
				URL:            "https://example.com/products/23434",
			},
		},
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").For("LineItems").ForIndex(1).On("UnitTaxAmount")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "95826"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "UnitTaxAmount"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Unit tax amount cannot be negative."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}

func TestTransactionWithLineItemsValidationErrorTooManyLineItems(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	lineItems := TransactionLineItemRequests{}
	for i := 0; i < 250; i++ {
		lineItems = append(lineItems, &TransactionLineItemRequest{
			Name:           "Name #1",
			Description:    "Description #1",
			Kind:           TransactionLineItemKindDebit,
			Quantity:       NewDecimal(10232, 4),
			UnitAmount:     NewDecimal(451232, 4),
			UnitTaxAmount:  NewDecimal(123, 2),
			UnitOfMeasure:  "gallon",
			TotalAmount:    NewDecimal(4515, 2),
			TaxAmount:      NewDecimal(455, 2),
			DiscountAmount: NewDecimal(102, 2),
			ProductCode:    "23434",
			CommodityCode:  "9SAASSD8724",
			URL:            "https://example.com/products/23434",
		})
	}

	_, err := testGateway.Transaction().Create(ctx, &TransactionRequest{
		Type:               "sale",
		Amount:             NewDecimal(1423, 2),
		PaymentMethodNonce: FakeNonceTransactable,
		LineItems:          lineItems,
	})

	if err == nil {
		t.Fatal("got no error, want error")
	}

	allValidationErrors := err.(*BraintreeError).All()
	if g, w := len(allValidationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}

	validationErrors := err.(*BraintreeError).For("Transaction").On("LineItems")
	if g, w := len(validationErrors), 1; g != w {
		t.Errorf("got %d errors, want %d", g, w)
	}
	if g, w := validationErrors[0].Code, "915157"; g != w {
		t.Errorf("got error code %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Attribute, "LineItems"; g != w {
		t.Errorf("got error attribute %q, want %q", g, w)
	}
	if g, w := validationErrors[0].Message, "Too many line items."; g != w {
		t.Errorf("got error message %q, want %q", g, w)
	}
}
