package braintree

import (
	"encoding/xml"
)

type ModificationRequest struct {
	Amount                *Decimal `xml:"amount,omitempty"`
	NumberOfBillingCycles int      `xml:"number-of-billing-cycles,omitempty"`
	Quantity              int      `xml:"quantity,omitempty"`
	NeverExpires          bool     `xml:"never-expires,omitempty"`
}

type AddModificationRequest struct {
	ModificationRequest
	InheritedFromID string `xml:"inherited-from-id,omitempty"`
}

type UpdateModificationRequest struct {
	ModificationRequest
	ExistingID string `xml:"existing-id,omitempty"`
}

type ModificationsRequest struct {
	Add               []AddModificationRequest
	Update            []UpdateModificationRequest
	RemoveExistingIDs []string
}

func (m ModificationsRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if len(m.Add) > 0 {
		if err := m.marshalXMLModifications(e, "add", m.Add); err != nil {
			return err
		}
	}
	if len(m.Update) > 0 {
		if err := m.marshalXMLModifications(e, "update", m.Update); err != nil {
			return err
		}
	}
	if len(m.RemoveExistingIDs) > 0 {
		if err := m.marshalXMLModifications(e, "remove", m.RemoveExistingIDs); err != nil {
			return err
		}
	}
	e.EncodeToken(start.End())
	return nil
}

func (m ModificationsRequest) marshalXMLModifications(e *xml.Encoder, name string, modifications interface{}) error {
	attr := []xml.Attr{{Name: xml.Name{Local: "type"}, Value: "array"}}
	start := xml.StartElement{Name: xml.Name{Local: name}, Attr: attr}

	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := e.EncodeElement(modifications, xml.StartElement{Name: xml.Name{Local: "modification"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}
