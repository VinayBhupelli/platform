package plugins

import "fmt"

type Plugin1 struct{}

func (p Plugin1) Method1() {
	fmt.Println("Plugin1: Method1")
}
