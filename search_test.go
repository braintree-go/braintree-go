package braintree

import (
	"testing"
)

func TestSearchXMLEncode(t *testing.T) {
	s := new(Search)

	f := s.AddTextField("customer-first-name")
	f.Is = "A"
	f.IsNot = "B"
	f.StartsWith = "C"
	f.EndsWidth = "D"
	f.Contains = "E"

	f2 := s.AddRangeField("amount")
	f2.Is = 15.01
	f2.Min = 10.01
	f2.Max = 20.01

	f3 := s.AddMultiField("status")
	f3.AddItem("authorized")
	f3.AddItem("submitted_for_settlement")

	xml := s.ToXML()
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

	if xml != expect {
		t.Fatal(xml)
	}
}
