// +build integration

package braintree

import (
	"context"
	"crypto/rand"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/braintree-go/braintree-go/testhelpers"
)

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestSubscriptionSimple(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(customer)

	g := testGateway.Subscription()

	// Create
	sub, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan",
	})

	t.Log("sub1", sub)

	if err != nil {
		t.Fatal(err)
	}
	if sub.Id == "" {
		t.Fatal("invalid subscription id")
	}
	if len(sub.StatusEvents) != 1 {
		t.Fatalf("expected one status event, got %d", len(sub.StatusEvents))
	}
	wantBalance := NewDecimal(0, 2)
	wantPrice := NewDecimal(1000, 2)
	event := sub.StatusEvents[0]
	if event.Status != SubscriptionStatusActive {
		t.Fatalf("expected status of status history event to be active, was %s", event.Status)
	}
	if event.CurrencyISOCode != "USD" {
		t.Fatalf("expected currency iso code of status history event to be USD, was %s", event.CurrencyISOCode)
	}
	if event.Balance.Cmp(wantBalance) != 0 {
		t.Fatalf("expected balance of status history event to be 0, was %s", event.Balance)
	}
	if event.Price.Cmp(wantPrice) != 0 {
		t.Fatalf("expected price of status history event to be 10, was %s", event.Price)
	}

	// Update
	b := make([]byte, 16)
	rand.Read(b)
	newId := fmt.Sprintf("%X", b[:])
	sub2, err := g.Update(ctx, sub.Id, &SubscriptionRequest{
		Id:     newId,
		PlanId: "test_plan_2",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	})

	t.Log("sub2", sub2)

	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != newId {
		t.Fatalf("expected subscription ID to change to %s but is %s", newId, sub2.Id)
	}
	if x := sub2.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}
	if len(sub2.StatusEvents) != 2 {
		t.Fatalf("expected two status events, got %d", len(sub2.StatusEvents))
	}
	for _, event := range sub2.StatusEvents {
		if event.Status != SubscriptionStatusActive {
			t.Fatalf("expected status of status history event to be active, was %s", event.Status)
		}
		if event.CurrencyISOCode != "USD" {
			t.Fatalf("expected currency iso code of status history event to be USD, was %s", event.CurrencyISOCode)
		}
		if event.Balance.Cmp(wantBalance) != 0 {
			t.Fatalf("expected balance of status history event to be 0, was %s", event.Balance)
		}
		if event.Price.Cmp(wantPrice) != 0 {
			t.Fatalf("expected price of status history event to be 10, was %s", event.Price)
		}
	}

	// Find
	sub3, err := g.Find(ctx, sub2.Id)
	if err != nil {
		t.Fatal(err)
	}
	if sub3.Id != sub2.Id {
		t.Fatal(sub3.Id)
	}

	// Cancel
	_, err = g.Cancel(ctx, sub2.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubscriptionAllFieldsWithBillingDayOfMonth(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(customer)

	g := testGateway.Subscription()

	// Create
	sub1, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken:    paymentMethod.GetToken(),
		PlanId:                "test_plan",
		MerchantAccountId:     testMerchantAccountId,
		BillingDayOfMonth:     testhelpers.IntPtr(15),
		NumberOfBillingCycles: testhelpers.IntPtr(2),
		Price:                 NewDecimal(100, 2),
		Descriptor: &Descriptor{
			Name:  "Company Name*Product 1",
			Phone: "0000000000",
			URL:   "example.com",
		},
	})

	t.Log("sub1", sub1)

	if err != nil {
		t.Fatal(err)
	}
	if sub1.Id == "" {
		t.Fatal("invalid subscription id")
	}
	if sub1.BillingDayOfMonth != "15" {
		t.Fatalf("got billing day of month %#v, want %#v", sub1.BillingDayOfMonth, "15")
	}
	if x := sub1.NeverExpires; x {
		t.Fatalf("got never expires %#v, want false", x)
	}
	if x := sub1.NumberOfBillingCycles; x == nil || *x != 2 {
		t.Fatalf("got number billing cycles %#v, want 2", x)
	}
	if x := sub1.Price; x == nil || x.Scale != 2 || x.Unscaled != 100 {
		t.Fatalf("got price %#v, want 1.00", x)
	}
	if x := sub1.TrialPeriod; x {
		t.Fatalf("got trial period %#v, want false", x)
	}
	if x := sub1.Status; x != SubscriptionStatusPending && x != SubscriptionStatusActive {
		t.Fatalf("got status %#v, want Pending or Active (it will be active if todays date matches the billing day of month)", x)
	}
	if x := sub1.Descriptor.Name; x != "Company Name*Product 1" {
		t.Fatalf("got descriptor name %#v, want Company Name*Product 1", x)
	}
	if x := sub1.Descriptor.Phone; x != "0000000000" {
		t.Fatalf("got descriptor phone %#v, want 0000000000", x)
	}
	if x := sub1.Descriptor.URL; x != "example.com" {
		t.Fatalf("got descriptor url %#v, want example.com", x)
	}
	if sub1.CreatedAt == nil {
		t.Fatal("expected createdAt to not be nil")
	}
	if sub1.UpdatedAt == nil {
		t.Fatal("expected updatedAt to not be nil")
	}

	// Update
	sub2, err := g.Update(ctx, sub1.Id, &SubscriptionRequest{
		PlanId: "test_plan_2",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	})

	t.Log("sub2", sub2)

	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != sub1.Id {
		t.Fatal(sub2.Id)
	}
	if x := sub2.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}

	// Find
	sub3, err := g.Find(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if sub3.Id != sub1.Id {
		t.Fatal(sub3.Id)
	}

	// Cancel
	sub4, err := g.Cancel(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if x := sub4.Status; x != SubscriptionStatusCanceled {
		t.Fatalf("got status %#v, want Canceled", x)
	}
}

func TestSubscriptionAllFieldsWithBillingDayOfMonthNeverExpires(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(customer)

	g := testGateway.Subscription()

	// Create
	sub1, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan",
		MerchantAccountId:  testMerchantAccountId,
		BillingDayOfMonth:  testhelpers.IntPtr(15),
		NeverExpires:       testhelpers.BoolPtr(true),
		Price:              NewDecimal(100, 2),
		Descriptor: &Descriptor{
			Name:  "Company Name*Product 1",
			Phone: "0000000000",
			URL:   "example.com",
		},
	})

	t.Logf("sub1 %#v", sub1)

	if err != nil {
		t.Fatal(err)
	}
	if sub1.Id == "" {
		t.Fatal("invalid subscription id")
	}
	if sub1.BillingDayOfMonth != "15" {
		t.Fatalf("got billing day of month %#v, want %#v", sub1.BillingDayOfMonth, "15")
	}
	if x := sub1.NeverExpires; !x {
		t.Fatalf("got never expires %#v, want true", x)
	}
	if x := sub1.NumberOfBillingCycles; x != nil {
		t.Fatalf("got number billing cycles %#v, didn't want", x)
	}
	if x := sub1.Price; x == nil || x.Scale != 2 || x.Unscaled != 100 {
		t.Fatalf("got price %#v, want 1.00", x)
	}
	if x := sub1.TrialPeriod; x {
		t.Fatalf("got trial period %#v, want false", x)
	}
	if x := sub1.Status; x != SubscriptionStatusPending && x != SubscriptionStatusActive {
		t.Fatalf("got status %#v, want Pending or Active (it will be active if todays date matches the billing day of month)", x)
	}
	if x := sub1.Descriptor.Name; x != "Company Name*Product 1" {
		t.Fatalf("got descriptor name %#v, want Company Name*Product 1", x)
	}
	if x := sub1.Descriptor.Phone; x != "0000000000" {
		t.Fatalf("got descriptor phone %#v, want 0000000000", x)
	}
	if x := sub1.Descriptor.URL; x != "example.com" {
		t.Fatalf("got descriptor url %#v, want example.com", x)
	}

	// Update
	sub2, err := g.Update(ctx, sub1.Id, &SubscriptionRequest{
		Id:     sub1.Id,
		PlanId: "test_plan_2",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	})

	t.Log("sub2", sub2)

	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != sub1.Id {
		t.Fatal(sub2.Id)
	}
	if x := sub2.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}

	// Find
	sub3, err := g.Find(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if sub3.Id != sub1.Id {
		t.Fatal(sub3.Id)
	}

	// Cancel
	sub4, err := g.Cancel(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if x := sub4.Status; x != SubscriptionStatusCanceled {
		t.Fatalf("got status %#v, want Canceled", x)
	}
}

func TestSubscriptionAllFieldsWithFirstBillingDate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(customer)

	g := testGateway.Subscription()

	// Create
	firstBillingDate := fmt.Sprintf("%d-12-31", time.Now().Year())
	sub1, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken:    paymentMethod.GetToken(),
		PlanId:                "test_plan",
		MerchantAccountId:     testMerchantAccountId,
		FirstBillingDate:      firstBillingDate,
		NumberOfBillingCycles: testhelpers.IntPtr(2),
		Price:                 NewDecimal(100, 2),
		Descriptor: &Descriptor{
			Name:  "Company Name*Product 1",
			Phone: "0000000000",
			URL:   "example.com",
		},
	})

	t.Log("sub1", sub1)

	if err != nil {
		t.Fatal(err)
	}
	if sub1.Id == "" {
		t.Fatal("invalid subscription id")
	}
	if sub1.BillingDayOfMonth != "31" {
		t.Fatalf("got billing day of month %#v, want %#v", sub1.BillingDayOfMonth, "31")
	}
	if sub1.FirstBillingDate != firstBillingDate {
		t.Fatalf("got first billing date %#v, want %#v", sub1.FirstBillingDate, firstBillingDate)
	}
	if x := sub1.NeverExpires; x {
		t.Fatalf("got never expires %#v, want false", x)
	}
	if x := sub1.NumberOfBillingCycles; x == nil {
		t.Fatalf("got number billing cycles nil, want 2")
	} else if *x != 2 {
		t.Fatalf("got number billing cycles %#v, want 2", *x)
	}
	if x := sub1.Price; x == nil || x.Scale != 2 || x.Unscaled != 100 {
		t.Fatalf("got price %#v, want 1.00", x)
	}
	if x := sub1.TrialPeriod; x {
		t.Fatalf("got trial period %#v, want false", x)
	}
	if x := sub1.Status; x != SubscriptionStatusPending {
		t.Fatalf("got status %#v, want Pending", x)
	}
	if x := sub1.Descriptor.Name; x != "Company Name*Product 1" {
		t.Fatalf("got descriptor name %#v, want Company Name*Product 1", x)
	}
	if x := sub1.Descriptor.Phone; x != "0000000000" {
		t.Fatalf("got descriptor phone %#v, want 0000000000", x)
	}
	if x := sub1.Descriptor.URL; x != "example.com" {
		t.Fatalf("got descriptor url %#v, want example.com", x)
	}

	// Update
	sub2, err := g.Update(ctx, sub1.Id, &SubscriptionRequest{
		Id:     sub1.Id,
		PlanId: "test_plan_2",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	})

	t.Log("sub2", sub2)

	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != sub1.Id {
		t.Fatal(sub2.Id)
	}
	if x := sub2.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}

	// Find
	sub3, err := g.Find(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if sub3.Id != sub1.Id {
		t.Fatal(sub3.Id)
	}

	// Cancel
	sub4, err := g.Cancel(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if x := sub4.Status; x != SubscriptionStatusCanceled {
		t.Fatalf("got status %#v, want Canceled", x)
	}
}

func TestSubscriptionAllFieldsWithFirstBillingDateNeverExpires(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(customer)

	g := testGateway.Subscription()

	// Create
	firstBillingDate := fmt.Sprintf("%d-12-31", time.Now().Year())
	sub1, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan",
		MerchantAccountId:  testMerchantAccountId,
		FirstBillingDate:   firstBillingDate,
		NeverExpires:       testhelpers.BoolPtr(true),
		Price:              NewDecimal(100, 2),
		Descriptor: &Descriptor{
			Name:  "Company Name*Product 1",
			Phone: "0000000000",
			URL:   "example.com",
		},
	})

	t.Log("sub1", sub1)

	if err != nil {
		t.Fatal(err)
	}
	if sub1.Id == "" {
		t.Fatal("invalid subscription id")
	}
	if sub1.BillingDayOfMonth != "31" {
		t.Fatalf("got billing day of month %#v, want %#v", sub1.BillingDayOfMonth, "31")
	}
	if sub1.FirstBillingDate != firstBillingDate {
		t.Fatalf("got first billing date %#v, want %#v", sub1.FirstBillingDate, firstBillingDate)
	}
	if x := sub1.NeverExpires; !x {
		t.Fatalf("got never expires %#v, want true", x)
	}
	if x := sub1.NumberOfBillingCycles; x != nil {
		t.Fatalf("got number billing cycles %#v, didn't want", x)
	}
	if x := sub1.Price; x == nil || x.Scale != 2 || x.Unscaled != 100 {
		t.Fatalf("got price %#v, want 1.00", x)
	}
	if x := sub1.TrialPeriod; x {
		t.Fatalf("got trial period %#v, want false", x)
	}
	if x := sub1.Status; x != SubscriptionStatusPending {
		t.Fatalf("got status %#v, want Pending", x)
	}
	if x := sub1.Descriptor.Name; x != "Company Name*Product 1" {
		t.Fatalf("got descriptor name %#v, want Company Name*Product 1", x)
	}
	if x := sub1.Descriptor.Phone; x != "0000000000" {
		t.Fatalf("got descriptor phone %#v, want 0000000000", x)
	}
	if x := sub1.Descriptor.URL; x != "example.com" {
		t.Fatalf("got descriptor url %#v, want example.com", x)
	}

	// Update
	sub2, err := g.Update(ctx, sub1.Id, &SubscriptionRequest{
		Id:     sub1.Id,
		PlanId: "test_plan_2",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	})

	t.Log("sub2", sub2)

	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != sub1.Id {
		t.Fatal(sub2.Id)
	}
	if x := sub2.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}

	// Find
	sub3, err := g.Find(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if sub3.Id != sub1.Id {
		t.Fatal(sub3.Id)
	}

	// Cancel
	sub4, err := g.Cancel(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if x := sub4.Status; x != SubscriptionStatusCanceled {
		t.Fatalf("got status %#v, want Canceled", x)
	}
}

func TestSubscriptionAllFieldsWithTrialPeriod(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(customer)

	g := testGateway.Subscription()

	// Create
	firstBillingDate := time.Now().In(testTimeZone).AddDate(0, 0, 7)
	sub1, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken:    paymentMethod.GetToken(),
		PlanId:                "test_plan",
		MerchantAccountId:     testMerchantAccountId,
		TrialPeriod:           testhelpers.BoolPtr(true),
		TrialDuration:         "7",
		TrialDurationUnit:     SubscriptionTrialDurationUnitDay,
		NumberOfBillingCycles: testhelpers.IntPtr(2),
		Price:                 NewDecimal(100, 2),
		Descriptor: &Descriptor{
			Name:  "Company Name*Product 1",
			Phone: "0000000000",
			URL:   "example.com",
		},
	})

	t.Log("sub1", sub1)

	if err != nil {
		t.Fatal(err)
	}
	if sub1.Id == "" {
		t.Fatal("invalid subscription id")
	}
	if sub1.BillingDayOfMonth != fmt.Sprintf("%d", firstBillingDate.Day()) {
		t.Fatalf("got billing day of month %#v, want %#v", sub1.BillingDayOfMonth, firstBillingDate.Day())
	}
	if sub1.FirstBillingDate != firstBillingDate.Format("2006-01-02") {
		t.Fatalf("got first billing date %#v, want %#v", sub1.FirstBillingDate, firstBillingDate)
	}
	if x := sub1.NeverExpires; x {
		t.Fatalf("got never expires %#v, want false", x)
	}
	if x := sub1.NumberOfBillingCycles; x == nil || *x != 2 {
		t.Fatalf("got number billing cycles %#v, want 2", x)
	}
	if x := sub1.Price; x == nil || x.Scale != 2 || x.Unscaled != 100 {
		t.Fatalf("got price %#v, want 1.00", x)
	}
	if x := sub1.TrialPeriod; !x {
		t.Fatalf("got trial period %#v, want false", x)
	}
	if sub1.TrialDuration != "7" {
		t.Fatalf("got trial duration %#v, want 7", sub1.TrialDuration)
	}
	if sub1.TrialDurationUnit != SubscriptionTrialDurationUnitDay {
		t.Fatalf("got trial duration unit %#v, want day", sub1.TrialDurationUnit)
	}
	if x := sub1.Status; x != SubscriptionStatusActive {
		t.Fatalf("got status %#v, want Active", x)
	}
	if x := sub1.Descriptor.Name; x != "Company Name*Product 1" {
		t.Fatalf("got descriptor name %#v, want Company Name*Product 1", x)
	}
	if x := sub1.Descriptor.Phone; x != "0000000000" {
		t.Fatalf("got descriptor phone %#v, want 0000000000", x)
	}
	if x := sub1.Descriptor.URL; x != "example.com" {
		t.Fatalf("got descriptor url %#v, want example.com", x)
	}

	// Update
	sub2, err := g.Update(ctx, sub1.Id, &SubscriptionRequest{
		Id:     sub1.Id,
		PlanId: "test_plan_2",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	})

	t.Log("sub2", sub2)

	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != sub1.Id {
		t.Fatal(sub2.Id)
	}
	if x := sub2.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}

	// Find
	sub3, err := g.Find(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if sub3.Id != sub1.Id {
		t.Fatal(sub3.Id)
	}

	// Cancel
	_, err = g.Cancel(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubscriptionAllFieldsWithTrialPeriodNeverExpires(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(customer)

	g := testGateway.Subscription()

	// Create
	firstBillingDate := time.Now().In(testTimeZone).AddDate(0, 0, 7)
	sub1, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan",
		MerchantAccountId:  testMerchantAccountId,
		TrialPeriod:        testhelpers.BoolPtr(true),
		TrialDuration:      "7",
		TrialDurationUnit:  SubscriptionTrialDurationUnitDay,
		NeverExpires:       testhelpers.BoolPtr(true),
		Price:              NewDecimal(100, 2),
		Descriptor: &Descriptor{
			Name:  "Company Name*Product 1",
			Phone: "0000000000",
			URL:   "example.com",
		},
	})

	t.Log("sub1", sub1)

	if err != nil {
		t.Fatal(err)
	}
	if sub1.Id == "" {
		t.Fatal("invalid subscription id")
	}
	if sub1.BillingDayOfMonth != fmt.Sprintf("%d", firstBillingDate.Day()) {
		t.Fatalf("got billing day of month %#v, want %#v", sub1.BillingDayOfMonth, firstBillingDate.Day())
	}
	if sub1.FirstBillingDate != firstBillingDate.Format("2006-01-02") {
		t.Fatalf("got first billing date %#v, want %#v", sub1.FirstBillingDate, firstBillingDate)
	}
	if x := sub1.NeverExpires; !x {
		t.Fatalf("got never expires %#v, want true", x)
	}
	if x := sub1.NumberOfBillingCycles; x != nil {
		t.Fatalf("got number billing cycles %#v, didn't want", x)
	}
	if x := sub1.Price; x == nil || x.Scale != 2 || x.Unscaled != 100 {
		t.Fatalf("got price %#v, want 1.00", x)
	}
	if x := sub1.TrialPeriod; !x {
		t.Fatalf("got trial period %#v, want false", x)
	}
	if sub1.TrialDuration != "7" {
		t.Fatalf("got trial duration %#v, want 7", sub1.TrialDuration)
	}
	if sub1.TrialDurationUnit != SubscriptionTrialDurationUnitDay {
		t.Fatalf("got trial duration unit %#v, want day", sub1.TrialDurationUnit)
	}
	if x := sub1.Status; x != SubscriptionStatusActive {
		t.Fatalf("got status %#v, want Active", x)
	}
	if x := sub1.Descriptor.Name; x != "Company Name*Product 1" {
		t.Fatalf("got descriptor name %#v, want Company Name*Product 1", x)
	}
	if x := sub1.Descriptor.Phone; x != "0000000000" {
		t.Fatalf("got descriptor phone %#v, want 0000000000", x)
	}
	if x := sub1.Descriptor.URL; x != "example.com" {
		t.Fatalf("got descriptor url %#v, want example.com", x)
	}

	// Update
	sub2, err := g.Update(ctx, sub1.Id, &SubscriptionRequest{
		Id:     sub1.Id,
		PlanId: "test_plan_2",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	})

	t.Log("sub2", sub2)

	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != sub1.Id {
		t.Fatal(sub2.Id)
	}
	if x := sub2.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}

	// Find
	sub3, err := g.Find(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if sub3.Id != sub1.Id {
		t.Fatal(sub3.Id)
	}

	// Cancel
	_, err = g.Cancel(ctx, sub1.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubscriptionModifications(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(customer)

	g := testGateway.Subscription()

	// Create
	sub, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan_2",
	})

	t.Log("sub1", sub)

	if err != nil {
		t.Fatal(err)
	}
	if sub.Id == "" {
		t.Fatal("invalid subscription id")
	}

	// Add AddOn
	sub2, err := g.Update(ctx, sub.Id, &SubscriptionRequest{
		Id: sub.Id,
		AddOns: &ModificationsRequest{
			Add: []AddModificationRequest{
				{
					InheritedFromID: "test_add_on",
					ModificationRequest: ModificationRequest{
						Amount:       NewDecimal(300, 2),
						Quantity:     1,
						NeverExpires: true,
					},
				},
			},
		},
		Discounts: &ModificationsRequest{
			Add: []AddModificationRequest{
				{
					InheritedFromID: "test_discount",
					ModificationRequest: ModificationRequest{
						Amount:                NewDecimal(100, 2),
						Quantity:              1,
						NumberOfBillingCycles: 2,
					},
				},
			},
		},
	})

	t.Log("sub2", sub2)

	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != sub.Id {
		t.Fatal(sub2.Id)
	}
	if x := sub2.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}
	if x := sub2.AddOns.AddOns; len(x) != 1 {
		t.Fatalf("got %d add ons, want 1 add on", len(x))
	}
	if x := sub2.AddOns.AddOns[0].Amount; x.String() != NewDecimal(300, 2).String() {
		t.Fatalf("got %v add on, want 3.00 add on", x)
	}
	if x := sub2.Discounts.Discounts; len(x) != 1 {
		t.Fatalf("got %d discounts, want 1 discount", len(x))
	}
	if x := sub2.Discounts.Discounts[0].Amount; x.String() != NewDecimal(100, 2).String() {
		t.Fatalf("got %v discount, want 1.00 discount", x)
	}
	if x := sub2.Discounts.Discounts[0].NumberOfBillingCycles; x != 2 {
		t.Fatalf("got %v number of billing cycles on discount, want 2 billing cycles", x)
	}
	if x := sub2.Discounts.Discounts[0].CurrentBillingCycle; x != 0 {
		t.Fatalf("got current billing cycle of %d on discount, want 0", x)
	}

	// Update AddOn
	sub3, err := g.Update(ctx, sub.Id, &SubscriptionRequest{
		Id: sub.Id,
		AddOns: &ModificationsRequest{
			Update: []UpdateModificationRequest{
				{
					ExistingID: "test_add_on",
					ModificationRequest: ModificationRequest{
						Amount: NewDecimal(150, 2),
					},
				},
			},
		},
		Discounts: &ModificationsRequest{
			RemoveExistingIDs: []string{
				"test_discount",
			},
		},
	})

	t.Log("sub3", sub3)

	if err != nil {
		t.Fatal(err)
	}
	if sub3.Id != sub.Id {
		t.Fatal(sub3.Id)
	}
	if x := sub3.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}
	if x := sub3.AddOns.AddOns; len(x) != 1 {
		t.Fatalf("got %d add ons, want 1 add on", len(x))
	}
	if x := sub3.AddOns.AddOns[0].Amount; x.String() != NewDecimal(150, 2).String() {
		t.Fatalf("got %v add on, want 1.50 add on", x)
	}
	if x := sub3.Discounts.Discounts; len(x) != 0 {
		t.Fatalf("got %d discounts, want 0 discounts", len(x))
	}

	// Cancel
	_, err = g.Cancel(ctx, sub3.Id)
	if err != nil {
		t.Fatal(err)
	}
}

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestSubscriptionTransactions(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(customer)

	g := testGateway.Subscription()

	// Create
	sub, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan",
		Options: &SubscriptionOptions{
			StartImmediately: true,
		},
	})

	t.Log("sub1", sub)

	if err != nil {
		t.Fatal(err)
	}
	if sub.Id == "" {
		t.Fatal("invalid subscription id")
	}

	// Find
	sub2, err := g.Find(ctx, sub.Id)
	if err != nil {
		t.Fatal(err)
	}
	if sub2.Id != sub.Id {
		t.Fatal(sub2.Id)
	}
	if x := sub2.PlanId; x != "test_plan" {
		t.Fatal(x)
	}
	if len(sub2.Transactions.Transaction) < 1 {
		t.Fatalf("Expected transactions slice not to be empty")
	}
	if x := sub2.Transactions.Transaction[0].PlanId; x != "test_plan" {
		t.Fatal(x)
	}
	if x := sub2.Transactions.Transaction[0].SubscriptionId; x != sub.Id {
		t.Fatal(x)
	}
	if x := sub2.Transactions.Transaction[0].SubscriptionDetails.BillingPeriodStartDate; x != sub.BillingPeriodStartDate {
		t.Fatal(x)
	}
	if x := sub2.Transactions.Transaction[0].SubscriptionDetails.BillingPeriodEndDate; x != sub.BillingPeriodEndDate {
		t.Fatal(x)
	}

	// Cancel
	_, err = g.Cancel(ctx, sub2.Id)
	if err != nil {
		t.Fatal(err)
	}
}

