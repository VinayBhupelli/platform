package plugins

import (
	"platform/notification"
)

type Paypal struct {
	Service2 notification.NotificationInterface
}

func (p Paypal) MakePayment() {
	p.Service2.SendNotification()
}
