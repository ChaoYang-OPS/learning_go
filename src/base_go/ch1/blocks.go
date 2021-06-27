package main

import "fmt"

func main() {
	//scope
	// Defines the range used by the identifier
	// In Go use {} Defines the range used by the identifier
	/*
		Principle of use: A child statement block can use an identifier in the parent statement block,
		and a parent cannot use a child
	*/
	outer := 1
	{
		fmt.Println(outer) // 1
		inner := 2
		// fmt.Print(inner)
		{
			inner2 := 3
			fmt.Println(outer, inner, inner2)
			outer := 2
			fmt.Println(outer, inner, inner2) // 2 2 3
		}
	}
}
