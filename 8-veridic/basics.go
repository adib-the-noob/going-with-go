package main

import "fmt"

func sums (nums ...int) int {
	sum := 0
	for num := range nums {
		sum += num
	}
	return sum
}

func forAnyFuckingType(args ...interface{}) any {
	for i, v := range args {
		fmt.Println("%v %v", i, v)
	}
}

func main() {
	res := sums(1, 3, 2, 3)
	fmt.Println(res)


	newStuff := forAnyFuckingType(2, 4, 5, "kireMama")
	fmt.Println(newStuff)
}