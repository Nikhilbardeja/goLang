package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func initRoutes() {
	Mux.HandleFunc("POST /dynamic-json", handleDynamicJson)
	Mux.HandleFunc("POST /static-json", handleStaticJson)
}

func handleDynamicJson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data map[string]any

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println(data)

}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleStaticJson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data User

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(data)
	w.WriteHeader(http.StatusCreated)
}
