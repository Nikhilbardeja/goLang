package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 32
	fmt.Printf("num's Type : %T\n", num)

	var data float64 = float64(num)
	fmt.Printf("data's Type : %T\n", data)

	var st string = strconv.Itoa(num)
	fmt.Printf("st: %T\n", st)

	var str string = "123456"
	var n, _ = strconv.Atoi(str)
	fmt.Printf("n: %T\n", n)

	var num1 int = int(data)
	fmt.Printf("num1: %T\n", num1)

	fl_str := "3.14"

	fl, _ := strconv.ParseFloat(fl_str, 64)

	fmt.Printf("fl: %T\n", fl)
}
