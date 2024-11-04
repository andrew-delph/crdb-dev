package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

var username = "postgres"
var password = "postgres"
var database = "test_script"

func main() {
	// Define the connection string
	// connStr := "postgres://username:password@localhost:5432/database_name"
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", username, password, database)

	// Connect to the database
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Define the SQL query to create a table
	query := `
		CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title VARCHAR(100) NOT NULL,
			author VARCHAR(100) NOT NULL,
			quantity INTEGER NOT NULL
		);
	`

	// Execute the query to create the table
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully.")
}
