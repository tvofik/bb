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

/* Get Books */
func GetBooks() ([]list.Item, error) {
	var books []list.Item
	db, err := openDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT book_id, title_short, title_full, chapters  FROM book_info")
	if err != nil {
		return books, err
	}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.id, &book.title, &book.fullTitle, &book.chapters)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}
	return books, err
}

/* Get Translation */
func GetTranslations() ([]list.Item, error) {
	var translations []list.Item
	db, err := openDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT abbreviation, version FROM bible_version_key")
	if err != nil {
		return translations, err
	}
	for rows.Next() {
		var translation Translation
		err := rows.Scan(&translation.abbreviation, &translation.name)
		if err != nil {
			return translations, err
		}
		translations = append(translations, translation)
	}
	return translations, err
}
