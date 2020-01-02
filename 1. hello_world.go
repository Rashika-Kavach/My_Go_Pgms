package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello World!", 17, true)
	fmt.Println(n)
	fmt.Println(err)
}
