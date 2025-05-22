package notification

import (
	"platform/plugins/notification"
)

type NotificationInterface interface {
	SendNotification()
}

func NewNotificationInterface() NotificationInterface {
	v := "plugin2"
	switch v {
	case "plugin1":
		return notification.NewAmazonSES()
	default:
		return notification.NewSendGrid()
	}
}
