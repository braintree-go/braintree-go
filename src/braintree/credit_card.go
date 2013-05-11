package braintree

type CreditCard struct {
  Number string
  ExpirationDate string
}

func NewTransactionCreditCardRequest(parent TransactionRequest) TransactionCreditCardRequest {
  return TransactionCreditCardRequest{parent, CreditCard{}}
}

type TransactionCreditCardRequest struct {
  parent TransactionRequest
  creditCard CreditCard
}

func (this TransactionCreditCardRequest) Number(number string) TransactionCreditCardRequest {
  this.creditCard.Number = number
  return this
}

func (this TransactionCreditCardRequest) ExpirationDate(date string) TransactionCreditCardRequest {
  this.creditCard.ExpirationDate = date
  return this
}

func (this TransactionCreditCardRequest) Done() TransactionRequest {
  return this.parent
}
