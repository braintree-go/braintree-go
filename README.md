# Braintree Go

A Go client library for [Braintree](https://www.braintreepayments.com), the payments company behind awesome companies like GitHub, Heroku, and 37signals.

This is *not* an official client library. Braintree maintains server-side libraries for [Ruby](https://www.github.com/braintree/braintree_ruby), [Python](https://www.github.com/braintree/braintree_python), [PHP](https://www.github.com/braintree/braintree_php), [Perl](https://www.github.com/braintree/braintree_perl), [Node](https://www.github.com/braintree/braintree_node), [C#](https://www.github.com/braintree/braintree_dotnet) and [Java](https://www.github.com/braintree/braintree_java), but not Go. This package implements the core functionality of the other client libraries, but it's missing a few advanced features.

With that said, this package contains more than enough to get you started accepting payments using Braintree. If there's a feature the other client libraries implement that you really need, open an issue (or better yet, a pull request).

### Usage

Setting up your credentials is easy.

```go
import braintree "github.com/lionelbarrow/braintree-go"
  
config := braintree.Configuration{
  environment: braintree.Sandbox,
  merchantId:  "my_merchant_id",
  publicKey:   "my_public_key",
  privateKey:  "my_private_key",
}

gateway := braintree.NewGateway(config)
```

So is creating your first transaction.

```go
transaction := braintree.Transaction{
  Type: "sale",
  Amount: 100,
  CreditCard: &braintree.CreditCard{
    Number:         41111111111111111,
    ExpirationDate: "05/14",
  },
}

result, err := gateway.Transaction().Create(transaction)
```

The create call returns an error when something mechanical goes wrong, such as receiving malformed XML or being unable to connect to the Braintree gateway. For semantic failures, such as your customer's credit card being expired, the result type has `Success()` and `Message()` methods.

```go
if err != nil {
  fmt.Println(err.Error()) 
} else if !result.Success() {
  fmt.Println(result.Message())
} else {
  fmt.Println("Transaction created! ID: " + result.Transaction().Id)
}
```

In addition to creating transactions, you can also tokenize credit card information for repeat or subscription billing using the `CreditCard` and `Customer` types. This package is completely compatible with [Braintree.js](https://www.braintreepayments.com/braintrust/braintree-js), so if you encrypt your customers' credit cards in the browser, you can pass them on to Braintree without ever seeing them yourself. This massively decreases your PCI scope.

### Installation

The usual. `go get github.com/lionelbarrow/braintree-go`

### Documentation

Braintree provides a [ton of documentation](https://www.braintreepayments.com/docs/ruby/guide/overview) on how to use their API. I recommend you use the Ruby documentation when following along, as the Ruby client library is broadly similar to this one.

### Testing

The integration tests run against an account I created in the [Braintree Sandbox](https://sandbox.braintreegateway.com/) for this package. This account has some non-default settings, so if you just plug your own sandbox account in, a few tests will break. I've noted what settings you need to change and why in comments above these tests.

### Liscense

The MIT License (MIT)

Copyright (c) 2013 Lionel Barrow

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
