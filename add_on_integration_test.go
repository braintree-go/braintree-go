package braintree

import (
	"testing"
)

func TestAddOn(t *testing.T) {
	addOns, err := testGateway.AddOn().All()

	if err != nil {
		t.Error(err)
	} else if len(addOns) != 1 {
		t.Fail()
	}

	addOn := addOns[0]

	t.Log(addOn)

	if addOn.Id != "test_add_on_id" {
		t.Fail()
	} else if addOn.Amount != 10 {
		t.Fail()
	} else if addOn.Kind != ModificationKindAddOn {
		t.Fail()
	} else if addOn.Name != "test_add_on_name" {
		t.Fail()
	} else if addOn.NeverExpires != true {
		t.Fail()
	} else if addOn.Description != "A test add on" {
		t.Fail()
	}
}
