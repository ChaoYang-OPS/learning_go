package main

import "fmt"

func main() {
	var nums [10]int
	var t2 [5]bool
	var t3 [3]string
	fmt.Printf("%T\n", nums) // [10]int
	fmt.Println(len(nums))   // 10
	fmt.Println(nums)        // [0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("%T\n", t2)
	fmt.Println(t2)        // [false false false false false]
	fmt.Printf("%q\n", t3) // ["" "" ""]
	nums = [10]int{10, 20, 30}
	fmt.Println(nums) // [10 20 30 0 0 0 0 0 0 0]
	nums = [10]int{0: 10, 9: 20}
	fmt.Println(nums) // [10 0 0 0 0 0 0 0 0 20]
	newNums := [...]int{1, 2, 3}
	fmt.Println(newNums) // [1 2 3]

	nums02 := [5]int{10, 20, 30}
	fmt.Printf("%T, %v\n", nums02, nums02)
	nums03 := [...]int{1, 2}
	fmt.Printf("%T, %v\n", nums03, nums03) // [2]int, [1 2]
	nums04 := [15]int{1: 10, 5: 20, 14: 30}
	fmt.Printf("%T, %v\n", nums04, nums04)

	// op
	nums05 := [3]int{1, 3, 4}
	nums06 := [...]int{2, 3, 4}
	fmt.Println(nums05 == nums06) // false
	fmt.Println(nums05 != nums06)

	// get arrays len
	fmt.Println(len(nums04))

	// index  0, 1, 2, len(array) - 1

	fmt.Println(nums04[0], nums04[1])

	nums04[0] = 1000
	fmt.Println(nums04)
	//

	// for i := 0; i < len(nums04); i++ {
	// 	fmt.Println(nums04[i])
	// }

	// for _, value := range nums04 {
	// 	fmt.Println(value)
	// }

	// slice
	s := "1235"
	fmt.Printf("%T\n", s[1:3])        //string
	fmt.Printf("%T\n", nums04[0:3:3]) // []int

	var marrays [3][2]int

	fmt.Println(marrays) // ]int[[0 0] [0 0] [0 0]]

}
