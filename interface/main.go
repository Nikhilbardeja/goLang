package main

import "fmt"

// import "fmt"

func main() {

}

type Animal interface {
	Eat()
	Run()
}

func Sleep(a Animal) {
	fmt.Println("Animal is sleeping")
}

type Cat struct {
	name string
}

func (i Cat) Eat() {
	fmt.Println("Cat is eating")
}

func (i Cat) Run() {
	fmt.Println("Cat is eating")
	fmt.Println("Cat is running")
}
