package main

import "fmt"

var z string //prints empty string
var y int    //prints zero
var x bool   //prints false

func main()  {
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
