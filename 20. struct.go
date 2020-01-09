package main

import "fmt"

type person struct {
	first string
	last string
}

//Embedding Struct

type details struct {
	person
	age int
}

func  main() {
	d1 := details{
		person : person{
			first: "Rashika",
			last: "Rajaraman",},
		age: 22,

	}

	p1 := person{
		first: "Vinaya",
		last: "Ravikumar",

	}
	fmt.Println(d1)
	fmt.Println(p1)

	fmt.Println(p1.first, p1.last)
	fmt.Println(d1.first, d1.last, d1.age)

}

