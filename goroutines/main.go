package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("1")
	go sayHello()
	sayHi()
	time.Sleep(2000 * time.Millisecond)
}

func sayHello() {
	fmt.Println("2")
	// time.Sleep(2000 * time.Millisecond)
	fmt.Println("3")
}

func sayHi() {
	fmt.Println("4")
}
