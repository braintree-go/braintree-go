package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/publicgoodsw/braintree-go"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatalln("USAGE: txstatus [command] txid")
	}

	var gateway = braintree.New(
		braintree.Sandbox,
		os.Getenv("BRAINTREE_MERCH_ID"),
		os.Getenv("BRAINTREE_PUB_KEY"),
		os.Getenv("BRAINTREE_PRIV_KEY"),
	)

	var showstatus bool

	cmd := strings.ToUpper(flag.Arg(0))
	txid := flag.Arg(1)
	fmt.Println("Transaction: " + txid)
	tx, err := gateway.Transaction().Find(txid)
	if tx == nil {
		fmt.Println("Transanction not found.")
	} else {
		fmt.Println("Transaction found.")
	}

	fmt.Println(cmd)

	switch cmd {
	case "STATUS":
		showstatus = true

	case "SUBMIT":
		showstatus = true
		if tx.Status == braintree.TxAuthorized {
			tx, err = gateway.Transaction().SubmitForSettlement(txid)
		} else {
			fmt.Println("Cannot settle transaction until transaction is authorized.")
		}

	case "HOLD":
		showstatus = true
		if (tx.Status == braintree.TxSubmittedForSettlement) || (tx.Status == braintree.TxAuthorized) {
			tx, err = gateway.Transaction().HoldInEscrow(txid)
		} else {
			fmt.Println("Cannot release transaction until transaction is authorized or submitted for settlement.")
		}

	case "RELEASE":
		showstatus = true
		if tx.EscrowStatus == braintree.TxEscrowHeld {
			tx, err = gateway.Transaction().ReleaseFromEscrow(txid)
		} else {
			fmt.Println("Cannot release transaction until held for escrow.")
		}

	case "CANCELRELEASE":
		showstatus = true
		if tx.EscrowStatus == braintree.TxEscrowReleasePending {
			tx, err = gateway.Transaction().CancelRelease(txid)
		} else {
			fmt.Println("Cannot cancel release until release is pending.")
		}

	case "VOID":
		showstatus = true
		if (tx.Status == braintree.TxSubmittedForSettlement) || (tx.Status == braintree.TxAuthorized) {
			tx, err = gateway.Transaction().Void(txid)
		} else {
			fmt.Println("Cannot void transaction until transaction is authorized or submitted for settlement.")
		}
	}

	if showstatus {
		if err == nil {
			fmt.Println("Status: " + tx.Status)
			fmt.Println("Escrow Status: " + tx.EscrowStatus)
		} else {
			fmt.Println(err)
		}
	}
}
