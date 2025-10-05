package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// We will define a lightweight version of addActor for sqlmock testing
func addActorWithDB(ctx context.Context, db *sql.DB, first, last, email, country string) (string, error) {
	const query = `
		INSERT INTO actor (first_name, last_name, email, country)
		VALUES ($1, $2, $3, $4)
		RETURNING actor_id;
	`

	var id string
	err := db.QueryRowContext(ctx, query, first, last, email, country).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("insert failed: %w", err)
	}
	return id, nil
}

func TestAddActor(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	expectedID := "123e4567-e89b-12d3-a456-426614174000"

	mock.ExpectQuery(`INSERT INTO actor`).
		WithArgs("John", "Doe", "john@example.com", "USA").
		WillReturnRows(sqlmock.NewRows([]string{"actor_id"}).AddRow(expectedID))

	got, err := addActorWithDB(context.Background(), db, "John", "Doe", "john@example.com", "USA")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != expectedID {
		t.Errorf("want %s, got %s", expectedID, got)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestAddActor_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectQuery(`INSERT INTO actor`).
		WithArgs("Alice", "Smith", "alice@example.com", "UK").
		WillReturnError(fmt.Errorf("duplicate email"))

	_, err := addActorWithDB(context.Background(), db, "Alice", "Smith", "alice@example.com", "UK")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
