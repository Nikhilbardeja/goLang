package main

import "fmt"

func main() {
	fun()
	fmt.Println("Ended")
}

func fun() {
	rec := func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic! Error: %v\n", r)
		}
	}

	defer rec()

	fmt.Println("Executing risky operation...")
	panic("Critical database failure!")
	fmt.Println("This will be skipped.")
}
