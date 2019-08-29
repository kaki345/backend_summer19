package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	//"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


type Users struct {
    ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type Error struct {
    Message string `json:"message"`
}

func errorInResponse(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
	return
}

func signup(w http.ResponseWriter, r *http.Request) {
	var users Users
	var error Error

	fmt.Println(r.Body)

	json.NewDecoder(r.Body).Decode(&users)

	if users.Username == "" {
		error.Message = "Please input user name."
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	if users.Email == "" {
		error.Message = "Please input your email."
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	if users.Password == "" {
		error.Message = "Please input password"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	fmt.Println(users)

}

func main() {
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/signup", signup)

}