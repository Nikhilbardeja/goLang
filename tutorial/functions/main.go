package main

import "fmt"

func main() {
	fmt.Println(add(8, 5))
}

func add(a int, b int) int {
	return a + b
}

// in param list if you add data type at the end
// all the param will be declared with same datatype
func add1(a, b int) int {
	return a + b
}

// need to define datatype of each return value
func add2(a, b int) (int, int) {
	return a + b, a - b
}

// result variable is already defined at return datatype definition
// and just a single return statment will automatically return the result variable
// the return variable's name should not match with param list's names
func add3(a, b int) (result int) {
	result = a + b
	return
}
