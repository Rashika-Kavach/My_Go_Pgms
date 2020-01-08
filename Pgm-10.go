package main

import "fmt"

func main() {

	//Creating a slice of a specific capacity
	//and then filling values in it.
	x := make([]string, 50,500)
	for i := 0; i < len(x); i++ {
		x[i] = fmt.Sprint(i)
	}
	fmt.Println(x)


	//Creating a slice of a slice on int
	mySlice := []int{1,2,3,4,5}
	fmt.Println(mySlice)
	mySlice1 := []int{6,7,8,9,0}
	fmt.Println(mySlice1)
	mySlice2 := [][]int{mySlice, mySlice1}
	fmt.Println(mySlice2)
}