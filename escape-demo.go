package main

import (
	"github.com/geremachek/escape"
	"fmt"
)

func main() {
	fmt.Print("Hello" + escape.Reg("1A"))
}
