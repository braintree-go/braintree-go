package braintree

import (
	"time"

	"github.com/lionelbarrow/braintree-go/nullable"
)

const (
	ModificationKindDiscount = "discount"
	ModificationKindAddOn    = "add_on"
)

type Modification struct {
	Id           string     `xml:"id,omitempty"`
	Amount       *Decimal   `xml:"amount,omitempty"`
	Description  string     `xml:"description,omitempty"`
	Kind         string     `xml:"kind,omitempty"`
	Name         string     `xml:"name,omitempty"`
	NeverExpires bool       `xml:"never-expires,omitempty"`
	Quantity     int        `xml:"quantity,omitempty"`
	UpdatedAt    *time.Time `xml:"updated_at,omitempty"`
}

type ModificationRequest struct {
	Amount                *Decimal            `xml:"amount,omitempty"`
	NumberOfBillingCycles *nullable.NullInt64 `xml:"number-of-billing-cycles,omitempty"`
	Quantity              int                 `xml:"quantity,omitempty"`
	NeverExpires          bool                `xml:"never-expires,omitempty"`
}

type AddModificationRequest struct {
	ModificationRequest
	InheritedFromId string `xml:"inheritedFromId,omitempty"`
}

type UpdateModificationRequest struct {
	ModificationRequest
	ExistingId string `xml:"existingId,omitempty"`
}

type ModificationsRequest struct {
	Add struct {
		Type string                   `xml:"type,attr"`
		Add  []AddModificationRequest `xml:"modification,omitempty"`
	} `xml:"add"`
	Update struct {
		Type   string                      `xml:"type,attr"`
		Update []UpdateModificationRequest `xml:"modification,omitempty"`
	} `xml:"update"`
	Remove struct {
		Type   string `xml:"type,attr"`
		Remove []string
	} `xml:"remove"`
}