// It is not possible to successfully retry a charge without manually creating
// a subcription with a card that will fail, waiting a day for it to be billed
// and fail which will cause the subscription to enter the PastDue status. This
// test instead attempts to retry a charge that is not PastDue and ensures the
// only errors returned is the status.
// Ref: https://developers.braintreepayments.com/guides/recurring-billing/overview#past-due
func TestSubscriptionRetryCharge(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("customer", customer)

	verifyCard := false
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
		Options: &PaymentMethodRequestOptions{
			VerifyCard: &verifyCard,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("payment method", paymentMethod)

	g := testGateway.Subscription()

	// Create Subscription
	sub1, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan",
		MerchantAccountId:  testMerchantAccountId,
		Price:              NewDecimal(100, 2),
		Options: &SubscriptionOptions{
			StartImmediately: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("sub1", sub1)
	defer func() {
		_, err = g.Cancel(ctx, sub1.Id)
		if err != nil {
			t.Error(err)
		}
	}()

	// Retry Charge
	err = testGateway.Subscription().RetryCharge(ctx, &SubscriptionTransactionRequest{
		SubscriptionID: sub1.Id,
		Amount:         NewDecimal(10, 2),
		Options: &SubscriptionTransactionOptionsRequest{
			SubmitForSettlement: true,
		},
	})
	if err == nil {
		t.Fatalf("Retry charge did not error, want error indicating Subscription status must be Past Due in order to retry.")
	}
	btErr, ok := err.(*BraintreeError)
	if !ok {
		t.Fatal(err)
	}
	validationErrs := btErr.All()
	if len(validationErrs) != 1 {
		t.Fatalf("got %d validation errors, want 1, validation errors: %#v", len(validationErrs), validationErrs)
	}
	wantValidationErr := ValidationError{
		Code:      "91531",
		Attribute: "Base",
		Message:   "Subscription status must be Past Due in order to retry.",
	}
	if validationErrs[0] != wantValidationErr {
		t.Errorf("got validation error %#v, want %#v", validationErrs[0], wantValidationErr)
	}
}

func TestSubscriptionSearchIDs(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}

	g := testGateway.Subscription()

	sub1, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan",
	})
	sub2, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan",
	})
	sub3, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan_2",
	})

	query := &SearchQuery{}
	f1 := query.AddTimeField("created-at")
	f1.Max = time.Now()
	f1.Min = time.Now().AddDate(0, 0, -1)
	f2 := query.AddTextField("plan-id")
	f2.Is = "test_plan"

	result, err := g.SearchIDs(ctx, query)
	if err != nil {
		t.Fatal(err)
	}

	if !testhelpers.StringSliceContains(result.IDs, sub1.Id) {
		t.Errorf("expected result.IDs to include %v", sub1.Id)
	}
	if !testhelpers.StringSliceContains(result.IDs, sub2.Id) {
		t.Errorf("expected result.IDs to include %v", sub2.Id)
	}
	if testhelpers.StringSliceContains(result.IDs, sub3.Id) {
		t.Errorf("expected result.Ids to not include %v", sub3.Id)
	}
}

