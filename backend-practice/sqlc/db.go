package main

import (
	"database/sql"
	"fmt"
	"log"

	"prac/db/sqlc"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
	Q  *sqlc.Queries
}

func NewDB(dsn string) *DB {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("ping failed: " + err.Error())
	}

	fmt.Println("db ping success")

	return &DB{
		DB: db,
		Q:  sqlc.New(db),
	}
}
