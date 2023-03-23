package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func main() {
	// Open the database file (creating it if it doesn't exist)
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the "users" table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Prompt the user for input
	var username, email string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)

	// Insert the user into the database
	result, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User created with ID", result.LastInsertId())

	// Retrieve all users from the database
	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the rows and print each user
	fmt.Println("All users:")
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", user.ID, user.Username, user.Email)
	}
	err = rows.Err()
	if err != nil {
	}
}
