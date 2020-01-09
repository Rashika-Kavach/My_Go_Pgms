package main

import "fmt"

func main() {
	foo1()
	bar1("Rashika")
	s1 := foo2("Rakshitha")
	fmt.Println(s1)
	x, y := bar2("Rashika", "Rajaraman")
	fmt.Println(x,y)

}

//Function with no parameters and return value
func foo1(){
	fmt.Println("This is the foo function")
}

//Function with parameters and no return value
func bar1(s string) {
	fmt.Println(s, "is 22 years old.")
}

//Function with parameters and return value
func foo2(s string) string {
	return fmt.Sprint(s, " is 16 years old.")
}

//Function with multiple parameters of different type and return type
func bar2(s string, t string) (string, bool) {
	a2 := fmt.Sprint(s, "\t", t, ", says hello\n")
	b2 := true
	return a2,b2
}