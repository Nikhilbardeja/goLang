package main

import "fmt"

func main() {
	marks := make(map[string]int)

	marks["raju"] = 40
	marks["golu"] = 50

	fmt.Println(marks)
	fmt.Println(marks["raju"])
	fmt.Println(marks["rajuu"])

	delete(marks, "raju")

	fmt.Println(marks)

	// how to check if key exists in the map

	grade, exists := marks["golu"]
	fmt.Println(grade, exists)

	grade, exists = marks["raju"]
	fmt.Println(grade, exists)

	for k, v := range marks {
		fmt.Println(k, v)
	}

	person := map[string]int{
		"raju": 10,
		"golu": 20,
	}

	fmt.Println(person)
}
