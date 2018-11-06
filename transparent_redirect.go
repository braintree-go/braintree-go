package braintree

type TransparentRedirectKind string

const (
	TransparentRedirectKindCreateCustomer      TransparentRedirectKind = "create_customer"
	TransparentRedirectKindUpdateCustomer      TransparentRedirectKind = "update_customer"
	TransparentRedirectKindCreatePaymentMethod TransparentRedirectKind = "create_payment_method"
	TransparentRedirectKindUpdatePaymentMethod TransparentRedirectKind = "update_payment_method"
	TransparentRedirectKindCreateTransaction   TransparentRedirectKind = "create_transaction"
)

type TransparentRedirectData struct {
	Kind        TransparentRedirectKind `url:"kind"`
	RedirectURL string                  `url:"redirect_url"`
	Transaction TransactionURLRequest   `url:"transaction,omitempty"`
}
