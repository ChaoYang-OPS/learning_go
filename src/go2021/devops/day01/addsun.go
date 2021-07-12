package main

import (
	"fmt"
	"os"
)

// Usage
var Usage = func() {
	fmt.Println("USEAGE: calcommnd [arguments]")
	fmt.Println("THe command is add ")
}

func main() {
	args := os.Args

	if args == nil || len(args) < 3 {
		Usage()
		return
	}
}
