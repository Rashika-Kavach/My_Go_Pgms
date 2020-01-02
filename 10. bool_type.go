package main

import "fmt"


var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 10
	b := 10
	fmt.Println(a == b)
	c := 23
	b = 15
	fmt.Println(c <= b)
}
