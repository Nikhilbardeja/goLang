package main

import "fmt"

func main() {
	var num int = 2

	var ptr *int = &num

	fmt.Println(ptr, *ptr)

	modify(ptr)

	fmt.Println(num)

}

func modify(ad *int) {
	*ad += 100
}
