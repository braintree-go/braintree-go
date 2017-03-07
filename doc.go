/*
Package braintree is a client library for Braintree.

API errors are intended to be consumed in two ways. One, they can be dealt with as a single unit:

    result, err := gateway.Create(transaction)
    err.Error() => "A top level error message"

Second, you can drill down to see specific error messages on a field-by-field basis:

    err.For("Transaction").On("Base")[0].Message => "A more specific error message"
*/
package braintree
