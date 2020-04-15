package braintree

import "strconv"

type ProcessorSettlementResponseCode int

func (rc ProcessorSettlementResponseCode) Int() int {
	return int(rc)
}

// UnmarshalText fills the response code with the integer value if the text contains one in string form. If the text is zero length, the response code's value is unchanged but unmarshaling is successful.
func (rc *ProcessorSettlementResponseCode) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return nil
	}

	n, err := strconv.Atoi(string(text))
	if err != nil {
		return err
	}

	*rc = ProcessorSettlementResponseCode(n)

	return nil
}

// MarshalText returns a string in bytes of the number, or nil in the case it is zero.
func (rc ProcessorSettlementResponseCode) MarshalText() ([]byte, error) {
	if rc == 0 {
		return nil, nil
	}
	return []byte(strconv.Itoa(int(rc))), nil
}
