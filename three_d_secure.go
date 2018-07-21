package braintree

type ThreeDSecureInfo struct {
	Status                 string `xml:"status"`
	Enrolled               string `xml:"enrolled"`
	LiabilityShiftPossible bool   `xml:"liability-shift-possible"`
	LiabilityShifted       bool   `xml:"liability-shifted"`
}
