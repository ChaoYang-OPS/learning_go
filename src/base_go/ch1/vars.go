package main

import "fmt"

var funOut string

// var funOut string = "Hello World"

func main() {
	// define var type string me
	var me string
	fmt.Println(me)
	// Variable assignment
	me = "Luckly"
	// print me
	fmt.Println(me)

	funOut = "Hello World"
	fmt.Println(funOut)

	var name, user string = "Luckly", "Lsunstack"
	fmt.Println(name, user)

	// var (
	// 	age    int
	// 	height float64
	// )
	// Recommended
	var (
		age    int     = 18
		height float64 = 1.72
	)
	fmt.Println(age, height) // 0 0
	// Recommended
	var (
		username = "Luckly"
		password = 12
	)
	fmt.Println(username, password)
	// Short declarations that can only be used inside a function
	isOk := true
	fmt.Println(isOk)

	isOk = false
	fmt.Println(isOk)
}
