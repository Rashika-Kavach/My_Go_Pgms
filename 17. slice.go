package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8}

	//Slicing a slice.
	fmt.Println(x[1:])
	fmt.Println(x[0:3])

	//Both the for loops perform the same operation.
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

	for i, v := range x {
		fmt.Println(i , v)
	}

	//Append to the slice x
	y := []int{12, 13, 14, 15, 16}
	x = append(x, y...)
	fmt.Println(x)

	//Append to the slice y
	y = append(y, 78, 79, 80, 81)
	fmt.Println(y)

	//Deleting from a Slice
	x = append(x[:1], x[3:]...)
	fmt.Println(x)
}
