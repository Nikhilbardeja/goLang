package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lines int
var words int
var chars int

func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error Opening the file", err)
	}
	defer file.Close()
	var scanner *bufio.Scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		lines++
		data := scanner.Text()

		words += countWords(data)
		chars += countChars(data)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error Opening the file", err)
	}

}

func countWords(data string) int {
	return len(strings.Fields(data))
}

func countChars(data string) int {
	return len(data)
}

func main1() {
	var fileName string = os.Args[1]

	readFile(fileName)

	fmt.Printf("lines: %v\n", lines)
	fmt.Printf("words: %v\n", words)
	fmt.Printf("chars: %v\n", chars)

}
