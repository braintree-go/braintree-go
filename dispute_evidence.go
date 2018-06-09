package braintree

import "time"

type DisputeEvidence struct {
	XMLName           string     `xml:"evidence"`
	Comment           string     `xml:"comment"`
	CreatedAt         *time.Time `xml:"created-at"`
	ID                string     `xml:"id"`
	SentToProcessorAt string     `xml:"sent-to-processor-at"`
	URL               string     `xml:"url"`
}
