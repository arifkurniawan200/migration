package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"migration/config"
)

var (
	db  *sql.DB
	err error
)

func NewDatabase(cfg config.Database) (*sql.DB, error) {
	dbUrl := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Name, cfg.Username, cfg.Password)
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		return db, err
	}

	if cfg.ActivePool {
		// Set the maximum and minimum connection pool size
		db.SetMaxIdleConns(cfg.MaxPool) // Maximum number of idle connections
		db.SetMaxOpenConns(cfg.MinPool) // Maximum number of open connections

	}
	// Test the database connection
	if err = db.Ping(); err != nil {
		return db, err
	}
	return db, err
}
