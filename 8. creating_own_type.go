package main

import "fmt"

var a int

//creating type called "mytype"
// which holds integer values as we have declared it to be int
type mytype int

//now b is of type mytype
var b mytype

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//a=b
	//fmt.Println(a)
	//fmt.Printf("%T\n",a)
}
