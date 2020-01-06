package main

import "fmt"

func main() {
	s := "Hello World!"
	fmt.Println(s)
	s = "Hello India"
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i, v := range s {
		fmt.Printf("At index position %d, we have %#U\n", i, v)
	}
}
