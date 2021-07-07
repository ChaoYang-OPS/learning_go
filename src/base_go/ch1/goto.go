package main

import "fmt"

func main() {

	// 1...100
	result := 0
	i := 1
START:
	if i > 100 {
		goto FOREND
	}
	result += i
	i++
	goto START
FOREND:
	fmt.Println(result)
BREAKEND:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j == 9 {
				break BREAKEND
			}
			fmt.Println(i, j, i*j)
		}
	}
}
