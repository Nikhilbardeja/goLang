package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 10
	arr[1] = 12

	fmt.Println(arr)

	var num = [3]int{100, 200}

	fmt.Println(num)

	fmt.Println(len(num))

	var a = []int{12, 24, 25}
	add(a, 5)
}

func add(arr []int, element int) (result []int) {
	arr = append(arr, element)
	return arr
}
