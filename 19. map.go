package main

import "fmt"

func main() {
	m := map[string]int{"James": 10, "Penny": 12}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Rashika"])

	v, ok := m["Rashika"]
	fmt.Println(v, ok)

	if v, ok := m["James"]; ok {
		fmt.Println(v, ok)
	}

	m["Rashika"] = 22
	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i,v)
	}

	if v, ok := m["Rashika"]; ok {
		fmt.Println(v)
		delete(m, "Rashika")
		fmt.Println(m)
	}
}

