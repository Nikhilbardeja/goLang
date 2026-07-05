package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// var i int = 0

	// for {
	// 	fmt.Println(i)
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	i++
	// }

	// num := []int{00, 100, 200, 300, 400, 500, 600, 700, 800, 900, 100}

	// for i, v := range num {
	// 	fmt.Println(i, v)
	// }

	name := "Raju"
	for i, v := range name {
		fmt.Printf("%d => %c\n", i, v)
	}

}
