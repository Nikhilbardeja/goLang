package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orrage,banna"

	var sl []string = strings.Split(data, ",")

	fmt.Println(sl)

	fmt.Println(strings.Count(data, ","))

	space := "     data    "
	fmt.Println(strings.TrimSpace(space))

	s1 := "Nacho"
	s2 := "Re!"

	result := strings.Join([]string{s1, s2}, " ")

	fmt.Printf("result: %v\n", result)
}
