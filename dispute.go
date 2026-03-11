package braintree

// Dispute object for a *Transaction*
type Dispute struct {
	Id                string   `xml:"id"`
	Amount            *Decimal `xml:"amount"`
	CurrencyISOCode   string   `xml:"currency-iso-code"`
	AmountWon         *Decimal `xml:"amount-won"`
	ReceivedDate      Time     `xml:"received-date"`
	DateWon           Time     `xml:"date-won"`
	DateOpened        Time     `xml:"date-opened"`
	Kind              string   `xml:"kind"`
	Reason            string   `xml:"reason"`
	ReasonCode        string   `xml:"reason-code"`
	ReasonDescription string   `xml:"reason-description"`
	ReplyByDate       string   `xml:"reply-by-date"`

	Transaction *Transaction `xml:"transaction"`

	Status string `xml:"status"`
}

type TransactionDetails struct {
	Id     string   `xml:"id"`
	Amount *Decimal `xml:"amount"`
}

type DisputeStatusHistory struct {
	Events []DisputeStatusEvent `xml:"status-event"`
}

type DisputeStatusEvent struct {
	Timestamp string `xml:"timestamp"`
	Status    string `xml:"status"`
}
