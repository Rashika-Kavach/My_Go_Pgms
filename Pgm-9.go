//Create an ARRAY which holds 5 VALUES of TYPE int
//Assign VALUES to each index position.
//Range over the array and print the values out.
//Using format printing print out the TYPE of the array
package main

import "fmt"

func main() {
	myArray := []int{1,2,3,4,5}
	fmt.Println(myArray)
	for i, v := range myArray {
		fmt.Println(i,v)
	}
	fmt.Printf("%T\n", myArray)

	//Illustrate Slicing.

	mySlice := []int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(mySlice[:5])
	fmt.Println(mySlice[5:])
	fmt.Println(mySlice[2:7])
	fmt.Println(mySlice[1:6])

	mySlice = append(mySlice, 52)
	fmt.Println(mySlice)
	mySlice = append(mySlice, 53, 54, 55)
	fmt.Println(mySlice)
	mySlice1 := []int{56,57,58,59,60}
	fmt.Println(mySlice1)
	mySlice = append(mySlice, mySlice1...)
	fmt.Println(mySlice)

	//Illustrate delete in Slice
	mySlice2 := []int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(mySlice2)
	mySlice3 := append(mySlice2[:3], mySlice2[6:]...)
	fmt.Println(mySlice3)


}

