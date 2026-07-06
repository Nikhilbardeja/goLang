package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reading()
}

// to read whole file at once may cause memory error
func reading() {
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Could not read file")
		return
	}
	fmt.Println(string(content))
}

func readingByBuffer() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	// create a buffer to read the file

	buffer := make([]byte, 1024)

	var result string
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading the file", err)
		}

		// join the string buffers
		result = strings.Join([]string{result, string(buffer[:n])}, "")

	}

	fmt.Println(result)
}

func writing() {

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating the file", err)
		return
	}
	content := "Nach ke dikhao re"
	bytes, errW := io.WriteString(file, content)
	if errW != nil {
		fmt.Println("Error while writing the file", errW)
		return
	}
	fmt.Println(bytes)
	defer file.Close()
}
