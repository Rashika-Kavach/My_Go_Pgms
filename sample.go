package main

import "fmt"

func main() {
	fmt.Println("First Statement")
	foo()
	for i:=0; i<=100; i++ {
		if i%2==0 {
			fmt.Println(i)
		}
	}
	bar()

}
func foo() {
	fmt.Println("Second Statement")
}

func bar() {
	fmt.Println("Third Statement")
}

//control flow
// 1) Sequencial flow
// 2) Iterative / Loop
// 3) Conditional flow