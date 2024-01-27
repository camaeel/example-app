package database

import (
	"database/sql"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

func SetupDriver() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	u, err := url.Parse(connStr)
	if err != nil {
		return nil, err
	}
	// connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	db, err := sql.Open(u.Scheme, connStr)
	return db, err
}
