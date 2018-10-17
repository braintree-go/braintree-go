/*
  This sample application shows how to use Braintree Go.
  To use it yourself, export your Braintree sandbox credentials as these environmental variables:

  export BRAINTREE_MERCH_ID={your-merchant-id}
  export BRAINTREE_PUB_KEY={your-public-key}
  export BRAINTREE_PRIV_KEY={your-private-key}

  For a list of testing values and expected behaviors, see
  https://developers.braintreepayments.com/guides/3d-secure/testing-go-live
*/

package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/braintree-go/braintree-go"
)

func getBT() *braintree.Braintree {
	bt := braintree.New(
		braintree.Sandbox,
		os.Getenv("BRAINTREE_MERCH_ID"),
		os.Getenv("BRAINTREE_PUB_KEY"),
		os.Getenv("BRAINTREE_PRIV_KEY"),
	)
	return bt
}

func showForm(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	bt := getBT()

	clientToken, err := bt.ClientToken().Generate(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t := template.Must(template.ParseFiles("form.html"))
	config := struct {
		ClientToken string
	}{
		ClientToken: clientToken,
	}
	err = t.Execute(w, config)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	bt := getBT()

	nonce := r.FormValue("nonce")

	pmn, _ := bt.PaymentMethodNonce().Find(ctx, nonce)

	tx := &braintree.TransactionRequest{
		Type:               "sale",
		Amount:             braintree.NewDecimal(1000, 2),
		PaymentMethodNonce: nonce,
		Options: &braintree.TransactionOptions{
			ThreeDSecure: &braintree.TransactionOptionsThreeDSecureRequest{Required: true},
		},
	}

	txn, err := bt.Transaction().Create(ctx, tx)

	t := template.Must(template.ParseFiles("result.html"))
	config := struct {
		PaymentMethodNonce *braintree.PaymentMethodNonce
		Error              error
		Transaction        *braintree.Transaction
	}{
		PaymentMethodNonce: pmn,
		Error:              err,
		Transaction:        txn,
	}
	err = t.Execute(w, config)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	http.HandleFunc("/", showForm)
	http.HandleFunc("/create_transaction", createTransaction)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
