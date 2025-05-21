package plugins

import "fmt"

type AmazonSES struct{}

func NewAmazonSES() AmazonSES {
	return AmazonSES{}
}

func (p AmazonSES) SendNotification() {
	fmt.Println("AmazonSES: sent Notification")
}
