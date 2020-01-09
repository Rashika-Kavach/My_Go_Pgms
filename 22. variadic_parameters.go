package main

import "fmt"

func main() {
	y := add(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("The total is", y)
}

func add(x ...int) int {
	total := 0
	for i, v := range x {
		total += v
		fmt.Printf("In the index position %d the value is %d, when added to the previous number gives %d\n", i, v, total)
	}
	return total
}