package testhelpers

import "regexp"

var ValidBIN = regexp.MustCompile(`^\d{6}$`).MatchString
var ValidLast4 = regexp.MustCompile(`^\d{4}$`).MatchString
var ValidExpiryMonth = regexp.MustCompile(`^\d{2}$`).MatchString
var ValidExpiryYear = regexp.MustCompile(`^\d{4}$`).MatchString
