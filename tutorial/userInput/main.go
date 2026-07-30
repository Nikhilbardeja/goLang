package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your Name?")
	// var name string

	// 	fmt.Scan(&name) // only reads till whitespace
	// 	fmt.Println("Name = ", name)

	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')

	// fmt.Println(name)

	fmt.Printf("Value %shas a Type of %T", name, name)
}
