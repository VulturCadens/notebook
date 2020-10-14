package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func countBooks(db *sql.DB) int {
	var count int

	_ = db.QueryRow("SELECT COUNT(*) FROM messages").Scan(&count)

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

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			author TEXT,
			title TEXT,
			created DATETIME
		);`)

	if err != nil {
		log.Fatal(err)
	}

	if countBooks(db) == 0 {
		if err = addBook(db, "Victor Hugo", "Les Misérables"); err != nil {
			log.Fatal(err)
		}
	}
}
