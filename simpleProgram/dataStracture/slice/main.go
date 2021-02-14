package main

import "fmt"

// slice
func main() {
	var students [3]string

	students[0] = "John"
	students[1] = "Jamal"
	students[2] = "Samad"

	x := students[0:2]

	fmt.Println(x)

}
