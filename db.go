package main

import (
	"database/sql"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	_ "github.com/mattn/go-sqlite3"
)

// Opens the database
func openDB() (*sql.DB, error) {
	// db, err := sql.Open("sqlite3", "./data/bible-sqlite.db")
	db, err := sql.Open("sqlite3", "/Users/tvofik/Developer/bible/data/bible-sqlite.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetBooks() ([]list.Item, error) {
	var books []list.Item
	db, err := openDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT title_short, title_full, chapters  FROM book_info")
	if err != nil {
		return books, err
	}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.title, &book.fullTitle, &book.chapters)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}
	return books, err
}
