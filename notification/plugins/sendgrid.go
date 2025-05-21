package plugins

import "fmt"

type SendGrid struct{}

func NewSendGrid() SendGrid {
	return SendGrid{}
}

func (p SendGrid) SendNotification() {
	fmt.Println("SendGrid: Sent Notification")
}
