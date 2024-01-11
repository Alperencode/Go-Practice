package main

import "fmt"

func main() {
	var arr []int = []int{2, 5, 4, 7, 6}
	fmt.Println(sumArr[int](arr))
	fmt.Println(isEmpty(arr))
}

func sumArr[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) <= 0
}
