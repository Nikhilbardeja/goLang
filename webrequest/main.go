package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	respone, err := http.Get("http://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println("Error making request to the server", err)
		return
	}
	defer respone.Body.Close()

	data, errr := io.ReadAll(respone.Body)

	if errr != nil {
		fmt.Println("Error while reading response", errr)
		return
	}

	var _ string = string(data) // to read data as a string

	var mapData []map[string]any
	errM := json.Unmarshal(data, &mapData)
	if errM != nil {
		fmt.Println("Error converting to the map", errM)
		return
	}
	for i, item := range mapData {
		fmt.Printf("Item %d: %v: %s\n", i, item["id"], item["title"])
	}

}
