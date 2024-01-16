package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "your_host"
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "your_database_name"
)

func main() {
	// Connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	// Create a table
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			email VARCHAR(100) NOT NULL
		);
	`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully")

	// Insert data into the table
	insertQuery := `
		INSERT INTO users (username, email) VALUES ($1, $2)
		RETURNING id;
	`

	var userID int
	err = db.QueryRow(insertQuery, "john_doe", "john@example.com").Scan(&userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted row with ID %d\n", userID)

	// Query data from the table
	selectQuery := `
		SELECT id, username, email FROM users;
	`

	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var username, email string
		err := rows.Scan(&id, &username, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", id, username, email)
	}
}
