package main

import "fmt"

func main() {
	for x := 0; x <= 100; x++ { //for + if + modulus
		if x%3 == 0 {
			fmt.Println(x)
		}
	}
	x := 11
	if x >= 25 { //if / else if / else
		fmt.Println(x)
	} else if x == 11 {
		fmt.Println(x)
	} else {
		fmt.Println(x)
	}
}
