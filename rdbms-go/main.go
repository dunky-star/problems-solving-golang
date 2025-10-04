package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool *pgxpool.Pool // Thread-Safe and Efficient Global Singleton
	o    sync.Once
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
	// Close pool when main exits, not inside Run()
	defer pool.Close()

	// Insert a new actor
	// actorID, err := addActor("Geoffrey", "Opiyo", "dunkygeoffrey39@gmail.com", "USA")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("New actor ID:", actorID)

	// Retrieve the actor
	actor, err := GetActor("dunkygeoffrey39@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Retrieved actor: %+v\n", actor)

	// Retrieve all actors
	actors, err := GetAllActors()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All actors:")
	for _, a := range actors {
		fmt.Printf("%+v\n", a)
	}
}

func Run() error {
	o.Do(func() {
		// DSN string
		dsn := "postgres://appuser:S3cure!@localhost:26257/bank?sslmode=require"

		// Configure pool with sensible defaults
		conn, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Fatal("Unable to parse config:", err)
		}

		// Optional: tweak pool settings (best practice for prod)
		conn.MaxConns = 4                         // max number of connections
		conn.MinConns = 2                         // keep warm connections
		conn.MaxConnLifetime = time.Hour          // recycle connections hourly
		conn.HealthCheckPeriod = 30 * time.Second // ping connections

		// Connect pool
		p, err := pgxpool.NewWithConfig(context.Background(), conn)
		if err != nil {
			log.Fatal("Unable to create pool:", err)
		}

		pool = p // assign to the global pool

		var now time.Time
		err = pool.QueryRow(context.Background(), "SELECT now()").Scan(&now)
		if err != nil {
			log.Fatal("Query failed: ", err)
		}

		fmt.Println("Connected, server time is:", now)
	})
	return nil
}

func addActor(firstName, lastName, email, country string) (string, error) {
	var actorID string

	query := `
		INSERT INTO actor (first_name, last_name, email, country)
		VALUES ($1, $2, $3, $4)
		RETURNING actor_id;
	`

	err := pool.QueryRow(context.Background(), query, firstName, lastName, email, country).Scan(&actorID)
	if err != nil {
		return "", fmt.Errorf("failed to insert actor: %w", err)
	}
	return actorID, nil
}

// Actor struct should match your table fields
type Actor struct {
	ActorID   string
	FirstName string
	LastName  string
	Email     string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetActor(email string) (*Actor, error) {
	query := `
		SELECT actor_id, first_name, last_name, email, country, created_at, updated_at
		FROM actor
		WHERE email = $1;
	`

	var a Actor
	err := pool.QueryRow(context.Background(), query, email).
		Scan(&a.ActorID, &a.FirstName, &a.LastName, &a.Email, &a.Country, &a.CreatedAt, &a.UpdatedAt)

	if err != nil {
		return nil, err // returns pgx.ErrNoRows if not found
	}
	return &a, nil
}

// GetAllActors retrieves all actors from the table
func GetAllActors() ([]Actor, error) {
	query := `
		SELECT actor_id, first_name, last_name, email, country, created_at, updated_at
		FROM actor
		ORDER BY created_at DESC;
	`

	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query actors: %w", err)
	}
	defer rows.Close()

	var actors []Actor

	for rows.Next() {
		var a Actor
		err := rows.Scan(
			&a.ActorID,
			&a.FirstName,
			&a.LastName,
			&a.Email,
			&a.Country,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan actor row: %w", err)
		}
		actors = append(actors, a)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error while iterating rows: %w", rows.Err())
	}

	return actors, nil
}
