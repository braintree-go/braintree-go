package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/publicgoodsw/braintree-go"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalln("must provide transaction id")
	}

	var gateway = braintree.New(
		braintree.Sandbox,
		os.Getenv("BRAINTREE_MERCH_ID"),
		os.Getenv("BRAINTREE_PUB_KEY"),
		os.Getenv("BRAINTREE_PRIV_KEY"),
	)

	txid := flag.Arg(0)
	fmt.Println(txid)

	tx, err := gateway.Transaction().Find(txid)

	if err == nil {
		fmt.Println("Status: " + tx.Status)
		fmt.Println("Escrow Status: " + tx.EscrowStatus)
	} else {
		fmt.Println(err)
	}
}
