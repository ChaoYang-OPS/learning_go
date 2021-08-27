package main

import (
	"fmt"
	"strconv"
)

func main1() {
	var num int = 10
	var num2 int8 = int8(num)
	var num3 int32 = int32(num2)
	var num4 int64 = int64(num3)
	fmt.Println(num2, num4)

	fmt.Printf("%T\n", num)  // int
	fmt.Printf("%T\n", num2) // int8 转换时，不会修改原来值与类型

	// var num1 int64 = 256
	// var num2 int8 = int8(num1)
	// fmt.Println(num2) // 0
}

// func main() {
// 	var num1 float32 = 32.5
// 	var num2 int = int(num1)
// 	fmt.Println(num2) // 32
// }

// 常规类型转字符串
func main2() {
	var mystr string

	var a int = 100

	var b float32 = 1000

	var c bool = false

	var d byte = 'A'

	mystr = fmt.Sprintf("%d", a)
	fmt.Printf("%T\n", mystr) // string

	mystr = fmt.Sprintf("你有%d个妹子,妹子给了你%f元,%t, 编号: %c", a, b, c, d)

	fmt.Println(mystr) // 你有100个妹子,妹子给了你1000.000000元,false, 编号: A

}

func main() {
	var a int = 100

	var b float32 = 1000

	var c bool = false

	var d byte = 'A'

	fmt.Println(strconv.FormatInt(int64(a), 10))
	fmt.Println(strconv.FormatFloat(float64(b), 'f', 10, 64))
	fmt.Println(strconv.FormatBool(c))
	fmt.Printf("%T\n", strconv.FormatBool(c)) // string
	fmt.Println(strconv.FormatInt(int64(d), 10))

}
