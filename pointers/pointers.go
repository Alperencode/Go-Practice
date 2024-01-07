package main

import "fmt"

func main() {
	var p *uint32 = new(uint32)
	fmt.Printf("Addr of p: %v, value of p: %v\n", p, *p)

	var i uint32 = 15
	p = &i
	fmt.Printf("Addr of p: %v, value of p: %v\n", p, *p)

	i = 50
	fmt.Printf("Addr of p: %v, value of p: %v\n", p, *p)
}
