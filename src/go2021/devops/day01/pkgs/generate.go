package main

import "fmt"

//go:generate go run generate.go
//go:generate go version
func main() {
	fmt.Println("Hello generate.....")
}

// go generate -x generate.go
// go run generate.go
// Hello generate.....
// go version
// go version go1.15.8 darwin/amd64
