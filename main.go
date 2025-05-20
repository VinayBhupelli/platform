// platform/main.go
package main

import (
	"platform/service1"
	"platform/service2"
)

func main() {
	service2Client := service2.NewService2Client()
	service1Client := service1.NewService1Client(service2Client)

	service1Client.MethodA()
}
