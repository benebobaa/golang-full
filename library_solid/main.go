package main

import "fmt"

func main() {

	fmt.Println("Library with Solid")
}

type User struct {
	ID   int
	Name string
}

type Book struct {
	ID    int
	Title string
	Year  int
}
