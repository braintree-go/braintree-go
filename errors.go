package braintree

type errorGroup interface {
	For(string) errorGroup
	On(string) []fieldError
}

type braintreeError struct {
	statusCode   int
	XMLName      string         `xml:"api-error-response"`
	Errors       responseErrors `xml:"errors"`
	ErrorMessage string         `xml:"message"`
}

func (e *braintreeError) Error() string {
	return e.ErrorMessage
}

func (e *braintreeError) StatusCode() int {
	return e.statusCode
}

func (e *braintreeError) All() []fieldError {
	baseErrors := e.Errors.TransactionErrors.ErrorList.Errors
	creditCardErrors := e.Errors.TransactionErrors.CreditCardErrors.ErrorList.Errors
	customerErrors := e.Errors.TransactionErrors.CustomerErrors.ErrorList.Errors
	allErrors := append(baseErrors, creditCardErrors...)
	allErrors = append(allErrors, customerErrors...)
	return allErrors
}

func (e *braintreeError) For(item string) errorGroup {
	switch item {
	default:
		return nil
	case "Transaction":
		return e.Errors.TransactionErrors
	}
}

func (e *braintreeError) On(item string) []fieldError {
	return []fieldError{}
}

type responseErrors struct {
	TransactionErrors responseError `xml:"transaction"`
}

type responseError struct {
	ErrorList        errorList  `xml:"errors"`
	CreditCardErrors errorBlock `xml:"credit-card"`
	CustomerErrors   errorBlock `xml:"customer"`
}

func (r responseError) For(item string) errorGroup {
	return nil
}

func (r responseError) On(item string) []fieldError {
	switch item {
	default:
		return []fieldError{}
	case "CreditCard":
		return r.CreditCardErrors.ErrorList.Errors
	}
}

type errorBlock struct {
	ErrorList errorList `xml:"errors"`
}

type errorList struct {
	Errors []fieldError `xml:"error"`
}

type fieldError struct {
	Code      string `xml:"code"`
	Attribute string `xml:"attribute"`
	Message   string `xml:"message"`
}
