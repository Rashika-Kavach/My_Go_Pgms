//At the package level scope, using the “var” keyword,
//create a VARIABLE with the IDENTIFIER “y”.
//The variable should be of the UNDERLYING TYPE of your custom TYPE “x”

//In func main print out the value of the variable “x”
//Print out the type of the variable “x”
//Assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
//Print out the value of the variable “x”

//Use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
//then use the “=” operator to ASSIGN that value to “y”
//Print out the value stored in “y”
//Print out the type of “y”

package main

import "fmt"

type meow int
var x meow
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 17
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}