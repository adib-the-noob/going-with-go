package main

import "fmt"

const PI = 3.13

func main() {
	fmt.Println("Bro! Its Go TIME!!!")

	variableStuff()
	fmt.Println("===============")
	variableMore()
	fmt.Println("===============")
	fmt.Println(PI)
	fmt.Println("===============")

	fmt.Println("f String in GO!")

	name := "adib-the-noob"
	age := 20

	fmt.Printf("Hello, there %v this side, a %v years old man! Good Luck!\n", name, age)
	fmt.Printf("The Type is %T\n", name)
	fmt.Printf("The another type is %T\n", age)
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
