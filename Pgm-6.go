//Write a program that prints a number in decimal, binary, and hex.
//Using the following operators,==, <=, >=, !=, <, >
//write expressions and assign their values to variables:
//Also, create a TYPED and UNTYPED variable.

package main

import "fmt"

const (
	cat int = 23 //typed constant
	dog     = 45 //untyped constant
)

func main() {
	x := 16
	fmt.Printf("%d\t%b\t%x\n", x, x, x)

	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(cat, dog)
}