func TestSubscriptionSearchPage(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}
	g := testGateway.Subscription()

	const subscriptionCount = 51
	expectedIDs := map[string]bool{}
	firstBillingDate := fmt.Sprintf("%d-12-31", time.Now().Year())
	for i := 0; i < subscriptionCount; i++ {
		sub, err := g.Create(ctx, &SubscriptionRequest{
			PaymentMethodToken:    paymentMethod.GetToken(),
			PlanId:                "test_plan_2",
			FirstBillingDate:      firstBillingDate,
			NumberOfBillingCycles: testhelpers.IntPtr(2),
		})
		if err != nil {
			t.Fatal(err)
		}
		expectedIDs[sub.Id] = true
	}

	t.Logf("expectedIDs = %v", expectedIDs)

	query := &SearchQuery{}
	f1 := query.AddTimeField("created-at")
	f1.Max = time.Now()
	f1.Min = time.Now().Add(-10 * time.Minute)
	f2 := query.AddTextField("plan-id")
	f2.Is = "test_plan_2"

	results, err := g.SearchIDs(ctx, query)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result.PageSize = %v", results.PageSize)
	t.Logf("result.PageCount = %v", results.PageCount)
	t.Logf("result.IDs = %d %v", len(results.IDs), results.IDs)

	if len(results.IDs) < subscriptionCount {
		t.Errorf("result.IDs = %v, want it to be more than %v", len(results.IDs), subscriptionCount)
	}

	for page := 1; page <= results.PageCount; page++ {
		results, err := g.SearchPage(ctx, query, results, page)
		if err != nil {
			t.Fatal(err)
		}
		for _, sub := range results.Subscriptions {
			if expectedIDs[sub.Id] {
				delete(expectedIDs, sub.Id)
			}
		}
	}

	if len(expectedIDs) > 0 {
		t.Fatalf("subscriptions not returned = %v", expectedIDs)
	}

	_, err = g.SearchPage(ctx, query, results, 0)
	t.Logf("%#v", err)
	if err == nil || !strings.Contains(err.Error(), "page 0 out of bounds") {
		t.Errorf("requesting page 0 should result in out of bounds error, but got %#v", err)
	}

	_, err = g.SearchPage(ctx, query, results, results.PageCount+1)
	t.Logf("%#v", err)
	if err == nil || !strings.Contains(err.Error(), fmt.Sprintf("page %d out of bounds", results.PageCount+1)) {
		t.Errorf("requesting page %d should result in out of bounds error, but got %v", results.PageCount+1, err)
	}
}

