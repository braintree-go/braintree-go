// +build integration

package braintree

import (
	"context"
	"testing"
)

func TestAddress(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	customer, err := testGateway.Customer().Create(ctx, &CustomerRequest{
		FirstName: "Jenna",
		LastName:  "Smith",
	})
	if err != nil {
		t.Fatal(err)
	}
	if customer.Id == "" {
		t.Fatal("invalid customer id")
	}

	addr := &AddressRequest{
		FirstName:          "Jenna",
		LastName:           "Smith",
		Company:            "Braintree",
		StreetAddress:      "1 E Main St",
		ExtendedAddress:    "Suite 403",
		Locality:           "Chicago",
		Region:             "Illinois",
		PostalCode:         "60622",
		CountryCodeAlpha2:  "US",
		CountryCodeAlpha3:  "USA",
		CountryCodeNumeric: "840",
		CountryName:        "United States of America",
	}

	addr2, err := testGateway.Address().Create(ctx, customer.Id, addr)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", addr)
	t.Logf("%+v\n", addr2)
	validateAddr(t, addr2, addr, customer)

	addr3 := &AddressRequest{
		FirstName:          "Al",
		LastName:           "Fredidandes",
		Company:            "Paypal",
		StreetAddress:      "1 W Main St",
		ExtendedAddress:    "Suite 402",
		Locality:           "Montreal",
		Region:             "Quebec",
		PostalCode:         "H1A",
		CountryCodeAlpha2:  "CA",
		CountryCodeAlpha3:  "CAN",
		CountryCodeNumeric: "124",
		CountryName:        "Canada",
	}
	addr4, err := testGateway.Address().Update(ctx, customer.Id, addr2.Id, addr3)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", addr3)
	t.Logf("%+v\n", addr4)
	validateAddr(t, addr4, addr3, customer)

	err = testGateway.Address().Delete(ctx, customer.Id, addr2.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func validateAddr(t *testing.T, addr *Address, addrRequest *AddressRequest, customer *Customer) {
	if addr.Id == "" {
		t.Fatal("generated id is empty")
	}
	if addr.CustomerId != customer.Id {
		t.Fatal("customer ids do not match")
	}
	if addr.FirstName != addrRequest.FirstName {
		t.Fatal("first names do not match")
	}
	if addr.LastName != addrRequest.LastName {
		t.Fatal("last names do not match")
	}
	if addr.Company != addrRequest.Company {
		t.Fatal("companies do not match")
	}
	if addr.StreetAddress != addrRequest.StreetAddress {
		t.Fatal("street addresses do not match")
	}
	if addr.ExtendedAddress != addrRequest.ExtendedAddress {
		t.Fatal("extended addresses do not match")
	}
	if addr.Locality != addrRequest.Locality {
		t.Fatal("localities do not match")
	}
	if addr.Region != addrRequest.Region {
		t.Fatal("regions do not match")
	}
	if addr.PostalCode != addrRequest.PostalCode {
		t.Fatal("postal codes do not match")
	}
	if addr.CountryCodeAlpha2 != addrRequest.CountryCodeAlpha2 {
		t.Fatal("country alpha2 codes do not match")
	}
	if addr.CountryCodeAlpha3 != addrRequest.CountryCodeAlpha3 {
		t.Fatal("country alpha3 codes do not match")
	}
	if addr.CountryCodeNumeric != addrRequest.CountryCodeNumeric {
		t.Fatal("country numeric codes do not match")
	}
	if addr.CountryName != addrRequest.CountryName {
		t.Fatal("country names do not match")
	}
	if addr.CreatedAt == nil {
		t.Fatal("generated created at is empty")
	}
	if addr.UpdatedAt == nil {
		t.Fatal("generated updated at is empty")
	}
}
