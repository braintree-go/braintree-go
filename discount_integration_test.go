package braintree

import (
	"reflect"
	"testing"
)

func TestDiscounts(t *testing.T) {
	discounts, err := testGateway.Discount().All()

	if err != nil {
		t.Error(err)
	} else if len(discounts) != 1 {
		t.Fail()
	}

	discount := discounts[0]

	t.Log(discount)

	if discount.Id != "test_discount" {
		t.Fail()
	} else if !reflect.DeepEqual(discount.Amount, NewDecimal(1000, 2)) {
		t.Fail()
	} else if discount.Kind != ModificationKindDiscount {
		t.Fail()
	} else if discount.Name != "test_discount_name" {
		t.Fail()
	} else if discount.NeverExpires != true {
		t.Fail()
	} else if discount.Description != "A test discount" {
		t.Fail()
	}
}
