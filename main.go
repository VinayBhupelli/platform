package main

import (
	"platform/notification"
	"platform/payment"
)

func main() {
	service2Client := notification.NewNotificationInterface()
	service1Client := payment.NewPaymentInterface(service2Client)

	service1Client.MakePayment()
}
