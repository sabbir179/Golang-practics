package main

import "fmt"

func main() {

	// primitive data types
	// rune, byte = alias
	// int, float32, string, bool

	// Composite Data Types Maps,
	// array,

	// var students [3]string
	// rules of  array
	// fixed lenght ,
	//same type

	// data set or data assign
	// students[0] = "Salam"
	// students[1] = "Scaymam"
	// students[2] = "Aalam"

	//fmt.Println(students)
	//fmt.Println(len(students))

	//data pull, data retrieve, data get
	//fmt.Println(students[2])

	// short hand way, string literals
	//In computer science, a literal is a notation for representing a fixed value in source code
	// implicit = ... in []
	students := [...]string{"Asmal", "Dhaka", "Bangladesh", "Canada"}
	fmt.Println(students)

}
