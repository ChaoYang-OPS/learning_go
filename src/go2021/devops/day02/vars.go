package main

import "fmt"

func main1() {
	//
	var n int // default 0

	n = 10
	fmt.Println(".....")
	fmt.Printf("有%d个前女友\n", n) // %d int
	n = 11
	fmt.Printf("有%d个前女友\n", n)
}

func main() {
	// var n int // defalut 0
	// var n = 1000
	n := 1000 // := auto and init
	fmt.Println(n)

	var girlFriend string = "Luckly"
	fmt.Println(girlFriend)

	var lucklyTall = 1.72 // lucklyTall can be change
	// 1.72 not change to left
	fmt.Printf("lucklyTall = %f\n", lucklyTall)

	var names = "Kubernetes"
	fmt.Printf("name=%s\n", names)

	// var a = 100
	// var b = 10
	// fmt.Println(a, b)
	// a, b = b, a
	// fmt.Println(a, b)

	// var N int // N declared but not used
	// a, b, c, _ := 1, 2, 3, 4 // _ cannot use _ as value
	// _, d, _ := 8, 9, 10
	// fmt.Println(a, b, c, d)
	a, b, c, d, _ := returngo()
	fmt.Println(a, b, c, d)
}

func returngo() (int, int, int, int, int) { // return five value
	return 1, 2, 3, 4, 5
}
