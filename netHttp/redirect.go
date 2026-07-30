package main

import "net/http"

func oldTemp(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/temp", http.StatusMovedPermanently)

	// w.Header().Set("Location", "/new-url")
	// w.WriteHeader(http.StatusMovedPermanently) // 301
}
