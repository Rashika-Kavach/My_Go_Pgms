package main

import "fmt"

const (
	a = 17
	b = 42.18
	c = "Hello"
)
func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
