package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("date") // 定义
	// bytes, err := cmd.Output()   // 执行
	// fmt.Println(string(bytes), err)
	output, err := cmd.StdoutPipe()
	cmd.Start()
	fmt.Println(err)
	io.Copy(os.Stdout, output)
	cmd.Wait()
}
