package main

import "fmt"

func main() {
	age := 5
	name := "Nikhil"
	height := 5.905

	fmt.Println(age, name, height) // autp add a space in between each argument

	fmt.Printf("Age = %.2f\n", height) // for formating the output in terminal

	fmt.Printf("Type of name = %T", height) // use %T for the type
}
