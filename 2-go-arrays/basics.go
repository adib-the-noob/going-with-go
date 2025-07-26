package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 3, 4}
	fmt.Println(arr1)

	fmt.Println("===========")
	var cars = [4]string{
		"Tesla",
		"Volvo",
		"Audi",
		"Ford",
	}
	fmt.Println(cars)
	fmt.Println(len(cars))
}
