package main

import "fmt"

func main() {
	age := 15
	var name = "Alperen"
	var nickname string = "Alperencode"
	var pi float32 = 3.14
	var programmer bool = true
	const language string = "Go"

	length := len(name)

	fmt.Println(name, nickname, age, programmer, language)
	fmt.Println(pi + float32(age))
	fmt.Println(length)
}
