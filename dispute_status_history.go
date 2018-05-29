package braintree

import "time"

type DisputeStatusHistoryEvent struct {
	XMLName          string     `xml:"status-history"`
	DisbursementDate string     `xml:"disbursement-date"`
	EffectiveDate    string     `xml:"effective-date"`
	Status           string     `xml:"status"`
	Timestamp        *time.Time `xml:"timestamp"`
}
