package main

import (
	"fmt"
	"log"
	"math/rand"
)

// print * *

// random [0,100]

// fmt.Println(rand.Int() % 100)
// fmt.Println(rand.Intn(100))

func main() {
	var inputNumber int
	fmt.Println("please input you number[0-100]")
	fmt.Scan(&inputNumber)
	fmt.Println("please input you number[0-100]: ", inputNumber)

	if inputNumber > 100 || inputNumber < 0 {
		log.Fatal("only 0~100")
	}

	codeNumber := rand.Intn(100)
}
