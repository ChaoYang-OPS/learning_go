package main

import "fmt"

func main() {
	// index ==> record

	result := 0
	for i := 1; i <= 100; i++ {
		result += i
	}
	fmt.Println("result=", result)

	result = 0
	i := 1
	for i <= 100 {
		result += i
		i++
	}
	fmt.Println("new result=", result)
	//Dead loop
	// i = 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	desc := "阳之幸福"
	for i, ch := range desc {
		// fmt.Println(i, ch)
		fmt.Printf("%d %T %q \n", i, ch, ch)
	}
}
