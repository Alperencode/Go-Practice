package main

import "fmt"

type Iphone struct {
	model string
}

func (iphone Iphone) call() {
	fmt.Println("Making a facetime call...")
}

type Android struct {
	model string
}

func (android Android) call() {
	fmt.Println("Making a whatsapp video call...")
}

type Phone interface {
	call()
}

func makeCall(phone Phone) {
	phone.call()
}

func main() {
	var iphone Iphone = Iphone{"Xs"}
	var android Android = Android{"Samsung"}

	makeCall(iphone)
	makeCall(android)
}
