// +build unit

package braintree

import (
	"encoding/xml"
	"testing"
)

var errorXML = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<api-error-response>
  <errors>
    <errors type="array"/>
    <transaction>
      <errors type="array">
        <error>
          <code>91560</code>
          <attribute type="symbol">base</attribute>
          <message>Transaction could not be held in escrow.</message>
        </error>
        <error>
          <code>81502</code>
          <attribute type="symbol">amount</attribute>
          <message>Amount is required.</message>
        </error>
        <error>
          <code>91526</code>
          <attribute type="symbol">custom_fields</attribute>
          <message>Custom field is invalid: store_me.</message>
        </error>
        <error>
          <code>91513</code>
          <attribute type="symbol">merchant_account_id</attribute>
          <message>Merchant account ID is invalid.</message>
        </error>
        <error>
          <code>915157</code>
          <attribute type="symbol">line_items</attribute>
          <message>Too many line items.</message>
        </error>
      </errors>
      <credit-card>
        <errors type="array">
          <error>
            <code>91708</code>
            <attribute type="symbol">base</attribute>
            <message>Cannot provide expiration_date if you are also providing expiration_month and expiration_year.</message>
          </error>
          <error>
            <code>81714</code>
            <attribute type="symbol">number</attribute>
            <message>Credit card number is required.</message>
          </error>
          <error>
            <code>81725</code>
            <attribute type="symbol">base</attribute>
            <message>Credit card must include either number or venmo_sdk_payment_method_code.</message>
          </error>
          <error>
            <code>81703</code>
            <attribute type="symbol">number</attribute>
            <message>Credit card type is not accepted by this merchant account.</message>
          </error>
        </errors>
      </credit-card>
      <customer>
        <errors type="array">
          <error>
            <code>81606</code>
            <attribute type="symbol">email</attribute>
            <message>Email is an invalid format.</message>
          </error>
        </errors>
      </customer>
      <line-items>
        <index-1>
          <errors type="array">
            <error>
              <code>95801</code>
              <attribute type="symbol">commodity_code</attribute>
              <message>Commodity code is too long.</message>
            </error>
          </errors>
        </index-1>
        <index-3>
          <errors type="array">
            <error>
              <code>95803</code>
              <attribute type="symbol">description</attribute>
              <message>Description is too long.</message>
            </error>
            <error>
              <code>95809</code>
              <attribute type="symbol">product_code</attribute>
              <message>Product code is too long.</message>
            </error>
          </errors>
        </index-3>
      </line-items>
    </transaction>
  </errors>
  <message>Everything is broken!</message>
</api-error-response>`)

func TestErrorsUnmarshalEverything(t *testing.T) {
	t.Parallel()

	apiErrors := &BraintreeError{}
	err := xml.Unmarshal(errorXML, apiErrors)
	if err != nil {
		t.Fatal("Error unmarshalling: " + err.Error())
	}

	allErrors := apiErrors.All()

	if g, w := len(allErrors), 13; g != w {
		t.Fatalf("got %d errors, want %d errors", g, w)
	}
}

func TestErrorsAccessors(t *testing.T) {
	t.Parallel()

	apiErrors := &BraintreeError{}
	err := xml.Unmarshal(errorXML, apiErrors)
	if err != nil {
		t.Fatal("Error unmarshalling: " + err.Error())
	}

	ccObjectErrors := apiErrors.For("Transaction").For("CreditCard")
	if g, w := ccObjectErrors.Object, "CreditCard"; g != w {
		t.Errorf("cc object, got %q, want %q", g, w)
	}

	ccErrors := apiErrors.For("Transaction").For("CreditCard").All()
	if len(ccErrors) != 4 {
		t.Error("Did not get the right credit card errors")
	}

	numberErrors := apiErrors.For("Transaction").For("CreditCard").On("Number")
	if len(numberErrors) != 2 {
		t.Error("Did not get the right number errors")
	}

	customerErrors := apiErrors.For("Transaction").For("Customer").All()
	if len(customerErrors) != 1 {
		t.Error("Did not get the right customer errors")
	}

	lineItemsErrors := apiErrors.For("Transaction").On("LineItems")
	if g, w := len(lineItemsErrors), 1; g != w {
		t.Errorf("line items, got %d, want %d", g, w)
	}

	lineItem1CommodityCodeErrors := apiErrors.For("Transaction").For("LineItems").ForIndex(1).On("CommodityCode")
	if g, w := len(lineItem1CommodityCodeErrors), 1; g != w {
		t.Errorf("line item 1, got %d, want %d", g, w)
	}
	if g, w := lineItem1CommodityCodeErrors[0].Code, "95801"; g != w {
		t.Errorf("line item 1, got %q, want %q", g, w)
	}
	if g, w := lineItem1CommodityCodeErrors[0].Attribute, "CommodityCode"; g != w {
		t.Errorf("line item 1, got %q, want %q", g, w)
	}
	if g, w := lineItem1CommodityCodeErrors[0].Message, "Commodity code is too long."; g != w {
		t.Errorf("line item 1, got %q, want %q", g, w)
	}

	lineItem3Errors := apiErrors.For("Transaction").For("LineItems").ForIndex(3).On("Description")
	if g, w := len(lineItem3Errors), 1; g != w {
		t.Errorf("line item 3, got %d, want %d", g, w)
	}

	baseErrors := apiErrors.For("Transaction").All()
	if g, w := len(baseErrors), 5; g != w {
		t.Errorf("transaction, got %d, want %d", g, w)
	}

	baseBaseErrors := apiErrors.For("Transaction").On("Base")
	if g, w := len(baseBaseErrors), 1; g != w {
		t.Errorf("transaction base, got %d, want %d", g, w)
	}
}

func TestErrorsNameSnakeToCamel(t *testing.T) {
	cases := []struct {
		Snake string
		Camel string
	}{
		{"amount", "Amount"},
		{"index_1", "Index1"},
		{"index_123", "Index123"},
		{"commodity_code", "CommodityCode"},
		{"description", "Description"},
	}

	for _, c := range cases {
		t.Run(c.Snake+"=>"+c.Camel, func(t *testing.T) {
			camel := errorNameSnakeToCamel(c.Snake)
			if g, w := camel, c.Camel; g == w {
				t.Logf("got %q, want %q", g, w)
			} else {
				t.Errorf("got %q, want %q", g, w)
			}
		})
	}
}

func TestErrorsNameKebabToCamel(t *testing.T) {
	cases := []struct {
		Kebab string
		Camel string
	}{
		{"amount", "Amount"},
		{"line-items", "LineItems"},
		{"index-1", "Index1"},
		{"index-123", "Index123"},
		{"commodity-code", "CommodityCode"},
		{"description", "Description"},
	}

	for _, c := range cases {
		t.Run(c.Kebab+"=>"+c.Camel, func(t *testing.T) {
			camel := errorNameKebabToCamel(c.Kebab)
			if g, w := camel, c.Camel; g == w {
				t.Logf("got %q, want %q", g, w)
			} else {
				t.Errorf("got %q, want %q", g, w)
			}
		})
	}
}
