package main

import (
	"fmt"
	"os/exec"
)

// func main() {

// 	cmd := exec.Command("pwd") // command
// 	cmd.Run()                  // run command
// }

func main() {
	// command := exec.Command("go", "version")
	command := exec.Command("date", "+%F")
	commandOut, _ := command.CombinedOutput() // get command output
	fmt.Printf("%s", commandOut)              // print result
}
