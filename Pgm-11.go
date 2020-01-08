//Create a map with a key of TYPE string which is a person’s “last_first” name,
//and a value of TYPE []string which stores their favorite things.
//Store three records in your map. Print out all of the values,
//along with their index position in the slice.

package main

import "fmt"

func main() {
	x := map[string][]string{"Rashika" : []string{"Rajaraman", "Ravindiri"},
		                     "Vinaya" : []string{"Ravikumar", "Mahitha"}}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i,v)
		for k, v1 := range v {
			fmt.Println(k, v1)
		}

	}
}
