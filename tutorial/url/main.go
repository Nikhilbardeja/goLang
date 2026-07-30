package main

import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "https://jsonplaceholder.typicode.com/posts?k=2"

	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Could not build URL", err)
		return
	}
	fmt.Println(parsedUrl) // it is has a type of url
	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.RawQuery)
	fmt.Println(parsedUrl.Host)

	//  convert url to string

	strUrl := parsedUrl.String()

	fmt.Println(strUrl)
}
