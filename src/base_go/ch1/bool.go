package main

import "fmt"

func main() {
	// Boolean type true/false  A value of zero represents false
	var zero bool
	isBoy := true
	isGirl := false

	fmt.Println(zero, isBoy, isGirl) //false true false
	// Logical operations(&&, ||, ! Get it back)
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	fmt.Println("---------------")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	fmt.Println("++++++++++++")
	fmt.Println("Get it back")
	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Println(!isBoy)
	fmt.Println(!isGirl)
	//Relationship operations(==,!=)

	fmt.Println(isBoy == isGirl)
	fmt.Println(isBoy != zero)
	fmt.Printf("%T\n", zero)  // bool
	fmt.Printf("%T\n", isBoy) // bool
}
