package payment

import (
	"platform/interfaces/notification"
)

type Stripe struct {
	Service2 notification.NotificationInterface
}

func NewStripe(service2 notification.NotificationInterface) Stripe {
	return Stripe{
		Service2: service2,
	}
}
func (p Stripe) MakePayment() {
	p.Service2.SendNotification()
}
