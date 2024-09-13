package database

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect() *sql.DB {
    // Define MySQL connection string
    dsn := "root:root@tcp(127.0.0.1:3306)/go_connection"

    // Open the connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }
    // defer db.Close()

    // Check the connection
    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected to MySQL!")

	return db
}

func InsertMySqlData(db *sql.DB, data string) {
	defer db.Close()

	// Insert data into the table
    query := "INSERT INTO website_data (website_name, data) VALUES (?, ?)"
    result, err := db.Exec(query, "Gitlab", data)
    if err != nil {
        log.Fatal(err)
    }

    // Get the last inserted ID
    lastInsertID, err := result.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }

    // Print the last inserted ID
    fmt.Printf("Successfully inserted user with ID: %d\n", lastInsertID)
}