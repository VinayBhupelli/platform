package notification

import (
	"platform/notification/plugins"
)

type NotificationInterface interface {
	SendNotification()
}

func NewNotificationInterface() NotificationInterface {
	v := "plugin2"
	switch v {
	case "plugin1":
		return plugins.NewAmazonSES()
	default:
		return plugins.NewSendGrid()
	}
}
