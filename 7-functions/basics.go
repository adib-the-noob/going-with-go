package main

import "fmt"

func main()  {
	fn := proccessIt()
	fmt.Println(fn(10)) // Output: 20
}


func proccessIt() func (a int) int {
	return func(a int) int {
		return a * 2
	}
}