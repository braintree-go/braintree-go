package braintree

import (
	"encoding/xml"
	"time"
)

// Date wraps a time object but handles deserializing dates returned from the Braintree API
// e.g. "2014-02-09"
type Date struct {
	*time.Time
}

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var v string
	dec.DecodeElement(&v, &start)

	parse, err := time.Parse("2006-01-02", v)
	if err != nil {
		return err
	}

	*d = Date{Time: &parse}
	return nil
}