func TestSubscriptionSearch(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}
	g := testGateway.Subscription()

	sub1, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan",
	})
	sub2, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan",
	})
	sub3, err := g.Create(ctx, &SubscriptionRequest{
		PaymentMethodToken: paymentMethod.GetToken(),
		PlanId:             "test_plan_2",
	})

	query := &SearchQuery{}
	f1 := query.AddTimeField("created-at")
	f1.Max = time.Now()
	f1.Min = time.Now().Add(-10 * time.Minute)
	f2 := query.AddTextField("plan-id")
	f2.Is = "test_plan"

	result, err := g.Search(ctx, query)
	if err != nil {
		t.Fatal(err)
	}

	if result.CurrentPageNumber != 1 {
		t.Errorf("expected page number to be 1, got %v", result.CurrentPageNumber)
	}
	if !testhelpers.StringSliceContains(result.TotalIDs, sub1.Id) {
		t.Errorf("expected subscription ids to contain %v", sub1.Id)
	}
	if !testhelpers.StringSliceContains(result.TotalIDs, sub2.Id) {
		t.Errorf("expected subscription ids to contain %v", sub2.Id)
	}
	if testhelpers.StringSliceContains(result.TotalIDs, sub3.Id) {
		t.Errorf("expected subscription ids to not contain %v", sub3.Id)
	}
}

