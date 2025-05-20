package service2

import (
	"platform/service2/plugins"
)

type Service2Client interface {
	Method1()
}

func NewService2Client() Service2Client {
	v := "plugin2"
	switch v {
	case "plugin1":
		return plugins.Plugin1{}
	default:
		return plugins.Plugin2{}
	}
}
