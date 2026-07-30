package main

import "fmt"

func main() {
	day := 1

	switch day {
	case 0, 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("UnKnown Day")
	}

	//  other syntax

	switch {
	case day == 0 || day == 1:
		fmt.Println("Monday")
	case day == 2:
		fmt.Println("Tuesday")
	case day == 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("UnKnown Day")
	}
}
