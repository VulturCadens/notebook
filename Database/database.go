package main

import (
	"database/sql" // https://golang.org/pkg/database/sql/
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3" // https://github.com/mattn/go-sqlite3
)

type book struct {
	author string
	title  string
	year   int
}

var books = []book{
	book{"Victor Hugo", "Les Misérables", 1862},
	book{"Victor Hugo", "The Hunchback of Notre-Dame", 1831},
	book{"Nikolai Gogol", "Dead Souls", 1842},
	book{"Fyodor Dostoevsky", "Crime and Punishment", 1866},
	book{"Leo Tolstoy", "War and Peace", 1869},
	book{"Leo Tolstoy", "Anna Karenina", 1877},
	book{"Alexandre Dumas", "The Count of Monte Cristo", 1844},
}

func createTable(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			author TEXT,
			title TEXT,
			year INTEGER,
			created DATETIME
		);`)

	if err != nil {
		log.Fatal(err)
	}
}

func countRows(db *sql.DB) int {
	var count int

	_ = db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)

	return count
}

func addRow(db *sql.DB, author, title string, year int) {
	statement, err := db.Prepare("INSERT INTO books(author, title, year, created) VALUES(?, ?, ?, ?)")

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	_, err = statement.Exec(author, title, year, time.Now())

	if err != nil {
		log.Fatal(err)
	}
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

	for _, book := range books {
		addRow(db, book.author, book.title, book.year)
	}

	fmt.Printf("There is %d rows in the database. \n", countRows(db))
}
