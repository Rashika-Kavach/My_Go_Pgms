package main

import "fmt"

func main() {
	x := 0
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue //skips rest of the for loop statements.
		}
		fmt.Println(x)
		//x++
	}
}
