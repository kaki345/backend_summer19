package mydatabase

import (
	"crypto/sha256"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func BuildDB() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func InitializeUsersTable(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS "Users" (
			"id" integer PRIMARY KEY AUTOINCREMENT, 
			"username" varchar(255) NOT NULL UNIQUE, 
			"fullname" varchar(255),
			"email" varchar(255) NOT NULL UNIQUE,
			"password" varchar(255) NOT NULL
		)`,
	)
	if err != nil {
		log.Fatal(err)
	}

}

func StoreUserData(db *sql.DB, u User) {
	cmdInsert := `INSERT INTO Users (username, fullname, email, password) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(cmdInsert, u.Username, u.Fullname, u.Email, u.Password)
	if err != nil {
		log.Fatal(err)
	}
}

func EncryptPass(pass string) string {
	encrypted := sha256.Sum256([]byte(pass))
	newpass := string(encrypted[:])
	return newpass
}