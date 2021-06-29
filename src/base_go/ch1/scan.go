package main

import "fmt"

func main() {
	var name string
	fmt.Println("please input name:")
	fmt.Scan(&name)
	fmt.Println("you input name:", name)

	var age int
	fmt.Println("please input age:")
	fmt.Scan(&age)
	fmt.Println(" you age is :", age)

	var height float64
	fmt.Println("please input height:")
	fmt.Scan(&height)
	fmt.Println("you input height is :", height)
}
