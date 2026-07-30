package main

import (
	"fmt"
	"log"
	"net/http"
)

var Mux *http.ServeMux = http.NewServeMux()

func main() {

	// Serve static files with prefix stripping
	fs := http.FileServer(http.Dir("./static"))
	Mux.Handle("/static/", http.StripPrefix("/static/", fs))

	Mux.HandleFunc("/", handelHelloWorld)

	initRoutes()
	initRoutesJson()
	initRender()

	if err := http.ListenAndServe(":5000", Mux); err != nil {
		log.Fatal("Error starting the server", err)
	}

}

func handelHelloWorld(wr http.ResponseWriter, r *http.Request) {

	var path string = r.URL.Path
	fmt.Println("handelHelloWorld hit by:", path)

	if path != "/" {
		http.NotFound(wr, r)
		return
	}

	wc, err := wr.Write([]byte("<h1>Hello World</h1>"))

	if err != nil {
		panic("Error writing the reponse")
	}
	fmt.Println("Bytes Written:::: ", wc)
}
