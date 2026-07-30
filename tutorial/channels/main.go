package main

import (
	"fmt"
	"strings"
)

func add(num1 int, num2 int, ch chan int) {
	var result int = num1 + num2

	ch <- result
	close(ch)
}

func append(str1 string, str2 string, ch chan string) {
	var result string = strings.Join([]string{str1, str2}, "")

	ch <- result
	close(ch)
}

func mainn() {
	var channelInt chan int = make(chan int)

	go add(2, 3, channelInt)

	var result int = <-channelInt

	fmt.Printf("result: %v\n", result)

	var channelStr chan string = make(chan string)

	go append("Nacho ", "re", channelStr)

	var resultStr string = <-channelStr

	fmt.Printf("resultStr: %v\n", resultStr)
}
