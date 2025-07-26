package main

import "fmt"

func main() {
	fmt.Println("Bro! Its Go TIME!!!")
	variableStuff()
	fmt.Println("===============")
	variableMore()
}

func variableStuff() {
	var student1 string = "John!"
	var student2 = "kireMama!"

	x := 2
	fmt.Println(x)
	fmt.Println(student1)
	fmt.Println(student2)

}

func variableMore() {
	var a, b, c, d int = 4, 3, 2, 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
