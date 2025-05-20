package plugins

import (
	"platform/service2"
)

type PluginB struct {
	Service2 service2.Service2Client
}

func (p PluginB) MethodA() {
	p.Service2.Method1()
}
