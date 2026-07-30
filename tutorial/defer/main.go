package main

import "fmt"

func main() {
	// defer fmt.Println("Starting the Program")
	// defer fmt.Println("Mid of The Program")
	// fmt.Println("End of The Program")

	fmt.Println(run())
}

func run() int { // will return 10
	num := 10

	defer func() { num += 5 }()
	return num
}

func run1() (num int) { // will return 15
	num = 10

	defer func() { num += 5 }()
	return num
}
