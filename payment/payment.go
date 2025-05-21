package payment

import (
	"platform/notification"
	"platform/payment/plugins"
)

type PaymentInterface interface {
	MakePayment()
}

func NewPaymentInterface(service2 notification.NotificationInterface) PaymentInterface {
	v := "pluginB"
	switch v {
	case "pluginB":
		return plugins.Stripe{
			Service2: service2, // inject dependency
		}
	default:
		return plugins.Paypal{
			Service2: service2, // inject dependency
		}
	}
}
