package main

import "fmt"

const (
	a1 = 17
	b1 = 42.18
	c = "Hello"
)

func main() {
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)
	fmt.Println(b1)
	fmt.Printf("%T\n", b1)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