func TestSubscriptionSearchNext(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{})
	if err != nil {
		t.Fatal(err)
	}
	paymentMethod, err := testGateway.PaymentMethod().Create(ctx, &PaymentMethodRequest{
		CustomerId:         customer.Id,
		PaymentMethodNonce: FakeNonceTransactable,
	})
	if err != nil {
		t.Fatal(err)
	}
	g := testGateway.Subscription()

	const subscriptionCount = 51
	expectedIDs := map[string]bool{}
	for i := 0; i < subscriptionCount; i++ {
		sub, err := g.Create(ctx, &SubscriptionRequest{
			PaymentMethodToken: paymentMethod.GetToken(),
			PlanId:             "test_plan_2",
		})
		if err != nil {
			t.Fatal(err)
		}
		expectedIDs[sub.Id] = true
	}

	t.Logf("expectedIDs = %v", expectedIDs)

	query := &SearchQuery{}
	f1 := query.AddTimeField("created-at")
	f1.Max = time.Now()
	f1.Min = time.Now().Add(-10 * time.Minute)
	f2 := query.AddTextField("plan-id")
	f2.Is = "test_plan_2"

	result, err := g.Search(ctx, query)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result.TotalItems = %v", result.TotalItems)
	t.Logf("result.TotalIDs = %v", result.TotalIDs)
	t.Logf("result.PageSize = %v", result.PageSize)

	if result.TotalItems < subscriptionCount {
		t.Errorf("result.TotalItems = %v, want it to be more than %v", result.TotalItems, subscriptionCount)
	}

	for {
		for _, sub := range result.Subscriptions {
			if expectedIDs[sub.Id] {
				delete(expectedIDs, sub.Id)
			}
		}

		result, err = g.SearchNext(ctx, query, result)
		if err != nil {
			t.Fatal(err)
		}
		if result == nil {
			break
		}
	}

	if len(expectedIDs) > 0 {
		t.Fatalf("subscriptions not returned = %v", expectedIDs)
	}
}
