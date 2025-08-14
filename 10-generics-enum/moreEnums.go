package main

import (
	"fmt"
)

func printSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func main() {
// 	printSlice([]int {1, 2, 3})
// 	printSlice([]string {"ARvik", "adb"})
// }


type stack[T any] struct {
	elements []T
}


func main() {
	myStack := stack[string]{
		elements: []string {"golang", "pyt"},
	}
	fmt.Println(myStack)
}