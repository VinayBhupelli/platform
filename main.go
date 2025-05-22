package main

import (
	"platform/interfaces/notification"
	"platform/interfaces/payment"

	"go.uber.org/dig"
)

func main() {
	// service2Client := notification.NewNotificationInterface()
	// service1Client := payment.NewPaymentInterface(service2Client)

	// service1Client.MakePayment()

	container := dig.New()
	container.Provide(notification.NewNotificationInterface)
	container.Provide(payment.NewPaymentInterface)

	container.Invoke(func(payment payment.PaymentInterface) {
		payment.MakePayment()
	})
	// methods will be invoked api call
}
