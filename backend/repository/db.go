package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(databaseURL string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * 60 * 60) // 5 hours

	log.Println("Database connected successfully")
	return db, nil
}

func GetDB() *sql.DB {
	return db
}