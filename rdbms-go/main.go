package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
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
	pool, err := pgxpool.NewWithConfig(context.Background(), conn)
	if err != nil {
		log.Fatal("Unable to create pool:", err)
	}
	defer pool.Close()

	var now time.Time
	err = pool.QueryRow(context.Background(), "SELECT now()").Scan(&now)
	if err != nil {
		log.Fatal("Query failed: ", err)
	}

	fmt.Println("Connected, server time is:", now)

}
