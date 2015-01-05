package braintree

import "testing"

// This test will fail unless you set up your Braintree sandbox account correctly. See TESTING.md for details.
func TestClientToken(t *testing.T) {
	g := testGateway.ClientToken()
	token, err := g.Generate()
	if err != nil {
		t.Fatalf("failed to generate client token: %s", err)
	}
	if len(token) == 0 {
		t.Fatalf("empty client token!")
	}
}
