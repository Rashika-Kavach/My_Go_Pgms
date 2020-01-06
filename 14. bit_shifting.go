package main

import "fmt"

const (
	_ = iota
	s = 1 << (iota * 10) //10 times shifting 1s
	t = 1 << (iota * 10) //20 times shifting 1s
	u = 1 << (iota * 10) //30 times shifting 1s
)

func main() {
	fmt.Printf("%d\t%b\n", s, s)
	fmt.Printf("%d\t%b\n", t, t)
	fmt.Printf("%d\t%b\n", u, u)
}
