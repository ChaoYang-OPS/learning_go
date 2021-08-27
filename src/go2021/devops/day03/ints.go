package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var data1 int = 10   // 十进制
	var data2 int = 0b10 // 2 进制
	var data3 int = 0o10 // 8 进制
	var data4 int = 0x10 // 16 进制

	fmt.Println(data1)                // 10
	fmt.Println(data2)                // 2
	fmt.Println(data3)                // 8
	fmt.Println(data4)                // 16
	fmt.Printf("二进制===>%b\n", data1)  // 1010
	fmt.Printf("八进制===>%o\n", data3)  // 10
	fmt.Printf("十六进制===>%x\n", data1) // a

	var myfloat float64 = 1.000001
	var myfloat1 float32 = 1.000001
	fmt.Println(unsafe.Sizeof(myfloat))  // 8
	fmt.Println(unsafe.Sizeof(myfloat1)) // 4

	// var tall float64 // remind float64
	// fmt.Scanf("%f", &tall)
	var tall = 8
	var tall1 = 8.8

	fmt.Printf("%T---%T\n", tall, tall1) // int---float64

}
