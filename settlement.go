package braintree

type Record struct {
	XMLName           string   `xml:"record" json:"record" bson:"record"`
	CardType          string   `xml:"card-type" json:"card-type" bson:"card-type"`
	Count             int      `xml:"count" json:"count" bson:"count"`
	MerchantAccountId string   `xml:"merchant-account_id" json:"merchantAccountId" bson:"merchantAccountId"`
	Kind              string   `xml:"kind" json:"kind" bson:"kind"`
	AmountSettled     *Decimal `xml:"amount-settled" json:"amount-settled" bson:"amount-settled"`
}

type XMLRecords struct {
	XMLName string   `xml:"records" json:"records" bson:"records"`
	Type    []Record `xml:"record" json:"record" bson:"record"`
}
type SettlementBatchSummary struct {
	XMLName string     `xml:"settlement-batch-summary" json:"settlementBatchSummary" bson:"settlementBatchSummary"`
	Records XMLRecords `xml:"records" json:"records" bson:"records"`
}

type Settlement struct {
	XMLName string `xml:"settlement_batch_summary" json:"settlementBatchSummary" bson:"settlementBatchSummary"`
	Date    string `xml:"settlement_date" json:"settlementDate" bson:"settlementDate"`
}
