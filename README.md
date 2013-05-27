# Braintree-Go

A Go client library for [Braintree](https://www.braintreepayments.com).

Note that this is *not* an official client library. Braintree officially only maintains server-side libraries for Ruby, Python, PHP, Perl, C# and Java. This package is not feature-complete compared to those libraries and will not receive updates from Braintree.

With that said, this package contains more than enough to get you started accepting payments using Braintree. If there's a feature the other client libraries implement that you really need, open an issue (or better yet, a pull request).

# Usage:

    import braintree "github.com/lionelbarrow/braintree-go"
  
    config := braintree.Configuration{
      environment: Sandbox,
      merchantId:  "my_merchant_id",
      publicKey:   "my_public_key",
      privateKey:  "my_private_key",
    }
    gateway := braintree.NewGateway(config)

    transaction := braintree.Transaction{
      Type: "sale",
      Amount: 100,
      CreditCard: &braintree.CreditCard{
        Number:         41111111111111111,
        ExpirationDate: "05/14",
      },
    }

    result, err := gateway.Create(transaction)

    if err != nil {
      fmt.Println(err.Error()) 
    } else if !result.Success() {
      fmt.Println(result.Message())
    } else {
      fmt.Println("Transaction created! ID: " + result.Transaction().Id)
    }

# Liscense

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
