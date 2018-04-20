package braintree

import (
	"time"
)

// StatusHistory contains information about the last 50 timestamps where
// something changed about the subscription.
type StatusHistory struct {
	StatusEvents []StatusEvent `xml:"status-event"`
}

// StatusEvent contains information about what and when something
// changed about the subscription.
type StatusEvent struct {
	Timestamp          time.Time `xml:"timestamp"`
	Balance            *Decimal  `xml:"balance"`
	Price              *Decimal  `xml:"price"`
	Status             string    `xml:"status"`
	SubscriptionSource string    `xml:"subscription-source"`
}
