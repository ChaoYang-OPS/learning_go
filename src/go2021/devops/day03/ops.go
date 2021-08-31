package main

import "fmt"

func main1() {
	a := 10
	b := 5
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
}
func main2() {
	fmt.Println("note" + "pad") // 字符串加法
}
func main3() {
	a := 3
	a++            // 自增1 a = a+1
	fmt.Println(a) // 4
	b := 2
	b--            // 自减1  n = b -1
	fmt.Println(b) // 1
}

func main() {
	var chinese int = 88
	var math int = 99
	var english int = 91
	var sum int
	var avg float32
	sum = chinese + math + english
	avg = float32(sum) / 3
	fmt.Println(sum, avg)
}
