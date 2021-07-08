package main

import "fmt"

func main() {

	var nums []int
	fmt.Printf("%T\n", nums) // []int

	fmt.Println(nums == nil) // true

	fmt.Printf("%#v , %d, %d\n", nums, len(nums), cap(nums)) // []int(nil) , 0, 0

	fmt.Printf("%#v\n", nums) // []int(nil)

	//
	nums = []int{1, 2, 3}
	fmt.Printf("%#v\n", nums)                                // []int{1, 2, 3}
	fmt.Printf("%#v , %d, %d\n", nums, len(nums), cap(nums)) // []int{1, 2, 3} , 3, 3
	var arrarys [10]int = [10]int{1, 2, 3, 4, 5}

	nums = arrarys[1:10]
	fmt.Printf("%#v , %d, %d\n", nums, len(nums), cap(nums)) // []int{2, 3, 4, 5, 0, 0, 0, 0, 0} , 9, 9

	// make func

	nums = make([]int, 3)
	fmt.Printf("%#v, %d, %d\n", nums, len(nums), cap(nums)) // []int{0, 0, 0}, 3, 3

	nums = make([]int, 3, 5)
	fmt.Printf("%#v, %d, %d\n", nums, len(nums), cap(nums)) // []int{0, 0, 0}, 3, 5

	// add delete modify get

	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])

	nums[0] = 1
	fmt.Println(nums[0])
	nums = append(nums, 1)
	fmt.Printf("%#v, %d, %d\n", nums, len(nums), cap(nums)) // []int{1, 0, 0, 1}, 4, 5
	nums = append(nums, 1)
	nums = append(nums, 1)
	fmt.Printf("%#v, %d, %d\n", nums, len(nums), cap(nums)) // []int{1, 0, 0, 1, 1, 1}, 6, 10

	// for _, v := range nums {
	// 	fmt.Println(v)
	// }

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	fmt.Printf("%T\n", nums[1:5]) // []int

	nums = make([]int, 3, 10)
	n := nums[1:3:10]

	// n cap - start
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n)) // []int []int{0, 0} 2 9

	n = nums[2:3]
	// src cap - start
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n)) // []int []int{0} 1 8

	nums = make([]int, 3, 5)

	nums02 := nums[1:3]
	fmt.Println(nums02)
	fmt.Println(nums)
	nums02[0] = 1
	fmt.Println(nums02, nums) // [1 0] [0 1 0]

	nums02 = append(nums02, 3)

	fmt.Println(nums02, nums) // [1 0 3] [0 1 0]

	nums = append(nums, 5)

	fmt.Println(nums02, nums) // [1 0 5] [0 1 0 5]

	nums = arrarys[:]

	fmt.Println(nums, arrarys) // [1 2 3 4 5 0 0 0 0 0] [1 2 3 4 5 0 0 0 0 0]

	nums[0] = 100

	fmt.Println(nums, arrarys) // [100 2 3 4 5 0 0 0 0 0] [100 2 3 4 5 0 0 0 0 0]

	//copy
	nums04 := []int{1, 2, 3}
	nums05 := []int{10, 20, 30, 40}

	// nums04 to nums05
	copy(nums05, nums04)
	fmt.Println(nums05) // [1 2 3 40]
	// nums05 to nums04
	copy(nums04, nums05)
	fmt.Println(nums04) // [1 2 3]

	// delete
	// delete index is 0 or end

	nums06 := []int{1, 2, 3, 4, 5}
	// delete index is 0
	fmt.Println(nums06[1:]) // [2 3 4 5]
	// delete end index
	fmt.Println(nums06[:len(nums06)-1]) // [1 2 3 4]

	// delete index 2

}
