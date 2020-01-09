package main

import "fmt"

//Store the values of type person in a map with the key of last name.
//Access each value in the map. Print out the values, ranging over the slice.

type person struct {
	first string
	last string
	icecream []string
}
func main() {
	p1 := person{
		first: "Rashika",
		last: "Rajaraman",
		icecream: []string{"Vanilla", "Butterscotch", "Strawberry"},
	}
	fmt.Println(p1)

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.icecream {
		fmt.Println(i, v)
	}

	m := map[string]person{
		p1.last: p1,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for k, v := range p1.icecream {
			fmt.Println(k,v)
		}
		fmt.Println("------------")
	}

}