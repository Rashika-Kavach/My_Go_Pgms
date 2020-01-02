//Using the short declaration operator, ASSIGN these VALUES to VARIABLES
//with the IDENTIFIERS “x” and “y” and “z”
//42, “James Bond”, and true

//Now print the values stored in those variables using,
//a single print statement and multiple print statements


package main
import "fmt"

func main() {
	x:= 42
	y:= "James Bond"
	z:= true

	//prints in a single statement
	fmt.Println("Printing in a Single Statement:", x,y,z)

	//prints in multiple lines
	fmt.Println("Printing in a Multiple Statement:")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

