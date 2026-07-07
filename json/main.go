package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var p Person = Person{Name: "Nikhil", Age: 20, IsAdult: true}

	// convert to json
	js, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling", err)
	}
	fmt.Println(string(js))

	// now decode

	var person Person // or datatype can also be []map[string]any to convert to map
	errUm := json.Unmarshal(js, &person)
	if errUm != nil {
		fmt.Println("Error Unmarshaling")
	}

	fmt.Println(person)
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}
