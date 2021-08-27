package main

import "fmt"

func main() {
	var complexValue1 complex64
	complexValue1 = 1.10 + 10i
	complexValue2 := 1.10 + 10i
	complexValue3 := complex(1.10, 10)
	// (1.1+10i)
	fmt.Println(complexValue1)
	fmt.Println(complexValue2)
	fmt.Println(complexValue3)
	fmt.Printf("%T\n", complexValue1) // complex64
	fmt.Printf("%T\n", complexValue2) // complex128

	var isok bool

	fmt.Println(isok) // false
	// if !isok {
	// 	fmt.Println(".....")
	// }
}
