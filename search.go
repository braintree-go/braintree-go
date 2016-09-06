package braintree

import (
	"encoding/xml"
)

type SearchQuery struct {
	XMLName string `xml:"search" json:"search" bson:"search"`
	Fields  []interface{}
}

type SearchResults struct {
	XMLName  string `xml:"search-results" json:"searchResults" bson:"searchResults"`
	PageSize string `xml:"page-size" json:"pageSize" bson:"pageSize"`
	Ids      struct {
		Item []string `xml:"item" json:"item" bson:"item"`
	} `xml:"ids" json:"ids" bson:"ids"`
}

type TextField struct {
	XMLName    xml.Name
	Is         string `xml:"is,omitempty" json:"is,omitempty" bson:"is,omitempty"`
	IsNot      string `xml:"is-not,omitempty" json:"isNot,omitempty" bson:"isNot,omitempty"`
	StartsWith string `xml:"starts-with,omitempty" json:"startsWith,omitempty" bson:"startsWith,omitempty"`
	EndsWith   string `xml:"ends-with,omitempty" json:"endsWith,omitempty" bson:"endsWith,omitempty"`
	Contains   string `xml:"contains,omitempty" json:"contains,omitempty" bson:"contains,omitempty"`
}

type RangeField struct {
	XMLName xml.Name
	Is      float64 `xml:"is,omitempty" json:"is,omitempty" bson:"is,omitempty"`
	Min     float64 `xml:"min,omitempty" json:"min,omitempty" bson:"min,omitempty"`
	Max     float64 `xml:"max,omitempty" json:"max,omitempty" bson:"max,omitempty"`
}

type MultiField struct {
	XMLName xml.Name
	Type    string   `xml:"type,attr" json:"type" bson:"type"`
	Items   []string `xml:"item" json:"item" bson:"item"`
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
