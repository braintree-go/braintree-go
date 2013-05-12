package braintree

type Transaction struct {
	Amount        int
	PaymentMethod CreditCard
}

type TransactionRequest struct {
	tx Transaction
}

func NewTransactionRequest() TransactionRequest { return TransactionRequest{Transaction{}} }

func (this TransactionRequest) Amount(amount int) TransactionRequest {
	this.tx.Amount = amount
	return this
}

func (this TransactionRequest) CreditCard() TransactionCreditCardRequest {
	return NewTransactionCreditCardRequest(this)
}
