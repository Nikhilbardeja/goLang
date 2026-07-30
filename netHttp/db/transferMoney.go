package main

import (
	"fmt"
	"net/http"
	"strings"
)

func transer(w http.ResponseWriter, r *http.Request) { // /user/transfer/<sendId>/<receiveId>
	var urlList []string = strings.Split(r.URL.Path, "/")
	var sendId string = urlList[3]
	var recvId string = urlList[4]

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(urlList) != 5 { // ["user", "transfer", "1", "2"]
		http.Error(w, "URL should be /user/transfer/<sendId>/<receiveId>", http.StatusBadRequest)
		return
	}

	tx, err := Db.Begin()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE users SET balance = balance - 100 WHERE id = ?", sendId)
	if err != nil {
		http.Error(w, "Failed to debit sender", 500)
		return
	}

	_, err = tx.Exec("UPDATE users SET balance = balance + 100 WHERE id = ?", recvId)
	if err != nil {
		http.Error(w, "Failed to credit receiver", 500)
		return // defer will rollback
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to trasfer money", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "<h1>Money transfered to account %s</h1>", recvId)
}
