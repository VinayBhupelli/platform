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
		return plugins.NewStripe(service2)
	default:
		return plugins.NewPaypal(service2)
	}
}
