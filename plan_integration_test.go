package braintree

import (
	"testing"
)

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestPlan(t *testing.T) {
	g := testGateway.Plan()
	plans, err := g.All()
	if err != nil {
		t.Fatal(err)
	}
	if len(plans) == 0 {
		t.Fatal(plans)
	}

	var plan *Plan
	for _, p := range plans {
		if p.Id == "test_plan" {
			plan = p
			break
		}
	}

	t.Log(plan)

	if plan == nil {
		t.Fatal("plan not found")
	}
	if x := plan.Id; x != "test_plan" {
		t.Fatal(x)
	}
	if x := plan.MerchantId; x == "" {
		t.Fatal(x)
	}
	if x := plan.BillingFrequency; !x.Valid || x.Int64 != 1 {
		t.Fatal(x)
	}
	if x := plan.CurrencyISOCode; x != "USD" {
		t.Fatal(x)
	}
	if x := plan.Description; x != "test_plan_desc" {
		t.Fatal(x)
	}
	if x := plan.Name; x != "test_plan_name" {
		t.Fatal(x)
	}
	if x := plan.NumberOfBillingCycles; !x.Valid || x.Int64 != 2 {
		t.Fatal(x)
	}
	if x := plan.Price; x.Cmp(NewDecimal(1000, 2)) != 0 {
		t.Fatal(x)
	}
	if x := plan.TrialDuration; !x.Valid || x.Int64 != 14 {
		t.Fatal(x)
	}
	if x := plan.TrialDurationUnit; x != "day" {
		t.Fatal(x)
	}
	if x := plan.TrialPeriod; !x.Valid || !x.Bool {
		t.Fatal(x)
	}
	if x := plan.CreatedAt; x == nil {
		t.Fatal(x)
	}
	if x := plan.UpdatedAt; x == nil {
		t.Fatal(x)
	}

	// Find
	plan2, err := g.Find("test_plan_2")
	if err != nil {
		t.Fatal(err)
	}
	if plan2.Id != "test_plan_2" {
		t.Fatal(plan2)
	}
}
