package main

import (
	"fmt"
)

func main() {
	var mapExample map[string]int = map[string]int{
		"Alperen": 21, "Goksu": 20, "Efe": 20, "Danyal": 21,
	}
	fmt.Println(mapExample["Alperen"])

	var age, exists = mapExample["Efe"]
	if exists {
		fmt.Println(age)
	}

	delete(mapExample, "Goksu")

	for name, age := range mapExample {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
}
