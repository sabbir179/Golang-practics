package main

import "fmt"

// make structure

// type Book struct {
// 	Title  string
// 	Author string
// 	ISBN   string
// 	Price  float32
// 	Pages  int
// }

func main() {

	// array
	// slice
	// map
	// struct
	// user defined data types or composite data type or
	// a struct is a collection of fields or properties.

	// assign value

	// var b1 Book
	// b1.Title = "An Introduction to programming in GO"
	// b1.Author = "CALEB DOXY"
	// b1.Price = 10.50
	// b1.Pages = 165

	// struct literals : In computer science, a literal is a notation for representing a fixed value in source code.
	// example: var name string = "Sabbir"

	// Annonymous stuct

	b1 := struct {
		Title  string
		Author string
		ISBN   string
		Price  float32
		Pages  int
	}{
		Title:  "An Introduction to programming in GO",
		Author: "CALEB DOXY",
		ISBN:   "987-1478355823",
		Price:  10.50,
		Pages:  165,
	}

	fmt.Println(b1)
	fmt.Println(b1.Title)
	fmt.Println(b1.Author)
	fmt.Println(b1.Price)
	fmt.Println(b1.Pages)
}
