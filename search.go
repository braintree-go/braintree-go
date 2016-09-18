package braintree

import (
	"encoding/xml"
	"time"
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
	Is      float64 `xml:"is,omitempty"`
	Min     float64 `xml:"min,omitempty"`
	Max     float64 `xml:"max,omitempty"`
}

type RangeDateField struct {
	XMLName xml.Name
	Is      time.Time `xml:"is,omitempty"`
	Min     time.Time `xml:"min,omitempty"`
	Max     time.Time `xml:"max,omitempty"`
}

func (d RangeDateField) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var err error
	format := "01/02/2006 15:04:05"
	err = e.EncodeToken(xml.StartElement{Name: d.XMLName})
	if err != nil {
		return err
	}

	if !d.Is.IsZero() {
		err = e.EncodeElement(d.Is.Format(format), xml.StartElement{Name: xml.Name{Local: "is"}})
		if err != nil {
			return err
		}
	}
	if !d.Min.IsZero() {
		err = e.EncodeElement(d.Min.Format(format), xml.StartElement{Name: xml.Name{Local: "min"}})
		if err != nil {
			return err
		}
	}
	if !d.Max.IsZero() {
		err = e.EncodeElement(d.Max.Format(format), xml.StartElement{Name: xml.Name{Local: "max"}})
		if err != nil {
			return err
		}
	}

	err = e.EncodeToken(xml.EndElement{Name: d.XMLName})
	return err
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

func (s *SearchQuery) AddRangeDateField(field string) *RangeDateField {
	f := &RangeDateField{XMLName: xml.Name{Local: field}}
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
