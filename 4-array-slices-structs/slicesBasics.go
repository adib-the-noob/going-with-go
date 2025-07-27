package main

import "fmt"

func main() {
	var nums = make([]int, 2, 4) // 2 -> initial length, 5 > initial capacity

	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(cap(nums))

	fmt.Println("========")

	nums = append(nums, 1)
	nums = append(nums, 5)
	nums = append(nums, 6)
	nums = append(nums, 9)

	// capacity will be double

	fmt.Println(cap(nums))

	fmt.Println("==========")

	var newNumber = make([]int, 0, 5)
	newNumber = append(newNumber, 4)

	var newNumber2 = make([]int, len(newNumber))
	copy(newNumber, newNumber2)

	fmt.Println(newNumber2, newNumber)
}
