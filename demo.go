package main

import (
	"fmt"
	"escape"
)

func main() {
	fmt.Println(escape.Vint(1) + "Hello, World" + escape.Vint(0))
}
