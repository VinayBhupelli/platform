package plugins

import "fmt"

type Plugin2 struct{}

func (p Plugin2) Method1() {
	fmt.Println("Plugin2: Method1")
}
