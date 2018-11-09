package braintree

type ProcessorResponseType string

const (
	ProcessorResponseTypeApproved     ProcessorResponseType = "approved"
	ProcessorResponseTypeSoftDeclined ProcessorResponseType = "soft_declined"
	ProcessorResponseTypeHardDeclined ProcessorResponseType = "hard_declined"
)
