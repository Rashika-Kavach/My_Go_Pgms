package main

import "fmt"

const m = iota //initializes to zero
const n = 17
const o = 13

const (
	d = iota //initializes to zero
	e = iota //doesn't initialize just increments
	f
)

func main() {
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
