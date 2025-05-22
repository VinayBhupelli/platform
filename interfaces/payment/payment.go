package payment

import (
	"platform/interfaces/notification"
	"platform/plugins/payment"
)

type PaymentInterface interface {
	MakePayment()
}

func NewPaymentInterface(service2 notification.NotificationInterface) PaymentInterface {
	v := "pluginB"
	switch v {
	case "pluginB":
		return payment.NewStripe(service2)
	default:
		return payment.NewPaypal(service2)
	}
}
