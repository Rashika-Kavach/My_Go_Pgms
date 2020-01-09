package main

import "fmt"

type person1 struct{
	first string
	last string
}

type details1 struct {
	person1
	age int
}

func main() {
	dt1 := details1 {
		person1: person1 {
			first: "Rashika",
			last: "Rajaraman",
		},
		age: 22,
	}

	dt2 := details1 {
		person1 : person1 {
			first: "Vinaya",
			last: "Ravikumar",
		},
		age: 22,
	}
	dt1.speak()
	dt2.speak()
}

func (s details1) speak() {
	fmt.Println(s)
}