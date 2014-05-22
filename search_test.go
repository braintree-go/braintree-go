package braintree

import (
	"encoding/xml"
	"testing"
)

func TestSearchXMLEncode(t *testing.T) {
	s := new(SearchQuery)

	f := s.AddTextField("customer-first-name")
	f.Is = "A"
	f.IsNot = "B"
	f.StartsWith = "C"
	f.EndsWith = "D"
	f.Contains = "E"

	f2 := s.AddRangeField("amount")
	f2.Is = 15.01
	f2.Min = 10.01
	f2.Max = 20.01

	f3 := s.AddMultiField("status")
	f3.Items = []string{
		"authorized",
		"submitted_for_settlement",
	}

	b, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	xmls := string(b)

	expect := `<search>
  <customer-first-name>
    <is>A</is>
    <is-not>B</is-not>
    <starts-with>C</starts-with>
    <ends-with>D</ends-with>
    <contains>E</contains>
  </customer-first-name>
  <amount>
    <is>15.01</is>
    <min>10.01</min>
    <max>20.01</max>
  </amount>
  <status type="array">
    <item>authorized</item>
    <item>submitted_for_settlement</item>
  </status>
</search>`

	if xmls != expect {
		t.Fatal(xmls)
	}
}

func TestSearchResultUnmarshal(t *testing.T) {
	xmls := `<search-results>
  <page-size type="integer">50</page-size>
  <ids type="array">
      <item>k658ww</item>
      <item>fd2h96</item>
  </ids>
</search-results>`

	var v SearchResults
	err := xml.Unmarshal([]byte(xmls), &v)
	if err != nil {
		t.Fatal(err)
	}

	if len(v.Ids.Item) != 2 {
		t.Fatal(v.Ids)
	}
	if x := v.Ids.Item[0]; x != "k658ww" {
		t.Fatal(x)
	}
	if x := v.Ids.Item[1]; x != "fd2h96" {
		t.Fatal(x)
	}
}
