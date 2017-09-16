package braintree

type TransactionCloneRequest struct {
	XMLName string                   `xml:"transaction-clone"`
	Amount  *Decimal                 `xml:"amount"`
	Options *TransactionCloneOptions `xml:"options"`
}

type TransactionCloneOptions struct {
	SubmitForSettlement bool `xml:"submit-for-settlement"`
}
