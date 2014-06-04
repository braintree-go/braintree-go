package braintree

import (
	"encoding/xml"
)

type TransactionGateway struct {
	*Braintree
}

// Create initiates a transaction.
func (g *TransactionGateway) Create(tx *Transaction) (*Transaction, error) {
	resp, err := g.execute("POST", "transactions", tx)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// SubmitForSettlement submits the transaction with the specified id for settlement.
// If the amount is omitted, the full amount is settled.
func (g *TransactionGateway) SubmitForSettlement(id string, amount ...float64) (*Transaction, error) {
	var tx *Transaction
	if len(amount) > 0 {
		tx = &Transaction{
			Amount: amount[0],
		}
	}
	resp, err := g.execute("PUT", "transactions/"+id+"/submit_for_settlement", tx)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// HoldInEscrow sets the transaction to be held for escrow.
// You can only hold a transaction in escrow if its status is authorized or submitted_for_settlement.
// Afterward, the transaction's escrow-status field should be "hold_pending".
// The Escrow Status will be "held" upon entering a "settled" status.
func (g *TransactionGateway) HoldInEscrow(id string) (*Transaction, error) {
	resp, err := g.execute("PUT", "transactions/"+id+"/hold_in_escrow", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// ReleaseFromEscrow will disburse to the Sub Merchant account (and Master Merchant account if there was a Service Fee)
// on the next business day. The Escrow Status of a transaction must be "held" before its funds can be released from escrow.
// A successful call to release from escrow will change the Escrow Status to "release_pending".
// The Escrow Status will change to "released" upon being disbursed the next business day:
func (g *TransactionGateway) ReleaseFromEscrow(id string) (*Transaction, error) {
	resp, err := g.execute("PUT", "transactions/"+id+"/release_from_escrow", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// CancelRelease stops the release of an escrowed transaction. y
// You can only cancel the release while the escrow-status is "release_pending".
// Upon a successful CancelRelease() the escrow-status will change back to "held".
func (g *TransactionGateway) CancelRelease(id string) (*Transaction, error) {
	resp, err := g.execute("PUT", "transactions/"+id+"/cancel_release", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Void voids the transaction with the specified id if it has a status of authorized or
// submitted_for_settlement. When the transaction is voided Braintree will do an authorization
// reversal if possible so that the customer won’t have a pending charge on their card
func (g *TransactionGateway) Void(id string) (*Transaction, error) {
	resp, err := g.execute("PUT", "transactions/"+id+"/void", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Find finds the transaction with the specified id.
func (g *TransactionGateway) Find(id string) (*Transaction, error) {
	resp, err := g.execute("GET", "transactions/"+id, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Search finds all transactions matching the search query.
func (g *TransactionGateway) Search(query *SearchQuery) (*TransactionSearchResult, error) {
	resp, err := g.execute("POST", "transactions/advanced_search", query)
	if err != nil {
		return nil, err
	}
	var v TransactionSearchResult
	err = xml.Unmarshal(resp.Body, &v)
	if err != nil {
		return nil, err
	}
	return &v, err
}
