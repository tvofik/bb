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

/* Get all the Translations available */
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

// Get the passage
// TODO Not sure which would work best
func (m Model) GetChapter(bookId, chapter, translation string) ([]Verse, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT v, t FROM %s WHERE b = ? AND c = ?`
	query = fmt.Sprintf(query, translation)

	rows, err := db.Query(query, bookId, chapter)
	if err != nil {
		return nil, err
	}

	var currentVerse Verse
	var fullChapter []Verse

	for rows.Next() {
		err := rows.Scan(&currentVerse.number, &currentVerse.text)
		if err != nil {
			return nil, err

		}
		fullChapter = append(fullChapter, currentVerse)
	}
	return fullChapter, nil
}

/* Helps gets the exact table for the translation selected */
func (m Model) GetCurrentTranslationTable(translation string) (string, error) {
	db, err := openDB()
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Get the table name of the Translation selected
	var tableName string
	row := db.QueryRow("SELECT table_name FROM bible_version_key WHERE abbreviation = ?", translation)
	err = row.Scan(&tableName)
	if err != nil {
		return "", err
	}
	return tableName, nil
}
