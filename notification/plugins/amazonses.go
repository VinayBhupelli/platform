package plugins

import "fmt"

type AmazonSES struct{}

func (p AmazonSES) SendNotification() {
	fmt.Println("AmazonSES: sent Notification")
}
