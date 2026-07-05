package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is grater than 5")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is smaller than 5")
	}
}
