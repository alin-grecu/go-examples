package main

import "fmt"

var name string

func init() {
	fmt.Println("This will get called on main initialization")
	name = "Elliot"
}

func main() {
	fmt.Println("My wonderful Go program.")
	fmt.Printf("Name: %s", name)
}
