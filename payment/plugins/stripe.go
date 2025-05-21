package plugins

import (
	"platform/notification"
)

type Stripe struct {
	Service2 notification.NotificationInterface
}

func (p Stripe) MakePayment() {
	p.Service2.SendNotification()
}
