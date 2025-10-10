package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Book struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Published int       `json:"published"`
	Pages     int       `json:"pages"`
	Genres    []string  `json:"genres"`
	Rating    float32   `json:"rating"`
	Version   int32     `json:"version"`
	CreatedAt time.Time `json:"created_at"`
}

type BookModel struct {
	DB *pgxpool.Pool
}

func (b *BookModel) Insert(book *Book) error {
	query := `
		INSERT INTO books (title, published, pages, genres, rating)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version;
	`

	// CockroachDB natively supports string arrays via []string
	err := b.DB.QueryRow(
		context.Background(),
		query,
		book.Title,
		book.Published,
		book.Pages,
		book.Genres, // no pq.Array() needed
		book.Rating,
	).Scan(&book.ID, &book.CreatedAt, &book.Version)

	if err != nil {
		return fmt.Errorf("insert book: %w", err)
	}

	return nil
}

func (b BookModel) Get(id string) (*Book, error) {
	query := `
		SELECT id, created_at, title, published, pages, genres, rating, version
		FROM books
		WHERE id = $1
	`

	var book Book

	err := b.DB.QueryRow(context.Background(), query, id).Scan(
		&book.ID,
		&book.CreatedAt,
		&book.Title,
		&book.Published,
		&book.Pages,
		&book.Genres,
		&book.Rating,
		&book.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, fmt.Errorf("query error: %w", err)
		}
	}

	return &book, nil
}

func (b BookModel) Update(book *Book) error {
	query := `
		UPDATE books
		SET title = $1,
		    published = $2,
		    pages = $3,
		    genres = $4,
		    rating = $5,
		    version = version + 1
		WHERE id = $6 AND version = $7
		RETURNING version
	`

	args := []interface{}{
		book.Title,
		book.Published,
		book.Pages,
		book.Genres, // CockroachDB supports arrays directly
		book.Rating,
		book.ID,
		book.Version,
	}

	// Return the incremented version (CockroachDB supports RETURNING natively)
	err := b.DB.QueryRow(context.Background(), query, args...).Scan(&book.Version)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("book with ID %s not found or version mismatch", book.ID)
		}
		return fmt.Errorf("failed to update book: %w", err)
	}

	return nil
}

func (b BookModel) Delete(id string) error {

	query := `
		DELETE FROM books
		WHERE id = $1
	`

	// Exec executes the query without returning rows
	result, err := b.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (b BookModel) GetAll() ([]*Book, error) {
	query := `
		SELECT 
			id, 
			created_at, 
			title, 
			published, 
			pages, 
			genres, 
			rating, 
			version
		FROM books
		ORDER BY id;
	`

	rows, err := b.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*Book

	for rows.Next() {
		var book Book

		err := rows.Scan(
			&book.ID,
			&book.CreatedAt,
			&book.Title,
			&book.Published,
			&book.Pages,
			&book.Genres,
			&book.Rating,
			&book.Version,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
