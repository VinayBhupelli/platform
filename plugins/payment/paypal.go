package payment

import (
	"platform/interfaces/notification"
)

type Paypal struct {
	Service2 notification.NotificationInterface
}

func NewPaypal(service2 notification.NotificationInterface) Paypal {
	return Paypal{
		Service2: service2,
	}
}
func (p Paypal) MakePayment() {
	p.Service2.SendNotification()
}
