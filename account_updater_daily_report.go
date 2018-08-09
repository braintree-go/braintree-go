package braintree

import "encoding/xml"

type AccountUpdaterDailyReport struct {
	XMLName    xml.Name `xml:"account-updater-daily-report"`
	ReportDate string   `xml:"report-date"`
	ReportURL  string   `xml:"report-url"`
}
