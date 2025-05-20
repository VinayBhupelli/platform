package plugins

import (
	"platform/service2"
)

type PluginA struct {
	Service2 service2.Service2Client
}

func (p PluginA) MethodA() {
	p.Service2.Method1()
}
