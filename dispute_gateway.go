package braintree

import (
	"context"
	"encoding/xml"
)

type EvidenceCategory string

const (
	EvidenceCategoryDeviceId                                   EvidenceCategory = "DEVICE_ID"
	EvidenceCategoryDeviceName                                 EvidenceCategory = "DEVICE_NAME"
	EvidenceCategoryPriorDigitalGoodsTransactionArn            EvidenceCategory = "PRIOR_DIGITAL_GOODS_TRANSACTION_ARN"
	EvidenceCategoryPriorDigitalGoodsTransactionDateTime       EvidenceCategory = "PRIOR_DIGITAL_GOODS_TRANSACTION_DATE_TIME"
	EvidenceCategoryDownloadDateTime                           EvidenceCategory = "DOWNLOAD_DATE_TIME"
	EvidenceCategoryGeographicalLocation                       EvidenceCategory = "GEOGRAPHICAL_LOCATION"
	EvidenceCategoryLegitPaymentsForSameMerchandise            EvidenceCategory = "LEGIT_PAYMENTS_FOR_SAME_MERCHANDISE"
	EvidenceCategoryMerchantWebsiteOrAppAccess                 EvidenceCategory = "MERCHANT_WEBSITE_OR_APP_ACCESS"
	EvidenceCategoryPriorNonDisputedTransactionArn             EvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_ARN"
	EvidenceCategoryPriorNonDisputedTransactionDateTime        EvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_DATE_TIME"
	EvidenceCategoryPriorNonDisputedTransactionEmailAddress    EvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_EMAIL_ADDRESS"
	EvidenceCategoryPriorNonDisputedTransactionIpAddress       EvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_IP_ADDRESS"
	EvidenceCategoryPriorNonDisputedTransactionPhoneNumber     EvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_PHONE_NUMBER"
	EvidenceCategoryPriorNonDisputedTransactionPhysicalAddress EvidenceCategory = "PRIOR_NON_DISPUTED_TRANSACTION_PHYSICAL_ADDRESS"
	EvidenceCategoryProfileSetupOrAppAccess                    EvidenceCategory = "PROFILE_SETUP_OR_APP_ACCESS"
	EvidenceCategoryProofOfAuthorizedSigner                    EvidenceCategory = "PROOF_OF_AUTHORIZED_SIGNER"
	EvidenceCategoryProofOfDeliveryEmpAddress                  EvidenceCategory = "PROOF_OF_DELIVERY_EMP_ADDRESS"
	EvidenceCategoryProofOfDelivery                            EvidenceCategory = "PROOF_OF_DELIVERY"
	EvidenceCategoryProofOfPossessionOrUsage                   EvidenceCategory = "PROOF_OF_POSSESSION_OR_USAGE"
	EvidenceCategoryPurchaserEmailAddress                      EvidenceCategory = "PURCHASER_EMAIL_ADDRESS"
	EvidenceCategoryPurchaserIpAddress                         EvidenceCategory = "PURCHASER_IP_ADDRESS"
	EvidenceCategoryPurchaserName                              EvidenceCategory = "PURCHASER_NAME"
	EvidenceCategoryRecurringTransactionArn                    EvidenceCategory = "RECURRING_TRANSACTION_ARN"
	EvidenceCategoryRecurringTransactionDateTime               EvidenceCategory = "RECURRING_TRANSACTION_DATE_TIME"
	EvidenceCategorySignedDeliveryForm                         EvidenceCategory = "SIGNED_DELIVERY_FORM"
	EvidenceCategorySignedOrderForm                            EvidenceCategory = "SIGNED_ORDER_FORM"
	EvidenceCategoryTicketProof                                EvidenceCategory = "TICKET_PROOF"
)

type DisputeGateway struct {
	*Braintree
}

type DisputeFileEvidenceRequest struct {
	XMLName    xml.Name         `xml:"evidence"`
	DocumentId string           `xml:"document-id"`
	Category   EvidenceCategory `xml:"category,omitempty"`
}

type DisputeTextEvidenceRequest struct {
	XMLName  xml.Name         `xml:"evidence"`
	Content  string           `xml:"comments"`
	Category EvidenceCategory `xml:"category,omitempty"`
}

func (g *DisputeGateway) AddFileEvidence(ctx context.Context, disputeId string, fileEvidenceRequest *DisputeFileEvidenceRequest) (*DisputeEvidence, error) {
	resp, err := g.executeVersion(ctx, "POST", "disputes/"+disputeId+"/evidence", fileEvidenceRequest, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.disputeEvidence()
	}
	return nil, &invalidResponseError{resp}
}

func (g *DisputeGateway) AddTextEvidence(ctx context.Context, disputeId string, textEvidenceRequest *DisputeTextEvidenceRequest) (*DisputeEvidence, error) {
	resp, err := g.executeVersion(ctx, "POST", "disputes/"+disputeId+"/evidence", textEvidenceRequest, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.disputeEvidence()
	}
	return nil, &invalidResponseError{resp}
}

func (g *DisputeGateway) Find(ctx context.Context, disputeId string) (*Dispute, error) {
	resp, err := g.executeVersion(ctx, "GET", "disputes/"+disputeId, nil, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.dispute()
	}
	return nil, &invalidResponseError{resp}
}

func (g *DisputeGateway) Accept(ctx context.Context, disputeId string) error {
	resp, err := g.executeVersion(ctx, "PUT", "disputes/"+disputeId+"/accept", nil, apiVersion4)
	if err != nil {
		return nil
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}

func (g *DisputeGateway) Finalize(ctx context.Context, disputeId string) (*Dispute, error) {
	resp, err := g.executeVersion(ctx, "PUT", "disputes/"+disputeId+"/finalize", nil, apiVersion4)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.dispute()
	}
	return nil, &invalidResponseError{resp}
}

func (g *DisputeGateway) RemoveEvidence(ctx context.Context, disputeId string, evidenceId string) error {
	resp, err := g.executeVersion(ctx, "DELETE", "disputes/"+disputeId+"/evidence/"+evidenceId, nil, apiVersion4)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}

func (g *DisputeGateway) Search(ctx context.Context, query *SearchQuery) ([]*Dispute, error) {
	resp, err := g.executeVersion(ctx, "POST", "disputes/advanced_search", query, apiVersion4)
	if err != nil {
		return nil, err
	}
	var v struct {
		XMLName  string     `xml:"disputes"`
		Disputes []*Dispute `xml:"dispute"`
	}
	err = xml.Unmarshal(resp.Body, &v)
	if err != nil {
		return nil, err
	}
	return v.Disputes, err
}
