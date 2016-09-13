package braintree

import (
	"encoding/xml"
)

type SearchQuery struct {
	XMLName string `xml:"search"`
	Fields  []interface{}
}

type SearchResults struct {
	XMLName  string `xml:"search-results"`
	PageSize string `xml:"page-size"`
	Ids      struct {
		Item []string `xml:"item"`
	} `xml:"ids"`
}

type TextField struct {
	XMLName    xml.Name
	Is         string `xml:"is,omitempty"`
	IsNot      string `xml:"is-not,omitempty"`
	StartsWith string `xml:"starts-with,omitempty"`
	EndsWith   string `xml:"ends-with,omitempty"`
	Contains   string `xml:"contains,omitempty"`
}

type RangeField struct {
	XMLName xml.Name
	Is      interface{} `xml:"is,omitempty"`
	Min     interface{} `xml:"min,omitempty"`
	Max     interface{} `xml:"max,omitempty"`
}

type MultiField struct {
	XMLName xml.Name
	Type    string   `xml:"type,attr"` // type=array
	Items   []string `xml:"item"`
}

func (s *SearchQuery) AddTextField(field string) *TextField {
	f := &TextField{XMLName: xml.Name{Local: field}}
	s.Fields = append(s.Fields, f)
	return f
}

func (s *SearchQuery) AddRangeField(field string) *RangeField {
	f := &RangeField{XMLName: xml.Name{Local: field}}
	s.Fields = append(s.Fields, f)
	return f
}

func (s *SearchQuery) AddMultiField(field string) *MultiField {
	f := &MultiField{
		XMLName: xml.Name{Local: field},
		Type:    "array",
	}
	s.Fields = append(s.Fields, f)
	return f
}
