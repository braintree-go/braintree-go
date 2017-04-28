package braintree

type GatewayRejectionReason string

const (
	GatewayRejectionReasonApplicationIncomplete GatewayRejectionReason = "application_incomplete"
	GatewayRejectionReasonAVS                                          = "avs"
	GatewayRejectionReasonAVSAndCVV                                    = "avs_and_cvv"
	GatewayRejectionReasonCVV                                          = "cvv"
	GatewayRejectionReasonDuplicate                                    = "duplicate"
	GatewayRejectionReasonFraud                                        = "fraud"
	GatewayRejectionReasonThreeDSecure                                 = "three_d_secure"
	GatewayRejectionReasonUnrecognized                                 = "unrecognized"
)
