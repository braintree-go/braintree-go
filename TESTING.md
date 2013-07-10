### Setup

In order to run the integration tests, you have to configure your sandbox account accordingly, or some tests might fail.

#### Environment

Before running the integration tests with `go test`, make sure the following environment variables are set

```
export BRAINTREE_MERCH_ID={your-merchant-id}
export BRAINTREE_PUB_KEY={your-public-key}
export BRAINTREE_PRIV_KEY={your-private-key}
```

When using Braintree Go in a production environment, we recommend that you continue to store these credentials in environment variables. See [the 12 Factor App](http://www.12factor.net/config) for details.

#### Sandbox settings

In your sandbox account go to `Settings > Processing > CVV` and enable the following

  1. `CVV does not match (when provided) (N)` to `For Any Transaction`
  2. `CVV is not verified (when provided) (U)` to `For Any Transaction`

Finally you also need to create a plan for recurring payments with id `test_plan`. Once you do all of these things, the integration tests should all pass.

**Test Plan 1 Setup**

```
Plan ID:                test_plan
Plan Name:              test_plan_name
Description:            test_plan_desc
Price:                  10
Currency:               USD

Include Trial Period:   YES
Duration:               14 days

Billing Cycle:          Every 1 Month
First Bill Date:        Specific Day - Last Day of the Month
End Date:               After 2 billing cycles
```

**Test Plan 2 Setup**

```
Plan ID:                test_plan_2
Plan Name:              test_plan_2_name
Price:                  20
Currency:               USD

Billing Cycle:          Every 1 Month
First Bill Date:        Immediately
End Date:               Never
```
