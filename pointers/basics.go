package main

import "fmt"

// func changeNum (num int) {
// 	num = 5
// 	fmt.Println("In ChangeNum", num)
// }

func changeNumByRef(num *int) {
	*num = 5
	fmt.Println(num)
}

func main() {
	num := 1
	changeNumByRef(&num)
	fmt.Println("Memory Address", &num)
	fmt.Println("After ChangeNum in Main", num)
}