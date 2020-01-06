package main

import "fmt"

//global initialization

var y = 10

//declares a variable of type int and assigns zero to z which is of type int.
//false for booleans, 0 for numeric types, "" for strings,
//and nil for pointers, functions, interfaces, slices, channels, and maps.

var z int

func main() {
	x := 17
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(y)
	fmt.Println(z)
}
