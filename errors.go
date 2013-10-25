/* API errors are intended to be consumed in two ways. One, they can be dealt with as a single unit:

result, err := gateway.Create(transaction)
err.Error() => "A top level error message"

Second, you can drill down to see specific error messages on a field-by-field basis:

err.For("Transaction").On("Base")[0].Message => "A more specific error message"
*/
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
	case "Base":
		return r.ErrorList.Errors
	case "Customer":
		return r.CustomerErrors.ErrorList.Errors
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
