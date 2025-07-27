package main

import "fmt"

func main() {
	m := make(map[string]string)

	// add elements
	m["name"] = "GoLang"
	m["version"] = "0.0.0.1"

	fmt.Println(m["name"], m["version"])

	fmt.Println("============")
	// new map with int and string
	newData := map[string]int{"price": 40, "age": 23}
	fmt.Println(newData)

	v, ok := newData["prices"]
	fmt.Println(v)
	if ok {
		fmt.Println("All Ok!")
	} else {
		fmt.Println("Not OK!")
	}
}
