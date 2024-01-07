package main

import (
	"errors"
	"fmt"
)

func main() {
	name := "Alperen"
	greet(name)

	num1, num2 := 5, 10
	var sum int = add(num1, num2)
	fmt.Println(sum)

	num2 = 0
	division, remainder, err := divide(num1, num2)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Division: %0.2f and remainder: %v\n", division, remainder)
	}

}

func greet(name string) {
	fmt.Println("Hello, " + name)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func divide(num1 int, num2 int) (float32, int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("Can't divide by zero\n")
		return 0, 0, err
	}
	return float32(num1) / float32(num2), num1 % num2, err
}
