package main

import "fmt"

var y = 17

func main() {
	//Format Printing
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)  //binary value of y
	fmt.Printf("%x\n", y)  //hexadecimal value of y
	fmt.Printf("%#x\n", y) // hexadecimal value prefixed with 0x

	y = 911
	fmt.Printf("%b\t%x\t%#x\n", y, y, y) //prints all the values of y in one line

	//stores all the values of y in a string
	s := fmt.Sprintf("%b\t%x\t%#x\n", y, y, y)
	fmt.Println(s)
}
