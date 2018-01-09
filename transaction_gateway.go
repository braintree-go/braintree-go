package braintree

import (
	"context"
	"encoding/xml"
	"fmt"
)

type TransactionGateway struct {
	*Braintree
}

// Create initiates a transaction.
func (g *TransactionGateway) Create(ctx context.Context, tx *TransactionRequest) (*Transaction, error) {
	resp, err := g.execute(ctx, "POST", "transactions", tx)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Clone clones a transaction.
func (g *TransactionGateway) Clone(ctx context.Context, id string, tx *TransactionCloneRequest) (*Transaction, error) {
	resp, err := g.execute(ctx, "POST", "transactions/"+id+"/clone", tx)
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
func (g *TransactionGateway) SubmitForSettlement(ctx context.Context, id string, amount ...*Decimal) (*Transaction, error) {
	var tx *TransactionRequest
	if len(amount) > 0 {
		tx = &TransactionRequest{
			Amount: amount[0],
		}
	}
	resp, err := g.execute(ctx, "PUT", "transactions/"+id+"/submit_for_settlement", tx)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Settle settles a transaction.
// This action is only available in the sandbox environment.
// Deprecated: use the Settle function on the TestingGateway instead. e.g. g.Testing().Settle(id).
func (g *TransactionGateway) Settle(ctx context.Context, id string) (*Transaction, error) {
	return g.Testing().Settle(ctx, id)
}

// Void voids the transaction with the specified id if it has a status of authorized or
// submitted_for_settlement. When the transaction is voided Braintree will do an authorization
// reversal if possible so that the customer won't have a pending charge on their card
func (g *TransactionGateway) Void(ctx context.Context, id string) (*Transaction, error) {
	resp, err := g.execute(ctx, "PUT", "transactions/"+id+"/void", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// CancelRelease cancels a pending release of a transaction with the given id from escrow.
func (g *TransactionGateway) CancelRelease(ctx context.Context, id string) (*Transaction, error) {
	resp, err := g.execute(ctx, "PUT", "transactions/"+id+"/cancel_release", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// ReleaseFromEscrow submits the transaction with the given id for release from escrow.
func (g *TransactionGateway) ReleaseFromEscrow(ctx context.Context, id string) (*Transaction, error) {
	resp, err := g.execute(ctx, "PUT", "transactions/"+id+"/release_from_escrow", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// HoldInEscrow holds the transaction with the given id for escrow.
func (g *TransactionGateway) HoldInEscrow(ctx context.Context, id string) (*Transaction, error) {
	resp, err := g.execute(ctx, "PUT", "transactions/"+id+"/hold_in_escrow", nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// A transaction can be refunded if it is settled or settling.
// If the transaction has not yet begun settlement, use Void() instead.
// If you do not specify an amount to refund, the entire transaction amount will be refunded.
func (g *TransactionGateway) Refund(ctx context.Context, id string, amount ...*Decimal) (*Transaction, error) {
	var tx *TransactionRequest
	if len(amount) > 0 {
		tx = &TransactionRequest{
			Amount: amount[0],
		}
	}
	resp, err := g.execute(ctx, "POST", "transactions/"+id+"/refund", tx)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.transaction()
	case 201:
		return resp.transaction()
	}
	return nil, &invalidResponseError{resp}
}

// Find finds the transaction with the specified id.
func (g *TransactionGateway) Find(ctx context.Context, id string) (*Transaction, error) {
	resp, err := g.execute(ctx, "GET", "transactions/"+id, nil)
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
// Per https://developers.braintreepayments.com/reference/general/searching/search-results/java a max of 20,000 results can be returned.
func (g *TransactionGateway) Search(ctx context.Context, query *SearchQuery) (*TransactionSearchResult, error) {
	// Get the ids of all transactions that match the search criteria.
	resp, err := g.execute(ctx, "POST", "transactions/advanced_search_ids", query)
	if err != nil {
		return nil, err
	}
	v := new(TransactionSearchResult)
	err = xml.Unmarshal(resp.Body, v)
	if err != nil {
		return nil, err
	}

	// Get transaction details and add them to the TransactionSearchResult. Only v.PageSize transactions can be fetched at one time.
	totalItems := len(v.Ids.Item)
	txs, err := fetchTransactions(ctx, g, query, v.Ids.Item, totalItems, v.PageSize)
	if err != nil {
		return nil, err
	}
	v.Transactions = txs

	// Add data for backward compatibility. CurrnetPageNumber is pretty meaningless though.
	v.CurrentPageNumber = 1
	v.TotalItems = totalItems

	return v, err
}

func fetchTransactions(ctx context.Context, g executer, query *SearchQuery, txIds []string, totalItems int, pageSize int) ([]*Transaction, error) {
	txs := make([]*Transaction, 0)
	ids := query.AddMultiField("ids")
	for i := 0; i < totalItems; i += pageSize {
		ids.Items = txIds[i:min(i+pageSize, totalItems)]
		res, fetchError := fetchTransactionPage(ctx, g, query)
		if fetchError != nil {
			return nil, fetchError
		}
		for _, tx := range res.Transactions {
			txs = append(txs, tx)
		}
	}
	if len(txs) != totalItems {
		return nil, fmt.Errorf("Unexpected number of transactions. Got %v, want %v.", len(txs), totalItems)
	}
	return txs, nil
}

type testOperationPerformedInProductionError struct {
	error
}

func (e *testOperationPerformedInProductionError) Error() string {
	return fmt.Sprint("Operation not allowed in production environment")
}

// FetchTransactionPage returns a page of transactions including details for a given query.
func fetchTransactionPage(ctx context.Context, g executer, query *SearchQuery) (*transactionDetailSearchResult, error) {
	resp, err := g.execute(ctx, "POST", "transactions/advanced_search", query)
	if err != nil {
		return nil, err
	}
	v := new(transactionDetailSearchResult)
	err = xml.Unmarshal(resp.Body, v)
	if err != nil {
		return nil, err
	}
	return v, err
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// TransactionDetailSearchResult is used to fetch a page of transactions with all details.
type transactionDetailSearchResult struct {
	XMLName           string         `xml:"credit-card-transactions"`
	CurrentPageNumber int            `xml:"current-page-number"`
	PageSize          int            `xml:"page-size"`
	TotalItems        int            `xml:"total-items"`
	Transactions      []*Transaction `xml:"transaction"`
}

// Executer provides an execute method.
type executer interface {
	execute(ctx context.Context, method, path string, xmlObj interface{}) (*Response, error)
}
