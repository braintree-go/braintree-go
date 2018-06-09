package braintree

import (
	"encoding/xml"
	"time"
)

type DisputeEvidence struct {
	XMLName           string     `xml:"evidence"`
	Comment           string     `xml:"comment"`
	CreatedAt         *time.Time `xml:"created-at"`
	ID                string     `xml:"id"`
	SentToProcessorAt string     `xml:"sent-to-processor-at"`
	URL               string     `xml:"url"`
}

type DisputeEvidenceCategory string

const (
	EvidenceCategoryDeviceId                                   DisputeEvidenceCategory = "DEVICE_ID"
	EvidenceCategoryDeviceName                                 DisputeEvidenceCategory = "DEVICE_NAME"
	EvidenceCategoryPriorDigitalGoodsTransactionArn            DisputeEvidenceCategory = "PRIOR_DIGITAL_GOODS_TRANSACTION_ARN"
	EvidenceCategoryPriorDigitalGoodsTransactionDateTime       DisputeEvidenceCategory = "PRIOR_DIGITAL_GOODS_TRANSACTION_DATE_TIME"
	EvidenceCategoryDownloadDateTime                           DisputeEvidenceCategory = "DOWNLOAD_DATE_TIME"
	EvidenceCategoryGeographicalLocation                       DisputeEvidenceCategory = "GEOGRAPHICAL_LOCATION"
	EvidenceCategoryLegitPaymentsForSameMerchandise            DisputeEvidenceCategory = "LEGIT_PAYMENTS_FOR_SAME_MERCHANDISE"
	EvidenceCategoryMerchantWebsiteOrAppAccess                 DisputeEvidenceCategory = "MERCHANT_WEBSITE_OR_APP_ACCESS"
	EvidenceCategoryPriorNonDisputedTransactionArn             DisputeEvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_ARN"
	EvidenceCategoryPriorNonDisputedTransactionDateTime        DisputeEvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_DATE_TIME"
	EvidenceCategoryPriorNonDisputedTransactionEmailAddress    DisputeEvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_EMAIL_ADDRESS"
	EvidenceCategoryPriorNonDisputedTransactionIpAddress       DisputeEvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_IP_ADDRESS"
	EvidenceCategoryPriorNonDisputedTransactionPhoneNumber     DisputeEvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_PHONE_NUMBER"
	EvidenceCategoryPriorNonDisputedTransactionPhysicalAddress DisputeEvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_PHYSICAL_ADDRESS"
	EvidenceCategoryProfileSetupOrAppAccess                    DisputeEvidenceCategory = "PROFILE_SETUP_OR_APP_ACCESS"
	EvidenceCategoryProofOfAuthorizedSigner                    DisputeEvidenceCategory = "PROOF_OF_AUTHORIZED_SIGNER"
	EvidenceCategoryProofOfDeliveryEmpAddress                  DisputeEvidenceCategory = "PROOF_OF_DELIVERY_EMP_ADDRESS"
	EvidenceCategoryProofOfDelivery                            DisputeEvidenceCategory = "PROOF_OF_DELIVERY"
	EvidenceCategoryProofOfPossessionOrUsage                   DisputeEvidenceCategory = "PROOF_OF_POSSESSION_OR_USAGE"
	EvidenceCategoryPurchaserEmailAddress                      DisputeEvidenceCategory = "PURCHASER_EMAIL_ADDRESS"
	EvidenceCategoryPurchaserIpAddress                         DisputeEvidenceCategory = "PURCHASER_IP_ADDRESS"
	EvidenceCategoryPurchaserName                              DisputeEvidenceCategory = "PURCHASER_NAME"
	EvidenceCategoryRecurringTransactionArn                    DisputeEvidenceCategory = "RECURRING_TRANSACTION_ARN"
	EvidenceCategoryRecurringTransactionDateTime               DisputeEvidenceCategory = "RECURRING_TRANSACTION_DATE_TIME"
	EvidenceCategorySignedDeliveryForm                         DisputeEvidenceCategory = "SIGNED_DELIVERY_FORM"
	EvidenceCategorySignedOrderForm                            DisputeEvidenceCategory = "SIGNED_ORDER_FORM"
	EvidenceCategoryTicketProof                                DisputeEvidenceCategory = "TICKET_PROOF"
)

type DisputeFileEvidenceRequest struct {
	XMLName    xml.Name                `xml:"evidence"`
	DocumentId string                  `xml:"document-upload-id"`
	Category   DisputeEvidenceCategory `xml:"category,omitempty"`
}

type DisputeTextEvidenceRequest struct {
	XMLName  xml.Name                `xml:"evidence"`
	Content  string                  `xml:"comments"`
	Category DisputeEvidenceCategory `xml:"category,omitempty"`
}
