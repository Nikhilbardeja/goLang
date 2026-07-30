package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func main() {
	godotenv.Load()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	var err error
	Db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}
	defer Db.Close()

	if err := Db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/users/", handlerUsers)
	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("/user/transfer/", transer)

	log.Println("Server :5000")
	http.ListenAndServe(":5000", mux)
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Json", 400)
		return
	}

	result, err := Db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(user)
}

func handlerUsers(w http.ResponseWriter, r *http.Request) {
	var urlList []string = strings.Split(r.URL.Path, "/")

	var id string = urlList[2]

	log.Println("User Id: ", id)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := Db.Query("SELECT id, name, email FROM users WHERE id = ?", id)
	if err != nil {
		log.Printf("DB query error : %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Printf("Scan error: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		users = append(users, u)

	}
	if err := rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}
