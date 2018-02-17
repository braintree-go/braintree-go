/*

  This sample application shows how to use Braintree Go to process
	subscriptions in a safe, PCI-compliant way. To use it yourself, export your
	Braintree sandbox credentials as these environmental variables:

   export BRAINTREE_MERCH_ID={your-merchant-id}
   export BRAINTREE_PUB_KEY={your-public-key}
   export BRAINTREE_PRIV_KEY={your-private-key}
   export BRAINTREE_CSE_KEY={your-cse-key}

  For a list of testing values and expected behaviors, see
  https://www.braintreepayments.com/docs/ruby/reference/sandbox

*/

package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/lionelbarrow/braintree-go"
)

func main() {
	bt := braintree.New(
		braintree.Sandbox,
		os.Getenv("BRAINTREE_MERCH_ID"),
		os.Getenv("BRAINTREE_PUB_KEY"),
		os.Getenv("BRAINTREE_PRIV_KEY"),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		clientToken, err := bt.ClientToken().Generate(ctx)
		if err != nil {
			log.Fatal(err)
		}
		t := template.Must(template.ParseFiles("form.html"))
		w.WriteHeader(http.StatusOK)
		err = t.Execute(w, clientToken)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		paymentMethodNonce := r.PostFormValue("payment_method_nonce")
		if paymentMethodNonce == "" {
			log.Fatal("Payment method nonce is empty")
		}
		// You can later search for the user by his ID
		// customer, err := bt.Customer().Find("CustomerID")
		customer, err := bt.Customer().Create(ctx, &braintree.CustomerRequest{
			// You can leave it empty, but, if you've got a user system, I recommend using the user's ID as the client ID
			// Or, createa a row for Braintree's customer ID
			ID: "<CustomerID>",
		})
		if err != nil {
			log.Fatal(err)
		}

		// We create a credit card that after generation, gives us the PaymentMethodToken that's needed in the Subscription.Create
		card, err := bt.CreditCard().Create(ctx, &braintree.CreditCard{
			// The created or existing customer ID
			CustomerId: customer.Id,
			// The nonce from the clinet side
			PaymentMethodNonce: paymentMethodNonce,
			Options: &braintree.CreditCardOptions{
				VerifyCard: func(b bool) *bool { return &b }(true),
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		// Create the subscription and make the user pay
		subscription, err := bt.Subscription().Create(ctx, &braintree.SubscriptionRequest{
			PlanId: "<YourPlanIDShouldGoHere>",
			// The payment method token generated by the CreditCard.Create
			PaymentMethodToken: card.Token,
		})
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(fmt.Sprintf("Success! Subscription #%s created with user ID %s", subscription.Id, customer.Id)))
		if err != nil {
			log.Fatal(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
