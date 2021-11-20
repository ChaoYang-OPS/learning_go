package main

import "fmt"

func main() {
	data := []string{"I", "am", "stupid", "weak"}
	fmt.Println(data)
	// 索引指向的是内存地址
	for k := range data {
		if k == 2 {
			data[k] = "smart"
		}
		if k == len(data)-1 {
			data[k] = "strong"
		}
	}
	fmt.Println(data)
	mySlice := []int{10, 20, 30, 40, 50}
	for _, value := range mySlice {
		value *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice) // mySlice [10 20 30 40 50]
	for index := range mySlice {
		mySlice[index] *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice) // mySlice [20 40 60 80 100]
}
