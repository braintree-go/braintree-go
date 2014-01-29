package braintree

import (
	"encoding/xml"
)

type SearchQuery struct {
	XMLName string `json:"search" xml:"search"`
	Fields  []interface{}
}

type SearchResults struct {
	XMLName  string `json:"search-results" xml:"search-results"`
	PageSize string `json:"page-size" xml:"page-size"`
	Ids      struct {
		Item []string `json:"item" xml:"item"`
	} `json:"ids" xml:"ids"`
}

type TextField struct {
	XMLName    xml.Name
	Is         string `json:"is,omitempty" xml:"is,omitempty"`
	IsNot      string `json:"is-not,omitempty" xml:"is-not,omitempty"`
	StartsWith string `json:"starts-with,omitempty" xml:"starts-with,omitempty"`
	EndsWidth  string `json:"ends-with,omitempty" xml:"ends-with,omitempty"`
	Contains   string `json:"contains,omitempty" xml:"contains,omitempty"`
}

type RangeField struct {
	XMLName xml.Name
	Is      float64 `json:"is,omitempty" xml:"is,omitempty"`
	Min     float64 `json:"min,omitempty" xml:"min,omitempty"`
	Max     float64 `json:"max,omitempty" xml:"max,omitempty"`
}

type MultiField struct {
	XMLName xml.Name
	Type    string   `json:"type,attr" xml:"type,attr"` // type=array
	Items   []string `json:"item" xml:"item"`
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

