package main

import "fmt"

func main() {
	var A = 2
	var B = A
	B = 3
	fmt.Println(A, B)

	// pointer
	C := &A
	// var C1 *int = &A
	fmt.Printf("%T\n", C) // *int

	fmt.Println(*C) // 2

	*C = 100

	fmt.Println(*C)                   // 100
	fmt.Printf("%T %v\n", C, C)       // *int 0xc000014058
	fmt.Printf("%T %v %p\n", C, C, C) // *int 0xc00018c008 0xc00018c008

}
