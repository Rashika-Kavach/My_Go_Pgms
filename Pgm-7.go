//Write a program that assigns an int to a variable
//Prints that int in decimal, binary, and hex
//Shifts the bits of that int over 1 position to the left, and assigns that to a variable
//Prints that variable in decimal, binary, and hex

package main

import "fmt"


func main() {
	l := 25
	fmt.Printf("%b\t%x\t%d\n", l, l, l)
	o := l << 1
	fmt.Printf("%b\t%x\t%d\n", o, o, o)
}

