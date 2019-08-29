package main

import (
	"net/http"
	"fmt"
	"log"
	//"github.com/mattn/go-sqlite3"
	"encoding/json"
	//. "./mydatabase"
)

var u User

func main () {
	db := BuildDB()
	InitializeUserDB(db)

	http.HandleFunc("/", index)								// Show the login page					
	http.HandleFunc("/login", login)						// Show the login page
	http.HandleFunc("/logout", logout)						// Logs the user out
	http.HandleFunc("/signup", signup)						// Show the signup page
	//http.HandleFunc("/signup_account", signupaccount)		// Create the user account
	//http.HandleFunc("/authenticate", authenticate)		// Authenticate the user given the email and password

	log.Fatal(http.ListenAndServe(":80", nil))
	fmt.Println("Server is Started...")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page!")
	fmt.Println("Access to HomePage")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Page!")
	fmt.Println("Access to LoginPage")

}

func logout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Signup Page!")
	fmt.Println("Access to SignupPage")

	fmt.Println(r.Body)

	json.NewDecoder(r.Body).Decode(&u)

	if u.Username == "" {
		error.Message = "Please input user name."
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	if u.Email == "" {
		error.Message = "Please input email."
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	if u.Password == "" {
		error.Message = "Please input password"
		errorInResponse(w, http.StatusBadRequest, error)
		return
	}

	u.Password = EncryptPass(u.Password)
	StoreUserData(db, u)

}
