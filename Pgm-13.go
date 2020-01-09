package main

import "fmt"


type vehicle struct{
	doors int
	color string
}

type truck struct{
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main(){
	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourwheel: true,
	}

	fmt.Println(t1)
	fmt.Println(t1.doors, t1.color, t1.fourwheel)
}
