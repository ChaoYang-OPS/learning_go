package main

import "fmt"

func main() {
	// The integer type
	var age int = 31
	fmt.Printf("%T %d\n", age, age)

	fmt.Println(0777) // 511

	// + - * / %
	fmt.Println(1 + 2)
	fmt.Println(1 - 2)
	fmt.Println(1 * 2)
	fmt.Println(1 / 2)
	fmt.Println(1 % 2)

	age++
	fmt.Println(age)

	age--
	fmt.Println(age)

	// > >= < <=  !=

	fmt.Println(2 == 3)
	fmt.Println(2 != 3)
	fmt.Println(2 >= 3)
	fmt.Println(2 <= 3)

	// Bit operations
	fmt.Println(7 & 2)  // 2
	fmt.Println(7 | 2)  // 7
	fmt.Println(7 ^ 2)  // 5
	fmt.Println(2 << 1) // 4
	fmt.Println(2 >> 1) // 1
	fmt.Println(7 &^ 2) // 5

	//assignment(=,+=,-=,*=,%-,&=,|=,^=,<<=,>>=,&^=)

	age = 1
	age += 3
	fmt.Println(age)

	var intA int = 10
	var uintB uint = 3

	// fmt.Println(intA + uintB) // invalid operation: intA + uintB (mismatched types int and uint)
	fmt.Println(intA + int(uintB)) // 13

	//fmt.Printf
	// byte, rune

	var a byte = 'A'
	var w rune = '中'

	fmt.Println(a, w) // 65 20013
	// %d ^%b %o %x %c
	fmt.Printf("%T\n", a) // uint8
	fmt.Printf("%c\n", a) // A
	fmt.Printf("%T\n", w) // int32
	fmt.Printf("%q\n", w) // '中'
	fmt.Printf("%05d %c\n", a, a)
	fmt.Printf("%-5d %c\n", a, a)

}
