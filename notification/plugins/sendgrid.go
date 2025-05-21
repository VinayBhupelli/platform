package plugins

import "fmt"

type SendGrid struct{}

func (p SendGrid) SendNotification() {
	fmt.Println("SendGrid: Sent Notification")
}
