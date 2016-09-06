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

type BraintreeError struct {
	statusCode      int
	XMLName         string           `xml:"api-error-response" json:"apiErrorResponse" bson:"apiErrorResponse"`
	Errors          responseErrors   `xml:"errors" json:"errors" bson:"errors"`
	ErrorMessage    string           `xml:"message" json:"message" bson:"message"`
	MerchantAccount *MerchantAccount `xml:",omitempty" json:",omitempty" bson:",omitempty"`
	Transaction     Transaction      `xml:"transaction" json:"transaction" bson:"transaction"`
}

func (e *BraintreeError) Error() string {
	return e.ErrorMessage
}

func (e *BraintreeError) StatusCode() int {
	return e.statusCode
}

func (e *BraintreeError) All() []FieldError {
	baseErrors := e.Errors.TransactionErrors.ErrorList.Errors
	creditCardErrors := e.Errors.TransactionErrors.CreditCardErrors.ErrorList.Errors
	customerErrors := e.Errors.TransactionErrors.CustomerErrors.ErrorList.Errors
	allErrors := append(baseErrors, creditCardErrors...)
	allErrors = append(allErrors, customerErrors...)
	return allErrors
}

func (e *BraintreeError) For(item string) errorGroup {
	switch item {
	default:
		return nil
	case "Transaction":
		return e.Errors.TransactionErrors
	}
}

func (e *BraintreeError) On(item string) []FieldError {
	return []FieldError{}
}

type responseErrors struct {
	TransactionErrors responseError `xml:"transaction" json:"transaction" bson:"transaction"`
}

type responseError struct {
	ErrorList        errorList  `xml:"errors" json:"errors" bson:"errors"`
	CreditCardErrors errorBlock `xml:"credit-card" json:"credit-card" bson:"credit-card"`
	CustomerErrors   errorBlock `xml:"customer" json:"customer" bson:"customer"`
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
	ErrorList errorList `xml:"errors" json:"errors" bson:"errors"`
}

type errorList struct {
	Errors FieldErrorList `xml:"error" json:"error" bson:"error"`
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
	Code      string `xml:"code" json:"code" bson:"code"`
	Attribute string `xml:"attribute" json:"attribute" bson:"attribute"`
	Message   string `xml:"message" json:"message" bson:"message"`
}
