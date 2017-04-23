package braintree

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHttpClientTimeout(t *testing.T) {
	t.Parallel()

	const gracePeriod = time.Second * 10

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second*60 + gracePeriod)
	}))
	env := NewEnvironment(server.URL)

	testCases := []struct {
		Braintree       *Braintree
		ExpectedTimeout time.Duration
	}{
		{
			Braintree:       New(env, "mid", "pubkey", "privkey"),
			ExpectedTimeout: time.Second * 60,
		},
		{
			Braintree:       NewWithHttpClient(env, "mid", "pubkey", "privkey", &http.Client{Timeout: time.Second * 10}),
			ExpectedTimeout: time.Second * 10,
		},
	}

	for _, tc := range testCases {
		finished := make(chan bool)
		go func() {
			_, err := tc.Braintree.Transaction().Create(&TransactionRequest{})
			if err == nil {
				t.Fatal("Expected timeout error, received no error")
			}
			if !strings.Contains(err.Error(), "Timeout") {
				t.Fatalf("Expected timeout error, received: %s", err)
			}
			finished <- true
		}()

		select {
		case <-finished:
			t.Logf("Timeout received as expected")
		case <-time.After(tc.ExpectedTimeout + gracePeriod):
			t.Fatalf("Timeout did not occur around %s, has been at least %s", tc.ExpectedTimeout, tc.ExpectedTimeout+gracePeriod)
		}
	}
}
