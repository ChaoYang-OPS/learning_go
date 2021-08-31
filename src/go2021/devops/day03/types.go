package main

import "fmt"

func main1() {
	// type类型的别名
	// %T 类型  %v数值 %#v展示数据格式，语法展示
	type ops int64
	var nums ops = 12
	fmt.Printf("%T | %#v\n", nums, nums)
	cmd := "TF"
	fmt.Printf("%T, %v %+v %#v\n", cmd, cmd, cmd, cmd)
	// string, TF TF "TF"
	var a rune = '一'
	fmt.Printf("%T, %v %+v %#v\n", a, a, a, a)
	// int32, 19968 19968 19968
	var a1 byte = 'a'
	fmt.Printf("%T, %v %+v %#v\n", a1, a1, a1, a1)
	// uint8, 97 97 0x61
}

// %T 类型， %t bool类型的展示
func main2() {
	var flag bool
	fmt.Printf("%T %t\n", flag, flag) // bool false
	flag = true
	fmt.Printf("%T %t\n", flag, flag) // bool true
}

func main3() {
	number := 123
	// %b 二进制
	fmt.Printf("%T %b\n", number, number)   // int 1111011
	fmt.Printf("%T %c\n", number, number)   // int {
	fmt.Printf("%T %d\n", number, number)   // int 123
	fmt.Printf("%T %8d\n", number, number)  // int      123
	fmt.Printf("%T %08d\n", number, number) // int 00000123
	fmt.Printf("%T %o\n", number, number)   // int 173
	fmt.Printf("%T %q\n", number, number)   // int '{'
	fmt.Printf("%T %x\n", number, number)   // int 7b 小写十六进制
	fmt.Printf("%T %X\n", number, number)   // int 7B 大写十六进制
	fmt.Printf("%T %U\n", number, number)   // int U+007B

}

func main() {
	arr := []byte{'x', 'y', 'z', 'z'}
	fmt.Printf("%s\n", "幸运女神在微笑")
	fmt.Printf("%q\n", "幸运女神在微笑")
	fmt.Printf("%x\n", "幸运女神在微笑")
	fmt.Printf("%X\n", "幸运女神在微笑")
	fmt.Printf("%T\n", "幸运女神在微笑")
	fmt.Printf("%T %s\n", arr, arr)
	fmt.Printf("%T %q\n", arr, arr)
	fmt.Printf("%T %s\n", arr, arr)
	fmt.Printf("%T %x\n", arr, arr)
	fmt.Printf("%T %X\n", arr, arr)
	/*
		幸运女神在微笑
		"幸运女神在微笑"
		e5b9b8e8bf90e5a5b3e7a59ee59ca8e5beaee7ac91
		E5B9B8E8BF90E5A5B3E7A59EE59CA8E5BEAEE7AC91
		string
		[]uint8 xyzz
		[]uint8 "xyzz"
		[]uint8 xyzz
		[]uint8 78797a7a
		[]uint8 78797A7A
	*/
}
