package main

import "fmt"

func main() {
	var age = 30

	age = age + 1
	fmt.Println("---", age)
	age = age + 1
	fmt.Println("+++", age)

	fmt.Print("----------")
	fmt.Print("++++++")

	fmt.Printf("\n%T, %s, %T, %d\n", "Lsunstack", "Lsunstack", 1, 1) // string, Lsunstack, int, 1
}
