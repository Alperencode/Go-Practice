package main

import (
	"fmt"
)

func main() {
	var nums [3]uint32 = [3]uint32{5, 10, 15}

	fmt.Println(nums[0])
	fmt.Println(nums[1:3])

	// if you not specify the length
	// it will be called "slice"
	var slice []int32 = []int32{2, 4, 6}
	slice = append(slice, 8)
	fmt.Println(slice)

	var slice2 []int32 = []int32{10, 12, 14}
	slice = append(slice, slice2...)
	fmt.Println(slice2)

	// use make() if you want to initialize it with capacity
}
