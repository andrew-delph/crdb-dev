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
var err error

func main() {
	// Define the connection string

	// Connect to the database
	// conn := postgresConn()
	conn := cockroachConn()
	defer conn.Close(context.Background())

	// test(conn)
	testIdentity(conn)

}

func test(conn *pgx.Conn) {
	defer println("finished test()")
	// Define the SQL query to create a table
	drop := `
		DROP SEQUENCE IF EXISTS test_sequence;
		`

	// Execute the query to create the table
	_, err = conn.Exec(context.Background(), drop)
	if err != nil {
		log.Fatal(err)
	}

	create := `
		CREATE SEQUENCE test_sequence MAXVALUE 33;
		`

	// Execute the query to create the table
	_, err = conn.Exec(context.Background(), create)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 4; i++ {
		// Increment a sequence and print the resulting value
		var seqValue int64
		err = conn.QueryRow(context.Background(), "SELECT nextval('test_sequence');").Scan(&seqValue)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sequence value: %d\n", seqValue)
	}
}

func testIdentity(conn *pgx.Conn) {

	drop := `
		DROP TABLE IF EXISTS test_identity;
		`

	// Execute the query to create the table
	_, err = conn.Exec(context.Background(), drop)
	if err != nil {
		log.Fatal(err)
	}

	create := `
		CREATE TABLE test_identity (id bigint GENERATED ALWAYS AS IDENTITY (MAXVALUE 10), temp int);
		`

	// Execute the query to create the table
	_, err = conn.Exec(context.Background(), create)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 4; i++ {
		_, err = conn.Exec(context.Background(), `INSERT INTO test_identity (temp) VALUES (1)`)
		if err != nil {
			log.Fatalf("Unable to insert row: %v\n", err)
		}
	}

	rows, err := conn.Query(context.Background(), `SELECT id FROM test_identity`)
	if err != nil {
		log.Fatalf("Unable to retrieve rows: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Fatalf("Unable to scan row: %v\n", err)
		}
		fmt.Printf("ID: %d\n", id)
	}

}

func postgresConn() *pgx.Conn {
	println("postgresConn")
	connStr := "postgres://postgres:postgres@localhost:5432/test_script"

	// Connect to the database
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func cockroachConn() *pgx.Conn {
	println("cockroachConn")
	connStr := "postgresql://root@127.0.0.1:26257/movr?options=-ccluster%3Ddemoapp&sslmode=disable"

	// Connect to the database
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
