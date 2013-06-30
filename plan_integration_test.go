package braintree

import (
	"testing"
)

func TestPlan(t *testing.T) {
	plans, err := testGateway.Plan().All()
	if err != nil {
		t.Fatal(err)
	}
	if len(plans) == 0 {
		t.Fatal(plans)
	}

	testPlanFound := false
	for _, p := range plans {
		if p.Id == "test_plan" {
			testPlanFound = true
			break
		}
	}
	if !testPlanFound {
		t.Fatal(plans)
	}
}
