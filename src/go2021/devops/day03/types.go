package main

import "fmt"

func main() {
	// type类型的别名

	type ops int64

	var nums ops = 12
	fmt.Printf("%T | %#v\n", nums, nums)
}
