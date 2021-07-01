package main

import "fmt"

func main() {
	fmt.Println("Lsunstack")

	fmt.Println(" please input Y or N")
	var yes string
	//Accept the value entered by the console
	fmt.Scan(&yes)

	if yes == "Y" || yes == "y" {
		fmt.Println("buy ")
	}

	fmt.Println("..............")
	if yes == "Y" || yes == "y" {
		fmt.Println(".....1.....")
	} else {
		fmt.Println("......10.......")
	}

	// var score int
	// fmt.Println("please input you score:")

}
