package main

import "fmt"

func main() {

	// var num = []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// num = append(num, append(num, num...)...)
	// num = append(num, 10)

	// fmt.Println(len(num))
	// fmt.Println(cap(num))
	// fmt.Println(num[1])
	// fmt.Printf("DataType = %T\n", num)
	// fmt.Println(num...)

	// num := make([]int, 2, 5)
	// fmt.Println(num)
	// fmt.Println(len(num))
	// fmt.Println(cap(num))

	num := make([]int, 3, 5)
	num = append(num, 4)
	num = append(num, 5)
	num = append(num, 6) // now it will auto double the capacity old was 5 now 10

	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))

}
