package main

import "fmt"

func main() {
	//Makes Slice of length 10 (Index 0 to 9) with capacity 100
	x := make([]int, 10, 11)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//Appends 13 to the slice at the index position 10
	x = append(x, 13)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//Appends 17 to the slice at the index position 11 and
	//doubles the array capacity as the added element occupies the position 11
	x = append(x, 17)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//Multi-dimensional Slice
	n := []string{"My", "name", "is", "Rashika"}
	fmt.Println(n)

	m := []string{"My", "friend's", "is", "Vinaya"}
	fmt.Println(m)

	o := [][]string{n,m}
	fmt.Println(o)
}
