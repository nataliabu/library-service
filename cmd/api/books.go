package main

import (
	"context"
	"log"

	"github.com/nataliabu/library-service/internal/database"
)

func listBooksDB(db *database.DB, ctx context.Context) ([]database.Book, error) {
	data := []database.Book{}
	query := `SELECT * FROM books;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int32
	var title string
	var author string
	var isbn string
	var issueYear int32
	var available bool
	for rows.Next() {
		err := rows.Scan(&id, &title, &author, &isbn, &issueYear, &available)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, database.Book{id, title, author, isbn, issueYear, available})
	}
	return data, nil
}

func getBookByIsbnDB(db *database.DB, ctx context.Context, isbn *string) (*database.Book, error) {
	query := `SELECT * FROM books
    WHERE isbn = $1 LIMIT 1;`
	var id int32
	var title string
	var author string
	var issueYear int32
	var available bool
	err := db.QueryRow(query, isbn).Scan(&id, &title, &author, isbn, &issueYear, &available)
	if err != nil {
		log.Fatal(err)
	}
	book := database.Book{id, title, author, *isbn, issueYear, available}
	return &book, nil
}

func getBookByIdDB(db *database.DB, ctx context.Context, id *int32) (*database.Book, error) {
	query := `SELECT * FROM books
    WHERE id = $1 LIMIT 1;`
	var title string
	var author string
	var isbn string
	var issueYear int32
	var available bool
	err := db.QueryRow(query, id).Scan(id, &title, &author, &isbn, &issueYear, &available)
	if err != nil {
		log.Fatal(err)
	}
	book := database.Book{*id, title, author, isbn, issueYear, available}
	return &book, nil
}

func addBookDB(db *database.DB, ctx context.Context, book *database.Book) (*database.Book, error) {
	query := `INSERT INTO books (title, author, isbn, issue_year, available)
	  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var pk int32
	err := db.QueryRow(query, book.Title, book.Author, book.Isbn, book.IssueYear, true).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	newBook := database.Book{
		ID:        pk,
		Title:     book.Title,
		Author:    book.Author,
		Isbn:      book.Isbn,
		IssueYear: book.IssueYear,
		Available: true,
	}
	return &newBook, nil
}

func removeBookDB(db *database.DB, ctx context.Context, id *int32) error {
	query := `DELETE FROM books WHERE id = $1;`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
