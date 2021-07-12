package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main1() {
	fmt.Println(os.Args[0]) // 第一个参数，是程序的执行路径
	// go run getcmd_line.go  go version  // 运用命令行
	cmd := exec.Command(os.Args[1], os.Args[2])
	out, _ := cmd.CombinedOutput() // 获取命令输出,_处理异常错误
	fmt.Printf("%s", out)
}

// add return a + b
func Add(a int, b int) int {
	return a + b
}

func main() {

	// 显示命令行
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])

	}
	// result := Add(1, 2)
	// log.Println(result)
}
