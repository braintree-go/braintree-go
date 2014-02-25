/* API errors are intended to be consumed in two ways. One, they can be dealt with as a single unit:

result, err := gateway.Create(transaction)
err.Error() => "A top level error message"

Second, you can drill down to see specific error messages on a field-by-field basis:

err.For("Transaction").On("Base")[0].Message => "A more specific error message"
*/
package braintree

import "strings"

type errorGroup interface {
	For(string) errorGroup
	On(string) []FieldError
}

type braintreeError struct {
	statusCode      int
	XMLName         string           `xml:"api-error-response"`
	Errors          responseErrors   `xml:"errors"`
	ErrorMessage    string           `xml:"message"`
	MerchantAccount *MerchantAccount `xml:",omitempty"`
}

func (e *braintreeError) Error() string {
	return e.ErrorMessage
}

func (e *braintreeError) StatusCode() int {
	return e.statusCode
}

func (e *braintreeError) All() []FieldError {
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

func (e *braintreeError) On(item string) []FieldError {
	return []FieldError{}
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
	switch item {
	default:
		return nil
	case "Base":
		return r.ErrorList.Errors
	case "Customer":
		return r.CustomerErrors.ErrorList.Errors
	case "CreditCard":
		return r.CreditCardErrors.ErrorList.Errors
	}
}

func (r responseError) On(item string) []FieldError {
	switch item {
	default:
		return []FieldError{}
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
	Errors FieldErrorList `xml:"error"`
}

type FieldErrorList []FieldError

func (f FieldErrorList) For(item string) errorGroup {
	return nil
}

func (f FieldErrorList) On(item string) []FieldError {
	errors := make([]FieldError, 0)
	for _, e := range f {
		if strings.ToLower(item) == e.Attribute {
			errors = append(errors, e)
		}
	}
	return errors
}

type FieldError struct {
	Code      string `xml:"code"`
	Attribute string `xml:"attribute"`
	Message   string `xml:"message"`
}
