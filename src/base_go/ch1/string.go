package main

import "fmt"

func main() {
	// ""  Can interpret strings
	// ``  native strings
	var name string = "Lsun\tstack"
	var desc string = `OPS\tK8S`
	fmt.Println(name) // Lsun	stack
	fmt.Println(desc) // OPS\tK8S

	fmt.Println("my name is " + "kubernetes")

	// (== != >  >= <=)
	fmt.Println("ab" == "ab")
	fmt.Println("ab" != "ac")

	s := "my name is "
	s += "Lsunstack"
	fmt.Println(s) // my name is Lsunstack

	// The contents of the string definition must be ASCII only
	desc = "abcdef"
	fmt.Println("%T %c\n", desc[0], desc[0])
	fmt.Printf("%T\n", desc[0:2])               // string
	fmt.Printf("%T %s\n", desc[0:2], desc[0:2]) // string ab
	fmt.Println(len(desc))                      // 6
	fmt.Printf("%T\n", len(desc))               // int
}
