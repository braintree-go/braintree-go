package braintree

import "time"

const (
	ModificationKindDiscount = "discount"
	ModificationKindAddOn    = "add_on"
)

type Modification struct {
	Id           string     `xml:"id,omitempty" json:"id,omitempty" bson:"id,omitempty"`
	Amount       *Decimal   `xml:"amount,omitempty" json:"amount,omitempty" bson:"amount,omitempty"`
	Description  string     `xml:"description,omitempty" json:"description,omitempty" bson:"description,omitempty"`
	Kind         string     `xml:"kind,omitempty" json:"kind,omitempty" bson:"kind,omitempty"`
	Name         string     `xml:"name,omitempty" json:"name,omitempty" bson:"name,omitempty"`
	NeverExpires bool       `xml:"never-expires,omitempty" json:"neverExpires,omitempty" bson:"neverExpires,omitempty"`
	Quantity     int        `xml:"quantity,omitempty" json:"quantity,omitempty" bson:"quantity,omitempty"`
	UpdatedAt    *time.Time `xml:"updated_at,omitempty" json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
