package braintree

import (
	"encoding/xml"
	"testing"
	"time"
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

	startDate := time.Date(2016, time.September, 11, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2016, time.September, 11, 23, 59, 59, 0, time.UTC)
	f3 := s.AddRangeDateField("settled-at")
	f3.Min = startDate
	f3.Max = endDate

	f4 := s.AddRangeDateField("created-at")
	f4.Min = startDate

	f5 := s.AddRangeDateField("authorization-expired-at")
	f5.Min = startDate

	f6 := s.AddRangeDateField("authorized-at")
	f6.Min = startDate

	f7 := s.AddRangeDateField("failed-at")
	f7.Min = startDate

	f8 := s.AddRangeDateField("gateway-rejected-at")
	f8.Min = startDate

	f9 := s.AddRangeDateField("processor-declined-at")
	f9.Min = startDate

	f10 := s.AddRangeDateField("submitted-for-settlement-at")
	f10.Min = startDate

	f11 := s.AddRangeDateField("voided-at")
	f11.Min = startDate

	f12 := s.AddRangeDateField("disbursement-date")
	f12.Min = startDate

	f13 := s.AddRangeDateField("dispute-date")
	f13.Min = startDate

	f14 := s.AddMultiField("status")
	f14.Items = []string{
		"authorized",
		"submitted_for_settlement",
		"settled",
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
  <settled-at>
    <min>09/11/2016 00:00:00</min>
    <max>09/11/2016 23:59:59</max>
  </settled-at>
  <created-at>
    <min>09/11/2016 00:00:00</min>
  </created-at>
  <authorization-expired-at>
    <min>09/11/2016 00:00:00</min>
  </authorization-expired-at>
  <authorized-at>
    <min>09/11/2016 00:00:00</min>
  </authorized-at>
  <failed-at>
    <min>09/11/2016 00:00:00</min>
  </failed-at>
  <gateway-rejected-at>
    <min>09/11/2016 00:00:00</min>
  </gateway-rejected-at>
  <processor-declined-at>
    <min>09/11/2016 00:00:00</min>
  </processor-declined-at>
  <submitted-for-settlement-at>
    <min>09/11/2016 00:00:00</min>
  </submitted-for-settlement-at>
  <voided-at>
    <min>09/11/2016 00:00:00</min>
  </voided-at>
  <disbursement-date>
    <min>09/11/2016 00:00:00</min>
  </disbursement-date>
  <dispute-date>
    <min>09/11/2016 00:00:00</min>
  </dispute-date>
  <status type="array">
    <item>authorized</item>
    <item>submitted_for_settlement</item>
    <item>settled</item>
  </status>
</search>`

	if xmls != expect {
		t.Fatal(expect)
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
