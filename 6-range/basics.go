package main

import "fmt"

func main() {
	nums := []int{6, 7, 5}
	sum := 0

	for i, num := range nums {
		sum += num
		fmt.Printf("%v %v\n", i, num)
	}
	fmt.Println(sum)
	fmt.Println("=======")

	data := map[string]string{"name": "adib", "role": "developer!"}
	for k, v := range data {
		fmt.Println(k, v)
	}
}
