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
		log.Fatalln("must provide merchant account id")
	}

	var gateway = braintree.New(
		braintree.Sandbox,
		os.Getenv("BRAINTREE_MERCH_ID"),
		os.Getenv("BRAINTREE_PUB_KEY"),
		os.Getenv("BRAINTREE_PRIV_KEY"),
	)

	maid := flag.Arg(0)
	fmt.Println(maid)

	merchantAccount, err := gateway.MerchantAccount().Find(maid)

	if err == nil {
		fmt.Println(merchantAccount.Status)
	} else {
		fmt.Println(err)
	}
}
