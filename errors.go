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
	statusCode   int
	XMLName      string         `json:"api-error-response" xml:"api-error-response"`
	Errors       responseErrors `json:"errors,omitempty" xml:"errors,omitempty"`
	Transaction  transaction    `json:"transaction,omitempty" xml:"transaction,omitempty"`
	ErrorMessage string         `json:"message,omitempty" xml:"message,omitempty"`
}

func (e *braintreeError) Error() string {
	return e.ErrorMessage
}

func (e *braintreeError) StatusCode() int {
	return e.statusCode
}

func (e *braintreeError) ResponseCode() int {
	return e.Transaction.ResponseCode
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
	TransactionErrors responseError `json:"transaction,omitempty" xml:"transaction,omitempty"`
}

type responseError struct {
	ErrorList        errorList  `json:"errors,omitempty" xml:"errors,omitempty"`
	CreditCardErrors errorBlock `json:"credit-card,omitempty" xml:"credit-card,omitempty"`
	CustomerErrors   errorBlock `json:"customer,omitempty" xml:"customer,omitempty"`	
}

type transaction struct {
	ResponseCode     int 	`json:"processor-response-code" xml:"processor-response-code"`
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
	ErrorList errorList `json:"errors,omitempty" xml:"errors,omitempty"`
}

type errorList struct {
	Errors FieldErrorList `json:"error,omitempty" xml:"error,omitempty"`
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
	Code      string `json:"code,omitempty" xml:"code"`
	Attribute string `json:"attribute,omitempty" xml:"attribute"`
	Message   string `json:"message,omitempty" xml:"message"`
}

