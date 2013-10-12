package braintree

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

type responseErrors struct {
	TransactionErrors responseError `xml:"transaction"`
}

type responseError struct {
	ErrorList errorList `xml:"errors"`
  CreditCardErrors errorBlock `xml:"credit-card"`
  CustomerErrors errorBlock `xml:"customer"`
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
