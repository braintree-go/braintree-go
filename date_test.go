package braintree

import (
	"encoding/xml"
	"testing"
)

func TestDateUnmarshalXML(t *testing.T) {
	date := &Date{}

	dateXML := []byte(`<?xml version="1.0" encoding="UTF-8"?><foo>2014-02-09</foo></xml>`)
	if err := xml.Unmarshal(dateXML, date); err != nil {
		t.Fatal(err)
	}

	if date.Format("2006-01-02") != "2014-02-09" {
		t.Fatalf("expected 2014-02-09 got %s", date)
	}
}
