package braintree

import (
	"testing"
)

func TestAddress(t *testing.T) {
	customer, err := testGateway.Customer().Create(&Customer{
		FirstName: "Jenna",
		LastName:  "Smith",
	})
	if err != nil {
		t.Fatal(err)
	}
	if customer.Id == "" {
		t.Fatal("invalid customer id")
	}

	addr := &Address{
		CustomerId:         customer.Id,
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

	addr2, err := testGateway.Address().Create(addr)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)
	t.Log(addr2)

	if addr2.Id == "" {
		t.Fatal()
	}
	if addr2.CustomerId != customer.Id {
		t.Fatal()
	}
	if addr2.FirstName != addr.FirstName {
		t.Fatal()
	}
	if addr2.LastName != addr.LastName {
		t.Fatal()
	}
	if addr2.Company != addr.Company {
		t.Fatal()
	}
	if addr2.StreetAddress != addr.StreetAddress {
		t.Fatal()
	}
	if addr2.ExtendedAddress != addr.ExtendedAddress {
		t.Fatal()
	}
	if addr2.Locality != addr.Locality {
		t.Fatal()
	}
	if addr2.Region != addr.Region {
		t.Fatal()
	}
	if addr2.PostalCode != addr.PostalCode {
		t.Fatal()
	}
	if addr2.CountryCodeAlpha2 != addr.CountryCodeAlpha2 {
		t.Fatal()
	}
	if addr2.CountryCodeAlpha3 != addr.CountryCodeAlpha3 {
		t.Fatal()
	}
	if addr2.CountryCodeNumeric != addr.CountryCodeNumeric {
		t.Fatal()
	}
	if addr2.CountryName != addr.CountryName {
		t.Fatal()
	}
	if addr2.CreatedAt == "" {
		t.Fatal()
	}
	if addr2.UpdatedAt == "" {
		t.Fatal()
	}

	err = testGateway.Address().Delete(customer.Id, addr2.Id)
	if err != nil {
		t.Fatal(err)
	}
}
