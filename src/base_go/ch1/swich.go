package main

import (
	"fmt"
	"log"
)

func main() {
	var yes string

	fmt.Print("Do you have a drink? Y/N")
	fmt.Scan(&yes)

	// switch yes {
	// case "y":
	// 	fmt.Println("buy one")
	// case "Y":
	// 	fmt.Println("buy two")
	// }
	switch yes {
	case "y", "Y":
		fmt.Println("buy one")
	default:
		fmt.Println("buy ten ")
	}

	var score int
	fmt.Print("Please input you score:")
	fmt.Scan(&score)

	if score > 100 || score < 0 {
		log.Fatal("You input score ERROR")
	}

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
