package main

import (
	"database/sql" // https://golang.org/pkg/database/sql/
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3" // https://github.com/mattn/go-sqlite3
)

func createTable(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			author TEXT,
			title TEXT,
			created DATETIME
		);`)

	if err != nil {
		log.Fatal(err)
	}
}

func countBooks(db *sql.DB) int {
	var count int

	_ = db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)

	return count
}

func addBook(db *sql.DB, author, title string) error {
	statement, err := db.Prepare("INSERT INTO books(author, title, created) VALUES(?, ?, ?)")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(author, title, time.Now())

	if err != nil {
		return err
	}

	return nil
}

func main() {
	os.Remove("./test.db")

	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.SetMaxOpenConns(1)

	createTable(db)

	if err := addBook(db, "Victor Hugo", "Les Misérables"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There is %d book(s). \n", countBooks(db))
}
