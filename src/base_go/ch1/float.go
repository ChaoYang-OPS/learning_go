package main

import "fmt"

func main() {
	// float32, float64
	var height float64
	fmt.Println(height)                   // 0
	fmt.Printf("%T %f\n", height, height) // float64, 0.000000

	var weight float64 = 13.05e1
	fmt.Println(weight)    // 130.5
	fmt.Println(1.1 + 1.2) // 2.3
	fmt.Println(1.1 - 0.5) // 0.6
	fmt.Println(1.1 * 0.5) // 0.55
	fmt.Println(1.1 / 0.5) // 2.2

	//(> >= < <= ~=)
	fmt.Println(1.1 > 1.2)
	fmt.Println(1.1 >= 1.2)
	fmt.Println(1.1 < 1.2)
	fmt.Println(1.1 <= 1.2)

	//(=, +=,-=,/=,*=)
	height += 0.1
	fmt.Println(height)
	fmt.Printf("%T %T\n", 1.1, 1.2) // float64 float64
	newHeight := 3.1415978202012313
	fmt.Printf("%5.3f\n", newHeight) // 3.142
	fmt.Printf("%v\n", newHeight)    // 3.1415978202012314
}
