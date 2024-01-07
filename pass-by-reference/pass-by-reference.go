package main

import "fmt"

func printArray(arr *[5]uint32) {
	fmt.Printf("Addr of array %p\n", arr)
	fmt.Println(*arr)
}

func main() {
	var arr [5]uint32 = [5]uint32{10, 20, 30, 40, 50}
	fmt.Printf("Addr of local array %p\n", &arr)
	printArray(&arr)
}
