package main

import "fmt"

// func main() {
// 	var tail int
// 	var heavy int
// 	fmt.Scanf("%d", &tail) // &取变量的地址，进行赋值
// 	fmt.Scanf("%d", &heavy)
// 	fmt.Println(&tail)  // 0xc0000b4008, address
// 	fmt.Println(&heavy) // 0xc0000b4010
// 	fmt.Printf("High:%d, Weight:%d\n", tail, heavy)
// }

func main() {
	var tail int
	var heavy int
	fmt.Scan(&tail)
	fmt.Scan(&heavy)
	fmt.Println(&tail)  // 0xc0000b4008, address
	fmt.Println(&heavy) // 0xc0000b4010
	// High:11, Weight:21 int int
	fmt.Printf("High:%d, Weight:%d %T %T\n", tail, heavy, tail, heavy)
}
