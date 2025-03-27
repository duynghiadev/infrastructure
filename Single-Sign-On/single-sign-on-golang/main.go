package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

// Fake user database (loaded from JSON)
var users []User

// Load users from JSON
func loadUsers() {
	file, _ := os.Open("users.json")
	defer file.Close()
	data, _ := io.ReadAll(file)
	json.Unmarshal(data, &users)
	fmt.Println("users:", users)
}

// Login handler (simulating SSO login)
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Fake login: Always return first user from JSON
	user := users[0]

	response := map[string]string{
		"message": "Login successful",
		"token":   user.Token,
	}
	json.NewEncoder(w).Encode(response)
}

// Profile handler (requires authentication)
func profileHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	for _, user := range users {
		if user.Token == token {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func main() {
	loadUsers()
	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler).Methods("GET")
	r.Handle("/profile", AuthMiddleware(http.HandlerFunc(profileHandler))).Methods("GET")

	fmt.Println("Server running on :9999")
	err := http.ListenAndServe(":9999", r)
	if err != nil {
		log.Fatal(err)
	}

}
