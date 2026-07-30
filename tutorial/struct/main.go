package main

import "fmt"

func main() {
	var p1 Person = Person{name: "Nikhil", age: 50} // direct assignmen

	var p2 Person
	p2.name = "Raju"
	p2.age = 25 // by variable accessing assignment

	var p3 *Person = new(Person) // by new keyword it returns a Pointer type
	p3.name = "Golu"
	p3.age = 12

	var c1 Contact
	c1.email = "bardjeanikhil@gmail.com"
	c1.number = 9999999999

	p1.contact = c1

	fmt.Println(p1, p2, p3)

}

type Person struct {
	name    string
	age     int
	contact Contact
}

type Contact struct {
	email  string
	number int
}
