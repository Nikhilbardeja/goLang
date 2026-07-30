package main

import (
	"encoding/json"
	"net/http"
)

func initRoutesJson() {
	Mux.HandleFunc("/send-json", sendJsonResponse)
}

func sendJsonResponse(w http.ResponseWriter, _ *http.Request) {
	var user map[string]any = map[string]any{"name": "Nikhil Bardeja", "age": 20}
	var response map[string]any = map[string]any{"status": "success", "message": "Reponse okay", "data": user}
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(response)

}
