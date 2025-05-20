package service1

import (
	"platform/service1/plugins"
	"platform/service2"
)

type Service1Client interface {
	MethodA()
}

func NewService1Client(service2 service2.Service2Client) Service1Client {
	v := "pluginB"
	switch v {
	case "pluginB":
		return plugins.PluginA{
			Service2: service2, // inject dependency
		}
	default:
		return plugins.PluginB{
			Service2: service2, // inject dependency
		}
	}
}
