package braintree

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestTransactionMarshalsExternalVaultTLID(t *testing.T) {
	tx := &Transaction{
		Type: "sale",
		ExternalVault: &ExternalVault{
			PreviousNetworkTransactionID:        "123456789012345",
			PreviousMastercardTransactionLinkID: "JavmPovbDUHu8aIAyjexxG",
			Status:                              ExternalVaultStatusVaulted,
		},
	}

	b, err := xml.Marshal(tx)
	if err != nil {
		t.Fatal(err)
	}

	want := "<external-vault>" +
		"<previous-network-transaction-id>123456789012345</previous-network-transaction-id>" +
		"<previous-mastercard-transaction-link-id>JavmPovbDUHu8aIAyjexxG</previous-mastercard-transaction-link-id>" +
		"<status>vaulted</status>" +
		"</external-vault>"
	if !strings.Contains(string(b), want) {
		t.Fatalf("expected %s in %s", want, b)
	}
}

func TestTransactionMarshalsWillVaultWithoutPreviousIDs(t *testing.T) {
	tx := &Transaction{
		Type:          "sale",
		ExternalVault: &ExternalVault{Status: ExternalVaultStatusWillVault},
	}

	b, err := xml.Marshal(tx)
	if err != nil {
		t.Fatal(err)
	}

	want := "<external-vault><status>will_vault</status></external-vault>"
	if !strings.Contains(string(b), want) {
		t.Fatalf("expected %s in %s", want, b)
	}
}

func TestTransactionUnmarshalsTLID(t *testing.T) {
	payload := `<transaction>
		<network-transaction-id>020261102133130</network-transaction-id>
		<mastercard-transaction-link-id>JavmPovbDUHu8aIAyjexxG</mastercard-transaction-link-id>
	</transaction>`

	var tx Transaction
	if err := xml.Unmarshal([]byte(payload), &tx); err != nil {
		t.Fatal(err)
	}

	if tx.NetworkTransactionID != "020261102133130" {
		t.Errorf("unexpected network transaction id %q", tx.NetworkTransactionID)
	}
	if tx.MastercardTransactionLinkID != "JavmPovbDUHu8aIAyjexxG" {
		t.Errorf("unexpected mastercard transaction link id %q", tx.MastercardTransactionLinkID)
	}
}

func TestTransactionUnmarshalsNilTLID(t *testing.T) {
	payload := `<transaction><mastercard-transaction-link-id nil="true"/></transaction>`

	var tx Transaction
	if err := xml.Unmarshal([]byte(payload), &tx); err != nil {
		t.Fatal(err)
	}

	if tx.MastercardTransactionLinkID != "" {
		t.Errorf("expected empty mastercard transaction link id, got %q", tx.MastercardTransactionLinkID)
	}
}

func TestCreditCardVerificationUnmarshalsTLID(t *testing.T) {
	payload := `<verification><mastercard-transaction-link-id>JavmPovbDUHu8aIAyjexxG</mastercard-transaction-link-id></verification>`

	var v CreditCardVerification
	if err := xml.Unmarshal([]byte(payload), &v); err != nil {
		t.Fatal(err)
	}

	if v.MastercardTransactionLinkID != "JavmPovbDUHu8aIAyjexxG" {
		t.Errorf("unexpected mastercard transaction link id %q", v.MastercardTransactionLinkID)
	}
}
