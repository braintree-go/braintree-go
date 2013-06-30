### Setup

In order to run the integration tests, you have to configure your sandbox account accordingly, or some tests might fail.

#### Environment

Before running the integration tests with `go test`, make sure the following environment variables are set

```
export BRAINTREE_MERCH_ID={your-merchant-id}
export BRAINTREE_PUB_KEY={your-public-key}
export BRAINTREE_PRIV_KEY={your-private-key}
```

#### Sandbox settings

In your sandbox account go to `Settings > Processing > CVV` and enable the following

  1. `CVV does not match (when provided) (N)` to `For Any Transaction`
  2. `CVV is not verified (when provided) (U)` to `For Any Transaction`

We also need to create a plan for recurring payments, with id `test_plan`

