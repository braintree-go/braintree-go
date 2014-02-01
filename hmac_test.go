package braintree

import (
	"testing"
)

func TestHmacerParseSignature(t *testing.T) {
	hmacer := newHmacer(testGateway)

	// Happy path
	realSignature, err := hmacer.parseSignature(testGateway.PublicKey + "|my_actual_signature")
	if err != nil {
		t.Fatal(err)
	} else if realSignature != "my_actual_signature" {
		t.Fatal("parseSignature returned wrong signature")
	}

	// Test hmacer rejects an incorrect public key
	_, err = hmacer.parseSignature("some_random_public_key|my_actual_signature")
	if err == nil {
		t.Fatal("Did not receive an error when the wrong public key was passed")
	}

	// Test hmacer rejects a signature-key pair with more than one pipe
	_, err = hmacer.parseSignature("some_random_public_key|some_other_stuff|my_actual_signature")
	if err == nil {
		t.Fatal("Did not receive an error when the wrong public key was passed")
	}

	// Test hmacer rejects a singature-key pair with no pipes
	_, err = hmacer.parseSignature("some_random_public_key&my_actual_signature")
	if err == nil {
		t.Fatal("Did not receive an error when the wrong public key was passed")
	}
}

func TestHmacerVerifySignature(t *testing.T) {
	gateway := New(Sandbox, "my_merchant_id", "my_public_key", "my_private_key")
	hmacer := newHmacer(gateway)
	signatureKeyPair := gateway.PublicKey + "|fa654fa4fe5537934960c483dbb0ee575d64b6ad"
	payload := "my_random_value"

	// Happy path
	verified, err := hmacer.verifySignature(signatureKeyPair, payload)

	if err != nil {
		t.Fatal(err)
	} else if !verified {
		t.Fatal("Did not verify correct signature")
	}

	// Test hmacer does not verify when the payload has been modified
	verified, err = hmacer.verifySignature(signatureKeyPair, payload+"a bad man in the middle")

	if err != nil {
		t.Fatal(err)
	} else if verified {
		t.Fatal("HMACer verified invalid signature.")
	}
}
