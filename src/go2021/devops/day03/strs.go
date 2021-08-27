package main

import (
	"fmt"
	"os/exec"
	"unsafe"
)

func main1() {
	var ch byte = 'A'

	fmt.Println(ch)              //  65
	fmt.Printf("%c, %d", ch, ch) // A, 65

}

func main2() {
	fmt.Println("Good")
	fmt.Printf("%c", '\n')
	fmt.Print("Work")
}

func main3() {
	var a rune = '我'             // rune用于汉字, byte use A
	fmt.Printf("%T, %v\n", a, a) // int32, 25105
	fmt.Printf("%c", a)          // 我
}

func main4() {
	var cmd string = "date"
	// fmt.Scanf("%s", &cmd) // input cmd
	fmt.Println("you will exec command:", cmd)
	exec.Command(cmd).Run()
}

// func main() {
// 	var cmd string = "uptime"
// 	fmt.Println("you will exec command:", cmd)

// 	exec.Command(cmd).Run()
// }

func main5() {
	var cmd1 = "cale"
	var cmd2 = "哈哈"

	fmt.Printf("%T %T", cmd1, cmd2)
}

func main6() {
	var mystr01 string = "hello"
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr01: %s\n", mystr01)
	fmt.Printf("mystr02: %s", mystr02)
}

func main7() {
	var country string = "hello,中国" // utf-8编写，英文字母和符合占用1个字节,汉字占三个字节
	fmt.Println(len(country))
}

func main8() {
	var ch = 'A'
	var chstr = "A"
	fmt.Printf("%T | %T\n", ch, chstr) // int32 | string
	fmt.Println(unsafe.Sizeof(ch))     // 4 int32
	fmt.Println(unsafe.Sizeof(chstr))  // 16
	fmt.Println(len(chstr))
	fmt.Printf("%c\n", chstr[0]) // A
	// chstr[0] = 'x'  // cannot assign to chstr[0]
	chstr = "XYZ" // 字符串可以修改跳到字符串地址，不可修改字符
	fmt.Println(chstr)
	fmt.Printf("%c %d\n", chstr[0], chstr[0]) // X 88
	fmt.Printf("%c %d\n", chstr[1], chstr[1]) // Y 89
}

func main() {
	var mystr01 string = "\\r\\n"
	var mystr02 string = `\r\n`  // 无转义字符串
	var mystr03 string = `您好啊!`  // 自带换行符
	fmt.Println(mystr01)         // \r\n
	fmt.Println(mystr02)         // \r\n
	fmt.Printf(" %q\n", mystr01) //  "\\r\\n" 原始字符串
	fmt.Println(mystr03)

}
