package braintree

type CVVResponseCode string

const (
	CVVMatches      CVVResponseCode = "M" // The CVV provided matches the information on file with the cardholder's bank.
	CVVDoesNotMatch CVVResponseCode = "N" // The CVV provided does not match the information on file with the cardholder's bank.

	// The card-issuing bank received the CVV but did not verify whether it was correct.
	// This typically happens if the processor declines an authorization before the bank evaluates the CVV.
	CVVNotVerified CVVResponseCode = "U"

	CVVNotProvided              CVVResponseCode = "I" // No CVV was provided.
	CVVIssuerDoesNotParticipate CVVResponseCode = "S" // The CVV was provided but the card-issuing bank does not participate in card verification.
	CVVNotApplicable            CVVResponseCode = "A" // The CVV was provided but this type of transaction does not support card verification.
)
