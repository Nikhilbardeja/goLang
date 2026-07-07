package main

import "fmt"

func mainCat() {
	var c Cat = Cat{name: "Kali Billi"}

	c.Run()

	fmt.Println(c.state)
	c.Stop()
	fmt.Println(c.state)
}

type Cat struct {
	name  string
	state string
}

func (cat *Cat) Run() {
	fmt.Println("Cat is running")
	cat.state = "Running"
}

func (cat *Cat) Stop() {
	fmt.Println("Cat stopped")
	cat.state = "Stopped"
}
