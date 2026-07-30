package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

var templates *template.Template

func init() {
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Failed to parse templates:", err)
	}
}

type PageData struct {
	Title   string
	Heading string
	Time    string
	Items   []string
}

func initRender() {
	Mux.HandleFunc("/temp", render)
}

func render(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var data PageData = PageData{
		Title:   "Go Web Server",
		Heading: "Welcome to our Go-powered server",
		Time:    time.Now().Format(time.RFC1123),
		Items:   []string{"Fast execution", "Memory safe", "Built-in concurrency"},
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.ExecuteTemplate(w, "temp.html", data); err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
