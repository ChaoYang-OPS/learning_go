package main

import (
	"fmt"
	"os/exec"
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

func main() {
	var cmd string = "uptime"
	fmt.Println("you will exec command:", cmd)

	exec.Command(cmd).Run()
}
