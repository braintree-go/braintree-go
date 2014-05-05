/*

  This sample application shows how to use Braintree Go in a safe, PCI-compliant way. To use it
   yourself, export your Braintree sandbox credentials as these environmental variables:

   export BRAINTREE_MERCH_ID={your-merchant-id}
   export BRAINTREE_PUB_KEY={your-public-key}
   export BRAINTREE_PRIV_KEY={your-private-key}
   export BRAINTREE_CSE_KEY={your-cse-key}

  For a list of testing values and expected behaviors, see
  https://www.braintreepayments.com/docs/ruby/reference/sandbox

*/

package main

import (
	"fmt"
	braintree "github.com/azoff/braintree-go"
	"html/template"
	"net/http"
	"os"
)

type BraintreeJS struct {
	Key template.HTML
}

func showForm(w http.ResponseWriter, r *http.Request) {
	config := BraintreeJS{Key: "'" + template.HTML(os.Getenv("BRAINTREE_CSE_KEY")) + "'"}
	t := template.Must(template.ParseFiles("form.html"))
	t.Execute(w, config)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	bt := braintree.New(
		braintree.Sandbox,
		os.Getenv("BRAINTREE_MERCH_ID"),
		os.Getenv("BRAINTREE_PUB_KEY"),
		os.Getenv("BRAINTREE_PRIV_KEY"),
	)

	tx := &braintree.Transaction{
		Type:   "sale",
		Amount: 100,
		CreditCard: &braintree.CreditCard{
			Number:          r.FormValue("number"),
			CVV:             r.FormValue("cvv"),
			ExpirationMonth: r.FormValue("month"),
			ExpirationYear:  r.FormValue("year"),
		},
	}

	_, err := bt.Transaction().Create(tx)

	if err == nil {
		fmt.Fprintf(w, "<h1>Success!</h1>")
	} else {
		fmt.Fprintf(w, "<h1>Something went wrong: "+err.Error()+"</h1>")
	}
}

func main() {
	http.HandleFunc("/", showForm)
	http.HandleFunc("/create_transaction", createTransaction)
	http.ListenAndServe(":8080", nil)
}
