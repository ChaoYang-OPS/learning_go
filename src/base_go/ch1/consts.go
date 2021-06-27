package main

import "fmt"

func main() {
	// Constants are not allowed to be modified
	// Constants must set default values
	// Constant full capital
	const NAME string = "Luckly"
	fmt.Println(NAME)

	// NAME = "Luckle" // cannot assign to NAME

	const STA = "UAT"

	const (
		age  = 10
		name = "Luckly"
	)
	const (
		age1  int    = 10
		name1 string = "Luckly"
	)
	fmt.Println(STA, age, name, name1, age1)

	const (
		C1 int = 1
		C2
		C3
	)
	fmt.Println(C1, C2, C3) // 1 1 1

	// Use scene enumeration
	const (
		E1 int = iota
		E2
		E3
	)
	fmt.Println(E1, E2, E3) // 0 1 2
	const (
		E4 int = iota
		E5
		E6
	)
	fmt.Println(E4, E5, E6) // 0 1 2
}
