package braintree

const (
	ModificationKindDiscount = "discount"
	ModificationKindAddOn    = "add_on"
)

type Modification struct {
	Id           string  `json:"id,omitempty" xml:"id,omitempty"`
	Amount       float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	Description  string  `json:"description,omitempty" xml:"description,omitempty"`
	Kind         string  `json:"kind,omitempty" xml:"kind,omitempty"`
	Name         string  `json:"name,omitempty" xml:"name,omitempty"`
	NeverExpires bool    `json:"never-expires,omitempty" xml:"never-expires,omitempty"`
	Quantity     int     `json:"quantity,omitempty" xml:"quantity,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

